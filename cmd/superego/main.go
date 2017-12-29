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
	"github.com/graphql-go/handler"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
		gqlAddr  = flag.String("graphql.addr", ":8081", "GraphQL listen address")
	)
	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	ctx := context.Background()
	projectID := os.Getenv("GCP_PROJECT_ID")
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		logger.Log("datastore: could not connect: %v", err)
	}
	defer client.Close()

	var (
		service     = service.New(client, logger)
		endpoints   = endpoint.New(service, logger)
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
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, httpHandler)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", *gqlAddr)
		errs <- http.ListenAndServe(*gqlAddr, gqlHandler)
	}()
	logger.Log("exit", <-errs)
}
