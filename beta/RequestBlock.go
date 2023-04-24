// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// BlockAccessActionRequestBuilder is request builder for BlockAccessAction
type BlockAccessActionRequestBuilder struct{ BaseRequestBuilder }

// Request returns BlockAccessActionRequest
func (b *BlockAccessActionRequestBuilder) Request() *BlockAccessActionRequest {
	return &BlockAccessActionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BlockAccessActionRequest is request for BlockAccessAction
type BlockAccessActionRequest struct{ BaseRequest }

// Get performs GET request for BlockAccessAction
func (r *BlockAccessActionRequest) Get(ctx context.Context) (resObj *BlockAccessAction, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BlockAccessAction
func (r *BlockAccessActionRequest) Update(ctx context.Context, reqObj *BlockAccessAction) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BlockAccessAction
func (r *BlockAccessActionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
