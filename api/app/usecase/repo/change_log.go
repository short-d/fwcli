package repo

import (
	"github.com/short-d/app-template/backend/app/entity"
)

// ChangeLog accesses changelog from storage, such as database.
type ChangeLog interface {
	GetChangeLog() ([]entity.Change, error)
}
