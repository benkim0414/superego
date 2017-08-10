package testing

import (
	"errors"
	"os"
	"testing"
)

var (
	noProjectID = errors.New("PROJECT_ID is not set")
)

type Context struct {
	ProjectID string
}

// SystemTestContext returns the test context.
// The test is skipped if the PROJECT_ID environment variable is not set.
func SystemTestContext(t *testing.T) Context {
	tc, err := newContext()
	if err == noProjectID {
		t.Skip(err)
	} else if err != nil {
		t.Fatal(err)
	}
	return tc
}

func newContext() (Context, error) {
	tc := Context{}

	tc.ProjectID = os.Getenv("PROJECT_ID")
	if tc.ProjectID == "" {
		return tc, noProjectID
	}

	return tc, nil
}
