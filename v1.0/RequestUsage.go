// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// UsageDetailsRequestBuilder is request builder for UsageDetails
type UsageDetailsRequestBuilder struct{ BaseRequestBuilder }

// Request returns UsageDetailsRequest
func (b *UsageDetailsRequestBuilder) Request() *UsageDetailsRequest {
	return &UsageDetailsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UsageDetailsRequest is request for UsageDetails
type UsageDetailsRequest struct{ BaseRequest }

// Get performs GET request for UsageDetails
func (r *UsageDetailsRequest) Get(ctx context.Context) (resObj *UsageDetails, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UsageDetails
func (r *UsageDetailsRequest) Update(ctx context.Context, reqObj *UsageDetails) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UsageDetails
func (r *UsageDetailsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
