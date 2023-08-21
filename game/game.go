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
	gptResultChan   chan string
	gptCalling      bool
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
		gptResultChan:   make(chan string),
		gptCalling:      false,
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

	// Call GPT when the button pressed
	if g.sendButton.IsClicked() && !g.gptCalling {
		go func() {
			fmt.Println("Call GPT")
			result, err := CallGpt()
			if err != nil {
				log.Printf("Error calling GPT: %v", err)
			}
			g.gptResultChan <- result
		}()
		g.gptCalling = true
	}

	select {
	case result := <-g.gptResultChan:
		fmt.Println("Received:", result)
		g.gptCalling = false
	default:
		// fmt.Println("Waiting for the response of GPT call")
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
