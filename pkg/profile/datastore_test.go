package profile

import (
	"context"
	"testing"

	"github.com/benkim0414/superego/internal/testutil"

	"cloud.google.com/go/datastore"
)

func TestDatastoreService(t *testing.T) {
	tc := testutil.SystemTestContext(t)
	ctx := context.Background()

	client, err := datastore.NewClient(
		ctx,
		tc.ProjectID,
	)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	s := newDatastoreService(client)

	p := &Profile{
		Email: "gunwoo@gunwoo.org",
	}

	got, err := s.PostProfile(ctx, p)
	if err != nil {
		t.Fatal(err)
	}
	if got.Email != p.Email {
		t.Errorf("PostProfile: got %q, want %q", got.Email, p.Email)
	}

	p.ID = got.ID
	got, err = s.GetProfile(ctx, p.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got.Email != p.Email {
		t.Errorf("GetProfile: got %q, want %q", got.Email, p.Email)
	}

	p.Email = "ben.kim@greenenergytrading.com.au"
	got, err = s.PutProfile(ctx, p.ID, p)
	if err != nil {
		t.Fatal(err)
	}
	if got.Email != p.Email {
		t.Errorf("PutProfile: got %q, want %q", got.Email, p.Email)
	}

	p.Email = "gunwoo@gunwoo.org"
	got, err = s.PatchProfile(ctx, p.ID, p)
	if err != nil {
		t.Fatal(err)
	}
	if got.Email != p.Email {
		t.Errorf("PatchProfile: got %q, want %q", got.Email, p.Email)
	}

	err = s.DeleteProfile(ctx, p.ID)
	if err != nil {
		t.Fatal(err)
	}
}
