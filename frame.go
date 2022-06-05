package tpx

import (
	"image"
	"io"
)

type Frame [][]Pixel // [row, x, width] [column, y, height]

func NewFrame(width int, height int) Frame {
	frame := make(Frame, width)
	for x := 0; x < width; x++ {
		frame[x] = make([]Pixel, height)
	}
	return frame
}

func NewFrameFromImage(file io.Reader) (Frame, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	frame := Frame{}
	for x := 0; x < height; x++ {
		frame = append(frame, []Pixel{})
		for y := 0; y < width; y++ {
			frame[x] = append(frame[x], NewPixelFromColor(img.At(x, y).RGBA()))
		}
	}
	return frame, nil
}
