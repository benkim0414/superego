package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/benkim0414/superego/endpoint"
	"github.com/benkim0414/superego/service"
	"github.com/go-kit/kit/log"
)

var handlerTests = []struct {
	method  string
	path    string
	profile *service.Profile
}{
	{
		method:  http.MethodPost,
		path:    "/api/v1/profiles/",
		profile: &service.Profile{ID: "gunwoo", Email: "gunwoo@gunwoo.org"},
	},
	{
		method:  http.MethodGet,
		path:    "/api/v1/profiles/gunwoo",
		profile: &service.Profile{ID: "gunwoo", Email: "gunwoo@gunwoo.org"},
	},
	{
		method:  http.MethodPut,
		path:    "/api/v1/profiles/gunwoo",
		profile: &service.Profile{ID: "gunwoo", Email: "ben.kim@greenenergytrading.com.au"},
	},
	{
		method:  http.MethodPatch,
		path:    "/api/v1/profiles/gunwoo",
		profile: &service.Profile{ID: "gunwoo", Email: "gunwoo@gunwoo.org"},
	},
	{
		method:  http.MethodDelete,
		path:    "/api/v1/profiles/gunwoo",
		profile: nil,
	},
}

func TestNewHTTPHandler(t *testing.T) {
	logger := log.NewNopLogger()
	endpoints := endpoint.New(service.FakeService, logger)
	handler := NewHTTPHandler(endpoints, logger)

	ts := httptest.NewServer(handler)
	defer ts.Close()

	for _, tt := range handlerTests {
		var body bytes.Buffer
		if tt.method != http.MethodGet || tt.method != http.MethodDelete {
			err := json.NewEncoder(&body).Encode(tt.profile)
			if err != nil {
				t.Fatal(err)
			}
		}

		req := httptest.NewRequest(tt.method, ts.URL+tt.path, &body)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)

		resp := w.Result()
		var response struct {
			Profile *service.Profile `json:"profile"`
			Err     error            `json:"err,omitempty"`
		}
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(response.Profile, tt.profile) {
			t.Errorf("%s %s got %v, want %v", tt.method, tt.path, response.Profile, tt.profile)
		}
	}
}

func TestDecodePostProfileRequest(t *testing.T) {
	ctx := context.Background()
	var buf bytes.Buffer
	r := httptest.NewRequest(http.MethodPost, "/api/v1/profiles/", &buf)

	request, err := decodePostProfileRequest(ctx, r)
	if request != nil {
		t.Errorf("decodePostProfileRequest: got %v, want %v", request, nil)
	}

	p := &service.Profile{Email: "gunwoo@gunwoo.org"}
	err = json.NewEncoder(&buf).Encode(p)
	if err != nil {
		t.Fatal(err)
	}
	r = httptest.NewRequest(http.MethodPost, "/api/v1/profiles/", &buf)

	request, err = decodePostProfileRequest(ctx, r)
	got := request.(endpoint.PostProfileRequest)
	if !reflect.DeepEqual(got.Profile, p) {
		t.Errorf("decodePostProfileRequest: got %v, want %v", got.Profile, p)
	}
}

func TestDecodeGetProfileRequest(t *testing.T) {
	ctx := context.Background()
	r := httptest.NewRequest(http.MethodGet, "/api/v1/profiles/", nil)

	request, err := decodeGetProfileRequest(ctx, r)
	if request != nil {
		t.Errorf("decodeGetProfileRequest: got %v, want %v", request, nil)
	}
	if err != ErrBadRouting {
		t.Errorf("decodeGetProfileRequest: got %v, want %v", err, ErrBadRouting)
	}
}

func TestDecodePutProfileRequest(t *testing.T) {
	ctx := context.Background()
	r := httptest.NewRequest(http.MethodPut, "/api/v1/profiles/", nil)

	request, err := decodePutProfileRequest(ctx, r)
	if request != nil {
		t.Errorf("decodePutProfileRequest: got %v, want %v", request, nil)
	}
	if err != ErrBadRouting {
		t.Errorf("decodePutProfileRequest: got %v, want %v", err, ErrBadRouting)
	}
}

func TestDecodePatchProfileRequest(t *testing.T) {
	ctx := context.Background()
	r := httptest.NewRequest(http.MethodPatch, "/api/v1/profiles/", nil)

	request, err := decodePatchProfileRequest(ctx, r)
	if request != nil {
		t.Errorf("decodePatchProfileRequest: got %v, want %v", request, nil)
	}
	if err != ErrBadRouting {
		t.Errorf("decodePatchProfileRequest: got %v, want %v", err, ErrBadRouting)
	}
}

func TestDecodeDeleteProfileRequest(t *testing.T) {
	ctx := context.Background()
	r := httptest.NewRequest(http.MethodDelete, "/api/v1/profiles/", nil)

	request, err := decodeDeleteProfileRequest(ctx, r)
	if request != nil {
		t.Errorf("decodeDeleteProfileRequest: got %v, want %v", request, nil)
	}
	if err != ErrBadRouting {
		t.Errorf("decodeDeleteProfileRequest: got %v, want %v", err, ErrBadRouting)
	}
}

type Response struct{ err error }

func (r Response) error() error { return r.err }

func TestEncodeResponse(t *testing.T) {
	want := struct {
		contentType string
	}{
		"application/json; charset=utf-8",
	}

	ctx := context.Background()
	w := httptest.NewRecorder()
	var response interface{} = &Response{errors.New("")}

	err := encodeResponse(ctx, w, response)
	if err != nil {
		t.Errorf("encodeResponse: got %v, want %v", err, nil)
	}

	response = struct{}{}

	err = encodeResponse(ctx, w, response)
	resp := w.Result()
	contentType := resp.Header.Get("Content-Type")
	if contentType != want.contentType {
		t.Errorf("encodeResponse: got %q, want %q", contentType, want.contentType)
	}
	if err != nil {
		t.Errorf("encodeResponse: got %v, want %v", err, nil)
	}
}

func TestEncodeError(t *testing.T) {
	want := struct {
		contentType string
		code        int
	}{
		"application/json; charset=utf-8",
		http.StatusInternalServerError,
	}

	ctx := context.Background()
	w := httptest.NewRecorder()
	var err error = nil
	defer func() {
		if r := recover(); r == nil {
			t.Error("encodceError could not be recovered with nil error")
		}
	}()

	encodeError(ctx, err, w)

	err = errors.New("")

	encodeError(ctx, err, w)
	resp := w.Result()
	contentType := resp.Header.Get("Content-Type")
	if contentType != want.contentType {
		t.Errorf("encodeError: got %q, want %q", contentType, want.contentType)
	}

	code := resp.StatusCode
	if code != want.code {
		t.Errorf("encodeError: got %d, want %v", code, want.code)
	}
}
