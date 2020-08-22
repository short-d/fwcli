package main

import (
	"github.com/short-d/app-template/backend/app/adapter/gqlapi"
	"github.com/short-d/app-template/backend/app/adapter/gqlapi/resolver"
	"github.com/short-d/app-template/backend/app/adapter/sqldb"
	"github.com/short-d/app-template/backend/app/usecase/changelog"
	"github.com/short-d/app-template/backend/dep"
	"github.com/short-d/app/fw/db"
	"github.com/short-d/app/fw/envconfig"
	"github.com/short-d/app/fw/filesystem"
	"github.com/short-d/app/fw/service"
)

func main() {
	// Load environmental variables
	env := dep.InjectEnv()
	env.AutoLoadDotEnvFile()
	envConfig := envconfig.NewEnvConfig(env)

	config := struct {
		DBHost     string `env:"DB_HOST" default:"localhost"`
		DBPort     int    `env:"DB_PORT" default:"5432"`
		DBUser     string `env:"DB_USER" default:"postgres"`
		DBPassword string `env:"DB_PASSWORD" default:"password"`
		DBName     string `env:"DB_NAME" default:"short"`
	}{}
	err := envConfig.ParseConfigFromEnv(&config)
	if err != nil {
		panic(err)
	}

	// Load connect to DB
	dbConfig := db.Config{
		Host:     config.DBHost,
		Port:     config.DBPort,
		User:     config.DBUser,
		Password: config.DBPassword,
		DbName:   config.DBName,
	}
	dbConnector := dep.InjectDBConnector()
	sqlDB, err := dbConnector.Connect(dbConfig)
	if err != nil {
		panic(err)
	}

	// Migrate DB tables
	dbMigrationTool := dep.InjectDBMigrationTool()
	err = dbMigrationTool.MigrateUp(sqlDB, "app/adapter/sqldb/migration")
	if err != nil {
		panic(err)
	}

	// Initialize GraphQL api
	fs := filesystem.NewLocal()
	changelogRepo := sqldb.NewChangeLogSQL(sqlDB)
	changeLog := changelog.NewChangeLog(changelogRepo)
	r := resolver.NewResolver(changeLog)
	schemaPath := "app/adapter/gqlapi/schema.graphql"
	api, err := gqlapi.NewGraphQLAPI(schemaPath, fs, r)
	if err != nil {
		panic(err)
	}

	// Start server
	graphqlService := service.
		NewGraphQLBuilder("Sample").
		Schema(api.Schema).
		Resolver(api.Resolver).
		Build()
	graphqlService.StartAndWait(8080)
}
