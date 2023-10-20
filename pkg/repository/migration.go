package repository

import "github.com/Vinicius-Santos-da-Silva/mongo-migrate/pkg/entity"

type MigrationRepository interface {
	Insert(rec *entity.VersionRecord) (*entity.VersionRecord, error)
	FindAll() ([]*entity.VersionRecord, error)
	FindOne(rec *entity.VersionRecord) (*entity.VersionRecord, error)
	CreateCollectionIfNotExists(name string) error
}
