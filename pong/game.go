package pong

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	Screen tcell.Screen
	Ball   Ball
}

func (g *Game) Run() {

	defStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite)
	g.Screen.SetStyle(defStyle)

	for {
		g.Screen.Clear()

		width, height := g.Screen.Size()
		g.Ball.CheckEdges(width, height)
		g.Ball.Update()

		g.Screen.SetContent(g.Ball.X, g.Ball.Y, g.Ball.Display(), nil, defStyle)

		time.Sleep(50 * time.Millisecond)
		g.Screen.Show()
	}
}
