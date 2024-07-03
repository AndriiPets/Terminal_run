package main

import (
	"github.com/AndriiPets/terminal_rouge/game"
)

func main() {
	g := game.NewGame(50, 12)

	RunTUI(g)
}
