package server

import (
	"log"
	"net/http"
)

func New(cfg Config, handler http.Handler, logger *log.Logger) http.Server {
	return http.Server{
		Addr:         cfg.Address,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimout,
		Handler:      handler,
		ErrorLog:     logger,
	}
}
