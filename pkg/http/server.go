package http

import (
	"context"
	"net/http"
	"time"

	"github.com/spf13/viper"
	"github.com/svbnbyrk/nba/config"
)

// Server
type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// New
func New(handler http.Handler) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  viper.GetDuration(config.HTTP_SERVER_TIMEOUT_READ),
		WriteTimeout: viper.GetDuration(config.HTTP_SERVER_TIMEOUT_WRITE),
		Addr:         viper.GetString(config.HTTP_SERVER_PORT),
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: viper.GetDuration(config.HTTP_SERVER_TIMEOUT_SHUTDOWN),
	}

	s.start()

	return s
}

func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
