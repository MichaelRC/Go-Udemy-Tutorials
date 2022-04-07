package main

import (
	"fmt"
	"os"
)

func main() {
	//prints out what is put into command line
	//creates a temporary file to pull the data from
	fmt.Println(os.Args)

	//will name the file what is placed in the 1st element
	os.Args[1]
}
