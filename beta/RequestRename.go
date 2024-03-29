// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// RenameActionRequestBuilder is request builder for RenameAction
type RenameActionRequestBuilder struct{ BaseRequestBuilder }

// Request returns RenameActionRequest
func (b *RenameActionRequestBuilder) Request() *RenameActionRequest {
	return &RenameActionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RenameActionRequest is request for RenameAction
type RenameActionRequest struct{ BaseRequest }

// Get performs GET request for RenameAction
func (r *RenameActionRequest) Get(ctx context.Context) (resObj *RenameAction, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RenameAction
func (r *RenameActionRequest) Update(ctx context.Context, reqObj *RenameAction) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RenameAction
func (r *RenameActionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
