package referee

import (
	"testing"

	"github.com/apurva304/virtualpingpong/domain"
)

var (
	ref = Referee{CurrentRound: 1,
		RegisteredPlayers: []domain.PlayerDetail{
			domain.PlayerDetail{
				Id:              "#1",
				Name:            "xyz",
				DefenceArrayLen: 5,
			},
			domain.PlayerDetail{
				Id:              "#2",
				Name:            "xyz",
				DefenceArrayLen: 5,
			},
			domain.PlayerDetail{
				Id:              "#3",
				Name:            "xyz",
				DefenceArrayLen: 5,
			},
			domain.PlayerDetail{
				Id:              "#4",
				Name:            "xyz",
				DefenceArrayLen: 5,
			},
			domain.PlayerDetail{
				Id:              "#5",
				Name:            "xyz",
				DefenceArrayLen: 5,
			},
			domain.PlayerDetail{
				Id:              "#6",
				Name:            "xyz",
				DefenceArrayLen: 5,
			},
			domain.PlayerDetail{
				Id:              "#7",
				Name:            "xyz",
				DefenceArrayLen: 5,
			},
			domain.PlayerDetail{
				Id:              "#8",
				Name:            "xyz",
				DefenceArrayLen: 5,
			},
		},
	}
)

func TestFirstRound(t *testing.T) {
	err := ref.FirstRound()
	if err != nil && err != ErrRequiredPlayerNotRegistered {
		t.Error(err)
	}

	if err == ErrRequiredPlayerNotRegistered {
		return
	}

	if len(ref.Round1Winners) != 4 {
		t.Error("Round 1 Winners not added")
	}

	if len(ref.GameScores) != 4 {
		t.Error("Game Score not added")
	}
}

func TestSecondRound(t *testing.T) {
	err := ref.SecondRound()
	if err != nil && err != ErrPreviousRoundWinnerNotFound {
		t.Error(err)
	}

	if err == ErrPreviousRoundWinnerNotFound {
		return
	}

	if len(ref.Round2Winners) != 2 {
		t.Error("Round 2 winners not added")
	}

	if len(ref.GameScores) != 6 {
		t.Error("Game Score not added")
	}
}

func TestFinalRound(t *testing.T) {
	err := ref.FinalRound()
	if err != nil && err != ErrPreviousRoundWinnerNotFound {
		t.Error(err)
	}

	if err == ErrPreviousRoundWinnerNotFound {
		return
	}

	emptyPlayDetail := domain.PlayerDetail{}

	if ref.Champian == emptyPlayDetail {
		t.Error("Champain not added")
	}

	if len(ref.GameScores) != 7 {
		t.Error("Game Score not added")
	}
}
