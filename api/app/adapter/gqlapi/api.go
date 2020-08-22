package gqlapi

import (
	"github.com/short-d/app-template/backend/app/adapter/gqlapi/resolver"
	"github.com/short-d/app/fw/filesystem"
	"github.com/short-d/app/fw/graphql"
)

// NewGraphQLAPI creates GraphQL API config
func NewGraphQLAPI(
	schemaPath string,
	fileSystem filesystem.FileSystem,
	r resolver.Resolver,
) (graphql.API, error) {
	buf, err := fileSystem.ReadFile(schemaPath)
	if err != nil {
		return graphql.API{}, err
	}
	return graphql.API{
		Schema:   string(buf),
		Resolver: &r,
	}, nil
}
