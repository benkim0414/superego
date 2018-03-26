package endpoint

import (
	"context"
	"testing"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
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

func TestInstrumentingMiddleware(t *testing.T) {
	duration := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "endpoint_middleware_test",
		Subsystem: "profile",
		Name:      "request_duration_seconds",
		Help:      "Request duration in seconds.",
	}, []string{"method", "success"})
	next := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return response, nil
	}

	middleware := InstrumentingMiddleware(duration.With("method", "profile"))(next)
	ctx := context.Background()
	got, err := middleware(ctx, nil)
	if got != nil {
		t.Errorf("InstrumentingMiddleware: got %v, want %v", got, nil)
	}
	if err != nil {
		t.Errorf("InstrumentingMiddleware: error should be nil, not %v", err)
	}
}
