package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/apurva304/virtualpingpong/domain"
	"github.com/apurva304/virtualpingpong/referee"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ref := referee.Referee{CurrentRound: 1}

	ref.RegisteredPlayers = append(ref.RegisteredPlayers, domain.PlayerDetail{
		Id:              "#1",
		Name:            "Joey",
		DefenceArrayLen: 7,
	})

	ref.RegisteredPlayers = append(ref.RegisteredPlayers, domain.PlayerDetail{
		Id:              "#2",
		Name:            "Monica",
		DefenceArrayLen: 6,
	})

	ref.RegisteredPlayers = append(ref.RegisteredPlayers, domain.PlayerDetail{
		Id:              "#3",
		Name:            "Chandler",
		DefenceArrayLen: 6,
	})

	ref.RegisteredPlayers = append(ref.RegisteredPlayers, domain.PlayerDetail{
		Id:              "#4",
		Name:            "Ross",
		DefenceArrayLen: 5,
	})

	ref.RegisteredPlayers = append(ref.RegisteredPlayers, domain.PlayerDetail{
		Id:              "#5",
		Name:            "Phoebe",
		DefenceArrayLen: 5,
	})

	ref.RegisteredPlayers = append(ref.RegisteredPlayers, domain.PlayerDetail{
		Id:              "#6",
		Name:            "Rachel",
		DefenceArrayLen: 6,
	})

	ref.RegisteredPlayers = append(ref.RegisteredPlayers, domain.PlayerDetail{
		Id:              "#7",
		Name:            "Sachin",
		DefenceArrayLen: 4,
	})

	ref.RegisteredPlayers = append(ref.RegisteredPlayers, domain.PlayerDetail{
		Id:              "#8",
		Name:            "Rohan",
		DefenceArrayLen: 5,
	})

	ref.FirstRound()
	ref.SecondRound()
	ref.FinalRound()
	fmt.Println(ref.Champian)
}
