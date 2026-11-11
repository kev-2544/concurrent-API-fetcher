package main

import (
	"fmt"
	"net/http"
	"time"
	
)

//endpoint struct rep the endpoint to be fetched
type Endpoint struct {
	URL  string
	Name string
}

//results strtuct rep the result of the fetch
type results struct {
	Name string
	url string
	Data []byte
	Err error
	Duration time.Duration
}

//fetcher is responsible for making the request to the endpoint and returning the result
type Fetcher struct {
	client *http.Client
		
	}



//NewFetcher creates and returns a new fetcher with a timeout
func NewFetcher(timeout time.Duration) *Fetcher {
	return &Fetcher{
		client: &http.Client{
			Timeout: timeout,
		},
	}

}

func main(){
	//define the api endpoints to be fetched
	endpoints := []Endpoint{
		{
			Name: "Users",
			URL:  "https://jsonplaceholder.typicode.com/users",
		},
		{
			Name: "Posts",
			URL:  "https://jsonplaceholder.typicode.com/posts",
		},
		{
			Name: "Comments",
			URL:  "https://jsonplaceholder.typicode.com/comments",
		},
	}

	//create our http fetcher with a timeout
	fetcher := NewFetcher(10 * time.Second)

	fmt.Println("Concurrent API fetcher started...")
	fmt.Println("----------------------------------------")

	fmt.Printf("Configured endpoints: %d\n", len(endpoints))
	fmt.Printf("HTTP client timeout: %s\n", fetcher.client.Timeout)
}
