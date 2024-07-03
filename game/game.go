package game

import (
	"math/rand"
	"strings"
)

type Game struct {
	screen_data    []string
	Width          int
	Height         int
	Player         char
	KeyPressed     bool
	Viewport       Vec2
	Gravity        int
	CurrentCommand string
	Obstacles      obstacles
	GameSpeed      int
	Counter        Counter
	Score          int
	PrevScore      int
}

const (
	WIDTH    int = 18
	HEIGHT   int = 12
	TILESIZE int = 16
)

func NewGame(w, h int) *Game {
	//18x12
	g := &Game{
		Width:    WIDTH,
		Height:   HEIGHT,
		Viewport: Vec2{X: w, Y: h},
		Player: char{
			X:  6,
			Y:  10,
			VY: 0,
			VX: 0,
		},
		KeyPressed: false,
		Gravity:    1,
		Obstacles:  newObstacles(),
		GameSpeed:  1,
		Counter: Counter{
			val:  1,
			size: 6,
		},
	}

	return g
}

var gravCouter = true

func (g *Game) applyGrav() {
	//Since gravity cannot be float
	//skip frames

	g.Player.VY++

}

func (g *Game) Evolve() {
	g.Counter.inc()
	ground := g.Viewport.Y - 2

	switch g.CurrentCommand {
	case "up":
		if g.Player.Y == ground {
			g.Player.VY -= 4
		}
	}
	g.applyGrav()

	g.Player.Y += g.Player.VY

	//Clamp player to the ground
	if g.Player.Y > ground {
		g.Player.Y = ground
		g.Player.VY = 0
	}

	//add score based on game speed
	if g.Counter.val == 6 {
		g.Score += g.GameSpeed
	}

	//increase game speed each 50 meters
	if g.Score-g.PrevScore > 50 {
		g.GameSpeed++
		g.PrevScore = g.Score
	}

	g.manageObstacles()

	//Reset button presses and commands
	g.KeyPressed = false
	g.CurrentCommand = "nil"

	g.renderGameScreen()
}

func (g *Game) GetScore() int {
	return g.Score
}

func (g *Game) manageObstacles() {
	ground := g.Viewport.Y - 2
	for _, obs := range g.Obstacles {
		//shift all obstacles right based on the currebt game speed
		obs.X -= g.GameSpeed

		if obs.X < 0 {
			g.Obstacles.remove()
		}
	}

	var rightmost *Vec2
	if len(g.Obstacles) > 0 {
		rightmost = g.Obstacles[len(g.Obstacles)-1].Vec2
	} else {
		rightmost = &Vec2{0, ground}
	}
	gap := g.Viewport.X - rightmost.X
	speefFactor := g.GameSpeed + 4

	//controlls frequency and heigth of obstacles
	if gap > 30+speefFactor || (testProbability(10) && gap > 15+speefFactor) {
		x := g.Viewport.X
		y := ground

		var h int
		switch {
		case testProbability(10):
			h = 3
		case testProbability(60):
			h = 2
		case testProbability(30):
			h = 1
		}

		g.Obstacles.add(newObstacle(h, &Vec2{x, y}))
		//fmt.Println(x,y) //!!
	}
}

func testProbability(percent int) bool {
	if rand.Intn(100) > 100-percent {
		return true
	}
	return false
}

func (g *Game) renderGameScreen() {
	viewport := make([]string, 0, g.Viewport.Y)

	for y := 0; y < g.Viewport.Y; y++ {
		var line strings.Builder
		var leftmost = 0
		for x := 0; x < g.Viewport.X; x++ {
			cellValue := ' '

			for _, o := range g.Obstacles[leftmost:] {
				if o.overlaps(Vec2{x, y}) {
					//fmt.Println("hit")
					cellValue = '&'
					leftmost++
					break
				}
			}

			if y == g.Viewport.Y-1 {
				cellValue = '#'
			}

			if g.Player.X == x && g.Player.Y == y {
				cellValue = 'p'
			}
			line.WriteRune(cellValue)
		}
		viewport = append(viewport, line.String())
	}

	g.screen_data = viewport
}

func (g *Game) ScreenData() []string {
	return g.screen_data
}

func (g *Game) Move(dir string) {
	//Locks keypress until current command is done
	if !g.KeyPressed {
		g.CurrentCommand = dir
	}
	g.KeyPressed = true
}
