package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	
	"urlShortener/server"
)


func main() {
	srv := startUrlServer("8080")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	registerSignalHandler(c,srv);
}


func registerSignalHandler(c chan os.Signal, urlServer *server.UrlServer) {
	signal := <-c
	fmt.Printf("Caught signal %s, gracefully shutting down servers", signal)
	urlServer.Stop()
}

func startUrlServer(port string) *server.UrlServer {
	urlServer := server.New(port)
	go func() {
		fmt.Printf("Starting server on port: %s !", port)
		err := urlServer.Start()
		if err != nil {
			fmt.Print("Failed to start app server")
		}
	}()

	return urlServer
}