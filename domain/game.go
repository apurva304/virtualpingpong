package domain

type Game interface {
	Run(oPlayer, dPlayer PlayerDetail) (gs GameScore)
}

type GameScore struct {
	Round         int
	Winner        PlayerDetail
	WinnerPoint   int
	RunnerUp      PlayerDetail
	RunnerUpPoint int
}
