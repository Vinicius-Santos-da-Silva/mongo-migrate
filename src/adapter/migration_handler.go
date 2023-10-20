package adapter

type MigrationHandler interface {
	GetVersion() uint64
	GetType() string
	GetName() string
	Up() error
	Down() error
}
