package choiceServices

import "rpssl_service/core/choiceModule/choiceModels"

type IChoiceService interface {
	GetChoices() []choiceModels.Choice
	GetChoice(id int) choiceModels.Choice
}

type ChoiceService struct {}


var choices = []choiceModels.Choice{
	{Name: "rock",Id: 1},
	{Name: "paper",Id: 2},
	{Name: "scissors",Id: 3},
	{Name: "spock", Id: 4},
	{Name: "lizard",Id: 5},
}
