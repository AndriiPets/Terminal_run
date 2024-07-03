package main

type GameInterface interface {
	Evolve()
	Move(string)
	ScreenData() []string
	GetScore() int
}
