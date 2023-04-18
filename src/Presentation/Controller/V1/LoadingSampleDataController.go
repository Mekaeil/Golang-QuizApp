package controllerV1

import (
	"QuizApp/config"
	datasource "QuizApp/data-source"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

type LoadingSampleDataController struct {
}

func (controller LoadingSampleDataController) Routes() []config.Route {
	return []config.Route{
		{
			Method:  echo.GET,
			Path:    "/sample",
			Handler: controller.GetSampleData,
		},
	}
}

func (controller LoadingSampleDataController) GetSampleData(c echo.Context) error {

	var sampleData map[string]string

	fileDataName := "sample.json"
	encodedIssuersData := datasource.GetFileData(fileDataName)
	json.Unmarshal([]byte(encodedIssuersData), &sampleData)

	return c.JSON(http.StatusOK, sampleData)
}
