package db

import (
	"context"
	"database/sql"
	"github.com/HarvinRaj/goldshop/configs"
	"github.com/HarvinRaj/goldshop/logger"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mSQL "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file" // imported for side effect
	"strings"
)

// NewMySQL ...
func NewMySQL(ctx context.Context, config *configs.Config) (*DB, error) {

	dbConfig := configMySQL(config)
	database, err := newDriver(config.DriverName, dbConfig)
	if err != nil {
		logger.ErrorLog.Error.Println(err)
		return nil, err
	}

	err = database.PingContext(ctx)
	if err != nil {
		logger.ErrorLog.Error.Printf("Failed to connect to DB, %v", err)
		return nil, err
	}

	logger.InfoLog.Connected.Println("Connected to MYSQL Database")

	if err = runMigrationUP(config, database); err != nil {
		logger.ErrorLog.Error.Println(err)
		return nil, err
	}

	logger.InfoLog.Connected.Println("MYSQL Database set-up complete")
	return &DB{db: database}, nil
}

// For MySQL Connection, connect to database, new driver
func newDriver(driverName string, dbConfig *mysql.Config) (db *sql.DB, err error) {

	db, err = sql.Open(driverName, dbConfig.FormatDSN())
	if err != nil {
		logger.ErrorLog.Error.Printf("Failed to Open DB connection, %v", err)
		return nil, err
	}

	return db, nil
}

// For MySQL credential from config
func configMySQL(config *configs.Config) *mysql.Config {

	dbConfig := &mysql.Config{
		User:                 config.DBUser,
		Passwd:               config.DBPassword,
		Net:                  config.Net,
		Addr:                 config.DBAddress,
		DBName:               config.DBName,
		AllowNativePasswords: config.AllowNativePasswords,
		ParseTime:            config.ParseTime,
	}

	return dbConfig
}

// To run up migration after database connection
func runMigrationUP(config *configs.Config, db *sql.DB) error {

	driver, err := mSQL.WithInstance(db, &mSQL.Config{})
	if err != nil {
		logger.ErrorLog.Error.Printf("Failed to make new migration instance, %v", err)
		return err
	}

	migrationPath := config.MigrateFiles

	if config.OSType == "windows" {
		migrationPath = "file://" + strings.ReplaceAll(migrationPath, "\\", "/")
	} else {
		migrationPath = "file://" + migrationPath
	}

	logger.InfoLog.Success.Println("Migration file path - ", migrationPath)

	m, err2 := migrate.NewWithDatabaseInstance(migrationPath, config.DBName, driver)
	if err2 != nil {
		logger.ErrorLog.Error.Printf("Failed to initialize new migration instance, %v", err2)
		return err2
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			logger.InfoLog.Info.Printf("No migration changes, %v", err)
			return nil
		}
		logger.ErrorLog.Error.Printf("Failed to make migrations changes, %v", err)
		return err
	}

	logger.InfoLog.Success.Println("Migration successful")
	return nil
}
