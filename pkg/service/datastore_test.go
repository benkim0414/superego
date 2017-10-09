package service

import (
	"context"
	"errors"
	"os"
	"testing"

	"cloud.google.com/go/datastore"
)

var (
	noProjectID = errors.New("GCP_PROJECT_ID is not set")
)

type Context struct {
	ProjectID string
}

func newContext() (Context, error) {
	tc := Context{}

	tc.ProjectID = os.Getenv("GCP_PROJECT_ID")
	if tc.ProjectID == "" {
		return tc, noProjectID
	}

	return tc, nil
}

// SystemTestContext returns the test context.
// The test is skipped if the GCP_PROJECT_ID environment variable is not set.
func SystemTestContext(t *testing.T) Context {
	tc, err := newContext()
	if err == noProjectID {
		t.Skip(err)
	} else if err != nil {
		t.Fatal(err)
	}
	return tc
}

func TestDatastoreService(t *testing.T) {
	tc := SystemTestContext(t)
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
