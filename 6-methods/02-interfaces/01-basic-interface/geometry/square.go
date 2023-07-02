package geometry

type Square struct {
	size float64
}

func NewSquare(size float64) *Square {

	return &Square{size}
}

func (s Square) Area() float64 {
	return s.size * s.size
}

func (s Square) Perimeter() float64 {
	return s.size * 4
}
