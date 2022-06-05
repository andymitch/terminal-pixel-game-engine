package tpx

type State string

type Object struct {
	States        []State
	CurrentState  State
	Sprites       map[State]*Sprite
	Start         func(scene *Scene)
	Update        func(scene *Scene)
	DX, DY        int
	X, Y          int
	Width, Height int
}

func (o *Object) AddState(state State, sprite *Sprite) {
	if o.Width < sprite.Width {
		o.Width = sprite.Width
	}
	if o.Height < sprite.Height {
		o.Height = sprite.Height
	}
	o.States = append(o.States, state)
	o.Sprites[state] = sprite
}

func (o *Object) AddStates(states []State, sprites []*Sprite) {
	for i, state := range states {
		o.AddState(state, sprites[i])
	}
}

func (o *Object) RemoveState(state State) {
	for i, s := range o.States {
		if s == state {
			o.States = append(o.States[:i], o.States[i+1:]...)
			delete(o.Sprites, state)
			return
		}
	}
}

func (o *Object) GetSprite() *Sprite {
	return o.Sprites[o.CurrentState]
}

func (o *Object) NextFrame() {
	sprite := o.Sprites[o.CurrentState]
	sprite.CurrentFrame++
	if sprite.CurrentFrame >= sprite.FrameCount {
		sprite.CurrentFrame = 0
	}
}
