package changelog

import (
	"github.com/short-d/app-template/backend/app/entity"
	"github.com/short-d/app-template/backend/app/usecase/repo"
)

// Persist retrieves change log from and saves changes to persistent data store.
type ChangeLog struct {
	changeLogRepo repo.ChangeLog
}

// GetChangeLog retrieves full ChangeLog from persistent data store.
func (p ChangeLog) GetChangeLog() ([]entity.Change, error) {
	return p.changeLogRepo.GetChangeLog()
}

// NewChangeLog creates ChangeLog
func NewChangeLog(
	changeLog repo.ChangeLog,
) ChangeLog {
	return ChangeLog{
		changeLogRepo: changeLog,
	}
}
