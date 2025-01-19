package db

import (
	"context"
	"database/sql"
	"github.com/HarvinRaj/goldshop/configs"
	"github.com/HarvinRaj/goldshop/logger"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mSQL "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file" //importing migrate/file
	"strings"
)

// NewDBConnection ...
func NewDBConnection(ctx context.Context, config *configs.Config) error {
	// you can add other connections here e.g. SQL, Postgres
	err := newMySQL(ctx, config)
	if err != nil {
		logger.ErrorLog.Println(err)
		return err
	}
	return nil
}

// Connection for MySQL
func newMySQL(ctx context.Context, config *configs.Config) error {

	dbConfig := configMySQL(config)
	db, err := newDriver(config.DriverName, dbConfig)
	if err != nil {
		logger.ErrorLog.Println(err)
		return err
	}

	err = db.PingContext(ctx)
	if err != nil {
		logger.ErrorLog.Printf("Failed to connect to DB, %v", err)
		return err
	}

	logger.InfoLog.Connected.Println("Connected to MYSQL Database")

	if err = runMigrationUP(config, db); err != nil {
		logger.ErrorLog.Println(err)
		return err
	}

	logger.InfoLog.Connected.Println("MYSQL Database set-up complete")
	return nil
}

// For MySQL Connection, connect to database, new driver
func newDriver(driverName string, dbConfig *mysql.Config) (db *sql.DB, err error) {

	db, err = sql.Open(driverName, dbConfig.FormatDSN())
	if err != nil {
		logger.ErrorLog.Printf("Failed to Open DB connection, %v", err)
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
		logger.ErrorLog.Printf("Failed to make new migration instance, %v", err)
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
		logger.ErrorLog.Printf("Failed to initialize new migration instance, %v", err)
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		logger.ErrorLog.Printf("Failed to make migrations changes, %v", err)
		return err
	}

	logger.InfoLog.Success.Println("Migration successful")
	return nil
}
