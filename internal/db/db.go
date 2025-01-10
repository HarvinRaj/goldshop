package db

import (
	"database/sql"
	"github.com/HarvinRaj/goldshop/configs"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mSQL "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file" //importing migrate/file
	"log"
	"strings"
)

func NewDBConnection(config *configs.Config) error {

	dbConfig := &mysql.Config{
		User:                 config.DBUser,
		Passwd:               config.DBPassword,
		Net:                  config.Net,
		Addr:                 config.DBAddress,
		DBName:               config.DBName,
		AllowNativePasswords: config.AllowNativePasswords,
		ParseTime:            config.ParseTime,
	}

	db, err := sql.Open(config.DriverName, dbConfig.FormatDSN())
	if err != nil {
		log.Fatalf("Failed to Open DB connection, %v", err)
		return err
	}

	if err = runMigration(config, db); err != nil {
		log.Fatalf("Failed to run migration,%v", err)
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to DB, %v", err)
		return err
	}

	log.Printf("Connected to DB")
	return nil
}

func runMigration(config *configs.Config, db *sql.DB) error {

	driver, err := mSQL.WithInstance(db, &mSQL.Config{})
	if err != nil {
		log.Fatalf("Failed to make migrate instance, %v", err)
		return err
	}

	migrationPath := config.MigrateFiles

	if config.OSType == "windows" {
		migrationPath = "file://" + strings.ReplaceAll(migrationPath, "\\", "/")
	} else {
		migrationPath = "file://" + migrationPath
	}

	log.Printf("Migration file path : %s", migrationPath)

	m, err := migrate.NewWithDatabaseInstance(migrationPath, config.DBName, driver)
	if err != nil {
		log.Fatalf("Failed to initialize new migration instance, %s", err)
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations, %s", err)
	}

	log.Printf("Migration succesful")
	return nil
}
