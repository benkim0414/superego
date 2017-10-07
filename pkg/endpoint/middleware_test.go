package endpoint

import (
	"context"
	"testing"

	"github.com/go-kit/kit/log"
)

func TestLoggingMiddleware(t *testing.T) {
	logger := log.NewNopLogger()
	next := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return response, nil
	}

	middleware := LoggingMiddleware(logger)(next)
	ctx := context.Background()
	got, err := middleware(ctx, nil)
	if got != nil {
		t.Errorf("LoggingMiddleware: got %v, want %v", got, nil)
	}
	if err != nil {
		t.Errorf("LoggingMiddleware: error should be nil, not %v", err)
	}
}
