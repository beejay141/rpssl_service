package scoreboardServices

import (
	"rpssl_service/config"
	"rpssl_service/core/scoreboardModule/scoreboardModels"
	"rpssl_service/utils"
)

func (service ScoreboardService) AddScores(score scoreboardModels.ScoreboardModel) error {

    var scores =  service.GetScores(score.Player)

	if len(scores) == 10 {
		scores = scores[:9]
	}

	//scores = append([]scoreboardModels.ScoreboardModel{score}, scores...)
	scores = append(scores,score)
	encodedData, err := utils.EncodeJSON(scores)
	if err != nil {
		return err
	}
	return service.RedisService.Set(score.Player,encodedData,config.GetInt("SCOREBOARD_EXP",0))
}
