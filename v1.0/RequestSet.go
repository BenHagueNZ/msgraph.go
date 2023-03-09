// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// SetRequestBuilder is request builder for Set
type SetRequestBuilder struct{ BaseRequestBuilder }

// Request returns SetRequest
func (b *SetRequestBuilder) Request() *SetRequest {
	return &SetRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SetRequest is request for Set
type SetRequest struct{ BaseRequest }

// Get performs GET request for Set
func (r *SetRequest) Get(ctx context.Context) (resObj *Set, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Set
func (r *SetRequest) Update(ctx context.Context, reqObj *Set) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Set
func (r *SetRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
