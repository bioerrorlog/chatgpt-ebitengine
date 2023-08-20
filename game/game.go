package game

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

type Game struct {
	input           *Input
	sendButton      *Button
	backgroundColor color.RGBA
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
		input:           i,
		sendButton:      b,
		backgroundColor: color.RGBA{53, 54, 65, 255},
	}
	return g, nil
}

func (g *Game) Update() error {
	if err := g.input.Update(); err != nil {
		return err
	}
	if err := g.sendButton.Update(); err != nil {
		return err
	}

	if g.sendButton.IsClicked() {
		fmt.Println("Button was clicked!")
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Background
	screen.Fill(g.backgroundColor)

	g.input.Draw(screen)
	g.sendButton.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
