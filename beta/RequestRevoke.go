// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// RevokeAppleVPPLicensesActionResultRequestBuilder is request builder for RevokeAppleVPPLicensesActionResult
type RevokeAppleVPPLicensesActionResultRequestBuilder struct{ BaseRequestBuilder }

// Request returns RevokeAppleVPPLicensesActionResultRequest
func (b *RevokeAppleVPPLicensesActionResultRequestBuilder) Request() *RevokeAppleVPPLicensesActionResultRequest {
	return &RevokeAppleVPPLicensesActionResultRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RevokeAppleVPPLicensesActionResultRequest is request for RevokeAppleVPPLicensesActionResult
type RevokeAppleVPPLicensesActionResultRequest struct{ BaseRequest }

// Get performs GET request for RevokeAppleVPPLicensesActionResult
func (r *RevokeAppleVPPLicensesActionResultRequest) Get(ctx context.Context) (resObj *RevokeAppleVPPLicensesActionResult, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RevokeAppleVPPLicensesActionResult
func (r *RevokeAppleVPPLicensesActionResultRequest) Update(ctx context.Context, reqObj *RevokeAppleVPPLicensesActionResult) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RevokeAppleVPPLicensesActionResult
func (r *RevokeAppleVPPLicensesActionResultRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
