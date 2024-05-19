# URL Shortener Service

This repository contains a URL Shortener service built with Golang. The service provides APIs to shorten URLs, redirect shortened URLs to the original URLs, and fetch the top 3 most shortened domains. The application is also Dockerized for easy deployment.

## Features

1. **URL Shortening API**: Accepts a URL as an argument and returns a shortened URL.
2. **Redirection API**: Redirects the shortened URL to the original URL.
3. **Metrics API**: Returns the top 3 domains that have been shortened the most.

## Getting Started

### Prerequisites

- Golang
- Docker

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/priyanshi2203/URL-Shortener.git
    cd url-shortener
    git checkout feature
    ```

2. Build and run the application using Docker:
    ```sh
    docker pull priyanshi2203/url
    docker run -p 8080:8080 priyanshi2203/url
    ```

    This will start the server on port 8080.

### APIs

#### 1. Shorten URL

- **Method**: POST
- **Endpoint**: `/shortly`
- **Parameters**: `url` (query parameter)

    Example:
    ```sh
    curl -X POST "http://localhost:8080/shortly?url=http://browserstack.com/swedcfrvgbthnjymkumyjnhgtsdcbfvdcdefrtgcvdxcfvcfvgdcfvgbcdfgcdfvgfcvg"
    ```

#### 2. Redirect to Original URL

- **Method**: GET
- **Endpoint**: `/shortgo/{code}`
- **Parameters**: `code` (path parameter, the 5-digit code returned by the shorten URL API)

    Example:
    ```sh
    curl -L "http://localhost:8080/shortgo/77516"
    ```

#### 3. Fetch Top 3 Domains

- **Method**: GET
- **Endpoint**: `/metrics`

    Example:
    ```sh
    curl "http://localhost:8080/metrics"
    ```

### Stopping the Server

To stop the server, press `Cmd+C` in the terminal where the server is running.

## Code Structure

The application is organized into the following packages:

- `cmd/main.go`: Entry point of the application.
- `internal/`: Contains the HTTP handlers for the APIs, common.go file containg the types and helpers.go .
- `server/`: Contains the URL Shortener server .

## Unit test

Run unit tests
```sh
cd internal
go test -run .

## Docker

The application is Dockerized and available on Docker Hub.

Docker Hub Repository: [priyanshi2203/url](https://hub.docker.com/repository/docker/priyanshi2203/url)

To pull the latest image:

```sh
docker pull priyanshi2203/url
