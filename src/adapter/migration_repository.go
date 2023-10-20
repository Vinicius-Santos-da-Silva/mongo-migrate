package adapter

import (
	entity "github.com/Vinicius-Santos-da-Silva/mongo-migrate/src/dto"
)

type MigrationRepository interface {
	Insert(rec *entity.VersionRecord) (*entity.VersionRecord, error)
	FindOne(rec *entity.VersionRecord) (*entity.VersionRecord, error)
	CreateCollectionIfNotExists(name string) error
}
