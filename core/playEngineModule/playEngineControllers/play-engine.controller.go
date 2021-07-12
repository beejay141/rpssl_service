package playEngineControllers

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"rpssl_service/core/choiceModule/choiceServices"
	"rpssl_service/core/playEngineModule/playEngineModels"
	"rpssl_service/core/playEngineModule/playEngineServices"
	"rpssl_service/core/scoreboardModule/scoreboardModels"
	"rpssl_service/core/scoreboardModule/scoreboardServices"
	"rpssl_service/infrastructures/services"
	"rpssl_service/utils"
)

type PlayEngineController struct {
    PlayEngineService  playEngineServices.IPlayEngineService
    ChoiceService      choiceServices.IChoiceService
    RandomNumberService services.IRandomNumberService
    ScoreboardService scoreboardServices.IScoreboardService
}

// controller action that process game result again computer
func (controller PlayEngineController) PlayAgainstComputer(ctx echo.Context) error {

	request := new(playEngineModels.PlayRequestModel)

	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request body.")
	}

	if err := utils.ValidateRequest(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	randomNumber, err := controller.RandomNumberService.GetRandomNumber(5)
	if err != nil {
		log.Fatalf("Error getting randmon number from API service, however 5 was used")
		randomNumber = 5
	}

	//if random number generated is equal to the player's choice then it;s tie
	if randomNumber == request.Player {
		result:= map[string]interface{}{
			"results":"tie",
			"player":randomNumber,
			"computer" :randomNumber,
		}
		return ctx.JSON(http.StatusOK, result)
	}

	//resultcomputerChoice := controller.ChoiceService.GetChoice(randomNumber)

	winnerChoice := controller.PlayEngineService.Play(request.Player,randomNumber)

	result := map[string]interface{}{
		"player": request.Player,
		"computer":randomNumber,
	}

	if winnerChoice == request.Player {
		result["results"]="win"
	}else {
		result["results"]="lose"
	}

	if request.PlayerId != "" {
		err := controller.ScoreboardService.AddScores(scoreboardModels.ScoreboardModel{
			Player: request.PlayerId,
			Choice: request.Player,
			ComputerChoice: randomNumber,
			Won: result["results"] == "win",
		})

		if err != nil {
			log.Printf("error saving score : %v",err)
		}
	}

	return ctx.JSON(http.StatusOK,result)
}