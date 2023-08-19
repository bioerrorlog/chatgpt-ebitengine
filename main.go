package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("chatgpt")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
