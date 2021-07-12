package scoreboardController

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"rpssl_service/core/scoreboardModule/scoreboardServices"
)

type ScoreboardController struct {
	ScoreboardService scoreboardServices.IScoreboardService
}

// controller action that returns player's scores
func (controller ScoreboardController) GetScores(ctx echo.Context) error {

	playerId := ctx.Param("id")

	scores := controller.ScoreboardService.GetScores(playerId)

	return ctx.JSON(http.StatusOK,scores)
}

// controller action that returns player's scores
func (controller ScoreboardController) ResetScores(ctx echo.Context) error {

	playerId := ctx.Param("id")

	err := controller.ScoreboardService.ResetScores(playerId)

	if err != nil {
		log.Printf("error resetting scores, %v",err)
		return ctx.JSON(http.StatusBadRequest,"scores not found")
	}

	return ctx.JSON(http.StatusOK,"your scoreboard has been cleared")
}
