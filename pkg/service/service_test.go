package service

import (
	"context"
	"reflect"
	"testing"

	"cloud.google.com/go/datastore"
	"github.com/benkim0414/superego/internal/testutil"
	"github.com/benkim0414/superego/pkg/profile"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func TestNew(t *testing.T) {
	tc := testutil.SystemTestContext(t)
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, tc.ProjectID)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	logger := log.NewNopLogger()

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "service_test",
		Subsystem: "profile",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "service_test",
		Subsystem: "profile",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	var svc Service
	svc = &service{
		profile.NewService(client),
	}
	svc = NewLoggingMiddleware(logger)(svc)
	svc = NewInstrumentingMiddleware(requestCount, requestLatency)(svc)

	got := New(client, logger, requestCount, requestLatency)
	if !reflect.DeepEqual(got, svc) {
		t.Errorf("New: got %v, want %v", got, svc)
	}
}
