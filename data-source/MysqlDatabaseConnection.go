package data_source

import (
	"QuizApp/config"
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

var onceDb sync.Once

var instance *gorm.DB

func GetMysqlDbInstance() *gorm.DB {
	onceDb.Do(func() {

		databaseConfig := config.GetMysqlDatabaseConfig()
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
			databaseConfig.Mysql.DbUsername,
			databaseConfig.Mysql.DbPassword,
			databaseConfig.Mysql.DbHost,
			databaseConfig.Mysql.DbDatabase,
			databaseConfig.Mysql.DbCharset,
		)
		//db, err := gorm.Open(mysql.Open(dsn))

		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN: dsn,
		}), &gorm.Config{})

		if err != nil {
			log.Fatalf("Could not connect to Mysql database :%v", err)
		}
		instance = db
	})

	return instance
}
