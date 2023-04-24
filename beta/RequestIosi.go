// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// IosiPadOSWebClipRequestBuilder is request builder for IosiPadOSWebClip
type IosiPadOSWebClipRequestBuilder struct{ BaseRequestBuilder }

// Request returns IosiPadOSWebClipRequest
func (b *IosiPadOSWebClipRequestBuilder) Request() *IosiPadOSWebClipRequest {
	return &IosiPadOSWebClipRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IosiPadOSWebClipRequest is request for IosiPadOSWebClip
type IosiPadOSWebClipRequest struct{ BaseRequest }

// Get performs GET request for IosiPadOSWebClip
func (r *IosiPadOSWebClipRequest) Get(ctx context.Context) (resObj *IosiPadOSWebClip, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IosiPadOSWebClip
func (r *IosiPadOSWebClipRequest) Update(ctx context.Context, reqObj *IosiPadOSWebClip) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IosiPadOSWebClip
func (r *IosiPadOSWebClipRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}