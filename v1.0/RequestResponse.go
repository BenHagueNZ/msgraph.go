// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ResponseStatusRequestBuilder is request builder for ResponseStatus
type ResponseStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns ResponseStatusRequest
func (b *ResponseStatusRequestBuilder) Request() *ResponseStatusRequest {
	return &ResponseStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ResponseStatusRequest is request for ResponseStatus
type ResponseStatusRequest struct{ BaseRequest }

// Get performs GET request for ResponseStatus
func (r *ResponseStatusRequest) Get(ctx context.Context) (resObj *ResponseStatus, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ResponseStatus
func (r *ResponseStatusRequest) Update(ctx context.Context, reqObj *ResponseStatus) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ResponseStatus
func (r *ResponseStatusRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
