package db

import (
	"database/sql"
	"github.com/HarvinRaj/goldshop/configs"
	"github.com/HarvinRaj/goldshop/logger"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mSQL "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file" //importing migrate/file
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
		logger.ErrorLog.Printf("Failed to Open DB connection, %v", err)
		return err
	}

	if err = runMigration(config, db); err != nil {
		logger.ErrorLog.Printf("Failed to run migration,%v", err)
		return err
	}

	logger.InfoLog.Println("Connected to MySQL Database")

	err = db.Ping()
	if err != nil {
		logger.ErrorLog.Printf("Failed to connect to DB, %v", err)
		return err
	}

	return nil
}

func runMigration(config *configs.Config, db *sql.DB) error {

	driver, err := mSQL.WithInstance(db, &mSQL.Config{})
	if err != nil {
		logger.ErrorLog.Printf("Failed to make new migration instance, %v", err)
		return err
	}

	migrationPath := config.MigrateFiles

	if config.OSType == "windows" {
		migrationPath = "file://" + strings.ReplaceAll(migrationPath, "\\", "/")
	} else {
		migrationPath = "file://" + migrationPath
	}

	logger.InfoLog.Println("Migration file path - ", migrationPath)

	m, err2 := migrate.NewWithDatabaseInstance(migrationPath, config.DBName, driver)
	if err2 != nil {
		logger.ErrorLog.Printf("Failed to initialize new migration instance, %s", err)
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		logger.ErrorLog.Printf("Failed to make migrations changes, %s", err)
	}

	logger.InfoLog.Println("Migration successful")
	return nil
}
