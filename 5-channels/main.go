package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://golang.okg",
		"http://yahoo.com",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	/*
		for {
			go checkLink( <-c, c )
		}
	*/

	// same as above, with addition of pause
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second) // pause 5 secs
			checkLink(link, c)
		}(l) // pass COPY of outer variable
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "could be down...")
		c <- link
		return
	}

	fmt.Println(link, "seems fine!")
	c <- link
}
