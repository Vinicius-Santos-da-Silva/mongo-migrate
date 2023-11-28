package migrate

import (
	"sort"

	repository "github.com/iamviniciuss/golang-migrations/src/adapter"
)

func migrationSort(migrations []Migration) {
	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version
	})
}

func hasVersion(migrations []Migration, current repository.MigrationHandler) bool {
	for _, m := range migrations {
		if m.Description == current.GetName() && m.Version == current.GetVersion() {
			return true
		}
	}
	return false
}
