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
	"google.golang.org/api/option"

	"github.com/benkim0414/superego"
	"github.com/go-kit/kit/log"
)

func main() {
	var (
		httpAddr            = flag.String("http.addr", ":8080", "HTTP listen address")
		credentialsFilename = flag.String(
			"credentials.filename",
			"superego-25e0162c5c59.json",
			"Google service account credentials file",
		)
	)
	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	ctx := context.Background()
	projectID := os.Getenv("DATASTORE_PROJECT_ID")
	client, err := datastore.NewClient(
		ctx,
		projectID,
		option.WithServiceAccountFile(*credentialsFilename),
	)
	if err != nil {
		logger.Log("datastore: could not connect: %v", err)
	}
	defer client.Close()

	var s superego.Service
	s = superego.NewDatastoreService(client)
	s = superego.LoggingMiddleware(logger)(s)

	var h http.Handler
	h = superego.MakeHTTPHandler(s, log.With(logger, "component", "HTTP"))

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, h)
	}()

	logger.Log("exit", <-errs)
}
