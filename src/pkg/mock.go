package migrate

import (
	entity "github.com/iamviniciuss/golang-migrations/src/dto"
	"github.com/stretchr/testify/mock"
)

type MigrationRepositoryMock struct {
	mock.Mock
}

func (m *MigrationRepositoryMock) Insert(rec *entity.VersionRecord) (*entity.VersionRecord, error) {
	args := m.Called(rec)
	return args.Get(0).(*entity.VersionRecord), args.Error(1)
}

func (m *MigrationRepositoryMock) FindOne(rec *entity.VersionRecord) (*entity.VersionRecord, error) {
	args := m.Called(rec)
	return args.Get(0).(*entity.VersionRecord), args.Error(1)
}

func (m *MigrationRepositoryMock) CreateCollectionIfNotExists(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

type MigrationHandlerMock struct {
	mock.Mock
}

func (m *MigrationHandlerMock) GetVersion() uint64 {
	args := m.Called()
	return args.Get(0).(uint64)
}

func (m *MigrationHandlerMock) GetType() string {
	args := m.Called()
	return args.String(0)
}

func (m *MigrationHandlerMock) GetName() string {
	args := m.Called()
	return args.String(0)
}

func (m *MigrationHandlerMock) Up() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MigrationHandlerMock) Down() error {
	args := m.Called()
	return args.Error(0)
}
