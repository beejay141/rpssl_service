package scoreboardServices



func (service ScoreboardService) ResetScores(player string) error {
	_,err := service.RedisService.Remove(player)
	return err
}
