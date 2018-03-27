// Package testutil provides test helpers for the superego repo.
package testutil

import (
	"errors"
	"os"
	"testing"
)

var (
	// noProjectID is returned when GCP_PROJECT_ID is not set.
	noProjectID = errors.New("GCP_PROJECT_ID is not set")
)

// Context represents a test context.
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
