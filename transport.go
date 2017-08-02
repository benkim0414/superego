package superego

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

// MakeHTTPHandler mounts all of the service endpoints into an http.Handler.
func MakeHTTPHandler(s Service, logger log.Logger) http.Handler {
	r := mux.NewRouter().PathPrefix("/api/v1/").Subrouter()
	e := MakeServerEndpoints(s)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	// POST		/api/v1/profiles/		adds another profile
	// GET		/api/v1/profiles/:id	retrieves the given profile by id
	// PUT		/api/v1/profiles/:id	post updated profile information about the profile
	// PATCH	/api/v1/profiles/:id	partial updated profile information
	// DELETE	/api/v1/profiles/:id	removes the given profile

	r.Methods("POST").Path("/profiles/").Handler(httptransport.NewServer(
		e.PostProfileEndpoint,
		decodePostProfileRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/profiles/{id}").Handler(httptransport.NewServer(
		e.GetProfileEndpoint,
		decodeGetProfileRequest,
		encodeResponse,
		options...,
	))
	r.Methods("PUT").Path("/profiles/{id}").Handler(httptransport.NewServer(
		e.PutProfileEndpoint,
		decodePutProfileRequest,
		encodeResponse,
		options...,
	))
	r.Methods("PATCH").Path("/profiles/{id}").Handler(httptransport.NewServer(
		e.PatchProfileEndpoint,
		decodePatchProfileRequest,
		encodeResponse,
		options...,
	))
	r.Methods("DELETE").Path("/profiles/{id}").Handler(httptransport.NewServer(
		e.DeleteProfileEndpoint,
		decodeDeleteProfileRequest,
		encodeResponse,
		options...,
	))
	return r
}

func decodePostProfileRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req postProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Profile); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetProfileRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return getProfileRequest{ID: id}, nil
}

func decodePutProfileRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	profile := &Profile{}
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		return nil, err
	}
	return putProfileRequest{ID: id, Profile: profile}, nil
}

func decodePatchProfileRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	profile := &Profile{}
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		return nil, err
	}
	return patchProfileRequest{ID: id, Profile: profile}, nil
}

func decodeDeleteProfileRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return deleteProfileRequest{ID: id}, nil
}

// errorer is implemented by all concrete response types that may contain
// errors. It allows us to change the HTTP response code without needing to
// trigger an endpoint (transport-level) error.
type errorer interface {
	error() error
}

// encodeResponse is the common method to encode all response types to the
// client.
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
