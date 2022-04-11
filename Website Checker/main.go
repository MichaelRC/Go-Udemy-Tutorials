package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		//make sure to include web protocol!
		//Go expects to see it and can use it to
		//differentiate between http and https
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		go checkLink(link)
	}

}

//checks if a link is receiving traffic
func checkLink(link string) {
	//we only care about the error returned
	//so the returned struct is left blank.
	_, err := http.Get(link)

	//if something went wrong
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	//if everything is working.
	fmt.Println(link, "is up!")
}
