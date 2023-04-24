// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// VersionActionRequestBuilder is request builder for VersionAction
type VersionActionRequestBuilder struct{ BaseRequestBuilder }

// Request returns VersionActionRequest
func (b *VersionActionRequestBuilder) Request() *VersionActionRequest {
	return &VersionActionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// VersionActionRequest is request for VersionAction
type VersionActionRequest struct{ BaseRequest }

// Get performs GET request for VersionAction
func (r *VersionActionRequest) Get(ctx context.Context) (resObj *VersionAction, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for VersionAction
func (r *VersionActionRequest) Update(ctx context.Context, reqObj *VersionAction) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for VersionAction
func (r *VersionActionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
