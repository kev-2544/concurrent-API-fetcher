package main

import (
	"fmt"
	"net/http"
	"time"
	"io"
	"context"

	
)

//endpoint struct rep the endpoint to be fetched
type Endpoint struct {
	URL  string
	Name string
}

//results strtuct rep the result of the fetch
type results struct {
	Name string
	URL  string
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

//Fetch makes a request to the endpoint and returns the result
//returns response with data or an error if any
func(f *Fetcher) Fetch(ctx context.Context, endpoint Endpoint) Result {

	//start the timer to measure the duration of the request
	start := time.Now()

	//prepare a result with endpoint info
	result := Result{
		Name: endpoint.Name,
		URL: endpoint.URL,
	}

	//make the request to the endpoint
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		endpoint.URL,
		nil,
		
	)
	if err != nil {
		result.Err = err
		return result
	}
	//execute the request
	resp, err := f.client.Do(req)

	if err != nil {
	result.Err = err
	result.Duration = time.Since(start)
	return result
    }

	//close the response body when done
	defer resp.Body.Close()

	//check the returned api status code
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		result.Err = fmt.Errorf("unexpected HTTP status %s", resp.Status)
		result.Duration = time.Since(start)
		return result
	}

	//read the response body into memory
	data, err := io.ReadAll(resp.Body)

	if err != nil {
		result.Err = err
		result.Duration = time.Since(start)
		return result

	}

	//store the data and duration in the result
	result.Data = data
	result.Duration = time.Since(start)
	return result
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
