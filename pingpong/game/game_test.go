package game

import (
	"testing"

	"github.com/apurva304/virtualpingpong/domain"
)

var (
	//mock data
	g = game{}
)

func TestGame(t *testing.T) {
	gs := g.Run(domain.PlayerDetail{Name: "abc", DefenceArrayLen: 7}, domain.PlayerDetail{Name: "xyz", DefenceArrayLen: 5})
	if gs.WinnerPoint < 5 {
		t.Error("Winners points must be greater then or equal to 5")
	}
}
