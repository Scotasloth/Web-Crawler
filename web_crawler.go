package main

import (
	"bufio"
	"fmt"      // Importing the fmt package for formatted I/O
	"io"       // Importing the io package for input/output utilities
	"log"      // Importing the log package for logging errors
	"net/http" // Importing the net/http package to make HTTP requests
	"os"
	"strings" // Importing the strings package for string manipulation

	"golang.org/x/net/html" // Importing the HTML parsing package
)

// getLinks fetches the URL and extracts all hyperlinks from the page
func getLinks(url string) ([]string, error) {
	// Make an HTTP GET request to the specified URL
	response, err := http.Get(url)
	if err != nil {
		return nil, err // Return nil and the error if the request fails
	}

	defer response.Body.Close() // Ensure the response body is closed after the function returns

	// Check if the response status code is OK (200)
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get URL: %s", url) // Return an error if not OK
	}

	// Read the entire response body into a byte slice
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err // Return nil and the error if reading the body fails
	}

	// Parse the HTML document from the byte slice
	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return nil, err // Return nil and the error if parsing fails
	}

	var links []string // Slice to hold extracted links

	// Define a recursive function to traverse the HTML nodes
	var f func(*html.Node)
	f = func(n *html.Node) {
		// Check if the node is an anchor element (<a>)
		if n.Type == html.ElementNode && n.Data == "a" {
			// Iterate through the attributes of the anchor element
			for _, attr := range n.Attr {
				if attr.Key == "href" { // Check for the "href" attribute
					links = append(links, attr.Val) // Add the link to the slice
				}
			}
		}

		// Recursively call f on each child node
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)            // Start the recursive function on the parsed HTML document
	return links, nil // Return the slice of links and nil for no errors
}

// main function - entry point of the program
func main() {
	url := "https://Google.co.uk" // Define the URL to fetch

	// Call getLinks and handle any potential errors
	links, err := getLinks(url)
	if err != nil {
		log.Fatal(err) // Log the error and exit if there was a problem
	}

	// Print the found links
	fmt.Print("Links Found: ")
	for _, link := range links {
		fmt.Print(link + " ") // Print each link, separated by a space
	}

	file, err := os.Create("Links.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, link := range links {
		_, err := writer.WriteString(link + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	writer.Flush()

	fmt.Print("All links have been written to links.txt")
}
