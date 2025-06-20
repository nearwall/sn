package handlers

import (
	"context"

	api "sn/api/rest/generated"
)

// PostCreatePost implements POST /post/create operation.
//
// POST /post/create
func (r *Resolver) PostCreatePost(ctx context.Context, req api.OptPostCreatePostReq) (api.PostCreatePostRes, error) {
	return &api.PostCreatePostServiceUnavailable{}, nil
}

// PostDeleteIDPut implements PUT /post/delete/{id} operation.
//
// PUT /post/delete/{id}
func (r *Resolver) PostDeleteIDPut(ctx context.Context, params api.PostDeleteIDPutParams) (api.PostDeleteIDPutRes, error) {
	return &api.PostDeleteIDPutServiceUnavailable{}, nil
}

// PostFeedGet implements GET /post/feed operation.
//
// GET /post/feed
func (r *Resolver) PostFeedGet(ctx context.Context, params api.PostFeedGetParams) (api.PostFeedGetRes, error) {
	return &api.PostFeedGetServiceUnavailable{}, nil
}

// PostGetIDGet implements GET /post/get/{id} operation.
//
// GET /post/get/{id}
func (r *Resolver) PostGetIDGet(ctx context.Context, params api.PostGetIDGetParams) (api.PostGetIDGetRes, error) {
	return &api.PostGetIDGetServiceUnavailable{}, nil
}

// PostUpdatePut implements PUT /post/update operation.
//
// PUT /post/update
func (r *Resolver) PostUpdatePut(ctx context.Context, req api.OptPostUpdatePutReq) (api.PostUpdatePutRes, error) {
	return &api.PostUpdatePutServiceUnavailable{}, nil
}
