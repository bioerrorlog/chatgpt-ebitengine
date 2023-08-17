package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 1920
	ScreenHeight = 1080
)

type Game struct{}

func NewGame() (*Game, error) {
	g := &Game{}

	return g, nil
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	game, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("chatgpt")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
