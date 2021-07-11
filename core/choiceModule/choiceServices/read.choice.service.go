package choiceServices

import "rpssl_service/core/choiceModule/choiceModels"

func (service ChoiceService) GetChoices() []choiceModels.Choice {
	return choices
}

func (service ChoiceService) GetChoice(id int) choiceModels.Choice {
	return choices[id-1]
}
