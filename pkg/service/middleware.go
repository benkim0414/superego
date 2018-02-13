package service

import (
	"github.com/go-kit/kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(Service) Service

// NewLoggingMiddleware takes a logger as a dependency
// and returns a ServiceMiddleware.
func NewLoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &LoggingMiddleware{logger, next}
	}
}

type LoggingMiddleware struct {
	Logger log.Logger
	Next   Service
}
