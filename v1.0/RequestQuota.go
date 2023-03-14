// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// QuotaRequestBuilder is request builder for Quota
type QuotaRequestBuilder struct{ BaseRequestBuilder }

// Request returns QuotaRequest
func (b *QuotaRequestBuilder) Request() *QuotaRequest {
	return &QuotaRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// QuotaRequest is request for Quota
type QuotaRequest struct{ BaseRequest }

// Get performs GET request for Quota
func (r *QuotaRequest) Get(ctx context.Context) (resObj *Quota, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Quota
func (r *QuotaRequest) Update(ctx context.Context, reqObj *Quota) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Quota
func (r *QuotaRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}