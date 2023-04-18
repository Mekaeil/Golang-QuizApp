package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type MysqlDbConnection struct {
	DbHost     string
	DbPort     string
	DbDatabase string
	DbUsername string
	DbPassword string
}

type DatabaseConfig struct {
	Mysql MysqlDbConnection
}

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Loading Environment file errors ocuured! Err: %s", err)
	}
}

func NewDatabase() DatabaseConfig {
	return DatabaseConfig{
		Mysql: MysqlDbConnection{
			DbHost:     os.Getenv("MYSQL_DB_HOST"),
			DbPort:     os.Getenv("MYSQL_DB_PORT"),
			DbDatabase: os.Getenv("MYSQL_DB_DATABASE"),
			DbUsername: os.Getenv("MYSQL_DB_USERNAME"),
			DbPassword: os.Getenv("MYSQL_DB_PASSWORD"),
		},
	}
}
