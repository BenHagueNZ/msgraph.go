// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// MobilityManagementPolicyRequestBuilder is request builder for MobilityManagementPolicy
type MobilityManagementPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobilityManagementPolicyRequest
func (b *MobilityManagementPolicyRequestBuilder) Request() *MobilityManagementPolicyRequest {
	return &MobilityManagementPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobilityManagementPolicyRequest is request for MobilityManagementPolicy
type MobilityManagementPolicyRequest struct{ BaseRequest }

// Get performs GET request for MobilityManagementPolicy
func (r *MobilityManagementPolicyRequest) Get(ctx context.Context) (resObj *MobilityManagementPolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MobilityManagementPolicy
func (r *MobilityManagementPolicyRequest) Update(ctx context.Context, reqObj *MobilityManagementPolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MobilityManagementPolicy
func (r *MobilityManagementPolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
