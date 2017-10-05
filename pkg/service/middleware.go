package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(Service) Service

// LoggingMiddleware takes a logger as a dependency
// and returns a ServiceMiddleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

func (mw loggingMiddleware) PostProfile(ctx context.Context, p *Profile) (profile *Profile, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "PostProfile", "id", p.ID, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.PostProfile(ctx, p)
}

func (mw loggingMiddleware) GetProfile(ctx context.Context, id string) (profile *Profile, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetProfile", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetProfile(ctx, id)
}

func (mw loggingMiddleware) PutProfile(ctx context.Context, id string, p *Profile) (profile *Profile, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "PutProfile", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.PutProfile(ctx, id, p)
}

func (mw loggingMiddleware) PatchProfile(ctx context.Context, id string, p *Profile) (profile *Profile, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "PatchProfile", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.PatchProfile(ctx, id, p)
}

func (mw loggingMiddleware) DeleteProfile(ctx context.Context, id string) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "DeleteProfile", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.DeleteProfile(ctx, id)
}