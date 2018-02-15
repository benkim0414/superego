package service

import (
	"cloud.google.com/go/datastore"
	"github.com/benkim0414/superego/pkg/profile"
	"github.com/go-kit/kit/log"
)

type Service interface {
	profile.Service
}

func New(client *datastore.Client, logger log.Logger) Service {
	var svc Service
	svc = &service{
		profile.NewService(client),
	}
	svc = NewLoggingMiddleware(logger)(svc)
	return svc
}

type service struct {
	profile.Service
}
