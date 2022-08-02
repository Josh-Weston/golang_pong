package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/josh-weston/golang_pong/game"
)

func main() {
	// initialize the screen
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	b := game.Ball{
		X:      1,
		Y:      1,
		Xspeed: 1,
		Yspeed: 1,
	}

	g := game.Game{
		Screen: screen,
		Ball:   b,
	}

	go g.Run()

	// prevents main from exiting
	for {
		// how we can listen to user events
		switch event := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				screen.Fini()
				os.Exit(0)
			}
		}
	}
}
