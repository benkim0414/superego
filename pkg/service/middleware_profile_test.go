package service

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/benkim0414/superego/pkg/profile"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func TestLoggingMiddlewarePostProfile(t *testing.T) {
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

func TestLoggingMiddlewareGetProfile(t *testing.T) {
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

func TestLoggingMiddlewarePutProfile(t *testing.T) {
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

func TestLoggingMiddlewarePatchProfile(t *testing.T) {
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

func TestLoggingMiddlewareDeleteProfile(t *testing.T) {
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

func TestInstrumentingMiddlewarePostProfile(t *testing.T) {
	namespace, subsystem := "test", "post_profile"
	mw := newTestInstrumentingMiddleware(namespace, subsystem)
	svc := mw(profile.FakeService)
	_, err := svc.PostProfile(context.Background(), &profile.Profile{
		Email: "gunwoo@gunwoo.org",
	})
	if err != nil {
		t.Fatal(err)
	}
	want, have := metric(namespace, subsystem, "request_count"), scrapePrometheus(t)
	if !strings.Contains(have, want) {
		t.Errorf("metric stanza not found or incorrect\n%s", have)
	}
}

func TestInstrumentingMiddlewareGetProfile(t *testing.T) {
	namespace, subsystem := "test", "get_profile"
	mw := newTestInstrumentingMiddleware(namespace, subsystem)
	svc := mw(profile.FakeService)
	_, err := svc.GetProfile(context.Background(), "")
	if err != nil {
		t.Fatal(err)
	}
	want, have := metric(namespace, subsystem, "request_count"), scrapePrometheus(t)
	if !strings.Contains(have, want) {
		t.Errorf("metric stanza not found or incorrect\n%s", have)
	}
}

func TestInstrumentingMiddlewarePutProfile(t *testing.T) {
	namespace, subsystem := "test", "put_profile"
	mw := newTestInstrumentingMiddleware(namespace, subsystem)
	svc := mw(profile.FakeService)
	_, err := svc.PutProfile(context.Background(), "", &profile.Profile{
		Email: "ben.kim@greenenergytrading.com.au",
	})
	if err != nil {
		t.Fatal(err)
	}
	want, have := metric(namespace, subsystem, "request_count"), scrapePrometheus(t)
	if !strings.Contains(have, want) {
		t.Errorf("metric stanza not found or incorrect\n%s", have)
	}
}

func TestInstrumentingMiddlewarePatchProfile(t *testing.T) {
	namespace, subsystem := "test", "patch_profile"
	mw := newTestInstrumentingMiddleware(namespace, subsystem)
	svc := mw(profile.FakeService)
	_, err := svc.PatchProfile(context.Background(), "", &profile.Profile{
		Email: "gunwoo@gunwoo.org",
	})
	if err != nil {
		t.Fatal(err)
	}
	want, have := metric(namespace, subsystem, "request_count"), scrapePrometheus(t)
	if !strings.Contains(have, want) {
		t.Errorf("metric stanza not found or incorrect\n%s", have)
	}
}

func TestInstrumentingMiddlewareDeleteProfile(t *testing.T) {
	namespace, subsystem := "test", "delete_profile"
	mw := newTestInstrumentingMiddleware(namespace, subsystem)
	svc := mw(profile.FakeService)
	err := svc.DeleteProfile(context.Background(), "")
	if err != nil {
		t.Fatal(err)
	}
	want, have := metric(namespace, subsystem, "request_count"), scrapePrometheus(t)
	if !strings.Contains(have, want) {
		t.Errorf("metric stanza not found or incorrect\n%s", have)
	}
}

func newTestInstrumentingMiddleware(namespace, subsystem string) Middleware {
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	return func(next Service) Service {
		return &InstrumentingMiddleware{
			RequestCount:   requestCount,
			RequestLatency: requestLatency,
			Next:           next,
		}
	}
}

func metric(namespace, subsystem, name string) string {
	return strings.Join([]string{namespace, subsystem, name}, "_")
}

// scrapePrometheus returns the test encoding of the current state of Prometheus.
func scrapePrometheus(t *testing.T) string {
	server := httptest.NewServer(stdprometheus.UninstrumentedHandler())
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	return strings.TrimSpace(string(buf))
}
