package game

import (
	"testing"

	"github.com/apurva304/virtualpingpong/domain"
)

var (
	//mock data
	g = Game{
		Id:    "#1",
		Round: 1,
		OPlayer: Participent{
			PlayerDetail: domain.PlayerDetail{Id: "#1", Name: "abc", DefenceArrayLen: 7},
		},
		DPlayer: Participent{
			PlayerDetail: domain.PlayerDetail{Id: "#2", Name: "xyz", DefenceArrayLen: 5},
		},
	}
)

func TestGame(t *testing.T) {
	_, w := g.Run()
	if w.Point < 5 {
		t.Error("Winners points must be greater then or equal to 5")
	}
}
