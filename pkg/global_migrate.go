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

// RegisteredMigrations returns all registered migrations.
func RegisteredMigrations() []Migration {
	ret := make([]Migration, len(globalMigrate.migrations))
	copy(ret, globalMigrate.migrations)
	return ret
}

// SetDatabase sets database for global migrate.
func SetDatabase(db *mongo.Database) {
	globalMigrate.db = db
}

// SetMigrationsCollection changes default collection name for migrations history.
func SetMigrationsCollection(name string) {
	globalMigrate.SetMigrationsCollection(name)
}

// Version returns current database version.
func Version() (uint64, string, error) {
	return globalMigrate.Version()
}

// Up performs "up" migration using registered migrations.
// Detailed description available in Migrate.Up().
func Up(n int) error {
	return globalMigrate.Up(n)
}

// Down performs "down" migration using registered migrations.
// Detailed description available in Migrate.Down().
func Down(n int) error {
	return globalMigrate.Down(n)
}
