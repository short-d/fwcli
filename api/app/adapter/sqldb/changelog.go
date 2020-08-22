package sqldb

import (
	"database/sql"
	"fmt"

	"github.com/short-d/app-template/backend/app/adapter/sqldb/table"
	"github.com/short-d/app-template/backend/app/entity"
	"github.com/short-d/app-template/backend/app/usecase/repo"
)

var _ repo.ChangeLog = (*ChangeLogSQL)(nil)

// ChangeLogSQL accesses ChangeLog information in change_log table through SQL.
type ChangeLogSQL struct {
	db *sql.DB
}

// GetChangeLog retrieves full changelog from change_log table.
func (c ChangeLogSQL) GetChangeLog() ([]entity.Change, error) {
	statement := fmt.Sprintf(`
SELECT "%s","%s","%s","%s" 
FROM "%s";`,
		table.ChangeLog.ColumnID,
		table.ChangeLog.ColumnTitle,
		table.ChangeLog.ColumnSummaryMarkdown,
		table.ChangeLog.ColumnReleasedAt,
		table.ChangeLog.TableName,
	)

	rows, err := c.db.Query(statement)
	if err != nil {
		return []entity.Change{}, err
	}

	// the consumer of GetChangeLog expects empty slice instead of `nil` if there are no records
	changeLog := []entity.Change{}
	for rows.Next() {
		change := entity.Change{}
		err = rows.Scan(&change.ID, &change.Title, &change.SummaryMarkdown, &change.ReleasedAt)
		if err != nil {
			return changeLog, err
		}
		change.ReleasedAt = change.ReleasedAt.UTC()
		changeLog = append(changeLog, change)
	}

	return changeLog, nil
}

// NewChangeLogSQL creates ChangeLogSQL
func NewChangeLogSQL(db *sql.DB) ChangeLogSQL {
	return ChangeLogSQL{
		db: db,
	}
}
