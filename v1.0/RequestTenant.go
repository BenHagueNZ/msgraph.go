// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// TenantAppManagementPolicyRequestBuilder is request builder for TenantAppManagementPolicy
type TenantAppManagementPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns TenantAppManagementPolicyRequest
func (b *TenantAppManagementPolicyRequestBuilder) Request() *TenantAppManagementPolicyRequest {
	return &TenantAppManagementPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TenantAppManagementPolicyRequest is request for TenantAppManagementPolicy
type TenantAppManagementPolicyRequest struct{ BaseRequest }

// Get performs GET request for TenantAppManagementPolicy
func (r *TenantAppManagementPolicyRequest) Get(ctx context.Context) (resObj *TenantAppManagementPolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TenantAppManagementPolicy
func (r *TenantAppManagementPolicyRequest) Update(ctx context.Context, reqObj *TenantAppManagementPolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TenantAppManagementPolicy
func (r *TenantAppManagementPolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TenantRelationshipRequestBuilder is request builder for TenantRelationship
type TenantRelationshipRequestBuilder struct{ BaseRequestBuilder }

// Request returns TenantRelationshipRequest
func (b *TenantRelationshipRequestBuilder) Request() *TenantRelationshipRequest {
	return &TenantRelationshipRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TenantRelationshipRequest is request for TenantRelationship
type TenantRelationshipRequest struct{ BaseRequest }

// Get performs GET request for TenantRelationship
func (r *TenantRelationshipRequest) Get(ctx context.Context) (resObj *TenantRelationship, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TenantRelationship
func (r *TenantRelationshipRequest) Update(ctx context.Context, reqObj *TenantRelationship) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TenantRelationship
func (r *TenantRelationshipRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
