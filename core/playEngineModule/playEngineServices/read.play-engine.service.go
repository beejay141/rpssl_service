package playEngineServices

// Read https://www.cs.drexel.edu/~jpopyack/Courses/CSP/Fa17/notes/CS140_RockPaperScissors_Revisited.pdf
// This simplifies RPSSL with this algorithm
func (service PlayEngineService) Play(player1 , player2 int) int {

	b:= player1-player2
	if mod(mod(b,5) ,2)== 1 {
		return player1
	}

	return player2
}

func mod(x,y int) int {
	if x<0 {
		return ((x%y)+y)%y
	}
	return x%y
}
