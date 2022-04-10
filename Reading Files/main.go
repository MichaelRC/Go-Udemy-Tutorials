package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fn, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(fn))
}
