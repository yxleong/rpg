package animations

type Animation struct {
	First        int
	Last         int
	Step         int
	SpeedInTips  float32
	frameCounter float32
	frame        int
}

func (a *Animation) Update() {
	a.frameCounter -= 1.0
	if a.frameCounter <= 0.0 {
		a.frameCounter = a.SpeedInTips
		a.frame += a.Step
		if a.frame > a.Last {
			a.frame = a.First
		} else if a.frame < a.First {
			a.frame = a.Last
		}
	}
}

func (a *Animation) Frame() int {
	return a.frame
}

func NewAnimation(first, last, step int, speed float32) *Animation {
	return &Animation{
		First:        first,
		Last:         last,
		Step:         step,
		SpeedInTips:  speed,
		frameCounter: speed,
		frame:        first,
	}
}
