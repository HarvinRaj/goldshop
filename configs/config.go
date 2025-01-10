package configs

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	OSType               string `json:"OSTYPE"`
	Port                 string `json:"PORT"`
	DBName               string `json:"DBNAME"`
	DBUser               string `json:"DBUSER"`
	DBPassword           string `json:"DBPASSWORD"`
	DBHost               string `json:"DBHOST"`
	DBPort               string `json:"DBPORT"`
	DBAddress            string `json:"DBADDRESS"`
	DriverName           string `json:"DRIVERNAME"`
	Net                  string `json:"NET"`
	AllowNativePasswords bool   `json:"ALLOWNATIVEPASSWORDS"`
	ParseTime            bool   `json:"PARSETIME"`
	MigrateFiles         string `json:"MIGRATEFILES"`
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read config file, %v", err)
		return nil, err
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &config, nil
}
