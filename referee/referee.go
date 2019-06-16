package referee

import (
	"errors"

	"github.com/apurva304/virtualpingpong/domain"
)

var (
	ErrRequiredPlayerNotRegistered = errors.New("Required Number of players not registered")
	ErrPreviousRoundWinnerNotFound = errors.New("Previous Round winner not found")
	ErrChampainShipNotCompleted    = errors.New("Champiainship Not Completed yet")
)

type referee struct {
	registeredPlayers []domain.PlayerDetail
	gameScores        []domain.GameScore
	currentRound      int
	round1Winners     []domain.PlayerDetail
	round2Winners     []domain.PlayerDetail
	champian          domain.PlayerDetail
	game              domain.Game
}

func NewReferee(game domain.Game) *referee {
	return &referee{
		game: game,
	}
}
func (ref *referee) firstRound() (err error) {
	if len(ref.registeredPlayers) < 8 {
		return ErrRequiredPlayerNotRegistered
	}

	for i := 0; i < len(ref.registeredPlayers); i += 2 {

		gs := ref.game.Run(ref.registeredPlayers[i], ref.registeredPlayers[i+1])
		gs.Round = ref.currentRound
		ref.gameScores = append(ref.gameScores, gs)
		ref.round1Winners = append(ref.round1Winners, gs.Winner)
	}
	ref.currentRound += 1
	return
}

func (ref *referee) secondRound() (err error) {
	if len(ref.round1Winners) < 4 {
		return ErrPreviousRoundWinnerNotFound
	}

	for i := 0; i < len(ref.round1Winners); i += 2 {
		gs := ref.game.Run(ref.round1Winners[i], ref.round1Winners[i+1])
		gs.Round = ref.currentRound
		ref.gameScores = append(ref.gameScores, gs)
		ref.round2Winners = append(ref.round2Winners, gs.Winner)
	}
	ref.currentRound += 1
	return
}
func (ref *referee) finalRound() (err error) {
	if len(ref.round2Winners) < 2 {
		return ErrPreviousRoundWinnerNotFound
	}

	for i := 0; i < len(ref.round2Winners); i += 2 {
		gs := ref.game.Run(ref.round2Winners[i], ref.round2Winners[i+1])
		gs.Round = ref.currentRound
		ref.gameScores = append(ref.gameScores, gs)
		ref.champian = gs.Winner
	}
	ref.currentRound += 1
	return
}
func (ref *referee) ListGamesScore() (fsb []domain.FinalScoreBoard, champain domain.PlayerDetail, err error) {
	if ref.currentRound <= 3 {
		err = ErrChampainShipNotCompleted
		return
	}

	for _, gs := range ref.gameScores {
		f := domain.FinalScoreBoard{
			Round:          gs.Round,
			WinnerName:     gs.Winner.Name,
			WinnerPoints:   gs.WinnerPoint,
			RunnerUpName:   gs.RunnerUp.Name,
			RunnerUpPoints: gs.RunnerUpPoint,
		}

		fsb = append(fsb, f)
	}
	champain = ref.champian

	return
}

func (ref *referee) RegisterPlayer(pd domain.PlayerDetail) {
	ref.registeredPlayers = append(ref.registeredPlayers, pd)
}

func (ref *referee) StartChampainShip() (err error) {
	if ref.currentRound == 0 {
		ref.currentRound = 1
	}
	err = ref.firstRound()
	if err != nil {
		return
	}

	err = ref.secondRound()
	if err != nil {
		return
	}

	err = ref.finalRound()
	if err != nil {
		return
	}

	return
}
