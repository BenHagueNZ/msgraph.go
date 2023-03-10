// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ProvisionedIdentityRequestBuilder is request builder for ProvisionedIdentity
type ProvisionedIdentityRequestBuilder struct{ BaseRequestBuilder }

// Request returns ProvisionedIdentityRequest
func (b *ProvisionedIdentityRequestBuilder) Request() *ProvisionedIdentityRequest {
	return &ProvisionedIdentityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ProvisionedIdentityRequest is request for ProvisionedIdentity
type ProvisionedIdentityRequest struct{ BaseRequest }

// Get performs GET request for ProvisionedIdentity
func (r *ProvisionedIdentityRequest) Get(ctx context.Context) (resObj *ProvisionedIdentity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ProvisionedIdentity
func (r *ProvisionedIdentityRequest) Update(ctx context.Context, reqObj *ProvisionedIdentity) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ProvisionedIdentity
func (r *ProvisionedIdentityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ProvisionedPlanRequestBuilder is request builder for ProvisionedPlan
type ProvisionedPlanRequestBuilder struct{ BaseRequestBuilder }

// Request returns ProvisionedPlanRequest
func (b *ProvisionedPlanRequestBuilder) Request() *ProvisionedPlanRequest {
	return &ProvisionedPlanRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ProvisionedPlanRequest is request for ProvisionedPlan
type ProvisionedPlanRequest struct{ BaseRequest }

// Get performs GET request for ProvisionedPlan
func (r *ProvisionedPlanRequest) Get(ctx context.Context) (resObj *ProvisionedPlan, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ProvisionedPlan
func (r *ProvisionedPlanRequest) Update(ctx context.Context, reqObj *ProvisionedPlan) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ProvisionedPlan
func (r *ProvisionedPlanRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
