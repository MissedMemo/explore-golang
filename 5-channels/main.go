package main

import (
	"fmt"
	"net/http"
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

	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "may be down!")
		c <- "could be down..."
		return
	}

	fmt.Println(link, "is up...")
	c <- "seems fine!"
}
