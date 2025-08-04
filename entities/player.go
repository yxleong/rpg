package entities

import "rpg-go/animations"

type PlayerState uint8

const (
	Down PlayerState = iota
	Up
	Left
	Right
)

type Player struct {
	*Sprite
	Health     uint
	Animations map[PlayerState]*animations.Animation
}

func (p *Player) ActiveAnimation(dx, dy int) *animations.Animation {
	if dx > 0 {
		return p.Animations[Right]
	} else if dx < 0 {
		return p.Animations[Left]
	} else if dy > 0 {
		return p.Animations[Down]
	} else if dy < 0 {
		return p.Animations[Up]
	}
	return nil
}
