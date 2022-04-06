package main

type shape interface {
	getArea() float64
}

type sqaure struct{
	sideLength int
}
type triangle struct{
	height int
	base int
}

func main() {

}

func (t triangle) getArea() float64 {
	a := 0.5 * int float(t.base) * int float(t.height)
	return float64(a) 
}

func (s sqaure) getArea() float64 {
	a := s.sideLength * s.sideLength
	return float64(a)
}