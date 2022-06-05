package tpx

import (
	"fmt"
	"io"
)

type Sprite struct {
	Animation          []Frame
	AnimationSpeed     int
	AnimationCountdown int
	Width              int
	Height             int
	FrameCount         int
	CurrentFrame       int
}

func NewSpriteFromFrames(frames []Frame, animationSpeed int) (Sprite, error) {
	if len(frames) == 0 {
		return Sprite{}, fmt.Errorf("no frames provided")
	}
	return Sprite{
		Animation:      frames,
		AnimationSpeed: animationSpeed,
		Width:          len(frames[0][0]),
		Height:         len(frames[0]),
		FrameCount:     len(frames),
		CurrentFrame:   0,
	}, nil
}

func NewSpriteFromFrame(frame Frame) (Sprite, error) {
	return NewSpriteFromFrames([]Frame{frame}, 0)
}

func NewSpriteFromImages(files []io.Reader, animationSpeed int) (Sprite, error) {
	frames := []Frame{}
	for _, file := range files {
		frame, err := NewFrameFromImage(file)
		if err != nil {
			return Sprite{}, err
		}
		frames = append(frames, frame)
	}
	if sprite, err := NewSpriteFromFrames(frames, animationSpeed); err != nil {
		return Sprite{}, err
	} else {
		return sprite, nil
	}
}

func NewSpriteFromImage(file io.Reader) (Sprite, error) {
	return NewSpriteFromImages([]io.Reader{file}, 0)
}
