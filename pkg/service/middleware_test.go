package service

import (
	"reflect"
	"testing"

	"github.com/benkim0414/superego/pkg/profile"
	"github.com/go-kit/kit/log"
)

func TestNewLoggingMiddleware(t *testing.T) {
	logger := log.NewNopLogger()
	want := &LoggingMiddleware{
		Next:   profile.FakeService,
		Logger: logger,
	}

	got := NewLoggingMiddleware(logger)(profile.FakeService)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewLoggingMiddleware: got %v, want %v", got, want)
	}
}
