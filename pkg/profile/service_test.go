package profile

import (
	"context"
	"reflect"
	"testing"

	"cloud.google.com/go/datastore"
	"github.com/benkim0414/superego/internal/testutil"
)

func TestNewService(t *testing.T) {
	tc := testutil.SystemTestContext(t)
	ctx := context.Background()

	client, err := datastore.NewClient(ctx, tc.ProjectID)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	s := newDatastoreService(client)

	got := NewService(client)
	if !reflect.DeepEqual(got, s) {
		t.Errorf("NewService: got %v, want %v", got, s)
	}
}
