package scoreboardRouter

import (
	"fmt"
	"github.com/labstack/echo"
	"rpssl_service/core/scoreboardModule/scoreboardController"
	"rpssl_service/core/scoreboardModule/scoreboardServices"
	"rpssl_service/infrastructures/services"
)


var scoreController = scoreboardController.ScoreboardController{
	ScoreboardService: scoreboardServices.ScoreboardService{
		RedisService: services.RedisService{},
	},
}

func RegisterRoute(router *echo.Group, root string) {
	router.GET(fmt.Sprintf( "%s/:id",root), scoreController.GetScores)

	router.DELETE(fmt.Sprintf( "%s/reset/:id",root), scoreController.ResetScores)
}
