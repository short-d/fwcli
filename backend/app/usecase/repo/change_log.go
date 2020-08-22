package repo

import (
	"fwcli/app/entity"
)

// ChangeLog accesses changelog from storage, such as database.
type ChangeLog interface {
	GetChangeLog() ([]entity.Change, error)
}
