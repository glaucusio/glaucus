package command

import "net/http"

type Service struct {
	*App
}

func (s *Service) Serve(h http.Handler) error {
	return http.ListenAndServe("0.0.0.0:8080", h)
}
