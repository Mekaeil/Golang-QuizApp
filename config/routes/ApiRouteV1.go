package routes

import (
	"QuizApp/config"
	controllerV1 "QuizApp/src/Presentation/Controller/V1"
	"github.com/labstack/echo/v4"
)

func DefineApiRouteV1(e *echo.Echo) {

	controllers := []config.Controller{
		controllerV1.LoadingSampleDataController{},
		controllerV1.QuestionController{},
	}
	var routes []config.Route
	for _, controller := range controllers {
		routes = append(routes, controller.Routes()...)
	}
	api := e.Group("/api/v1")
	for _, route := range routes {
		switch route.Method {
		case echo.POST:
			{
				api.POST(route.Path, route.Handler, route.Middleware...)
				break
			}
		case echo.GET:
			{
				api.GET(route.Path, route.Handler, route.Middleware...)
				break
			}
		case echo.DELETE:
			{
				api.DELETE(route.Path, route.Handler, route.Middleware...)
				break
			}
		case echo.PUT:
			{
				api.PUT(route.Path, route.Handler, route.Middleware...)
				break
			}
		case echo.PATCH:
			{
				api.PATCH(route.Path, route.Handler, route.Middleware...)
				break
			}
		}
	}
}
