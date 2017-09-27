package service

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/go-kit/kit/log"
)

// Service is a simple CRUD interface for user profiles.
type Service interface {
	PostProfile(ctx context.Context, p *Profile) (*Profile, error)
	GetProfile(ctx context.Context, id string) (*Profile, error)
	PutProfile(ctx context.Context, id string, p *Profile) (*Profile, error)
	PatchProfile(ctx context.Context, id string, p *Profile) (*Profile, error)
	DeleteProfile(ctx context.Context, id string) error
}

// New returns a datastore service with all of the expected middlewares wired in.
func New(client *datastore.Client, logger log.Logger) Service {
	var s Service
	s = newDatastoreService(client)
	s = LoggingMiddleware(logger)(s)
	return s
}
