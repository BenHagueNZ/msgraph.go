// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ConnectionInfoRequestBuilder is request builder for ConnectionInfo
type ConnectionInfoRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConnectionInfoRequest
func (b *ConnectionInfoRequestBuilder) Request() *ConnectionInfoRequest {
	return &ConnectionInfoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConnectionInfoRequest is request for ConnectionInfo
type ConnectionInfoRequest struct{ BaseRequest }

// Get performs GET request for ConnectionInfo
func (r *ConnectionInfoRequest) Get(ctx context.Context) (resObj *ConnectionInfo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConnectionInfo
func (r *ConnectionInfoRequest) Update(ctx context.Context, reqObj *ConnectionInfo) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConnectionInfo
func (r *ConnectionInfoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}