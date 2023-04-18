package data_source

//import (
//	"QuizApp/config"
//	"gorm.io/driver/mysql"
//	_ "gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"log"
//	"sync"
//)
//
//var onceDb sync.Once
//
//var instance *gorm.DB
//
//func GetInstance() *gorm.DB {
//
//	onceDb.Do(func() {
//		//databaseConfig := config.NewDatabase()
//		//
//		//// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
//		//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
//		//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//
//		//db, err := gorm.Open(mysql.New(mysql.Config{
//		//	DSN: "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local", // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
//		//}), &gorm.Config{})
//
//		//db, err := gorm.Open("mysql", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
//		//	databaseConfig.Mysql.DbHost,
//		//	databaseConfig.Mysql.DbPort,
//		//	databaseConfig.Mysql.DbUsername,
//		//	databaseConfig.Mysql.DbDatabase,
//		//	databaseConfig.Mysql.DbPassword,
//		//))
//		if err != nil {
//			log.Fatalf("Could not connect to database :%v", err)
//		}
//		instance = db
//	})
//	return instance
//}
