package main

import (
	data_source "QuizApp/data-source"
	"fmt"
	"github.com/joho/godotenv"
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
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	data_source.GetInstance()

	fileDataName := "sample.json"
	sampleData := data_source.GetFileData(fileDataName)
	fmt.Println(sampleData)

	e.Logger.Fatal(e.Start(":1323"))
}
