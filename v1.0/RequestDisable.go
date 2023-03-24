// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// DisableAndDeleteUserApplyActionRequestBuilder is request builder for DisableAndDeleteUserApplyAction
type DisableAndDeleteUserApplyActionRequestBuilder struct{ BaseRequestBuilder }

// Request returns DisableAndDeleteUserApplyActionRequest
func (b *DisableAndDeleteUserApplyActionRequestBuilder) Request() *DisableAndDeleteUserApplyActionRequest {
	return &DisableAndDeleteUserApplyActionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DisableAndDeleteUserApplyActionRequest is request for DisableAndDeleteUserApplyAction
type DisableAndDeleteUserApplyActionRequest struct{ BaseRequest }

// Get performs GET request for DisableAndDeleteUserApplyAction
func (r *DisableAndDeleteUserApplyActionRequest) Get(ctx context.Context) (resObj *DisableAndDeleteUserApplyAction, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DisableAndDeleteUserApplyAction
func (r *DisableAndDeleteUserApplyActionRequest) Update(ctx context.Context, reqObj *DisableAndDeleteUserApplyAction) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DisableAndDeleteUserApplyAction
func (r *DisableAndDeleteUserApplyActionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}