package migrate

import "github.com/iamviniciuss/golang-migrations/src/adapter"

type Migration struct {
	Version     uint64
	Description string
	Handler     adapter.MigrationHandler
}
