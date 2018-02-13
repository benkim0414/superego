package profile

import (
	"context"
	"reflect"
	"testing"
)

func TestFakeServicePostProfile(t *testing.T) {
	ctx := context.Background()
	p := &Profile{ID: "gunwoo", Email: "gunwoo@gunwoo.org"}

	got, err := FakeService.PostProfile(ctx, p)
	if !reflect.DeepEqual(got, p) {
		t.Errorf("PostProfile: got %v, want %v", got, p)
	}
	if err != nil {
		t.Errorf("PostProfile: error should be nil, not %v", err)
	}
}

func TestFakeServiceGetProfile(t *testing.T) {
	ctx := context.Background()
	p := &Profile{ID: "gunwoo", Email: "gunwoo@gunwoo.org"}

	got, err := FakeService.GetProfile(ctx, p.ID)
	if !reflect.DeepEqual(got, p) {
		t.Errorf("GetProfile: got %v, want %v", got, p)
	}
	if err != nil {
		t.Errorf("GetProfile: error should be nil, not %v", err)
	}

	got, err = FakeService.GetProfile(ctx, "invalid")
	if !reflect.DeepEqual(got, &Profile{}) {
		t.Errorf("GetProfile: profile should be empty, not %v", got)
	}
	if err != ErrNoSuchEntity {
		t.Errorf("GetProfile: got %v, want %v", err, ErrNoSuchEntity)
	}
}

func TestFakeServicePutProfile(t *testing.T) {
	ctx := context.Background()
	p := &Profile{ID: "gunwoo", Email: "benkim@greenenergytrading.com.au"}

	got, err := FakeService.PutProfile(ctx, p.ID, p)
	if !reflect.DeepEqual(got, p) {
		t.Errorf("PutProfile: got %v, want %v", got, p)
	}
	if err != nil {
		t.Errorf("PutProfile: error should be nil, not %v", err)
	}
}

func TestFakeServicePatchProfile(t *testing.T) {
	ctx := context.Background()
	p := &Profile{
		ID:          "gunwoo",
		DisplayName: "benkim0414",
		Email:       "gunwoo@gunwoo.org",
		ImageURL:    "https://octodex.github.com/images/codercat.jpg",
		AboutMe:     "Codercat",
	}

	got, err := FakeService.PatchProfile(ctx, p.ID, p)
	if !reflect.DeepEqual(got, p) {
		t.Errorf("PatchProfile: got %v, want %v", got, p)
	}
	if err != nil {
		t.Errorf("PatchProfile: error should be nil, not %v", err)
	}

	got, err = FakeService.PatchProfile(ctx, "invalid", p)
	if !reflect.DeepEqual(got, &Profile{}) {
		t.Errorf("PatchProfile: profile should be empty, not %v", got)
	}
	if err != ErrNoSuchEntity {
		t.Errorf("PatchProfile: got %v, want %v", err, ErrNoSuchEntity)
	}
}

func TestFakeServiceDeleteProfile(t *testing.T) {
	ctx := context.Background()
	id := "gunwoo"

	err := FakeService.DeleteProfile(ctx, id)
	if err != nil {
		t.Errorf("DeleteProfile: error should be nil, not %v", err)
	}

	err = FakeService.DeleteProfile(ctx, "invalid")
	if err != ErrNoSuchEntity {
		t.Errorf("GetProfile: got %v, want %v", err, ErrNoSuchEntity)
	}
}
