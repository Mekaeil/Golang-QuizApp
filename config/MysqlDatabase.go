package config

import (
	"os"
)

type MysqlDbConnection struct {
	DbHost     string
	DbPort     string
	DbDatabase string
	DbUsername string
	DbPassword string
	DbCharset  string
}

type DatabaseConfig struct {
	Mysql MysqlDbConnection
}

func GetMysqlDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Mysql: MysqlDbConnection{
			DbHost:     os.Getenv("MYSQL_DB_HOST"),
			DbPort:     os.Getenv("MYSQL_DB_PORT"),
			DbDatabase: os.Getenv("MYSQL_DB_DATABASE"),
			DbUsername: os.Getenv("MYSQL_DB_USERNAME"),
			DbPassword: os.Getenv("MYSQL_DB_PASSWORD"),
			DbCharset:  os.Getenv("MYSQL_DB_CHARSET"),
		},
	}
}
