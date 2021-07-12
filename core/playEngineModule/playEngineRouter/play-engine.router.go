package playEngineRouter

import (
	"fmt"
	"github.com/labstack/echo"
	"rpssl_service/core/choiceModule/choiceServices"
	"rpssl_service/core/playEngineModule/playEngineControllers"
	"rpssl_service/core/playEngineModule/playEngineServices"
	"rpssl_service/core/scoreboardModule/scoreboardServices"
	"rpssl_service/infrastructures/helpers"
	"rpssl_service/infrastructures/services"
)

var playEngineController = playEngineControllers.PlayEngineController{
	PlayEngineService: playEngineServices.PlayEngineService{},
	ChoiceService: choiceServices.ChoiceService{},
	RandomNumberService: services.RandomNumberService{
		RestRequestHelperService: helpers.RestRequestHelper{},
	},
	ScoreboardService: scoreboardServices.ScoreboardService{
		RedisService: services.RedisService{},
	},
}


func RegisterRoute(router *echo.Group, root string) {
	router.POST(fmt.Sprintf( "%s/play",root), playEngineController.PlayAgainstComputer)
}
