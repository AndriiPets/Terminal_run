package game

import (
	"strings"
)

type Game struct {
	screen_data []string
	Width int
	Height int
	Player char
	KeyPressed bool
	Viewport Vec2
	Gravity int
	CurrentCommand string
	Obstacles obstacles
}

const (
	WIDTH int = 18
	HEIGHT int = 12
	TILESIZE int = 16
)

func NewGame(w, h int) *Game {
	//18x12
	g := &Game{
		Width: WIDTH,
		Height: HEIGHT,
		Viewport: Vec2{X: w, Y: h},
		Player: char{
			X: 6,
			Y: 10,
			VY: 0,
			VX: 0,
		},
		KeyPressed: false,
		Gravity: 1,
		Obstacles: newObstacles(),
	}

	return g
}

func (g *Game) Evolve() {
	ground := g.Viewport.Y - 2
	//apply gravity
	g.Player.VY += g.Gravity

	switch g.CurrentCommand {
	case "up":
		if g.Player.Y == ground {
			g.Player.VY -= 4
		}
	}

	g.Player.Y += g.Player.VY

	//Clamp player to the ground
	if g.Player.Y > ground {
		g.Player.Y = ground
		g.Player.VY = 0
	}

	g.manageObstacles()
	

	//Reset button presses and commands
	g.KeyPressed = false
	g.CurrentCommand = "nil"
	
	g.renderGameScreen()
}

func (g *Game) manageObstacles() {
	ground := g.Viewport.Y - 2
	for _, obs := range g.Obstacles {
		obs.X --

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

	if gap > 5 {
		x := g.Viewport.X
		y := ground

		g.Obstacles.add(newObstacle(2, &Vec2{x, y}))
		//fmt.Println(x,y) //!!
	}
}

func (g *Game) renderGameScreen() {
	viewport := make([]string, 0, g.Viewport.Y)

	for y := 0; y < g.Viewport.Y; y++ {
		var line strings.Builder
		var leftmost = 0
		for x := 0; x < g.Viewport.X; x++ {
			cellValue := ' '

			for _, o := range g.Obstacles[leftmost:] {
				if o.overlaps(Vec2{x,y}) {
					//fmt.Println("hit")
					cellValue = '&'
					leftmost ++
					break
				}
			}

			if y == g.Viewport.Y - 1 {
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