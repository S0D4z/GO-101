package AreaCircle

type AreaCircle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * (c.Radius * c.Radius)
}