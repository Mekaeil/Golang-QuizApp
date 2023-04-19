package main

import (
	"QuizApp/config"
	"QuizApp/config/routes"
	datasource "QuizApp/data-source"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
	time "time"
)

func init() {

	loadingEnvironmentVariables()
	setOutputLog()

}

func main() {
	app := echo.New()
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	datasource.GetInstance()

	/*
	   |--------------------------------------------------------------------------
	   | CORS middleware for API endpoint.
	   |--------------------------------------------------------------------------
	   |
	*/

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST},
	}))

	// ROUTE VERSION 1
	routes.DefineApiRouteV1(app)

	app.Logger.Fatal(app.Start(":1323"))
}

func setOutputLog() {

	var logFileName string

	logChannel := os.Getenv("LOG_CHANNEL")
	switch logChannel {
	case string(config.DAILY):
		timeNow := time.Now()
		dateTimeString := fmt.Sprintf("%d-%d-%d", timeNow.Year(), timeNow.Month(), timeNow.Day())
		logFileName = fmt.Sprintf("log-%s.log", dateTimeString)
	default:
		logFileName = "log.log"
	}

	filePath := fmt.Sprintf("./storage/log/%s", logFileName)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.Printf("We are logging in Golang Based on LOG CHANNEL: %s!", logChannel)
}

func loadingEnvironmentVariables() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Loading Environment file errors ocuured! Err: %s", err)
	}
}
