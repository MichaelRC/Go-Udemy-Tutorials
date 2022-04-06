package main

import "fmt"

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
	tri := triangle{5, 6}
	squ := sqaure{8}

	printArea(tri)
	printArea(squ)
}

/*
	STRUCT FUNCTIONS
*/
func (t triangle) getArea() float64 {
	a := 0.5 * t.base * t.height
	return a
}

func (s sqaure) getArea() float64 {
	a := s.sideLength * s.sideLength
	return a
}

/*
	INTERFACE FUNCTIONS
*/

func printArea(s shape) {
	fmt.Println(s.getArea())
}
