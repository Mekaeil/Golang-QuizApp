package data_source

import (
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func GetInstance() *gorm.DB {

	var instance *gorm.DB

	switch driver := os.Getenv("DATABASE_DRIVER"); driver {
	case string(MYSQL):
		instance = GetMysqlDbInstance()

	default:
		log.Fatalf("could not find the dtabase driver :%v", driver)
	}

	return instance
}
