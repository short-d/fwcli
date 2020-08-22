package repo

import (
	"github.com/short-d/app-template/backend/app/entity"
)

var _ ChangeLog = (*ChangeLogFake)(nil)

// ChangeLogFake represents in memory implementation of ChangeLog repository
type ChangeLogFake struct {
	changeLog []entity.Change
}

// GetChangeLog fetches full ChangeLog from memory
func (c ChangeLogFake) GetChangeLog() ([]entity.Change, error) {
	return c.changeLog, nil
}

// NewChangeLogFake creates ChangeLogFake
func NewChangeLogFake(changeLog []entity.Change) ChangeLogFake {
	return ChangeLogFake{
		changeLog: changeLog,
	}
}
