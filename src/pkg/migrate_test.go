package migrate

import (
	"errors"
	"testing"
	"time"

	entity "github.com/iamviniciuss/golang-migrations/src/dto"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MigrationManagerTestSuite struct {
	suite.Suite
	value int
	repo  *MigrationRepositoryMock
}

// SetupSuite é executado uma vez antes de todos os testes da suíte.
func (suite *MigrationManagerTestSuite) SetupSuite() {
	// Configuração global
	suite.value = 42
}

// SetupTest é executado antes de cada teste.
func (suite *MigrationManagerTestSuite) SetupTest() {
	// Configuração específica para cada teste
	suite.repo = new(MigrationRepositoryMock)
}

// TearDownTest é executado depois de cada teste.
func (suite *MigrationManagerTestSuite) TearDownTest() {
	// Limpeza específica para cada teste
}

// TearDownSuite é executado uma vez após todos os testes da suíte.
func (suite *MigrationManagerTestSuite) TearDownSuite() {
	// Limpeza global
}

func (suite *MigrationManagerTestSuite) TestBasicTestUpgradeMigrationVersion() {
	migrateHandler := &MigrationHandlerMock{}
	migrateHandler.On("GetVersion").Return(uint64(2))
	migrateHandler.On("GetName").Return("BasicTest")
	migrateHandler.On("GetType").Return("migration")
	migrateHandler.On("Up").Return(nil)

	suite.repo.On("CreateCollectionIfNotExists", mock.Anything).Return(nil)
	suite.repo.On("Insert", mock.Anything).Return(&entity.VersionRecord{}, nil)
	suite.repo.
		On("FindOne", mock.Anything).
		Return(&entity.VersionRecord{Type: "migration", Version: 1, Description: "BasicTestDesc", Timestamp: time.Now()}, nil)

	migration_manager := NewMigrate(suite.repo)
	migration_manager.Register(migrateHandler)
	migration_manager.Up(1)

	migrateHandler.AssertNumberOfCalls(suite.T(), "Up", 1)
	suite.repo.AssertNumberOfCalls(suite.T(), "Insert", 1)
}

func (suite *MigrationManagerTestSuite) TestBasicTestNotUpgradeMigrationVersionWhenIsMinor() {
	newMigrationMinorVersion := uint64(1)
	currentMigrationMinorVersion := uint64(3)

	migrateHandler := &MigrationHandlerMock{}
	migrateHandler.On("GetVersion").Return(newMigrationMinorVersion)
	migrateHandler.On("GetName").Return("BasicTest")
	migrateHandler.On("GetType").Return("migration")
	migrateHandler.On("Up").Return(nil)

	suite.repo.On("CreateCollectionIfNotExists", mock.Anything).Return(nil)
	suite.repo.On("Insert", mock.Anything).Return(&entity.VersionRecord{}, nil)
	suite.repo.
		On("FindOne", mock.Anything).
		Return(&entity.VersionRecord{Type: "migration", Version: currentMigrationMinorVersion, Description: "BasicTestDesc", Timestamp: time.Now()}, nil)

	migration_manager := NewMigrate(suite.repo)
	migration_manager.Register(migrateHandler)
	migration_manager.Up(1)

	migrateHandler.AssertNumberOfCalls(suite.T(), "Up", 0)
	suite.repo.AssertNumberOfCalls(suite.T(), "Insert", 0)
}

func (suite *MigrationManagerTestSuite) TestBasicTestNotUpgradeMigrationVersionWhenHappensError() {
	newMigrationMinorVersion := uint64(3)
	currentMigrationMinorVersion := uint64(3)

	migrateHandler := &MigrationHandlerMock{}
	migrateHandler.On("GetVersion").Return(newMigrationMinorVersion)
	migrateHandler.On("GetName").Return("BasicTest")
	migrateHandler.On("GetType").Return("migration")
	migrateHandler.On("Up").Return(errors.New("force error on Up method."))

	suite.repo.On("CreateCollectionIfNotExists", mock.Anything).Return(nil)
	suite.repo.On("Insert", mock.Anything).Return(&entity.VersionRecord{}, nil)
	suite.repo.
		On("FindOne", mock.Anything).
		Return(&entity.VersionRecord{Type: "migration", Version: currentMigrationMinorVersion, Description: "BasicTestDesc", Timestamp: time.Now()}, nil)

	migration_manager := NewMigrate(suite.repo)
	migration_manager.Register(migrateHandler)
	migration_manager.Up(1)

	migrateHandler.AssertNumberOfCalls(suite.T(), "Up", 0)
	suite.repo.AssertNumberOfCalls(suite.T(), "Insert", 0)
}

func (suite *MigrationManagerTestSuite) TestBasicTestNotUpgradeMigrationVersionWhenHappensErrorOnVersion() {
	newMigrationMinorVersion := uint64(3)
	currentMigrationMinorVersion := uint64(3)

	migrateHandler := &MigrationHandlerMock{}
	migrateHandler.On("GetVersion").Return(newMigrationMinorVersion)
	migrateHandler.On("GetName").Return("BasicTest")
	migrateHandler.On("GetType").Return("migration")
	migrateHandler.On("Up").Return(nil)

	suite.repo.On("CreateCollectionIfNotExists", mock.Anything).Return(errors.New("force error on CreateCollectionIfNotExists"))
	suite.repo.On("Insert", mock.Anything).Return(&entity.VersionRecord{}, nil)
	suite.repo.
		On("FindOne", mock.Anything).
		Return(&entity.VersionRecord{Type: "migration", Version: currentMigrationMinorVersion, Description: "BasicTestDesc", Timestamp: time.Now()}, nil)

	migration_manager := NewMigrate(suite.repo)
	migration_manager.Register(migrateHandler)
	migration_manager.Up(1)

	migrateHandler.AssertNumberOfCalls(suite.T(), "Up", 0)
	suite.repo.AssertNumberOfCalls(suite.T(), "Insert", 0)
}

func (suite *MigrationManagerTestSuite) TestBasicTestNotUpgradeMigrationVersionWhenHappensErrorOnUp() {
	newMigrationMinorVersion := uint64(3)
	currentMigrationMinorVersion := uint64(2)

	migrateHandler := &MigrationHandlerMock{}
	migrateHandler.On("GetVersion").Return(newMigrationMinorVersion)
	migrateHandler.On("GetName").Return("BasicTest")
	migrateHandler.On("GetType").Return("migration")
	migrateHandler.On("Up").Return(errors.New("force error on Up method."))

	suite.repo.On("CreateCollectionIfNotExists", mock.Anything).Return(nil)
	suite.repo.On("Insert", mock.Anything).Return(&entity.VersionRecord{}, nil)
	suite.repo.
		On("FindOne", mock.Anything).
		Return(&entity.VersionRecord{Type: "migration", Version: currentMigrationMinorVersion, Description: "BasicTestDesc", Timestamp: time.Now()}, nil)

	migration_manager := NewMigrate(suite.repo)
	migration_manager.Register(migrateHandler)
	migration_manager.Up(1)

	migrateHandler.AssertNumberOfCalls(suite.T(), "Up", 1)
	suite.repo.AssertNumberOfCalls(suite.T(), "Insert", 0)
}

func (suite *MigrationManagerTestSuite) TestBasicTestNotUpgradeMigrationVersionWhenHappensErrorOnUp3() {
	newMigrationMinorVersion := uint64(3)
	currentMigrationMinorVersion := uint64(2)

	migrateHandler := &MigrationHandlerMock{}
	migrateHandler.On("GetVersion").Return(newMigrationMinorVersion)
	migrateHandler.On("GetName").Return("BasicTest")
	migrateHandler.On("GetType").Return("migration")
	migrateHandler.On("Up").Return(nil)

	suite.repo.On("CreateCollectionIfNotExists", mock.Anything).Return(nil)
	insertError := errors.New("force error on Insert")
	suite.repo.On("Insert", mock.Anything).Return(&entity.VersionRecord{}, insertError)
	suite.repo.
		On("FindOne", mock.Anything).
		Return(&entity.VersionRecord{Type: "migration", Version: currentMigrationMinorVersion, Description: "BasicTestDesc", Timestamp: time.Now()}, nil)

	migration_manager := NewMigrate(suite.repo)
	migration_manager.Register(migrateHandler)
	upError := migration_manager.Up(1)

	migrateHandler.AssertNumberOfCalls(suite.T(), "Up", 1)
	suite.repo.AssertNumberOfCalls(suite.T(), "Insert", 1)
	suite.Equal(insertError.Error(), upError.Error())
}

func (suite *MigrationManagerTestSuite) TestBasicTestNotExecuteWhenThereIsNotMigrations() {
	migrateHandler := &MigrationHandlerMock{}

	migration_manager := NewMigrate(suite.repo)
	migration_manager.Up(0)

	migrateHandler.AssertNumberOfCalls(suite.T(), "Up", 0)
	suite.repo.AssertNumberOfCalls(suite.T(), "Insert", 0)
}

func TestMigrationManagerTestSuite(t *testing.T) {
	suite.Run(t, new(MigrationManagerTestSuite))
}
