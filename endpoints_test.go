package superego

import (
	"context"
	"reflect"
	"testing"
)

func TestMakePostProfileEndpoint(t *testing.T) {
	e := MakePostProfileEndpoint(s)

	ctx := context.Background()
	p := &Profile{Email: "gunwoo@gunwoo.org"}
	req := postProfileRequest{
		Profile: p,
	}
	resp, err := e(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	got := resp.(postProfileResponse)
	if !reflect.DeepEqual(got.Profile, req.Profile) {
		t.Errorf("PostProfileEndpoint: got %v, want %v", got.Profile, req.Profile)
	}

	err = got.error()
	if err != nil {
		t.Errorf("postProfileResponse.error(): got %v, want %v", err, nil)
	}
}

func TestMakeGetProfileEndpoint(t *testing.T) {
	e := MakeGetProfileEndpoint(s)

	ctx := context.Background()
	p := &Profile{Email: "gunwoo@gunwoo.org"}
	req := getProfileRequest{
		ID: p.ID,
	}
	resp, err := e(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	got := resp.(getProfileResponse)
	if !reflect.DeepEqual(got.Profile, p) {
		t.Errorf("GetProfileEndpoint: got %v, want %v", got.Profile, p)
	}
	err = got.error()
	if err != nil {
		t.Errorf("getProfileResponse.error(): got %v, want %v", err, nil)
	}
}

func TestMakePutProfileEndpoint(t *testing.T) {
	e := MakePutProfileEndpoint(s)

	ctx := context.Background()
	p := &Profile{Email: "ben.kim@greenenergytrading.com.au"}
	req := putProfileRequest{
		ID:      p.ID,
		Profile: p,
	}
	resp, err := e(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	got := resp.(putProfileResponse)
	if !reflect.DeepEqual(got.Profile, p) {
		t.Errorf("PutProfileEndpoint: got %v, want %v", got.Profile, p)
	}
	err = got.error()
	if err != nil {
		t.Errorf("putProfileResponse.error(): got %v, want %v", err, nil)
	}
}

func TestMakePatchProfileEndpoint(t *testing.T) {
	e := MakePatchProfileEndpoint(s)

	ctx := context.Background()
	p := &Profile{Email: "gunwoo@gunwoo.org"}
	req := patchProfileRequest{
		ID:      p.ID,
		Profile: p,
	}
	resp, err := e(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	got := resp.(patchProfileResponse)
	if !reflect.DeepEqual(got.Profile, p) {
		t.Errorf("PatchProfileEndpoint: got %v, want %v", got.Profile, p)
	}
	err = got.error()
	if err != nil {
		t.Errorf("patchProfileResponse.error(): got %v, want %v", err, nil)
	}
}

func TestDeleteProfileEndpoinit(t *testing.T) {
	e := MakeDeleteProfileEndpoint(s)

	ctx := context.Background()
	req := deleteProfileRequest{
		ID: "",
	}
	resp, err := e(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	got := resp.(deleteProfileResponse)
	if got.Err != nil {
		t.Errorf("DeleteProfileEndpoint: got %v, want %v", got.Err, nil)
	}
	err = got.error()
	if err != nil {
		t.Errorf("deleteProfileResponse.error(): got %v, want %v", err, nil)
	}
}
