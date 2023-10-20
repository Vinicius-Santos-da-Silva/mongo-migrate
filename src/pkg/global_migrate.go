package migrate

import (
	"fmt"

	"github.com/Vinicius-Santos-da-Silva/mongo-migrate/src/repository"
)

var globalMigrate = NewMigrate(nil)

func internalRegister(migrationHandler MigrationHandler, skip int) error {
	if hasVersion(globalMigrate.migrations, migrationHandler.GetVersion()) {
		return fmt.Errorf("migration with version %v already registered", migrationHandler.GetVersion())
	}

	globalMigrate.migrations = append(globalMigrate.migrations, Migration{
		Version:     migrationHandler.GetVersion(),
		Description: migrationHandler.GetName(),
		Handler:     migrationHandler,
	})
	return nil
}

func Register(migrationHandler MigrationHandler) error {
	return internalRegister(migrationHandler, 2)
}

func RegisteredMigrations() []Migration {
	ret := make([]Migration, len(globalMigrate.migrations))
	copy(ret, globalMigrate.migrations)
	return ret
}

func SetRepository(repo repository.MigrationRepository) {
	globalMigrate.migrationRepository = repo
}

func SetMigrationsCollection(name string) {
	globalMigrate.SetMigrationsCollection(name)
}

func Up(n int) error {
	return globalMigrate.Up(n)
}

func Down(n int) error {
	return globalMigrate.Down(n)
}
