package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	dsn := "docker:password@tcp(docker-db)/quiz_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(db.Config)

	e.Logger.Fatal(e.Start(":1323"))
}
