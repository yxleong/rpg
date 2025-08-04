package spritesheet

import "image"

type SpriteSheet struct {
	WidthInTiles  int
	HeightInTiles int
	Tizesize      int
}

func (s *SpriteSheet) Rect(index int) image.Rectangle {
	x := (index % s.WidthInTiles) * s.Tizesize
	y := (index / s.WidthInTiles) * s.Tizesize
	return image.Rect(x, y, x+s.Tizesize, y+s.Tizesize)
}

func NewSpriteSheet(w, h, t int) *SpriteSheet {
	return &SpriteSheet{
		WidthInTiles:  w,
		HeightInTiles: h,
		Tizesize:      t,
	}
}
