package profile

import (
	"context"

	"cloud.google.com/go/datastore"
)

// Service is a simple CRUD interface for user profiles.
type Service interface {
	PostProfile(ctx context.Context, p *Profile) (*Profile, error)
	GetProfile(ctx context.Context, id string) (*Profile, error)
	PutProfile(ctx context.Context, id string, p *Profile) (*Profile, error)
	PatchProfile(ctx context.Context, id string, p *Profile) (*Profile, error)
	DeleteProfile(ctx context.Context, id string) error
}

// NewService returns a datastore service with all of the expected middlewares wired in.
func NewService(client *datastore.Client) Service {
	return newDatastoreService(client)
}
