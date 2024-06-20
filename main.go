package main

import (
	"github.com/AndriiPets/terminal_rouge/game"
)

func main() {
	g := game.NewGame(18, 12)

	RunTUI(g)
}