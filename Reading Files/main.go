package main

import (
	"fmt"
	"io"
	"os"
)

/* My Code
func main() {
	fn, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(fn))
}
*/

//Error checking function
func check(e error) {
	if e != nil {
		fmt.Println("Error:", e)
		panic(e)
	}
}

func main() {
	//grabs the second element (in the first index)
	//that is input from the terminal command.
	//Opens the file called, and places it in
	//veriable f. if error, returned in err.
	f, err := os.Open(os.Args[1])
	check(err)

	//takes the files and prints it to the command line.
	//**os.Stdout puts it on the command line.
	//f is the variable holding the file.
	io.Copy(os.Stdout, f)
}
