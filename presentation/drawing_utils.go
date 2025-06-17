package presentation

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

func drawRect(screen *ebiten.Image, x, y, w, h int, c color.Color) {
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			screen.Set(x+i, y+j, c)
		}
	}
}

func drawBorder(screen *ebiten.Image, x, y, w, h int, c color.Color, thickness int) {
	// Top border
	drawRect(screen, x, y, w, thickness, c)
	// Bottom border
	drawRect(screen, x, y+h-thickness, w, thickness, c)
	// Left border
	drawRect(screen, x, y, thickness, h, c)
	// Right border
	drawRect(screen, x+w-thickness, y, thickness, h, c)
}