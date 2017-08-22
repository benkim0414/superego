package superego

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/go-kit/kit/log"
)

func TestLoggingMiddleware(t *testing.T) {
	logger := log.NewNopLogger()
	want := &loggingMiddleware{
		next:   s,
		logger: logger,
	}

	got := LoggingMiddleware(logger)(s)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("LoggingMiddleware: got %v, want %v", got, want)
	}
}

func TestPostProfile(t *testing.T) {
	var buf bytes.Buffer
	logger := log.NewLogfmtLogger(&buf)

	mw := LoggingMiddleware(logger)(s)

	ctx := context.Background()
	p := &Profile{Email: "gunwoo@gunwoo.org"}
	got, err := mw.PostProfile(ctx, p)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, p) {
		t.Errorf("PostProfile: got %v, want %v", got, p)
	}
}

func TestGetProfile(t *testing.T) {
	var buf bytes.Buffer
	logger := log.NewLogfmtLogger(&buf)

	mw := LoggingMiddleware(logger)(s)

	ctx := context.Background()
	p := &Profile{ID: "", Email: "gunwoo@gunwoo.org"}
	got, err := mw.GetProfile(ctx, p.ID)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, p) {
		t.Errorf("GetProfile: got %v, want %v", got, p)
	}
}

func TestPutProfile(t *testing.T) {
	var buf bytes.Buffer
	logger := log.NewLogfmtLogger(&buf)

	mw := LoggingMiddleware(logger)(s)

	ctx := context.Background()
	p := &Profile{ID: "", Email: "ben.kim@greenenergytrading.com.au"}
	got, err := mw.PutProfile(ctx, p.ID, p)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, p) {
		t.Errorf("PutProfile: got %v, want %v", got, p)
	}
}

func TestPatchProfile(t *testing.T) {
	var buf bytes.Buffer
	logger := log.NewLogfmtLogger(&buf)

	mw := LoggingMiddleware(logger)(s)

	ctx := context.Background()
	p := &Profile{ID: "", Email: "gunwoo@gunwoo.org"}
	got, err := mw.PatchProfile(ctx, p.ID, p)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, p) {
		t.Errorf("PatchProfile: got %v, want %v", got, p)
	}
}

func TestDeleteProfile(t *testing.T) {
	var buf bytes.Buffer
	logger := log.NewLogfmtLogger(&buf)

	mw := LoggingMiddleware(logger)(s)

	ctx := context.Background()
	id := ""
	err := mw.DeleteProfile(ctx, id)
	if err != nil {
		t.Errorf("DeleteProfile: got %v, want %v", err, nil)
	}
}
