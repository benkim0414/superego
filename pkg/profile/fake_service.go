package profile

import (
	"context"
	"errors"
	"sync"
)

var (
	ErrNoSuchEntity = errors.New("no such entity")
)

// fakeService is a simple fake service for testing.
type fakeService struct {
	mu       sync.RWMutex
	profiles map[string]*Profile
}

var FakeService = &fakeService{profiles: map[string]*Profile{}}

func (f *fakeService) PostProfile(_ context.Context, p *Profile) (*Profile, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.profiles[p.ID] = p
	return p, nil
}

func (f *fakeService) GetProfile(_ context.Context, id string) (*Profile, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()

	p, ok := f.profiles[id]
	if !ok {
		return &Profile{}, ErrNoSuchEntity
	}
	return p, nil
}

func (f *fakeService) PutProfile(_ context.Context, id string, p *Profile) (*Profile, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.profiles[p.ID] = p
	return p, nil
}

func (f *fakeService) PatchProfile(_ context.Context, id string, p *Profile) (*Profile, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	existing, ok := f.profiles[id]
	if !ok {
		return &Profile{}, ErrNoSuchEntity
	}

	if p.DisplayName != "" {
		existing.DisplayName = p.DisplayName
	}
	if p.Email != "" {
		existing.Email = p.Email
	}
	if p.ImageURL != "" {
		existing.ImageURL = p.ImageURL
	}
	if p.AboutMe != "" {
		existing.AboutMe = p.AboutMe
	}

	f.profiles[id] = existing
	return p, nil
}

func (f *fakeService) DeleteProfile(_ context.Context, id string) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	if _, ok := f.profiles[id]; !ok {
		return ErrNoSuchEntity
	}
	delete(f.profiles, id)
	return nil
}
