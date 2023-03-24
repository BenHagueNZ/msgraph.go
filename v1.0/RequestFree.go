// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// FreeBusyErrorRequestBuilder is request builder for FreeBusyError
type FreeBusyErrorRequestBuilder struct{ BaseRequestBuilder }

// Request returns FreeBusyErrorRequest
func (b *FreeBusyErrorRequestBuilder) Request() *FreeBusyErrorRequest {
	return &FreeBusyErrorRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// FreeBusyErrorRequest is request for FreeBusyError
type FreeBusyErrorRequest struct{ BaseRequest }

// Get performs GET request for FreeBusyError
func (r *FreeBusyErrorRequest) Get(ctx context.Context) (resObj *FreeBusyError, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for FreeBusyError
func (r *FreeBusyErrorRequest) Update(ctx context.Context, reqObj *FreeBusyError) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for FreeBusyError
func (r *FreeBusyErrorRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
