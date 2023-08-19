package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

type Game struct {
	input *Input
}

func NewGame() (*Game, error) {
	i, err := NewInput()
	if err != nil {
		log.Fatal(err)
	}

	g := &Game{
		input: i,
	}
	return g, nil
}

func (g *Game) Update() error {
	g.input.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.input.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
