package migrate

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type collectionSpecification struct {
	Name string `bson:"name"`
	Type string `bson:"type"`
}

type versionRecord struct {
	Type        string    `bson:"type"`
	Version     uint64    `bson:"version"`
	Description string    `bson:"description,omitempty"`
	Timestamp   time.Time `bson:"timestamp"`
}

const defaultMigrationsCollection = "migrations"

const AllAvailable = -1

type Migrate struct {
	db                   *mongo.Database
	migrations           []Migration
	migrationsCollection string
}

func NewMigrate(db *mongo.Database, migrations ...Migration) *Migrate {
	internalMigrations := make([]Migration, len(migrations))
	copy(internalMigrations, migrations)
	return &Migrate{
		db:                   db,
		migrations:           internalMigrations,
		migrationsCollection: defaultMigrationsCollection,
	}
}

func (m *Migrate) SetMigrationsCollection(name string) {
	m.migrationsCollection = name
}

func (m *Migrate) isCollectionExist(name string) (isExist bool, err error) {
	collections, err := m.getCollections()
	if err != nil {
		return false, err
	}

	for _, c := range collections {
		if name == c.Name {
			return true, nil
		}
	}
	return false, nil
}

func (m *Migrate) createCollectionIfNotExist(name string) error {
	exist, err := m.isCollectionExist(name)
	if err != nil {
		return err
	}
	if exist {
		return nil
	}

	command := bson.D{bson.E{Key: "create", Value: name}}
	err = m.db.RunCommand(nil, command).Err()
	if err != nil {
		return err
	}

	return nil
}

func (m *Migrate) getCollections() (collections []collectionSpecification, err error) {
	filter := bson.D{bson.E{Key: "type", Value: "collection"}}
	options := options.ListCollections().SetNameOnly(true)

	cursor, err := m.db.ListCollections(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}

	if cursor != nil {
		defer func(cursor *mongo.Cursor) {
			curErr := cursor.Close(context.TODO())
			if curErr != nil {
				if err != nil {
					err = errors.Wrapf(curErr, "migrate: get collection failed: %s", err.Error())
				} else {
					err = curErr
				}
			}
		}(cursor)
	}

	for cursor.Next(context.TODO()) {
		var collection collectionSpecification

		err := cursor.Decode(&collection)
		if err != nil {
			return nil, err
		}

		collections = append(collections, collection)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return
}

func (m *Migrate) Version() (uint64, string, error) {
	if err := m.createCollectionIfNotExist(m.migrationsCollection); err != nil {
		return 0, "", err
	}

	filter := bson.D{{}}
	sort := bson.D{bson.E{Key: "_id", Value: -1}}
	options := options.FindOne().SetSort(sort)

	// find record with greatest id (assuming it`s latest also)
	result := m.db.Collection(m.migrationsCollection).FindOne(context.TODO(), filter, options)
	err := result.Err()
	switch {
	case err == mongo.ErrNoDocuments:
		return 0, "", nil
	case err != nil:
		return 0, "", err
	}

	var rec versionRecord
	if err := result.Decode(&rec); err != nil {
		return 0, "", err
	}

	return rec.Version, rec.Description, nil
}

func (m *Migrate) SetVersion(version uint64, description string, typing string) error {
	rec := versionRecord{
		Version:     version,
		Timestamp:   time.Now().UTC(),
		Description: description,
		Type:        typing,
	}

	_, err := m.db.Collection(m.migrationsCollection).InsertOne(context.TODO(), rec)
	if err != nil {
		return err
	}

	return nil
}

func (m *Migrate) Up(n int) error {
	currentVersion, _, err := m.Version()
	if err != nil {
		return err
	}
	if n <= 0 || n > len(m.migrations) {
		n = len(m.migrations)
	}
	migrationSort(m.migrations)

	for i, p := 0, 0; i < len(m.migrations) && p < n; i++ {
		migration := m.migrations[i]
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
	currentVersion, _, err := m.Version()
	if err != nil {
		return err
	}
	if n <= 0 || n > len(m.migrations) {
		n = len(m.migrations)
	}
	migrationSort(m.migrations)

	for i, p := len(m.migrations)-1, 0; i >= 0 && p < n; i-- {
		migration := m.migrations[i]
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
