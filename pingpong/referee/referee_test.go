package referee

import (
	"testing"

	"github.com/apurva304/virtualpingpong/domain"
	"github.com/apurva304/virtualpingpong/game"
)

var (
	ref = referee{currentRound: 1,
		game: game.NewGame(),
	}

	pds = []domain.PlayerDetail{
		domain.PlayerDetail{
			Name:            "xyz",
			DefenceArrayLen: 5,
		},
		domain.PlayerDetail{
			Name:            "xyz",
			DefenceArrayLen: 5,
		},
		domain.PlayerDetail{
			Name:            "xyz",
			DefenceArrayLen: 5,
		},
		domain.PlayerDetail{
			Name:            "xyz",
			DefenceArrayLen: 5,
		},
		domain.PlayerDetail{
			Name:            "xyz",
			DefenceArrayLen: 5,
		},
		domain.PlayerDetail{
			Name:            "xyz",
			DefenceArrayLen: 5,
		},
		domain.PlayerDetail{
			Name:            "xyz",
			DefenceArrayLen: 5,
		},
		domain.PlayerDetail{
			Name:            "xyz",
			DefenceArrayLen: 5,
		},
	}
)

func TestRegisterPlayer(t *testing.T) {
	for _, pd := range pds {
		ref.RegisterPlayer(pd)
	}

	if len(ref.registeredPlayers) != 8 {
		t.Error("Expect 8 player to register")
	}
}
func TestFirstRound(t *testing.T) {
	err := ref.firstRound()
	if err != nil && err != ErrRequiredPlayerNotRegistered {
		t.Error(err)
	}

	if err == ErrRequiredPlayerNotRegistered {
		return
	}

	if len(ref.round1Winners) != 4 {
		t.Error("Round 1 Winners not added")
	}

	if len(ref.gameScores) != 4 {
		t.Error("Game Score not added")
	}
}

func TestSecondRound(t *testing.T) {
	err := ref.secondRound()
	if err != nil && err != ErrPreviousRoundWinnerNotFound {
		t.Error(err)
	}

	if err == ErrPreviousRoundWinnerNotFound {
		return
	}

	if len(ref.round2Winners) != 2 {
		t.Error("Round 2 winners not added")
	}

	if len(ref.gameScores) != 6 {
		t.Error("Game Score not added")
	}
}

func TestFinalRound(t *testing.T) {
	err := ref.finalRound()
	if err != nil && err != ErrPreviousRoundWinnerNotFound {
		t.Error(err)
	}

	if err == ErrPreviousRoundWinnerNotFound {
		return
	}

	emptyPlayerDetail := domain.PlayerDetail{}

	if ref.champian == emptyPlayerDetail {
		t.Error("Champain not added")
	}

	if len(ref.gameScores) != 7 {
		t.Error("Game Score not added")
	}
}

func TestListFinalScore(t *testing.T) {
	fsb, champ, err := ref.ListGamesScore()
	if err != nil && err != ErrChampainShipNotCompleted {
		t.Error(err)
	}

	if len(fsb) != 7 {
		t.Error("Expect Length of final Score board to be 7")
	}

	emptyPlayerDetail := domain.PlayerDetail{}
	if champ == emptyPlayerDetail {
		t.Error("Incorrect champain")
	}
}
