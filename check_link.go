package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	links := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%s link is down!\n", link)
		os.Exit(1)
	}
	fmt.Printf("%s link is up!\n", link)
}
