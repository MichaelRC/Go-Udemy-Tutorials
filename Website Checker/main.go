package main

import (
	"fmt"
	"net/http"
	"time"
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

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//l is short for link
	for l := range c {
		go func() {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}()
	}
}

//checks if a link is receiving traffic
func checkLink(link string, c chan string) {
	//we only care about the error returned
	//so the returned struct is left blank.
	_, err := http.Get(link)

	//if something went wrong
	if err != nil {
		fmt.Println(link, "might be down!")
		//send link adress in to channel
		c <- link
		return
	}
	//if everything is working.
	fmt.Println(link, "is up!")
	//send link adress in to channel
	c <- link
}
