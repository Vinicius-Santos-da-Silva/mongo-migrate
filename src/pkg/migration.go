package migrate

import (
	"sort"

	"go.mongodb.org/mongo-driver/mongo"
)

type MigrationFunc func(db *mongo.Database) error

type Migration struct {
	Version     uint64
	Description string
	Handler     MigrationHandler
}

type MigrationHandler interface {
	GetVersion() uint64
	GetType() string
	GetName() string
	Up() error
	Down() error
}

func migrationSort(migrations []Migration) {
	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version
	})
}

func hasVersion(migrations []Migration, version uint64) bool {
	for _, m := range migrations {
		if m.Version == version {
			return true
		}
	}
	return false
}
