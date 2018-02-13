package service

import (
	"reflect"
	"testing"

	"github.com/go-kit/kit/log"
)

func TestNewLoggingMiddleware(t *testing.T) {
	logger := log.NewNopLogger()
	want := &loggingMiddleware{
		next:   FakeService,
		logger: logger,
	}

	got := NewLoggingMiddleware(logger)(FakeService)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewLoggingMiddleware: got %v, want %v", got, want)
	}
}
