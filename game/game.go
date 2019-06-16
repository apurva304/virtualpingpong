package game

import (
	"math/rand"

	"github.com/apurva304/virtualpingpong/domain"
)

type Participent struct {
	domain.PlayerDetail
	DefenceMap map[int]struct{} // for defence player, function is searching here so map instead of array
	Selected   int              // for offensive player
	Point      int
}

type PlayerScore struct {
	PlayerId string
	Point    int
	IsWinner bool
}

type Game struct {
	Id      string
	Round   int
	OPlayer Participent //offensive player
	DPlayer Participent // defensive player
}
type GameScore struct {
	Id           string
	Round        int
	PlayerScores []PlayerScore
	WinnerName   string
}

func NewGame(id string, round int, oPlayer domain.PlayerDetail, dPlayer domain.PlayerDetail) *Game {
	return &Game{
		Id:      id,
		Round:   round,
		OPlayer: Participent{PlayerDetail: oPlayer},
		DPlayer: Participent{PlayerDetail: dPlayer},
	}
}
func (g *Game) Run() (gs GameScore, winner Participent) {

	for g.OPlayer.Point < 5 && g.DPlayer.Point < 5 {

		g.OPlayer.Selected = rand.Intn(10) + 1                                   // rand.Intn generates number for [0,n) i.e. [0,10)
		g.DPlayer.DefenceMap = make(map[int]struct{}, g.DPlayer.DefenceArrayLen) // clear the previous map and initialize it if not done

		for len(g.DPlayer.DefenceMap) != g.DPlayer.DefenceArrayLen {
			g.DPlayer.DefenceMap[rand.Intn(10)+1] = struct{}{}
		}
		if _, ok := g.DPlayer.DefenceMap[g.OPlayer.Selected]; ok {
			g.DPlayer.Point += 1
			g.OPlayer, g.DPlayer = g.DPlayer, g.OPlayer // switch roles
		} else {
			g.OPlayer.Point += 1
		}
	}

	oplayerScore := PlayerScore{
		PlayerId: g.OPlayer.Id,
		Point:    g.OPlayer.Point,
		IsWinner: g.OPlayer.Point >= 5,
	}

	dplayerScore := PlayerScore{
		PlayerId: g.DPlayer.Id,
		Point:    g.DPlayer.Point,
		IsWinner: g.DPlayer.Point >= 5,
	}

	gs.PlayerScores = append(gs.PlayerScores, oplayerScore, dplayerScore)
	gs.Id = g.Id
	gs.Round = g.Round

	gs.WinnerName = g.DPlayer.Name
	winner = g.DPlayer

	if oplayerScore.IsWinner {
		gs.WinnerName = g.OPlayer.Name
		winner = g.OPlayer
	}
	return
}
