// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// IdleSessionSignOutRequestBuilder is request builder for IdleSessionSignOut
type IdleSessionSignOutRequestBuilder struct{ BaseRequestBuilder }

// Request returns IdleSessionSignOutRequest
func (b *IdleSessionSignOutRequestBuilder) Request() *IdleSessionSignOutRequest {
	return &IdleSessionSignOutRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IdleSessionSignOutRequest is request for IdleSessionSignOut
type IdleSessionSignOutRequest struct{ BaseRequest }

// Get performs GET request for IdleSessionSignOut
func (r *IdleSessionSignOutRequest) Get(ctx context.Context) (resObj *IdleSessionSignOut, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IdleSessionSignOut
func (r *IdleSessionSignOutRequest) Update(ctx context.Context, reqObj *IdleSessionSignOut) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IdleSessionSignOut
func (r *IdleSessionSignOutRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}