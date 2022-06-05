package tpx

import (
	"fmt"

	tcell "github.com/gdamore/tcell/v2"
)

func GetScreen() (tcell.Screen, error) {
	return tcell.NewScreen()
}

type Position struct {
	x, y int
}

//type Sprite [][]*Color
type RenderSprite [][]*Pixel

func (c *Color) Pixel() *Pixel {
	if c == nil {
		return nil
	}
	p := Pixel(fmt.Sprintf("%d,%d,%d", c.r, c.g, c.b))
	return &p
}

func (p *Pixel) Color() *Color {
	if p == nil {
		return nil
	}
	var c Color
	fmt.Sscanf(string(*p), "%d,%d,%d", &c.r, &c.g, &c.b)
	return &c
}

type Asset struct {
	Sprite   Sprite
	Width    int
	Height   int
	Position Position
}

func (a *Asset) Render() RenderSprite {
	rs := RenderSprite{}
	for _, row := range a.Sprite {
		rs = append(rs, []*Pixel{})
		for _, col := range row {
			rs[len(rs)-1] = append(rs[len(rs)-1], col.Pixel())
		}
	}
	return rs
}

type Screen struct {
	screen     tcell.Screen
	Assets     []*Asset
	background *Asset
	width      int
	height     int
}

func (s *Screen) Render() {

}

func (s *Screen) SetBackground(sp *Sprite) {
	if sp == nil {
		return
	}
	s.screen.Clear()
}
