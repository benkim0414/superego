package service

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
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

// LoggingMiddleware takes a logger as a dependency and returns a ServiceMiddleware.
type LoggingMiddleware struct {
	Logger log.Logger
	Next   Service
}

// InstrumentingMiddleware returns a service middleware that record statistics
// about service's runtime behavior.
func NewInstrumentingMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram) Middleware {
	return func(next Service) Service {
		return &InstrumentingMiddleware{
			RequestCount:   requestCount,
			RequestLatency: requestLatency,
			Next:           next,
		}
	}
}

// InstrumentingMiddleware returns a service middleware that instruments
// the number of requests received and total duration of requests.
type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	Next           Service
}
