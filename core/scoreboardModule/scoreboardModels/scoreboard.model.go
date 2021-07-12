package scoreboardModels

type ScoreboardModel struct {
	Player string `json:"player"`
	Choice int `json:"choice"`
	ComputerChoice int `json:"computerChoice"`
	Won bool `json:"won"`
}
