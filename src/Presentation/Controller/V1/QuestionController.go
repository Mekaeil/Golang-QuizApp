package controllerV1

import (
	"QuizApp/config"
	datasource "QuizApp/data-source"
	"QuizApp/src/Infrastructure/Persistence/Model"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type (
	QuestionController struct{}

	CreateQuestionRequest struct {
		Title        string `json:"title" validate:"required"`
		QuestionType string `json:"question_type" validate:"required"`
		Status       bool   `json:"status" validate:"required"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (controller QuestionController) Routes() []config.Route {
	return []config.Route{
		{
			Method:  echo.POST,
			Path:    "/question/create",
			Handler: controller.CreateQuestion,
		},
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (controller QuestionController) CreateQuestion(ctx echo.Context) error {

	params := new(CreateQuestionRequest)

	if err := ctx.Bind(params); err != nil {
		log.Fatal(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	//if err := ctx.Validate(params); err != nil {
	//	log.Fatal(err)
	//	return ctx.JSON(http.StatusBadRequest, err)
	//}

	dbInstance := datasource.GetInstance()

	var question Model.Question

	question.Title = params.Title
	question.QuestionType = params.QuestionType
	question.Status = params.Status

	dbInstance.Create(&question)

	return ctx.JSON(http.StatusOK, question)
}
