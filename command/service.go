package command

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

type Service struct {
	*App

	srv     *http.Server
	srvOnce sync.Once
}

func (s *Service) Serve(h http.Handler) error {
	l, err := s.listen("tcp", "0.0.0.0:8080")
	if err != nil {
		return fmt.Errorf("failure listening: %w", err)
	}

	srv := &http.Server{
		Handler:           h,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      30 * time.Second,
		BaseContext:       func(net.Listener) context.Context { return s.Context },
	}

	s.Group.Go(func() error { return s.serve(srv, l) })
	s.Group.Go(func() error { <-s.Context.Done(); return s.shutdown(srv) })

	return s.wait()
}

func (s *Service) listen(network, addr string) (net.Listener, error) {
	return new(net.ListenConfig).Listen(s.Context, network, addr)
}

func (s *Service) serve(srv *http.Server, l net.Listener) error {
	return srv.Serve(l)
}

func (s *Service) shutdown(srv *http.Server) error {
	ctx, cancel := WithSignals(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	return srv.Shutdown(ctx)
}

func (s *Service) wait() error {
	switch err := s.Group.Wait(); {
	case errors.Is(err, http.ErrServerClosed):
		return nil
	default:
		return err
	}
}
