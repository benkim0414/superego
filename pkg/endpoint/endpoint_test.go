package endpoint

import (
	"context"
	"reflect"
	"testing"

	"github.com/benkim0414/superego/pkg/profile"
)

func TestMakePostProfileEndpoint(t *testing.T) {
	e := MakePostProfileEndpoint(profile.FakeService)

	ctx := context.Background()
	p := &profile.Profile{Email: "gunwoo@gunwoo.org"}
	req := PostProfileRequest{
		Profile: p,
	}
	resp, err := e(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	got := resp.(PostProfileResponse)
	if !reflect.DeepEqual(got.Profile, req.Profile) {
		t.Errorf("PostProfileEndpoint: got %v, want %v", got.Profile, req.Profile)
	}

	err = got.error()
	if err != nil {
		t.Errorf("postProfileResponse.error(): got %v, want %v", err, nil)
	}
}

func TestMakeGetProfileEndpoint(t *testing.T) {
	e := MakeGetProfileEndpoint(profile.FakeService)

	ctx := context.Background()
	p := &profile.Profile{Email: "gunwoo@gunwoo.org"}
	req := GetProfileRequest{
		ID: p.ID,
	}
	resp, err := e(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	got := resp.(GetProfileResponse)
	if !reflect.DeepEqual(got.Profile, p) {
		t.Errorf("GetProfileEndpoint: got %v, want %v", got.Profile, p)
	}
	err = got.error()
	if err != nil {
		t.Errorf("getProfileResponse.error(): got %v, want %v", err, nil)
	}
}

func TestMakePutProfileEndpoint(t *testing.T) {
	e := MakePutProfileEndpoint(profile.FakeService)

	ctx := context.Background()
	p := &profile.Profile{Email: "ben.kim@greenenergytrading.com.au"}
	req := PutProfileRequest{
		ID:      p.ID,
		Profile: p,
	}
	resp, err := e(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	got := resp.(PutProfileResponse)
	if !reflect.DeepEqual(got.Profile, p) {
		t.Errorf("PutProfileEndpoint: got %v, want %v", got.Profile, p)
	}
	err = got.error()
	if err != nil {
		t.Errorf("putProfileResponse.error(): got %v, want %v", err, nil)
	}
}

func TestMakePatchProfileEndpoint(t *testing.T) {
	e := MakePatchProfileEndpoint(profile.FakeService)

	ctx := context.Background()
	p := &profile.Profile{Email: "gunwoo@gunwoo.org"}
	req := PatchProfileRequest{
		ID:      p.ID,
		Profile: p,
	}
	resp, err := e(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	got := resp.(PatchProfileResponse)
	if !reflect.DeepEqual(got.Profile, p) {
		t.Errorf("PatchProfileEndpoint: got %v, want %v", got.Profile, p)
	}
	err = got.error()
	if err != nil {
		t.Errorf("patchProfileResponse.error(): got %v, want %v", err, nil)
	}
}

func TestDeleteProfileEndpoinit(t *testing.T) {
	e := MakeDeleteProfileEndpoint(profile.FakeService)

	ctx := context.Background()
	req := DeleteProfileRequest{
		ID: "",
	}
	resp, err := e(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	got := resp.(DeleteProfileResponse)
	if got.Err != nil {
		t.Errorf("DeleteProfileEndpoint: got %v, want %v", got.Err, nil)
	}
	err = got.error()
	if err != nil {
		t.Errorf("deleteProfileResponse.error(): got %v, want %v", err, nil)
	}
}
