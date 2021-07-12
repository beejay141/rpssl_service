package scoreboardServices

import (
	"rpssl_service/core/scoreboardModule/scoreboardModels"
	"rpssl_service/infrastructures/services"
)

type ScoreboardService struct {
	RedisService services.IRedisService
}

type IScoreboardService interface {
	AddScores(score scoreboardModels.ScoreboardModel) error
	GetScores(player string) []scoreboardModels.ScoreboardModel
	ResetScores(player string) error
}

