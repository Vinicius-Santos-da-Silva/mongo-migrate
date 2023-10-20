package migrate

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
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

func SetDatabase(db *mongo.Database) {
	globalMigrate.db = db
}

func SetRepository(repo MigrationRepository) {
	globalMigrate.migrationRepository = repo
}

func SetMigrationsCollection(name string) {
	globalMigrate.SetMigrationsCollection(name)
}

func Version() (uint64, string, error) {
	return globalMigrate.Version()
}

func Up(n int) error {
	return globalMigrate.Up(n)
}

func Down(n int) error {
	return globalMigrate.Down(n)
}
