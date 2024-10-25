package main

import (
	"fmt"
	"log"
	"net/http"
)

func getLinks(url string) ([]string, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Failed to get URL: %s", url)
	}

	var links []string
	return links, nil
}

func main() {
	url := "Testsite"

	links, err := getLinks(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Links Found: ")
	for _, link := range links {
		fmt.Print(link)
	}
}
