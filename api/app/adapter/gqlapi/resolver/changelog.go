package resolver

import (
	"github.com/short-d/app-template/backend/app/entity"
)

// ChangeLog retrieves full change log and the time when the user viewed it.
type ChangeLog struct {
	changeLog []Change
}

// Changes retrieves full change log
func (c ChangeLog) Changes() []Change {
	return c.changeLog
}

func newChangeLog(changeLog []entity.Change) ChangeLog {
	var changes []Change
	for _, v := range changeLog {
		changes = append(changes, newChange(v))
	}

	return ChangeLog{changeLog: changes}
}
