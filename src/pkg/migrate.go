package migrate

import (
	"fmt"
	"time"

	repository "github.com/iamviniciuss/golang-migrations/src/adapter"
	entity "github.com/iamviniciuss/golang-migrations/src/dto"
)

const defaultMigrationsCollection = "migrations"

const AllAvailable = -1

type Migrate struct {
	migrationRepository  repository.MigrationRepository
	migrations           []Migration
	migrationsCollection string
}

func NewMigrate(migrationRepository repository.MigrationRepository, migrations ...Migration) *Migrate {
	internalMigrations := make([]Migration, len(migrations))
	copy(internalMigrations, migrations)
	return &Migrate{
		migrations:           internalMigrations,
		migrationsCollection: defaultMigrationsCollection,
		migrationRepository:  migrationRepository,
	}
}

func (m *Migrate) SetMigrationsCollection(name string) {
	m.migrationsCollection = name
}

func (m *Migrate) Version(rec *entity.VersionRecord) (uint64, string, error) {
	if err := m.migrationRepository.CreateCollectionIfNotExists(m.migrationsCollection); err != nil {
		return 0, "", err
	}

	rec, err := m.migrationRepository.FindOne(rec)

	if err != nil {
		return 0, "", nil
	}

	return rec.Version, rec.Description, nil
}

func (m *Migrate) SetVersion(version uint64, description string, typing string) error {
	rec := &entity.VersionRecord{
		Version:     version,
		Timestamp:   time.Now().UTC(),
		Description: description,
		Type:        typing,
	}

	_, err := m.migrationRepository.Insert(rec)

	return err
}

func (m *Migrate) Up(n int) error {

	if n <= 0 || n > len(m.migrations) {
		n = len(m.migrations)
	}
	migrationSort(m.migrations)

	for i, p := 0, 0; i < len(m.migrations) && p < n; i++ {
		migration := m.migrations[i]
		currentVersion, _, err := m.Version(&entity.VersionRecord{
			Type:        migration.Handler.GetType(),
			Version:     migration.Version,
			Description: migration.Handler.GetName(),
		})

		if err != nil {
			return err
		}

		if migration.Version <= currentVersion || migration.Handler == nil {
			continue
		}
		p++
		if err := migration.Handler.Up(); err != nil {
			return err
		}
		if err := m.SetVersion(migration.Version, migration.Description, migration.Handler.GetType()); err != nil {
			return err
		}
	}
	return nil
}

func (m *Migrate) Down(n int) error {

	if n <= 0 || n > len(m.migrations) {
		n = len(m.migrations)
	}
	migrationSort(m.migrations)

	for i, p := len(m.migrations)-1, 0; i >= 0 && p < n; i-- {
		migration := m.migrations[i]

		currentVersion, _, err := m.Version(&entity.VersionRecord{
			Type:        migration.Handler.GetType(),
			Version:     migration.Version,
			Description: migration.Handler.GetName(),
		})

		if err != nil {
			return err
		}

		if migration.Version > currentVersion || migration.Handler == nil {
			continue
		}
		p++
		if err := migration.Handler.Down(); err != nil {
			return err
		}

		var prevMigration Migration
		if i == 0 {
			prevMigration = Migration{Version: 0}
		} else {
			prevMigration = m.migrations[i-1]
		}
		if err := m.SetVersion(prevMigration.Version, prevMigration.Description, migration.Handler.GetType()); err != nil {
			return err
		}
	}
	return nil
}

func (mig *Migrate) internalRegister(migrationHandler repository.MigrationHandler, skip int) error {
	if hasVersion(mig.migrations, migrationHandler) {
		return fmt.Errorf("migration with version %v already registered", migrationHandler.GetVersion())
	}

	mig.migrations = append(mig.migrations, Migration{
		Version:     migrationHandler.GetVersion(),
		Description: migrationHandler.GetName(),
		Handler:     migrationHandler,
	})
	return nil
}

func (mig *Migrate) Register(migrationHandler repository.MigrationHandler) error {
	return mig.internalRegister(migrationHandler, 2)
}
