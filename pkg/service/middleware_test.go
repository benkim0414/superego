package service

import (
	"reflect"
	"testing"

	"github.com/benkim0414/superego/pkg/profile"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func TestNewLoggingMiddleware(t *testing.T) {
	logger := log.NewNopLogger()
	want := &LoggingMiddleware{
		Logger: logger,
		Next:   profile.FakeService,
	}

	got := NewLoggingMiddleware(logger)(profile.FakeService)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewLoggingMiddleware: got %v, want %v", got, want)
	}
}

func TestNewInstrumentingMiddleware(t *testing.T) {
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "middleware_test",
		Subsystem: "profile",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "middleware_test",
		Subsystem: "profile",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	want := &InstrumentingMiddleware{
		RequestCount:   requestCount,
		RequestLatency: requestLatency,
		Next:           profile.FakeService,
	}

	got := NewInstrumentingMiddleware(requestCount, requestLatency)(profile.FakeService)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewInstrumentingMiddleware: got %v, want %v", got, want)
	}
}
