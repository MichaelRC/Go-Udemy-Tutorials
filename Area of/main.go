package main

/*
	INTERFACE DECLERATIONS
*/
type shape interface {
	getArea() float64
}

/*
	STRUCT DECLERATIONS
*/
type sqaure struct {
	sideLength float64
}
type triangle struct {
	height float64
	base   float64
}

/*
	MAIN
*/
func main() {

}

/*
	OTHER FUNCTIONS
*/
func (t triangle) getArea() float64 {
	a := 0.5 * t.base * t.height
	return a
}

func (s sqaure) getArea() float64 {
	a := s.sideLength * s.sideLength
	return a
}
