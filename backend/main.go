package main

import (
	"fwcli/app/adapter/routes"
	"fwcli/dep"

	"github.com/short-d/app/fw/envconfig"
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
		DBName     string `env:"DB_NAME" default:"fwcli"`
	}{}
	err := envConfig.ParseConfigFromEnv(&config)
	if err != nil {
		panic(err)
	}

	// Start server
	routingService := service.
		NewRoutingBuilder("fwCLI").
		Routes(routes.NewRoutes()).
		Build()
	routingService.StartAndWait(8080)
}
