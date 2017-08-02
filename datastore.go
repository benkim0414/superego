package superego

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
)

const (
	// datastore entity kind for Profile
	profileKind = "Profile"
)

type datastoreService struct {
	client *datastore.Client
}

func NewDatastoreService(client *datastore.Client) Service {
	return &datastoreService{client: client}
}

func (s *datastoreService) PostProfile(ctx context.Context, p *Profile) (*Profile, error) {
	key := datastore.IncompleteKey(profileKind, nil)
	key, err := s.client.Put(ctx, key, p)
	if err != nil {
		return nil, fmt.Errorf("datastore: could not put Profile: %v", err)
	}
	p.ID = key.Encode()
	return p, nil
}

func (s *datastoreService) GetProfile(ctx context.Context, id string) (*Profile, error) {
	key, err := datastore.DecodeKey(id)
	if err != nil {
		return nil, fmt.Errorf("datastore: invalid Profile id: %v", err)
	}
	profile := &Profile{}
	err = s.client.Get(ctx, key, profile)
	if err == datastore.ErrNoSuchEntity {
		return nil, fmt.Errorf("datastore: could not get Profile: %v", err)
	}
	profile.ID = id
	return profile, nil
}

func (s *datastoreService) PutProfile(ctx context.Context, id string, p *Profile) (*Profile, error) {
	key, err := datastore.DecodeKey(id)
	if err != nil {
		return nil, fmt.Errorf("datastore: invalid Profile id: %v", err)
	}
	_, err = s.client.Put(ctx, key, p)
	if err != nil {
		return nil, fmt.Errorf("datastore: could not put Profile: %v", err)
	}
	return p, nil
}

func (s *datastoreService) PatchProfile(ctx context.Context, id string, p *Profile) (*Profile, error) {
	key, err := datastore.DecodeKey(id)
	if err != nil {
		return nil, fmt.Errorf("datastore: invalid Profile id: %v", err)
	}
	profile := &Profile{}
	err = s.client.Get(ctx, key, profile)
	if err == datastore.ErrNoSuchEntity {
		return nil, fmt.Errorf("datastore: could not get Profile: %v", err)
	}

	// assume that it's not possible to PATCH the ID, and that it's not
	// possible to PATCH any field to its zero value. That is, the zero value
	// means not specified.
	if p.DisplayName != "" {
		profile.DisplayName = p.DisplayName
	}
	if p.Email != "" {
		profile.Email = p.Email
	}
	if p.ImageURL != "" {
		profile.ImageURL = p.ImageURL
	}
	if p.AboutMe != "" {
		profile.AboutMe = p.AboutMe
	}

	_, err = s.client.Put(ctx, key, profile)
	if err != nil {
		return nil, fmt.Errorf("datastore: could not put Profile: %v", err)
	}
	return profile, nil
}

func (s *datastoreService) DeleteProfile(ctx context.Context, id string) error {
	key, err := datastore.DecodeKey(id)
	if err != nil {
		return fmt.Errorf("datastore: invalid Profile id: %v", err)
	}
	err = s.client.Delete(ctx, key)
	if err != nil {
		return fmt.Errorf("datastore: could not delete Profile: %v", err)
	}
	return nil
}
