// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// MFADetailRequestBuilder is request builder for MFADetail
type MFADetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns MFADetailRequest
func (b *MFADetailRequestBuilder) Request() *MFADetailRequest {
	return &MFADetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MFADetailRequest is request for MFADetail
type MFADetailRequest struct{ BaseRequest }

// Get performs GET request for MFADetail
func (r *MFADetailRequest) Get(ctx context.Context) (resObj *MFADetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MFADetail
func (r *MFADetailRequest) Update(ctx context.Context, reqObj *MFADetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MFADetail
func (r *MFADetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
