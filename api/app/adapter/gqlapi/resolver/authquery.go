package resolver

import (
	"github.com/short-d/app-template/backend/app/usecase/changelog"
)

// AuthQuery represents GraphQL query resolver that acts differently based
// on the identify of the user
type AuthQuery struct {
	authToken *string
	changeLog changelog.ChangeLog
}

// ChangeLog retrieves full ChangeLog from persistent storage
func (v AuthQuery) ChangeLog() (ChangeLog, error) {
	changeLog, err := v.changeLog.GetChangeLog()
	if err != nil {
		return ChangeLog{}, err
	}

	return newChangeLog(changeLog), err
}

func newAuthQuery(
	authToken *string,
	changeLog changelog.ChangeLog,
) AuthQuery {
	return AuthQuery{
		authToken: authToken,
		changeLog: changeLog,
	}
}
