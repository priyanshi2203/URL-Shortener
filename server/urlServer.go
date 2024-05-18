package server

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"urlShortener/internal"
)

type UrlServer struct{
	ctx context.Context
	Port string
	server *http.Server
	serverLock            sync.RWMutex
}

// New creates a new instance of the server.
func New(port string) *UrlServer {
	return &UrlServer{
		Port:                  port,
	}
}

// Start starts the server and listens for incoming requests.
func (s *UrlServer) Start() error {
	ctx := context.Background();

	shortener := &internal.URLShortener{
		Urls: make(map[string]string),
		UrlHashes: make(map[string]string),
	}

	mux := http.NewServeMux()
	server := &http.Server{
        Addr:    fmt.Sprintf(":%s",s.Port),
        Handler: mux,
    }

	s.serverLock.Lock()
	s.server =server
	s.ctx=ctx
	s.serverLock.Unlock()


	mux.HandleFunc("/shortly", shortener.HandleShorten)
	mux.HandleFunc("/shortgo/", shortener.HandleRedirect)

	err := server.ListenAndServe()
	if(err!=nil){
		return err;
	}

	return nil
}

// Stop gracefully stops the GRPC server.
func (s *UrlServer) Stop() {
	s.serverLock.RLock()
	defer s.serverLock.RUnlock()
	if s.server == nil {
		return
	}

	s.server.Shutdown(s.ctx);
}