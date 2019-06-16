package domain

type Referee interface {
	RegisterPlayer(pd PlayerDetail)
	StartChampainShip() (err error)
	ListGamesScore(fsb []FinalScoreBoard, champain PlayerDetail, err error)
}

type FinalScoreBoard struct {
	Round          int
	WinnerName     string
	WinnerPoints   int
	RunnerUpName   string
	RunnerUpPoints int
}
