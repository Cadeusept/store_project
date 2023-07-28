package handler

import (
	"context"
	"net/http"
	"time"
)

type HandlerServer struct {
	httpServer *http.Server
}

func (s *HandlerServer) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MByte
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *HandlerServer) Shutdown(ctx context.Context) {
	s.httpServer.Shutdown(ctx)
}
