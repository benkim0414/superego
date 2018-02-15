package service

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/benkim0414/superego/pkg/profile"
	"github.com/go-kit/kit/log"
)

func TestMiddlewarePostProfile(t *testing.T) {
	var buf bytes.Buffer
	logger := log.NewLogfmtLogger(&buf)

	mw := NewLoggingMiddleware(logger)(profile.FakeService)

	ctx := context.Background()
	p := &profile.Profile{Email: "gunwoo@gunwoo.org"}
	got, err := mw.PostProfile(ctx, p)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, p) {
		t.Errorf("PostProfile: got %v, want %v", got, p)
	}
}

func TestMiddlewareGetProfile(t *testing.T) {
	var buf bytes.Buffer
	logger := log.NewLogfmtLogger(&buf)

	mw := NewLoggingMiddleware(logger)(profile.FakeService)

	ctx := context.Background()
	p := &profile.Profile{ID: "", Email: "gunwoo@gunwoo.org"}
	got, err := mw.GetProfile(ctx, p.ID)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, p) {
		t.Errorf("GetProfile: got %v, want %v", got, p)
	}
}

func TestMiddlewarePutProfile(t *testing.T) {
	var buf bytes.Buffer
	logger := log.NewLogfmtLogger(&buf)

	mw := NewLoggingMiddleware(logger)(profile.FakeService)

	ctx := context.Background()
	p := &profile.Profile{ID: "", Email: "ben.kim@greenenergytrading.com.au"}
	got, err := mw.PutProfile(ctx, p.ID, p)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, p) {
		t.Errorf("PutProfile: got %v, want %v", got, p)
	}
}

func TestMiddlewarePatchProfile(t *testing.T) {
	var buf bytes.Buffer
	logger := log.NewLogfmtLogger(&buf)

	mw := NewLoggingMiddleware(logger)(profile.FakeService)

	ctx := context.Background()
	p := &profile.Profile{ID: "", Email: "gunwoo@gunwoo.org"}
	got, err := mw.PatchProfile(ctx, p.ID, p)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, p) {
		t.Errorf("PatchProfile: got %v, want %v", got, p)
	}
}

func TestMiddlewareDeleteProfile(t *testing.T) {
	var buf bytes.Buffer
	logger := log.NewLogfmtLogger(&buf)

	mw := NewLoggingMiddleware(logger)(profile.FakeService)

	ctx := context.Background()
	id := ""
	err := mw.DeleteProfile(ctx, id)
	if err != nil {
		t.Errorf("DeleteProfile: got %v, want %v", err, nil)
	}
}
