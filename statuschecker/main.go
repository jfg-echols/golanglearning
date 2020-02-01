package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkSite(site string, c chan string) {
	time.Sleep(5 * time.Second)
	_, err := http.Get(site)
	if err != nil {
		fmt.Println(site, "might be down")
		c <- site
		return
	}
	fmt.Println(site, "is up")
	c <- site

}

func main() {
	sites := []string{
		"http://google.com",
		"http://golang.org",
		"http://hjtyrguy6gf.org",
		"http://reddit.com",
	}
	//channels are used to communicate between subroutines
	channel := make(chan string)
	for _, site := range sites {
		//go keyword runs this function in a new go routine
		//go scheduler uses 1 cpu by default - then runs 1 per core of a cpu

		go checkSite(site, channel)
	}
	// for i := 0; i < len(sites); i++ {
	for s := range channel {
		go func(site string) {
			time.Sleep(5 * time.Second)
			//s is defined outside of the function literal, so s keeps looking to the most recent s
			checkSite(site, channel)
		}(s)
	}
	// fmt.Println()
}
