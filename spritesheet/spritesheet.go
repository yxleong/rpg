package spritesheet

import "image"

type Spritesheet struct {
	WidthInTiles  int
	HeightInTiles int
	Tizesize      int
}

func (s *Spritesheet) Rect(index int) image.Rectangle {
	x := (index % s.WidthInTiles) * s.Tizesize
	y := (index / s.WidthInTiles) * s.Tizesize
	return image.Rect(x, y, x+s.Tizesize, y+s.Tizesize)
}

func NewSpriteSheet(w, h, t int) *Spritesheet {
	return &Spritesheet{
		WidthInTiles:  w,
		HeightInTiles: h,
		Tizesize:      t,
	}
}
