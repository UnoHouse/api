package http

import (
	"net/http"
	"os"
	"time"
)

func NewServer(handler http.Handler) *http.Server {

	return &http.Server{
		Addr:         os.Getenv("LISTEN"),
		Handler:      handler,
		WriteTimeout: time.Duration(30) * time.Second,
		ReadTimeout:  time.Duration(30) * time.Second,
	}
}
