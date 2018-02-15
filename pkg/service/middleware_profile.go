package service

import (
	"context"
	"time"

	"github.com/benkim0414/superego/pkg/profile"
)

func (mw LoggingMiddleware) PostProfile(ctx context.Context, p *profile.Profile) (profile *profile.Profile, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log("method", "PostProfile", "id", p.ID, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.Next.PostProfile(ctx, p)
}

func (mw LoggingMiddleware) GetProfile(ctx context.Context, id string) (profile *profile.Profile, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log("method", "GetProfile", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.Next.GetProfile(ctx, id)
}

func (mw LoggingMiddleware) PutProfile(ctx context.Context, id string, p *profile.Profile) (profile *profile.Profile, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log("method", "PutProfile", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.Next.PutProfile(ctx, id, p)
}

func (mw LoggingMiddleware) PatchProfile(ctx context.Context, id string, p *profile.Profile) (profile *profile.Profile, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log("method", "PatchProfile", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.Next.PatchProfile(ctx, id, p)
}

func (mw LoggingMiddleware) DeleteProfile(ctx context.Context, id string) (err error) {
	defer func(begin time.Time) {
		mw.Logger.Log("method", "DeleteProfile", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.Next.DeleteProfile(ctx, id)
}
