package container

type Orchestrator interface {
	ApplyConfig(configDir string) error
}
