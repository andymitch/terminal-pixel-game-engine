package tpx

import (
	"time"

	tcell "github.com/gdamore/tcell/v2"
)

type Scene struct {
	backgroundColor Pixel
	background      Sprite
	objects         []*Object
	screen          tcell.Screen
	FPS             int
}

func (s *Scene) SetBackgroundColor(c *Pixel) {
	s.backgroundColor = *c
}

func (s *Scene) SetBackground(sprite *Sprite) {
	s.background = *sprite
}

func (s *Scene) RemoveBackground() {
	s.background = Sprite{}
}

func (s *Scene) AddObject(o *Object) {
	s.objects = append(s.objects, o)
}

func (s *Scene) RemoveObject(o *Object) {
	for i, obj := range s.objects {
		if obj == o {
			s.objects = append(s.objects[:i], s.objects[i+1:]...)
			return
		}
	}
}

func overlayCell(c1, c2 Pixel) Pixel {
	a1 := int(c1.A * 255)
	a2 := int(c2.A * 255)
	a := a1 + (a2 * (255 - a1) / 255)
	r := (c1.R*a1 + c2.R*a2*(255-a1)/255) / a
	g := (c1.G*a1 + c2.G*a2*(255-a1)/255) / a
	b := (c1.B*a1 + c2.B*a2*(255-a1)/255) / a
	return Pixel{r, g, b, float32(a / 255)}
}

func handleRenderUpdate(s *Scene, o *Object) {
	sprite := o.GetSprite()

	newFrame := sprite.AnimationSpeed > 0 && sprite.AnimationCountdown <= 0
	moved := o.DX > 0 || o.DY > 0

	if newFrame {
		o.NextFrame()
		sprite.AnimationCountdown = sprite.AnimationSpeed
	}
	if sprite.AnimationSpeed > 0 {
		sprite.AnimationCountdown--
	}
	if moved || newFrame {
		old_x0, old_y0, old_x1, old_y1 := o.X, o.Y, o.X+o.Width, o.Y+o.Height
		new_x0, new_y0, new_x1, new_y1 := old_x0+o.DX, old_y0+o.DY, old_x1+o.DX, old_y1+o.DY
		o.DX, o.DY = 0, 0
		if old_y0%2 == 1 {
			old_y0--
		}
		if old_y1%2 == 0 {
			old_y1++
		}

		old_patch := NewFrame(old_x1-old_x0, old_y1-old_y0)
		new_patch := NewFrame(new_x1-new_x0, new_y1-new_y0)
		sprite = o.GetSprite()

		overlappedObjects := []*Object{}
		for _, obj := range s.objects {
			if obj == o {
				continue
			}
			if obj.X < old_x1 || obj.X+obj.Width >= old_x0 || obj.Y < old_y1 || obj.Y+obj.Height > old_y0 {
				overlappedObjects = append(overlappedObjects, obj)
			}
		}

		// patch background
		for x := old_x0; x < old_x1; x++ {
			for y := old_y0; y < old_y1; y++ {
				old_patch[x-old_x0][y-old_y0] = s.background.Animation[0][x][y]
			}
		}

		// patch overlapped objects
		for _, obj := range overlappedObjects {
			frame := obj.GetSprite().Animation[obj.GetSprite().CurrentFrame]
			for x := old_x0; x < old_x1; x++ {
				for y := old_y0; y < old_y1; y++ {
					old_patch[x-old_x0][y-old_y0] = overlayCell(old_patch[x-old_x0][y-old_y0], frame[x-obj.X][y-obj.Y])
				}
			}
		}

		// patch object
		for x := new_x0; x < new_x1; x++ {
			for y := new_y0; y < new_y1; y++ {
				new_patch[x-new_x0][y-new_y0] = sprite.Animation[sprite.CurrentFrame][x-o.X][y-o.Y]
			}
		}

		// TODO: overlap new_patch onto old_patch

		// TODO: render patches on screen
	}
}

func (s *Scene) start() {
	s.screen.Clear()
	// TODO: render background
	for _, obj := range s.objects {
		obj.Start(s)
		// TODO: render obj
	}
	s.screen.Show()
}

func (s *Scene) update() {
	for range time.Tick(time.Second / time.Duration(s.FPS)) {
		for _, obj := range s.objects {
			go obj.Update(s) // probably need to use channels
			go handleRenderUpdate(s, obj)
		}
		s.screen.Show()
	}
}

func (s *Scene) handleCommands() {
	for range time.Tick(time.Second / time.Duration(s.FPS)) {
		ev := s.screen.PollEvent()
		if ev == nil {
			break
		}
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				s.screen.Fini()
				return
			}
		}
	}
}

func (s *Scene) Run() {
	s.start()
	go s.handleCommands()
	go s.update()
}
