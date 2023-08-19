// Reference
// https://github.com/hajimehoshi/ebiten/blob/main/examples/typewriter/main.go

package main

import (
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

type Input struct {
	runes   []rune
	text    string
	counter int
}

func NewInput() (*Input, error) {
	i := &Input{
		text:    "Type on the keyboard:\n",
		counter: 0,
	}
	return i, nil
}

func (i *Input) Update() error {
	// Add runes that are input by the user by AppendInputChars.
	// Note that AppendInputChars result changes every frame, so you need to call this
	// every frame.
	i.runes = ebiten.AppendInputChars(i.runes[:0])
	i.text += string(i.runes)

	// Adjust the string to be at most 10 lines.
	ss := strings.Split(i.text, "\n")
	if len(ss) > 10 {
		i.text = strings.Join(ss[len(ss)-10:], "\n")
	}

	// If the enter key is pressed, add a line break.
	if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyNumpadEnter) {
		i.text += "\n"
	}

	// If the backspace key is pressed, remove one character.
	if repeatingKeyPressed(ebiten.KeyBackspace) {
		if len(i.text) >= 1 {
			i.text = i.text[:len(i.text)-1]
		}
	}

	i.counter++
	return nil
}

func (i *Input) Draw(screen *ebiten.Image) {
	// Blink the cursor.
	t := i.text
	if i.counter%60 < 30 {
		t += "_"
	}

	text.Draw(screen, i.text, basicfont.Face7x13, 20, 60, color.White)
}

// repeatingKeyPressed return true when key is pressed considering the repeat state.
func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}
