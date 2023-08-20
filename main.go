package main

import (
	"log"

	"github.com/bioerrorlog/chatpgt-ebitengine/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("chatgpt")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
