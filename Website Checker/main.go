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

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//Iterate through each element in links
	//and print out what is sent through the channel.
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
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
		//this this string into the channel.
		c <- "Might be down I think."
		return
	}
	//if everything is working.
	fmt.Println(link, "is up!")
	//send this string into the channel.
	c <- "Yep it's up."
}
