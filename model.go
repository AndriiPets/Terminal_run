package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type App struct {
	Game GameInterface
	Keys KeyMap
}

type TickMsg time.Time

func RunTUI(game GameInterface) {
	app := &App{
		Game: game,
		Keys: KeyMap{
			Up: key.NewBinding(key.WithKeys("k", "up", " ", "w"), key.WithHelp("↑/k/w/space", "jump")),
			Quit: key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q/ctrl+c", "quit")),
		},
	}

	p := tea.NewProgram(app)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
	}
}

func (app *App) tick() tea.Cmd {
	return tea.Tick(time.Second/5, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (app *App) Init() tea.Cmd {
	return app.tick()
}

func (app *App) evolve() (tea.Model, tea.Cmd) {
	app.Game.Evolve()
	return app, app.tick()
}

func (app *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	//User pressed a key
	case tea.KeyMsg:

		switch {

		case key.Matches(msg, app.Keys.Quit):
			fmt.Println()
			return app, tea.Quit

		case key.Matches(msg, app.Keys.Up):
			
			app.Game.Move("up")
		}
	
	case TickMsg:
		return app.evolve()
	}

	return app, nil
}

func (app *App) View() string {
	var sb strings.Builder
	frame := app.Game.ScreenData()

	sb.WriteString(strings.Join(frame, "\n"))

	return sb.String()
}