package adapter

import (
	entity "github.com/iamviniciuss/golang-migrations/src/dto"
)

type MigrationRepository interface {
	Insert(rec *entity.VersionRecord) (*entity.VersionRecord, error)
	FindOne(rec *entity.VersionRecord) (*entity.VersionRecord, error)
	CreateCollectionIfNotExists(name string) error
}
