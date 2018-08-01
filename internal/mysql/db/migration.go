package db

import (
	"database/sql"
	"path"
	"runtime"
	"strings"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
)

type Migration struct {
	Migrate *migrate.Migrate
}

func (this *Migration) Up() ([]error, bool) {
	err := this.Migrate.Up()
	if err != nil {
		return []error{err}, false
	}

	return []error{}, true
}

func (this *Migration) Down() ([]error, bool) {
	err := this.Migrate.Down()
	if err != nil {
		return []error{err}, false
	}

	return []error{}, true
}

func RunMigration(dbURI string) (*Migration, error) {
	_, filename, _, _ := runtime.Caller(0)

	migrationPath := path.Join(path.Dir(filename), "migrations")

	dataPath := []string{}
	dataPath = append(dataPath, "file://")
	dataPath = append(dataPath, migrationPath)

	pathToMigrate := strings.Join(dataPath, "")

	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, err
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{DatabaseName: "inventory_db"})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		pathToMigrate,
		"mysql",
		driver,
	)

	return &Migration{
		Migrate: m,
	}, nil
}
