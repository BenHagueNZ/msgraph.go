// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// LookupColumnRequestBuilder is request builder for LookupColumn
type LookupColumnRequestBuilder struct{ BaseRequestBuilder }

// Request returns LookupColumnRequest
func (b *LookupColumnRequestBuilder) Request() *LookupColumnRequest {
	return &LookupColumnRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// LookupColumnRequest is request for LookupColumn
type LookupColumnRequest struct{ BaseRequest }

// Get performs GET request for LookupColumn
func (r *LookupColumnRequest) Get(ctx context.Context) (resObj *LookupColumn, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for LookupColumn
func (r *LookupColumnRequest) Update(ctx context.Context, reqObj *LookupColumn) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for LookupColumn
func (r *LookupColumnRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
