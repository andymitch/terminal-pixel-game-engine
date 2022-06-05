package tpx

import "testing"

func TestGetScreen(t *testing.T) {
	if _, e := GetScreen(); e != nil {
		t.Error(e)
	}
}

func TestRender(t *testing.T) {
	a := Asset{
		Sprite: Sprite{
			{&Color{255, 0, 0}, &Color{255, 255, 0}, &Color{0, 255, 0}},
			{&Color{255, 255, 0}, nil, &Color{0, 255, 255}},
			{&Color{0, 255, 0}, &Color{0, 255, 255}, &Color{0, 0, 255}},
		},
		Width:    3,
		Height:   3,
		Position: Position{0, 0},
	}
	rs := a.Render()
	t.Log(rs)
	if len(rs) != 3 {
		t.Error("expected 3 rows")
	}
	if len(rs[0]) != 3 {
		t.Error("expected 3 columns got", len(rs[0]))
	}
	if *rs[0][0] != "255,0,0" {
		t.Error("expected 255,0,0 got", rs[0][0])
	}
}
