package main

import (
	"fmt"
	"net/http"
)

func main() {
	var website string
	fmt.Println("Enter the website name you want to check the status on")
	fmt.Scanln(&website)

	c := make(chan string) 

	go checkLink(website, c)

	fmt.Println(<-c)

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- link + " is down"
	} else {
		c <- link + " is up"
	}
}
