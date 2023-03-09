// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// DeleteUserFromSharedAppleDeviceActionResultRequestBuilder is request builder for DeleteUserFromSharedAppleDeviceActionResult
type DeleteUserFromSharedAppleDeviceActionResultRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeleteUserFromSharedAppleDeviceActionResultRequest
func (b *DeleteUserFromSharedAppleDeviceActionResultRequestBuilder) Request() *DeleteUserFromSharedAppleDeviceActionResultRequest {
	return &DeleteUserFromSharedAppleDeviceActionResultRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeleteUserFromSharedAppleDeviceActionResultRequest is request for DeleteUserFromSharedAppleDeviceActionResult
type DeleteUserFromSharedAppleDeviceActionResultRequest struct{ BaseRequest }

// Get performs GET request for DeleteUserFromSharedAppleDeviceActionResult
func (r *DeleteUserFromSharedAppleDeviceActionResultRequest) Get(ctx context.Context) (resObj *DeleteUserFromSharedAppleDeviceActionResult, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeleteUserFromSharedAppleDeviceActionResult
func (r *DeleteUserFromSharedAppleDeviceActionResultRequest) Update(ctx context.Context, reqObj *DeleteUserFromSharedAppleDeviceActionResult) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeleteUserFromSharedAppleDeviceActionResult
func (r *DeleteUserFromSharedAppleDeviceActionResultRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
