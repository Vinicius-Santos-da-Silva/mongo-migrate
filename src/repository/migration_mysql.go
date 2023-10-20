package repository

import (
	"database/sql"
	"fmt"

	entity "github.com/Vinicius-Santos-da-Silva/mongo-migrate/src/dto"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
)

const migrationsPath = "file://path/to/your/migrations"

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
	err := erm.db.QueryRow("SELECT version, description, type FROM migrations WHERE version = ? AND description = ? AND type = ?", record.Version, record.Description, record.Type).Scan(&output.Version, &output.Description, &output.Type)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, nil
	case err != nil:
		return nil, err
	}
	return &output, nil
}

func (erm *MigrationRepositoryMySQL) CreateCollectionIfNotExists(name string) error {
	createTableQuery := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id INT AUTO_INCREMENT PRIMARY KEY,
			nome VARCHAR(255),
			idade INT
		);
	`, tableName)

	_, err := erm.db.Exec(createTableQuery)
	if err != nil {
		return err
	}

	fmt.Println("Tabela criada ou já existente.")

	return nil
}

const (
	username  = "root"
	password  = ""
	hostname  = "localhost"
	port      = 3306
	dbName    = "wallet"
	tableName = "migrations"
)

func NewConnection() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, hostname, port, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
