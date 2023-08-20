package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

type Game struct {
	input      *Input
	sendButton *Button
}

func NewGame() (*Game, error) {
	i, err := NewInput()
	if err != nil {
		log.Fatal(err)
	}

	b, err := NewButton(500, 400, 100, 50)
	if err != nil {
		log.Fatal(err)
	}

	g := &Game{
		input:      i,
		sendButton: b,
	}
	return g, nil
}

func (g *Game) Update() error {
	g.input.Update()
	g.sendButton.Update()

	if g.sendButton.IsClicked() {
		fmt.Println("Button was clicked!")
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.input.Draw(screen)
	g.sendButton.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
