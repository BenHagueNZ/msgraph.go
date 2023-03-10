// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// RecordingInfoRequestBuilder is request builder for RecordingInfo
type RecordingInfoRequestBuilder struct{ BaseRequestBuilder }

// Request returns RecordingInfoRequest
func (b *RecordingInfoRequestBuilder) Request() *RecordingInfoRequest {
	return &RecordingInfoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RecordingInfoRequest is request for RecordingInfo
type RecordingInfoRequest struct{ BaseRequest }

// Get performs GET request for RecordingInfo
func (r *RecordingInfoRequest) Get(ctx context.Context) (resObj *RecordingInfo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RecordingInfo
func (r *RecordingInfoRequest) Update(ctx context.Context, reqObj *RecordingInfo) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RecordingInfo
func (r *RecordingInfoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
