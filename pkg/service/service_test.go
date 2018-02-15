package service

import (
	"context"
	"reflect"
	"testing"

	"cloud.google.com/go/datastore"
	"github.com/benkim0414/superego/internal/testutil"
	"github.com/benkim0414/superego/pkg/profile"
	"github.com/go-kit/kit/log"
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

	svc := &service{
		profile.NewService(client),
	}
	want := NewLoggingMiddleware(logger)(svc)

	got := New(client, logger)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("New: got %v, want %v", got, want)
	}
}
