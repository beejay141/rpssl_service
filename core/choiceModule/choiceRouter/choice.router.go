package choiceRouter

import (
	"fmt"
	"github.com/labstack/echo"
	"rpssl_service/core/choiceModule/choiceControllers"
	"rpssl_service/core/choiceModule/choiceServices"
	"rpssl_service/infrastructures/helpers"
	"rpssl_service/infrastructures/services"
)

var choiceController = choiceControllers.ChoiceController{
	ChoiceService: choiceServices.ChoiceService{},
	RandomNumberService: services.RandomNumberService{
		RestRequestHelperService: helpers.RestRequestHelper{},
		},
    }

func RegisterRoute(router *echo.Group, root string) {
	router.GET(fmt.Sprintf( "%s/choices",root), choiceController.GetAllChoices)
	router.GET(fmt.Sprintf( "%s/choice",root), choiceController.GetRandomChoice)
}
