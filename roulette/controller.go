package roulette

func postController(bets []Bet, spinWheelFunc SpinWheelFunc) (int, int) {
	winningNumber := spinWheelFunc()
	winnings := 0
	for _, bet := range bets {
		winnings += calculateWinnings(bet, winningNumber)
	}
	return winningNumber, winnings
}
