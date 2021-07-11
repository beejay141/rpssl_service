package playEngineServices

type IPlayEngineService interface {
	Play(player1 , player2 int) int
}

type PlayEngineService struct {}

