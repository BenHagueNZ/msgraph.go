// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// DowngradeJustificationRequestBuilder is request builder for DowngradeJustification
type DowngradeJustificationRequestBuilder struct{ BaseRequestBuilder }

// Request returns DowngradeJustificationRequest
func (b *DowngradeJustificationRequestBuilder) Request() *DowngradeJustificationRequest {
	return &DowngradeJustificationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DowngradeJustificationRequest is request for DowngradeJustification
type DowngradeJustificationRequest struct{ BaseRequest }

// Get performs GET request for DowngradeJustification
func (r *DowngradeJustificationRequest) Get(ctx context.Context) (resObj *DowngradeJustification, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DowngradeJustification
func (r *DowngradeJustificationRequest) Update(ctx context.Context, reqObj *DowngradeJustification) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DowngradeJustification
func (r *DowngradeJustificationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
