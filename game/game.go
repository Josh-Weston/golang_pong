package game

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	Screen tcell.Screen
	Ball   Ball
}

func (g *Game) Run() {
	// set the screen's style
	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorBlack)
	g.Screen.SetStyle(defStyle)

	// this is similar to the draw() function in p5.js
	for {
		g.Screen.Clear()
		g.Ball.Update()
		g.Screen.SetContent(g.Ball.X, g.Ball.Y, g.Ball.Display(), nil, defStyle)

		// show the content
		g.Screen.Show()
		time.Sleep(40 * time.Millisecond)

		width, height := g.Screen.Size()
		g.Ball.CheckEdges(width, height)
	}
}

type Ball struct {
	X      int
	Y      int
	Xspeed int
	Yspeed int
}

func (b *Ball) Display() rune {
	return '\u25CF'
}

func (b *Ball) Update() {
	b.X += b.Xspeed
	b.Y += b.Yspeed
}

func (b *Ball) CheckEdges(maxWidth int, maxHeight int) {
	if b.X <= 0 || b.X >= maxWidth {
		b.Xspeed *= -1 // change direction
	}
	if b.Y <= 0 || b.Y >= maxHeight {
		b.Yspeed *= -1 // change direction
	}
}
