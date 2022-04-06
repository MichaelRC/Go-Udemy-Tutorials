package main

import "fmt"

func main() {
	//variable num is a int slice with elements 0-10
	num := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//iterates through num to check each
	//if each element is even or odd
	//n%2 if 0 is Even, if not is Odd.
	for _, n := range num {
		if n%2 == 0 {
			fmt.Println(n, "is Even")
		} else {
			fmt.Println(n, "is Odd")
		}
	}
}
