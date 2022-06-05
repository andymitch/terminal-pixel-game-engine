package tpx

type Pixel struct {
	R, G, B int
	A       float32
}

func NewPixelFromColor(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), float32(a) / float32(257*255)}
}
