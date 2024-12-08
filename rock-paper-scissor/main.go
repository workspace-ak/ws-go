package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

type Symbols int

const (
	Rock    Symbols = iota // beats scissor, (scissors + 1) % 3 = 0
	Paper                  // beats rock, (rock + 1) % 3 = 1
	Scissor                // beats paper, (paper + 1) % 3 = 2
)

func (s Symbols) String() string {
	return [...]string{"Rock", "Paper", "Scissor"}[s]
}

type Player struct {
	name   string
	points int
	choice Symbols
}

type GameRound struct {
	Plyr *Player
	Comp *Player
}

func (g *GameRound) Result() string {
	res := ""
	switch {
	case g.Plyr.choice == g.Comp.choice:
		res = "Draw"
	case g.Plyr.choice == (g.Comp.choice+1)%3:
		res = fmt.Sprintf("%s Wins 😃", g.Plyr.name)
		g.Plyr.points++
	default:
		res = fmt.Sprintf("%s Wins 😒", g.Comp.name)
		g.Comp.points++
	}
	return res
}

func (g *GameRound) FinalResult() string {
	res := ""
	switch {
	case g.Plyr.points > g.Comp.points:
		res = fmt.Sprintf("%s has WON the game 🏆", g.Plyr.name)
	case g.Plyr.points < g.Comp.points:
		res = fmt.Sprintf("%s has WON the game 🏆", g.Comp.name)
	default:
		res = "Game is a DRAW..."
	}
	return res
}

func playGame() {
	rand.Seed(uint64(time.Now().UnixNano()))
	EndCondition := true

	var pl_name, comp_name string
	fmt.Println("[-] What's your name:")
	fmt.Scanln(&pl_name)
	fmt.Println("[-] With whom you want to play:")
	fmt.Scanln(&comp_name)
	fmt.Printf("[-] %s is playing rock paper scissor with %s\n", pl_name, comp_name)

	p := Player{pl_name, 0, 0}
	c := Player{comp_name, 0, 0}
	rounds := 0

	for EndCondition {
		rounds++
		var pl_choice Symbols
		var play string
		fmt.Println("[-] Choose rock(0), paper(1) or scissor(2): ")
		fmt.Scanln(&pl_choice)

		p.choice = pl_choice
		c.choice = Symbols(rand.Intn(3))
		fmt.Printf("[-] %s has choosen: %s\n", c.name, c.choice)

		round := GameRound{&p, &c}

		// round := GameRound{
		// 	Plyr: Player{"AK", 0, pl_choice},
		// 	Comp: Player{"PK", 0, Symbols(rand.Intn(3))},
		// }
		res := round.Result()

		fmt.Printf("[-] Round result: %s\n", res)

		fmt.Println("[-] Continue game (y/n): ")
		fmt.Scanln(&play)

		if play == "n" {
			EndCondition = false
			fmt.Println(round.FinalResult())
		}
	}

	fmt.Println("[-] Thanks for playing Rock Paper Scissor", pl_name)
	fmt.Println("===*===*===*===*===*===*===*===*====*===*===")
	fmt.Println("[-] Rounds:", rounds)
	fmt.Printf("[-] Final Points:\n\t%s: %d\n\t%s: %d\n", pl_name, p.points, comp_name, c.points)

}

func main() {
	playGame()
}
