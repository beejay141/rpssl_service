package choiceControllers

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"rpssl_service/core/choiceModule/choiceServices"
	"rpssl_service/infrastructures/services"
)

type ChoiceController struct {
	ChoiceService choiceServices.IChoiceService
	RandomNumberService services.IRandomNumberService
}

// controller action that returns all available choices
func (controller ChoiceController) GetAllChoices(ctx echo.Context) error {

	choices := controller.ChoiceService.GetChoices()

	return ctx.JSON(http.StatusOK,choices)
}

// controller action that randomly selects a choice and returns it
func (controller ChoiceController) GetRandomChoice(ctx echo.Context) error {

	number,err := controller.RandomNumberService.GetRandomNumber(5)
	if err != nil {
		log.Fatalf("An error occurred while getting ramdom number :%s",err.Error())
		number = 5
	}

	choice := controller.ChoiceService.GetChoice(number)

	return ctx.JSON(http.StatusOK, choice)
}

