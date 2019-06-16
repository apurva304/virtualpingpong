package game

import (
	"math/rand"

	"github.com/apurva304/virtualpingpong/domain"
)

type participent struct {
	domain.PlayerDetail
	defenceMap map[int]struct{} // for defence player, function is searching here so map instead of array
	selected   int              // for offensive player
	point      int
}

type game struct {
	oPlayer participent //offensive player
	dPlayer participent // defensive player
}

func NewGame() *game {
	return &game{}
}
func (g *game) Run(oPlayer, dPlayer domain.PlayerDetail) (gs domain.GameScore) {

	g.oPlayer = participent{PlayerDetail: oPlayer}
	g.dPlayer = participent{PlayerDetail: dPlayer}
	for g.oPlayer.point < 5 && g.dPlayer.point < 5 {

		g.oPlayer.selected = rand.Intn(10) + 1                                   // rand.Intn generates number for [0,n) i.e. [0,10)
		g.dPlayer.defenceMap = make(map[int]struct{}, g.dPlayer.DefenceArrayLen) // clear the previous map and initialize it if not done

		for len(g.dPlayer.defenceMap) != g.dPlayer.DefenceArrayLen {
			g.dPlayer.defenceMap[rand.Intn(10)+1] = struct{}{}
		}
		if _, ok := g.dPlayer.defenceMap[g.oPlayer.selected]; ok {
			g.dPlayer.point += 1
			g.oPlayer, g.dPlayer = g.dPlayer, g.oPlayer // switch roles
		} else {
			g.oPlayer.point += 1
		}
	}

	var winner, runnerUp participent

	if g.oPlayer.point >= 5 {
		winner = g.oPlayer
		runnerUp = g.dPlayer
	} else {
		winner = g.dPlayer
		runnerUp = g.oPlayer
	}

	gs.Winner = winner.PlayerDetail
	gs.WinnerPoint = winner.point

	gs.RunnerUp = runnerUp.PlayerDetail
	gs.RunnerUpPoint = runnerUp.point
	return
}
