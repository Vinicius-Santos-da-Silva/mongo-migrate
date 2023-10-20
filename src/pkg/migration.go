package migrate

import "github.com/Vinicius-Santos-da-Silva/mongo-migrate/src/adapter"

type Migration struct {
	Version     uint64
	Description string
	Handler     adapter.MigrationHandler
}
