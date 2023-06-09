package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"rpssl_service/core/choiceModule/choiceRouter"
	"rpssl_service/core/playEngineModule/playEngineRouter"
	"rpssl_service/core/scoreboardModule/scoreboardRouter"
)

// register routes function to be used by negroni to mount routes
func RegisterRoutes() *echo.Echo {
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.CORS())

	rootRouter := router.Group("")

	// register choices router
	choiceRouter.RegisterRoute(rootRouter,"")

	// register play-engine router
	playEngineRouter.RegisterRoute(rootRouter,"")

	// register score-board router
	scoreboardRouter.RegisterRoute(rootRouter,"/score-board")

	return router
}
