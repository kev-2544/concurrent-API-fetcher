
package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestFetchSuccess checks whether Fetch correctly handles
// a successful HTTP response.
func TestFetchSuccess(t *testing.T) {

	// Create a mock HTTP server.
	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello from API"))
		},
	))
	defer server.Close()

	fetcher := NewFetcher(2 * time.Second)

	endpoint := Endpoint{
		Name: "Test API",
		URL:  server.URL,
	}

	result := fetcher.Fetch(context.Background(), endpoint)

	if result.Err != nil {
		t.Fatalf("expected no error, got %v", result.Err)
	}

	if string(result.Data) != "Hello from API" {
		t.Errorf("unexpected response: %s", result.Data)
	}
}

// TestFetchHTTPError checks whether Fetch detects
// unsuccessful HTTP status codes.
func TestFetchHTTPError(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		},
	))
	defer server.Close()

	fetcher := NewFetcher(2 * time.Second)

	endpoint := Endpoint{
		Name: "Failing API",
		URL:  server.URL,
	}

	result := fetcher.Fetch(context.Background(), endpoint)

	if result.Err == nil {
		t.Fatal("expected an HTTP error, got nil")
	}
}

// TestFetchContextCancellation checks whether an API request
// respects context cancellation.
func TestFetchContextCancellation(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Wait until the client cancels the request.
			<-r.Context().Done()
		},
	))
	defer server.Close()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		50*time.Millisecond,
	)
	defer cancel()

	fetcher := NewFetcher(2 * time.Second)

	endpoint := Endpoint{
		Name: "Slow API",
		URL:  server.URL,
	}

	result := fetcher.Fetch(ctx, endpoint)

	if result.Err == nil {
		t.Fatal("expected a context cancellation error")
	}
}
