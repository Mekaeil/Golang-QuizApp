package main

import (
	"QuizApp/bootstrap/application"
	"QuizApp/config/routes"
	datasource "QuizApp/data-source"
	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	"net/http"
)

func init() {

	application.LoadingEnvironmentVariables()

	application.SetOutputLog()
}

func main() {
	app := echo.New()
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// CREATING DB CONNECTION
	datasource.GetInstance()

	/*
	   |--------------------------------------------------------------------------
	   | CORS middleware for API endpoint.
	   |--------------------------------------------------------------------------
	   |
	*/

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// INJECTING ROUTE VERSION 1
	routes.DefineApiRouteV1(app)

	app.Logger.Fatal(app.Start(":1323"))
}
