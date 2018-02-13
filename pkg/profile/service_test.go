package profile

import (
	"context"
	"reflect"
	"testing"

	"cloud.google.com/go/datastore"
)

func TestNewService(t *testing.T) {
	tc := SystemTestContext(t)
	ctx := context.Background()

	client, err := datastore.NewClient(ctx, tc.ProjectID)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	s := newDatastoreService(client)

	got := NewService(client)
	if !reflect.DeepEqual(got, s) {
		t.Errorf("New: got %v, want %v", got, s)
	}
}
