package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"cloud.google.com/go/datastore"

	"github.com/benkim0414/superego/pkg/endpoint"
	"github.com/benkim0414/superego/pkg/graphql"
	"github.com/benkim0414/superego/pkg/service"
	"github.com/benkim0414/superego/pkg/transport"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/graphql-go/handler"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var (
		promAddr = flag.String("prom.addr", ":8079", "Prometheus listen address")
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
		gqlAddr  = flag.String("graphql.addr", ":8081", "GraphQL listen address")
	)
	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	fieldKeys := []string{"method", "error"}
	var requestCount metrics.Counter
	requestCount = kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "superego",
		Subsystem: "profile",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	var requestLatency metrics.Histogram
	requestLatency = kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "superego",
		Subsystem: "profile",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	var duration metrics.Histogram
	duration = kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "superego",
		Subsystem: "profile",
		Name:      "request_duration_seconds",
		Help:      "Request duration in seconds.",
	}, []string{"method", "success"})
	http.DefaultServeMux.Handle("/metrics", promhttp.Handler())

	ctx := context.Background()
	projectID := os.Getenv("GCP_PROJECT_ID")
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		logger.Log("datastore: could not connect: %v", err)
	}
	defer client.Close()

	var (
		service     = service.New(client, logger, requestCount, requestLatency)
		endpoints   = endpoint.New(service, logger, duration)
		httpHandler = transport.NewHTTPHandler(endpoints, logger)
	)

	schema, err := graphql.NewSchema(service)
	if err != nil {
		logger.Log("graphql: could not create new schema: %v", err)
	}

	var gqlHandler = handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "Prometheus/HTTP", "addr", *promAddr)
		errs <- http.ListenAndServe(*promAddr, http.DefaultServeMux)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, httpHandler)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", *gqlAddr)
		errs <- http.ListenAndServe(*gqlAddr, gqlHandler)
	}()
	logger.Log("exit", <-errs)
}
