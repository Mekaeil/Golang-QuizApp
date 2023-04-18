package main

import (
	"QuizApp/config/routes"
	datasource "QuizApp/data-source"
	"github.com/joho/godotenv"
	middleware "github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Loading Environment file errors ocuured! Err: %s", err)
	}
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
