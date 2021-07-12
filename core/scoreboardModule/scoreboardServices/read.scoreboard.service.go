package scoreboardServices

import (
	"log"
	"rpssl_service/core/scoreboardModule/scoreboardModels"
	"rpssl_service/utils"
)

func (service ScoreboardService) GetScores(player string) []scoreboardModels.ScoreboardModel {
	storedScore,err := service.RedisService.Get(player).Bytes()
	var scores =  []scoreboardModels.ScoreboardModel{}
	if err == nil {
		err = utils.DecodeJSON(storedScore,&scores)
		if err != nil {
			log.Println(err.Error())
		}
	}

	return scores
}
