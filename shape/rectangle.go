package shapes

type Rectangle struct {
	Width  float32
	Length float32
}

func (Rectangle *Rectangle) Area() float32 {
	return Rectangle.Width * Rectangle.Length
}

func (Rectangle *Rectangle) Perimeter() float32 {
	return 2 * (Rectangle.Width + Rectangle.Length)
}
