package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/apurva304/virtualpingpong/referee"

	"github.com/apurva304/virtualpingpong/game"

	"github.com/apurva304/virtualpingpong/domain"
)

var (
	pds = []domain.PlayerDetail{
		domain.PlayerDetail{
			Name:            "Joey",
			DefenceArrayLen: 7,
		},

		domain.PlayerDetail{
			Name:            "Monica",
			DefenceArrayLen: 6,
		},

		domain.PlayerDetail{
			Name:            "Chandler",
			DefenceArrayLen: 6,
		},

		domain.PlayerDetail{
			Name:            "Ross",
			DefenceArrayLen: 5,
		},

		domain.PlayerDetail{
			Name:            "Phoebe",
			DefenceArrayLen: 5,
		},

		domain.PlayerDetail{
			Name:            "Rachel",
			DefenceArrayLen: 6,
		},

		domain.PlayerDetail{
			Name:            "Sachin",
			DefenceArrayLen: 4,
		},

		domain.PlayerDetail{
			Name:            "Rohan",
			DefenceArrayLen: 5,
		},
	}
)

func main() {
	rand.Seed(time.Now().UnixNano())
	t := time.Now()

	game := game.NewGame()
	referee := referee.NewReferee(game)

	for _, pd := range pds {
		referee.RegisterPlayer(pd)
	}

	referee.StartChampainShip()

	fsb, champ, err := referee.ListGamesScore()
	if err != nil {
		panic(err)
	}

	fmt.Println("game", "round", "winner", "winner's Points", "RunnerUp", "RunnerUp's Points")
	for i, f := range fsb {
		fmt.Println(i+1, f.Round, f.WinnerName, f.WinnerPoints, f.RunnerUpName, f.RunnerUpPoints)
	}
	fmt.Println("\nChampain of the Game is ", champ.Name, "\n")
	fmt.Println("Time taken ", time.Since(t))
}
