package superego

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/go-kit/kit/log"
)

func TestMakeHTTPHandler(t *testing.T) {
	logger := log.NewNopLogger()
	h := MakeHTTPHandler(s, logger)

	ts := httptest.NewServer(h)
	defer ts.Close()

	var buf bytes.Buffer
	p := &Profile{Email: "gunwoo@gunwoo.org"}
	err := json.NewEncoder(&buf).Encode(p)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, ts.URL+"/api/v1/profiles/", &buf)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)

	resp := w.Result()
	var body postProfileResponse
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(body.Profile, p) {
		t.Errorf("got %v, want %v", body.Profile, p)
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

	p := &Profile{Email: "gunwoo@gunwoo.org"}
	err = json.NewEncoder(&buf).Encode(p)
	if err != nil {
		t.Fatal(err)
	}
	r = httptest.NewRequest(http.MethodPost, "/api/v1/profiles/", &buf)

	request, err = decodePostProfileRequest(ctx, r)
	got := request.(postProfileRequest)
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
		t.Errorf("encodeResponse:error: got %v, want %v", err, nil)
	}

	response = struct{}{}

	err = encodeResponse(ctx, w, response)
	resp := w.Result()
	contentType := resp.Header.Get("Content-Type")
	if contentType != want.contentType {
		t.Errorf("encodeResponse:Content-Type: got %q, want %q", contentType, want.contentType)
	}
	if err != nil {
		t.Errorf("encodeResponse:error: got %v, want %v", err, nil)
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
		t.Errorf("encodeError:Content-Type got %q, want %q", contentType, want.contentType)
	}

	code := resp.StatusCode
	if code != want.code {
		t.Errorf("encodeError:StatusCode: got %d, want %v", code, want.code)
	}
}
