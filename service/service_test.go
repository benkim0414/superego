package service

import (
	"context"
	"reflect"
	"testing"

	"cloud.google.com/go/datastore"
	"github.com/go-kit/kit/log"
)

func TestNew(t *testing.T) {
	tc := SystemTestContext(t)
	ctx := context.Background()

	client, err := datastore.NewClient(ctx, tc.ProjectID)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	logger := log.NewNopLogger()
	s := newDatastoreService(client)
	s = LoggingMiddleware(logger)(s)

	got := New(client, logger)
	if !reflect.DeepEqual(got, s) {
		t.Errorf("New: got %v, want %v", got, s)
	}
}
