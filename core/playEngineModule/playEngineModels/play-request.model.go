package playEngineModels

type PlayRequestModel struct {
	Player int `json:"player" validate:"required,min=1,max=5"`
}
