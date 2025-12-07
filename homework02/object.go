package main

type Rectangle struct {
	Width  float64
	Height float64
}

func (R *Rectangle) Area() float64 {
	return R.Width * R.Height
}

func (R *Rectangle) Perimeter() float64 {
	return 2 * (R.Width + R.Height)
}

type Circle struct {
	Radius float64
}

func (C *Circle) Area() float64 {
	return 3.14 * C.Radius * C.Radius
}

func (C *Circle) Perimeter() float64 {
	return 2 * 3.14 * C.Radius
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}

	println("Rectangle Area:", rect.Area())
	println("Rectangle Perimeter:", rect.Perimeter())

	println("Circle Area:", circle.Area())
	println("Circle Perimeter:", circle.Perimeter())
}
