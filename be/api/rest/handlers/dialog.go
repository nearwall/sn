package handlers

import (
	"context"

	api "sn/api/rest/generated"
)

// DialogUserIDListGet implements GET /dialog/{user_id}/list operation.
//
// GET /dialog/{user_id}/list
func (r *Resolver) DialogUserIDListGet(ctx context.Context, params api.DialogUserIDListGetParams) (api.DialogUserIDListGetRes, error) {
	return nil, nil
}

// DialogUserIDSendPost implements POST /dialog/{user_id}/send operation.
//
// POST /dialog/{user_id}/send
func (r *Resolver) DialogUserIDSendPost(
	ctx context.Context,
	req api.OptDialogUserIDSendPostReq,
	params api.DialogUserIDSendPostParams) (api.DialogUserIDSendPostRes, error) {
	return nil, nil
}
