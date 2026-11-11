
# Concurrent API Fetcher

A Go project exploring concurrent HTTP requests, goroutines,
channels, context cancellation, and error handling.

## Overview

The application fetches data from multiple API endpoints
concurrently and collects their responses.

It uses JSONPlaceholder as its public test API.

## Features

- Concurrent HTTP requests using goroutines
- Fan-out / fan-in concurrency pattern
- Buffered channels for result communication
- sync.WaitGroup for worker coordination
- select for channel operations and cancellation
- Context-based request cancellation and deadlines
- HTTP status validation
- Independent error handling
- Request duration tracking
- Automated tests using httptest

## Technologies

- Go
- net/http
- context
- sync
- channels
- httptest

## API Endpoints

- https://jsonplaceholder.typicode.com/users
- https://jsonplaceholder.typicode.com/posts
- https://jsonplaceholder.typicode.com/comments

## Running the application

```bash
go run .
```

## Running tests

```bash
go test -v ./...
```

## Architecture

The application follows a fan-out/fan-in pattern.

1. Main goroutine initializes API endpoints.
2. Each endpoint is assigned to a separate goroutine.
3. Goroutines execute HTTP requests concurrently.
4. Results are sent through a shared channel.
5. A WaitGroup tracks worker completion.
6. A coordinator closes the results channel safely.
7. Main goroutine collects and displays results.

## Learning Objectives

This project explores Go concurrency primitives and
their application in backend development.
