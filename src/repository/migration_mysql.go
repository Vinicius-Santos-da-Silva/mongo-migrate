package repository

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	entity "github.com/iamviniciuss/golang-migrations/src/dto"
	"github.com/pkg/errors"
)

type MigrationRepositoryMySQL struct {
	db *sql.DB
}

func NewMigrationRepositoryMySQL(db *sql.DB) *MigrationRepositoryMySQL {
	return &MigrationRepositoryMySQL{db}
}

func (erm *MigrationRepositoryMySQL) Insert(rec *entity.VersionRecord) (*entity.VersionRecord, error) {
	_, err := erm.db.Exec("INSERT INTO migrations (version, description, type) VALUES (?, ?, ?)", rec.Version, rec.Description, rec.Type)
	if err != nil {
		return rec, err
	}
	return rec, nil
}

func (erm *MigrationRepositoryMySQL) FindOne(record *entity.VersionRecord) (*entity.VersionRecord, error) {
	var output entity.VersionRecord
	err := erm.db.QueryRow("SELECT version, description, type FROM "+tableName+" WHERE version = ? AND description = ? AND type = ?", record.Version, record.Description, record.Type).Scan(&output.Version, &output.Description, &output.Type)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return &entity.VersionRecord{}, nil
	case err != nil:
		return nil, err
	}
	return &output, nil
}

func (erm *MigrationRepositoryMySQL) CreateCollectionIfNotExists(name string) error {

	createTableQuery := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			_id int NOT NULL AUTO_INCREMENT,
			type varchar(255) DEFAULT NULL,
			description varchar(255) DEFAULT NULL,
			version int DEFAULT NULL,
			timestamp datetime DEFAULT NULL,
			PRIMARY KEY (_id)
		);
	`, tableName)

	_, err := erm.db.Exec(createTableQuery)
	if err != nil {
		return err
	}
	return nil
}

var (
	username  = os.Getenv("MIGRATION_USERNAME")
	password  = os.Getenv("MIGRATION_PASSWORD")
	hostname  = os.Getenv("MIGRATION_HOSTNAME")
	port      = os.Getenv("MIGRATION_PORT")
	dbName    = os.Getenv("MIGRATION_DB")
	tableName = os.Getenv("MIGRATION_TABLE")
)

func NewConnection() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, hostname, port, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
