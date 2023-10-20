package migrate

import "sort"

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
