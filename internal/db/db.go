package db

import (
	"database/sql"
	"fmt"
	"github.com/HarvinRaj/goldshop/configs"
	"github.com/go-sql-driver/mysql"
	"log"
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

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to DB, %v", err)
		return err
	}

	fmt.Println("Connected to DB")
	return nil
}
