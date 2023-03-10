// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ToneInfoRequestBuilder is request builder for ToneInfo
type ToneInfoRequestBuilder struct{ BaseRequestBuilder }

// Request returns ToneInfoRequest
func (b *ToneInfoRequestBuilder) Request() *ToneInfoRequest {
	return &ToneInfoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ToneInfoRequest is request for ToneInfo
type ToneInfoRequest struct{ BaseRequest }

// Get performs GET request for ToneInfo
func (r *ToneInfoRequest) Get(ctx context.Context) (resObj *ToneInfo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ToneInfo
func (r *ToneInfoRequest) Update(ctx context.Context, reqObj *ToneInfo) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ToneInfo
func (r *ToneInfoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
