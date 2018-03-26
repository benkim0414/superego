package service

import (
	"cloud.google.com/go/datastore"
	"github.com/benkim0414/superego/pkg/profile"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
)

// Service describes a service that manage profiles.
type Service interface {
	profile.Service
}

// New returns a basic Service with all of the expected middlewares wired in.
func New(client *datastore.Client, logger log.Logger, requestCount metrics.Counter, requestLatency metrics.Histogram) Service {
	var svc Service
	svc = &service{
		profile.NewService(client),
	}
	svc = NewLoggingMiddleware(logger)(svc)
	svc = NewInstrumentingMiddleware(requestCount, requestLatency)(svc)
	return svc
}

type service struct {
	profile.Service
}
