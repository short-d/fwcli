//+build wireinject

package dep

import (
	"github.com/google/wire"
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
