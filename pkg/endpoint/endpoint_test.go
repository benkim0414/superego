package endpoint

import (
	"context"
	"reflect"
	"testing"

	"github.com/benkim0414/superego/pkg/profile"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func TestNew(t *testing.T) {
	logger := log.NewNopLogger()
	duration := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "endpoint_test",
		Subsystem: "profile",
		Name:      "request_duration_seconds",
		Help:      "Request duration in seconds.",
	}, []string{"method", "success"})

	postProfileEndpoint := MakePostProfileEndpoint(profile.FakeService)
	postProfileEndpoint = LoggingMiddleware(log.With(logger, "method", "PostProfile"))(postProfileEndpoint)
	postProfileEndpoint = InstrumentingMiddleware(duration.With("method", "PostProfile"))(postProfileEndpoint)

	getProfileEndpoint := MakeGetProfileEndpoint(profile.FakeService)
	getProfileEndpoint = LoggingMiddleware(log.With(logger, "method", "GetProfile"))(getProfileEndpoint)
	getProfileEndpoint = InstrumentingMiddleware(duration.With("method", "GetProfile"))(getProfileEndpoint)

	putProfileEndpoint := MakePutProfileEndpoint(profile.FakeService)
	putProfileEndpoint = LoggingMiddleware(log.With(logger, "method", "PutProfile"))(putProfileEndpoint)
	putProfileEndpoint = InstrumentingMiddleware(duration.With("method", "PutProfile"))(putProfileEndpoint)

	patchProfileEndpoint := MakePatchProfileEndpoint(profile.FakeService)
	patchProfileEndpoint = LoggingMiddleware(log.With(logger, "method", "PatchProfile"))(patchProfileEndpoint)
	patchProfileEndpoint = InstrumentingMiddleware(duration.With("method", "PatchProfile"))(patchProfileEndpoint)

	deleteProfileEndpoint := MakeDeleteProfileEndpoint(profile.FakeService)
	deleteProfileEndpoint = LoggingMiddleware(log.With(logger, "method", "DeleteProfile"))(deleteProfileEndpoint)
	deleteProfileEndpoint = InstrumentingMiddleware(duration.With("method", "DeleteProfile"))(deleteProfileEndpoint)

	endpoints := New(profile.FakeService, logger, duration)
	ctx := context.Background()
	var req interface{}
	req = PostProfileRequest{
		Profile: &profile.Profile{Email: "gunwoo@gunwoo.org"},
	}
	want, _ := postProfileEndpoint(ctx, req)
	got, _ := endpoints.PostProfileEndpoint(ctx, req)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Endpoints.PostProfileEndpoint: got %v, want %v", got, want)
	}

	req = GetProfileRequest{
		ID: "",
	}
	want, _ = getProfileEndpoint(ctx, req)
	got, _ = endpoints.GetProfileEndpoint(ctx, req)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Endpoints.GetProfileEndpoint: got %v, want %v", got, want)
	}

	req = PutProfileRequest{
		ID:      "",
		Profile: &profile.Profile{Email: "ben.kim@greenenergytrading.com.au"},
	}
	want, _ = putProfileEndpoint(ctx, req)
	got, _ = endpoints.PutProfileEndpoint(ctx, req)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Endpoints.PutProfileEndpoint: got %v, want %v", got, want)
	}

	req = PatchProfileRequest{
		ID:      "",
		Profile: &profile.Profile{Email: "gunwoo@gunwoo.org"},
	}
	want, _ = patchProfileEndpoint(ctx, req)
	got, _ = endpoints.PatchProfileEndpoint(ctx, req)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Endpoints.PatchProfileEndpoint: got %v, want %v", got, want)
	}

	req = DeleteProfileRequest{
		ID: "",
	}
	_, want = deleteProfileEndpoint(ctx, req)
	_, got = endpoints.DeleteProfileEndpoint(ctx, req)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Endpoints.DeleteProfileEndpoint: got %v, want %v", got, want)
	}
}

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
