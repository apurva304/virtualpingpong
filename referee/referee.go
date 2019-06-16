package referee

import (
	"errors"
	"strconv"

	"github.com/apurva304/virtualpingpong/domain"
	"github.com/apurva304/virtualpingpong/game"
)

var (
	ErrRequiredPlayerNotRegistered = errors.New("Required Number of players not registered")
	ErrPreviousRoundWinnerNotFound = errors.New("Previous Round winner not found")
)

type Referee struct {
	RegisteredPlayers []domain.PlayerDetail
	GameScores        []game.GameScore
	CurrentRound      int
	Round1Winners     []game.Participent
	Round2Winners     []game.Participent
	Champian          domain.PlayerDetail
}

func (ref *Referee) FirstRound() (err error) {
	if len(ref.RegisteredPlayers) < 8 {
		return ErrRequiredPlayerNotRegistered
	}

	for i := 0; i < len(ref.RegisteredPlayers); i += 2 {

		g := game.NewGame("#"+strconv.Itoa(i), ref.CurrentRound, ref.RegisteredPlayers[i], ref.RegisteredPlayers[i+1])
		gs, w := g.Run()
		ref.GameScores = append(ref.GameScores, gs)
		ref.Round1Winners = append(ref.Round1Winners, w)
	}
	ref.CurrentRound += 1
	return
}

func (ref *Referee) SecondRound() (err error) {
	if len(ref.Round1Winners) < 4 {
		return ErrPreviousRoundWinnerNotFound
	}

	for i := 0; i < len(ref.Round1Winners); i += 2 {
		g := game.NewGame("#"+strconv.Itoa(i+(len(ref.GameScores)/2)), ref.CurrentRound, ref.Round1Winners[i].PlayerDetail, ref.Round1Winners[i+1].PlayerDetail)
		gs, w := g.Run()
		ref.GameScores = append(ref.GameScores, gs)
		ref.Round2Winners = append(ref.Round2Winners, w)
	}
	ref.CurrentRound += 1
	return
}
func (ref *Referee) FinalRound() (err error) {
	if len(ref.Round2Winners) < 2 {
		return ErrPreviousRoundWinnerNotFound
	}

	for i := 0; i < len(ref.Round2Winners); i += 2 {
		g := game.NewGame("#"+strconv.Itoa(i+(len(ref.GameScores)/2)), ref.CurrentRound, ref.Round2Winners[i].PlayerDetail, ref.Round2Winners[i+1].PlayerDetail)
		gs, w := g.Run()
		ref.GameScores = append(ref.GameScores, gs)
		ref.Champian = domain.PlayerDetail{w.Id, w.Name, w.DefenceArrayLen}
	}
	ref.CurrentRound += 1
	return
}
