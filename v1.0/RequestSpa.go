// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// SpaApplicationRequestBuilder is request builder for SpaApplication
type SpaApplicationRequestBuilder struct{ BaseRequestBuilder }

// Request returns SpaApplicationRequest
func (b *SpaApplicationRequestBuilder) Request() *SpaApplicationRequest {
	return &SpaApplicationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SpaApplicationRequest is request for SpaApplication
type SpaApplicationRequest struct{ BaseRequest }

// Get performs GET request for SpaApplication
func (r *SpaApplicationRequest) Get(ctx context.Context) (resObj *SpaApplication, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SpaApplication
func (r *SpaApplicationRequest) Update(ctx context.Context, reqObj *SpaApplication) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SpaApplication
func (r *SpaApplicationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
