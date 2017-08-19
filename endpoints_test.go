package superego

import (
	"context"
	"reflect"
	"testing"
)

var (
	s   = &FakeService{profiles: map[string]*Profile{}}
	ctx = context.Background()
	p   = &Profile{Email: "gunwoo@gunwoo.org"}
)

func TestMakePostProfileEndpoint(t *testing.T) {
	e := MakePostProfileEndpoint(s)

	req := postProfileRequest{
		Profile: p,
	}
	resp, err := e(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	got := resp.(postProfileResponse)
	if !reflect.DeepEqual(got.Profile, req.Profile) {
		t.Errorf("got %v, want %v", got.Profile, req.Profile)
	}
}

func TestMakeGetProfileEndpoint(t *testing.T) {
	e := MakeGetProfileEndpoint(s)

	req := getProfileRequest{
		ID: "",
	}
	resp, err := e(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	got := resp.(getProfileResponse)
	if !reflect.DeepEqual(got.Profile, p) {
		t.Errorf("got %v, want %v", got.Profile, p)
	}
}

func TestMakePutProfileEndpoint(t *testing.T) {
	e := MakePutProfileEndpoint(s)

	p.Email = "ben.kim@greenenergytrading.com.au"
	req := putProfileRequest{
		ID:      "",
		Profile: p,
	}
	resp, err := e(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	got := resp.(putProfileResponse)
	if !reflect.DeepEqual(got.Profile, p) {
		t.Errorf("got %v, want %v", got.Profile, p)
	}
}

func TestMakePatchProfileEndpoint(t *testing.T) {
	e := MakePatchProfileEndpoint(s)

	p.Email = "gunwoo@gunwoo.org"
	req := patchProfileRequest{
		ID:      "",
		Profile: p,
	}
	resp, err := e(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	got := resp.(patchProfileResponse)
	if !reflect.DeepEqual(got.Profile, p) {
		t.Errorf("got %v, want %v", got.Profile, p)
	}
}

func TestDeleteProfileEndpoinit(t *testing.T) {
	e := MakeDeleteProfileEndpoint(s)

	req := deleteProfileRequest{
		ID: "",
	}
	resp, err := e(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	got := resp.(deleteProfileResponse)
	if got.Err != nil {
		t.Errorf("got %v, want %v", got.Err, nil)
	}
}
