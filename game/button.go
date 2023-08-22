package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

type Button struct {
	x, y          int
	width, height int
	color         color.RGBA
	colorClicked  color.RGBA
	buttonText    string
	isClicked     bool
}

func NewButton(x, y, width, height int, buttonText string) (*Button, error) {
	b := &Button{
		x:            x,
		y:            y,
		width:        width,
		height:       height,
		buttonText:   buttonText,
		color:        color.RGBA{R: 171, G: 104, B: 255, A: 255},
		colorClicked: color.RGBA{R: 171 - 20, G: 104 - 20, B: 255 - 20, A: 255},
	}
	return b, nil
}

func (b *Button) Update() error {
	x, y := ebiten.CursorPosition()

	// Check the cursor position
	if x >= b.x && x <= b.x+b.width && y >= b.y && y <= b.y+b.height {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			b.isClicked = true
		} else {
			b.isClicked = false
		}
	} else {
		b.isClicked = false
	}

	return nil
}

func (b *Button) Draw(screen *ebiten.Image) {
	fillColor := b.color
	if b.isClicked {
		fillColor = b.colorClicked
	}

	buttonImage := ebiten.NewImage(b.width, b.height)
	buttonImage.Fill(fillColor)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.x), float64(b.y))
	screen.DrawImage(buttonImage, op)

	// Draw the button text
	textX := b.x + b.width/2 - 15
	textY := b.y + b.height/2 + 5
	text.Draw(screen, b.buttonText, basicfont.Face7x13, textX, textY, color.White)
}

func (b *Button) IsClicked() bool {
	return b.isClicked
}
