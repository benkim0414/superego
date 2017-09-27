package endpoint

import (
	"context"

	"github.com/benkim0414/superego/pkg/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
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

// New returns an Endpoints struct where each endpoint
// invokes the corresponding method on the provided service.
func New(s service.Service, logger log.Logger) Endpoints {
	var postProfileEndpoint endpoint.Endpoint
	postProfileEndpoint = MakePostProfileEndpoint(s)
	postProfileEndpoint = LoggingMiddleware(log.With(logger, "method", "PostProfile"))(postProfileEndpoint)

	var getProfileEndpoint endpoint.Endpoint
	getProfileEndpoint = MakeGetProfileEndpoint(s)
	getProfileEndpoint = LoggingMiddleware(log.With(logger, "method", "GetProfile"))(getProfileEndpoint)

	var putProfileEndpoint endpoint.Endpoint
	putProfileEndpoint = MakePutProfileEndpoint(s)
	putProfileEndpoint = LoggingMiddleware(log.With(logger, "method", "PutProfile"))(putProfileEndpoint)

	var patchProfileEndpoint endpoint.Endpoint
	patchProfileEndpoint = MakePatchProfileEndpoint(s)
	patchProfileEndpoint = LoggingMiddleware(log.With(logger, "method", "PatchProfile"))(patchProfileEndpoint)

	var deleteProfileEndpoint endpoint.Endpoint
	deleteProfileEndpoint = MakeDeleteProfileEndpoint(s)
	deleteProfileEndpoint = LoggingMiddleware(log.With(logger, "method", "DeleteProfile"))(deleteProfileEndpoint)

	return Endpoints{
		PostProfileEndpoint:   postProfileEndpoint,
		GetProfileEndpoint:    getProfileEndpoint,
		PutProfileEndpoint:    putProfileEndpoint,
		PatchProfileEndpoint:  patchProfileEndpoint,
		DeleteProfileEndpoint: deleteProfileEndpoint,
	}
}

// MakePostProfileEndpoint returns an endpoint via the passed service.
func MakePostProfileEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(PostProfileRequest)
		p, e := s.PostProfile(ctx, req.Profile)
		return PostProfileResponse{Profile: p, Err: e}, nil
	}
}

// MakeGetProfileEndpoint returns an endpoint via the passed service.
func MakeGetProfileEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetProfileRequest)
		p, e := s.GetProfile(ctx, req.ID)
		return GetProfileResponse{Profile: p, Err: e}, nil
	}
}

// MakePutProfileEndpoint returns an endpoint via the passed service.
func MakePutProfileEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(PutProfileRequest)
		p, e := s.PutProfile(ctx, req.ID, req.Profile)
		return PutProfileResponse{Profile: p, Err: e}, nil
	}
}

// MakePatchProfileEndpoint returns an endpoint via the passed service.
func MakePatchProfileEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(PatchProfileRequest)
		p, e := s.PatchProfile(ctx, req.ID, req.Profile)
		return PatchProfileResponse{Profile: p, Err: e}, nil
	}
}

// MakeDeleteProfileEndpoint returns an endpoint via the passed service.
func MakeDeleteProfileEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteProfileRequest)
		e := s.DeleteProfile(ctx, req.ID)
		return DeleteProfileResponse{Err: e}, nil
	}
}

type PostProfileRequest struct {
	Profile *service.Profile `json:"profile"`
}

type PostProfileResponse struct {
	Profile *service.Profile `json:"profile"`
	Err     error            `json:"err,omitempty"`
}

func (r PostProfileResponse) error() error { return r.Err }

type GetProfileRequest struct {
	ID string `json:"id"`
}

type GetProfileResponse struct {
	Profile *service.Profile `json:"profile,omitempty"`
	Err     error            `json:"err,omitempty"`
}

func (r GetProfileResponse) error() error { return r.Err }

type PutProfileRequest struct {
	ID      string           `json:"id"`
	Profile *service.Profile `json:"profile"`
}

type PutProfileResponse struct {
	Profile *service.Profile `json:"profile,omitempty"`
	Err     error            `json:"err,omitempty"`
}

func (r PutProfileResponse) error() error { return r.Err }

type PatchProfileRequest struct {
	ID      string           `json:"id"`
	Profile *service.Profile `json:"profile"`
}

type PatchProfileResponse struct {
	Profile *service.Profile `json:"profile,omitempty"`
	Err     error            `json:"err,omitempty"`
}

func (r PatchProfileResponse) error() error { return r.Err }

type DeleteProfileRequest struct {
	ID string
}

type DeleteProfileResponse struct {
	Err error `json:"err,omitempty"`
}

func (r DeleteProfileResponse) error() error { return r.Err }
