package game

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

type Game struct {
	input              *Input
	output             string
	sendButton         *Button
	gptResultChan      chan string
	gptCalling         bool
	backgroundColor    color.RGBA
	backgroundColor2nd color.RGBA
	counter            int
}

func NewGame() (*Game, error) {
	i, err := NewInput()
	if err != nil {
		log.Fatal(err)
	}

	b, err := NewButton(550, ScreenHeight/2-60, 50, 50, "Send")
	if err != nil {
		log.Fatal(err)
	}

	g := &Game{
		input:              i,
		sendButton:         b,
		gptResultChan:      make(chan string),
		gptCalling:         false,
		backgroundColor:    color.RGBA{53, 54, 65, 255},
		backgroundColor2nd: color.RGBA{68, 70, 84, 255},
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

	// Call GPT when the button clicked
	if g.sendButton.isClicked && !g.gptCalling {
		go func() {
			fmt.Println("Call GPT")
			fmt.Println(g.input.text)
			result, err := CallGpt(g.input.text)
			if err != nil {
				log.Printf("Error calling GPT: %v", err)
			}
			g.gptResultChan <- result
		}()
		g.gptCalling = true
		g.output = ""
	}
	select {
	case result := <-g.gptResultChan:
		fmt.Println("Received:", result)
		g.output = result
		g.gptCalling = false
	default:
	}

	// Loading texts...
	if g.gptCalling {
		if g.counter%60 < 20 {
			g.output = ""
		} else if g.counter%80 < 40 {
			g.output = " ."
		} else if g.counter%80 < 60 {
			g.output = " . ."
		} else {
			g.output = " . . ."
		}
	}

	g.counter++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Background
	screen.Fill(g.backgroundColor)

	// Background for the lower half
	lowerHalfImage := ebiten.NewImage(ScreenWidth, ScreenHeight/2)
	lowerHalfImage.Fill(g.backgroundColor2nd)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, float64(ScreenHeight)/2)
	screen.DrawImage(lowerHalfImage, op)

	g.input.Draw(screen, 80, 60)
	g.sendButton.Draw(screen)

	// GPT response
	text.Draw(screen, g.output, basicfont.Face7x13, 80, ScreenHeight/2+60, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
