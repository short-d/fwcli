package resolver

import "fwcli/app/usecase/changelog"

// Resolver contains GraphQL request handlers.
type Resolver struct {
	Query
	Mutation
}

// NewResolver creates a new GraphQL resolver.
func NewResolver(changeLog changelog.ChangeLog) Resolver {
	return Resolver{
		Query:    newQuery(changeLog),
		Mutation: newMutation(),
	}
}
