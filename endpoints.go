package superego

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints collects all of the endpoints that compose a profile service.
// It's meant to be used as a helper struct, to collect all of the endpoints
// into a single parameter.
type Endpoints struct {
	PostProfileEndpoint   endpoint.Endpoint
	GetProfileEndpoint    endpoint.Endpoint
	PutProfileEndpoint    endpoint.Endpoint
	PatchProfileEndpoint  endpoint.Endpoint
	DeleteProfileEndpoint endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct where each endpoint
// invokes the corresponding method on the provided service.
func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		PostProfileEndpoint:   MakePostProfileEndpoint(s),
		GetProfileEndpoint:    MakeGetProfileEndpoint(s),
		PutProfileEndpoint:    MakePutProfileEndpoint(s),
		PatchProfileEndpoint:  MakePatchProfileEndpoint(s),
		DeleteProfileEndpoint: MakeDeleteProfileEndpoint(s),
	}
}

// MakePostProfileEndpoint returns an endpoint via the passed service.
func MakePostProfileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postProfileRequest)
		p, e := s.PostProfile(ctx, req.Profile)
		return postProfileResponse{Profile: p, Err: e}, nil
	}
}

// MakeGetProfileEndpoint returns an endpoint via the passed service.
func MakeGetProfileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getProfileRequest)
		p, e := s.GetProfile(ctx, req.ID)
		return getProfileResponse{Profile: p, Err: e}, nil
	}
}

// MakePutProfileEndpoint returns an endpoint via the passed service.
func MakePutProfileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(putProfileRequest)
		p, e := s.PutProfile(ctx, req.ID, req.Profile)
		return putProfileResponse{Profile: p, Err: e}, nil
	}
}

// MakePatchProfileEndpoint returns an endpoint via the passed service.
func MakePatchProfileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(patchProfileRequest)
		p, e := s.PatchProfile(ctx, req.ID, req.Profile)
		return patchProfileResponse{Profile: p, Err: e}, nil
	}
}

// MakeDeleteProfileEndpoint returns an endpoint via the passed service.
func MakeDeleteProfileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteProfileRequest)
		e := s.DeleteProfile(ctx, req.ID)
		return deleteProfileResponse{Err: e}, nil
	}
}

type postProfileRequest struct {
	Profile *Profile `json:"profile"`
}

type postProfileResponse struct {
	Profile *Profile `json:"profile"`
	Err     error    `json:"err,omitempty"`
}

func (r postProfileResponse) error() error { return r.Err }

type getProfileRequest struct {
	ID string `json:"id"`
}

type getProfileResponse struct {
	Profile *Profile `json:"profile,omitempty"`
	Err     error    `json:"err,omitempty"`
}

func (r getProfileResponse) error() error { return r.Err }

type putProfileRequest struct {
	ID      string   `json:"id"`
	Profile *Profile `json:"profile"`
}

type putProfileResponse struct {
	Profile *Profile `json:"profile,omitempty"`
	Err     error    `json:"err,omitempty"`
}

func (r putProfileResponse) error() error { return r.Err }

type patchProfileRequest struct {
	ID      string   `json:"id"`
	Profile *Profile `json:"profile"`
}

type patchProfileResponse struct {
	Profile *Profile `json:"profile,omitempty"`
	Err     error    `json:"err,omitempty"`
}

func (r patchProfileResponse) error() error { return r.Err }

type deleteProfileRequest struct {
	ID string
}

type deleteProfileResponse struct {
	Err error `json:"err,omitempty"`
}

func (r deleteProfileResponse) error() error { return r.Err }
