//+build wireinject

package dep

import (
	"github.com/google/wire"
	"github.com/short-d/app/fw/db"
	"github.com/short-d/app/fw/env"
)

// InjectEnv creates Environment with configured dependencies.
func InjectEnv() env.Env {
	wire.Build(
		wire.Bind(new(env.Env), new(env.GoDotEnv)),
		env.NewGoDotEnv,
	)
	return env.GoDotEnv{}
}

// InjectDBConnector creates DBConnector with configured dependencies.
func InjectDBConnector() db.Connector {
	wire.Build(
		wire.Bind(new(db.Connector), new(db.PostgresConnector)),
		db.NewPostgresConnector,
	)
	return db.PostgresConnector{}
}

// InjectDBMigrationTool creates DBMigrationTool with configured dependencies.
func InjectDBMigrationTool() db.MigrationTool {
	wire.Build(
		wire.Bind(new(db.MigrationTool), new(db.PostgresMigrationTool)),
		db.NewPostgresMigrationTool,
	)
	return db.PostgresMigrationTool{}
}
