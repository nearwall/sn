package handlers

import (
	"context"

	api "sn/api/rest/generated"
)

// FriendDeleteUserIDPut implements PUT /friend/delete/{user_id} operation.
//
// PUT /friend/delete/{user_id}
func (r *Resolver) FriendDeleteUserIDPut(ctx context.Context, params api.FriendDeleteUserIDPutParams) (api.FriendDeleteUserIDPutRes, error) {
	return &api.FriendDeleteUserIDPutServiceUnavailable{}, nil
}

// FriendSetUserIDPut implements PUT /friend/set/{user_id} operation.
//
// PUT /friend/set/{user_id}
func (r *Resolver) FriendSetUserIDPut(ctx context.Context, params api.FriendSetUserIDPutParams) (api.FriendSetUserIDPutRes, error) {
	return &api.FriendSetUserIDPutServiceUnavailable{}, nil
}
