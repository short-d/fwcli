package resolver

import "github.com/short-d/app-template/backend/app/usecase/changelog"

// Query represents GraphQL query resolver
type Query struct {
	changeLog changelog.ChangeLog
}

// AuthQueryArgs represents possible parameters for AuthQuery endpoint
type AuthQueryArgs struct {
	AuthToken *string
}

// AuthQuery extracts user information from authentication token
func (q Query) AuthQuery(args *AuthQueryArgs) (*AuthQuery, error) {
	authQuery := newAuthQuery(args.AuthToken, q.changeLog)
	return &authQuery, nil
}

func newQuery(changeLog changelog.ChangeLog) Query {
	return Query{changeLog: changeLog}
}
