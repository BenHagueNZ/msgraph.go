// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ExtendRemoteHelpSessionResponseRequestBuilder is request builder for ExtendRemoteHelpSessionResponse
type ExtendRemoteHelpSessionResponseRequestBuilder struct{ BaseRequestBuilder }

// Request returns ExtendRemoteHelpSessionResponseRequest
func (b *ExtendRemoteHelpSessionResponseRequestBuilder) Request() *ExtendRemoteHelpSessionResponseRequest {
	return &ExtendRemoteHelpSessionResponseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ExtendRemoteHelpSessionResponseRequest is request for ExtendRemoteHelpSessionResponse
type ExtendRemoteHelpSessionResponseRequest struct{ BaseRequest }

// Get performs GET request for ExtendRemoteHelpSessionResponse
func (r *ExtendRemoteHelpSessionResponseRequest) Get(ctx context.Context) (resObj *ExtendRemoteHelpSessionResponse, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ExtendRemoteHelpSessionResponse
func (r *ExtendRemoteHelpSessionResponseRequest) Update(ctx context.Context, reqObj *ExtendRemoteHelpSessionResponse) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ExtendRemoteHelpSessionResponse
func (r *ExtendRemoteHelpSessionResponseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
