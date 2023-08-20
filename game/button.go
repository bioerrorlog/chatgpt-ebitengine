package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Button struct {
	x, y          int
	width, height int
	isClicked     bool
}

func NewButton(x, y, width, height int) (*Button, error) {
	b := &Button{
		x:      x,
		y:      y,
		width:  width,
		height: height,
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
	buttonImage := ebiten.NewImage(b.width, b.height)
	buttonColor := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	buttonImage.Fill(buttonColor)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.x), float64(b.y))
	screen.DrawImage(buttonImage, op)
}

func (b *Button) IsClicked() bool {
	return b.isClicked
}
