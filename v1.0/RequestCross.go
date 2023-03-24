// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// CrossTenantAccessPolicyRequestBuilder is request builder for CrossTenantAccessPolicy
type CrossTenantAccessPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns CrossTenantAccessPolicyRequest
func (b *CrossTenantAccessPolicyRequestBuilder) Request() *CrossTenantAccessPolicyRequest {
	return &CrossTenantAccessPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CrossTenantAccessPolicyRequest is request for CrossTenantAccessPolicy
type CrossTenantAccessPolicyRequest struct{ BaseRequest }

// Get performs GET request for CrossTenantAccessPolicy
func (r *CrossTenantAccessPolicyRequest) Get(ctx context.Context) (resObj *CrossTenantAccessPolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CrossTenantAccessPolicy
func (r *CrossTenantAccessPolicyRequest) Update(ctx context.Context, reqObj *CrossTenantAccessPolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CrossTenantAccessPolicy
func (r *CrossTenantAccessPolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CrossTenantAccessPolicyConfigurationDefaultRequestBuilder is request builder for CrossTenantAccessPolicyConfigurationDefault
type CrossTenantAccessPolicyConfigurationDefaultRequestBuilder struct{ BaseRequestBuilder }

// Request returns CrossTenantAccessPolicyConfigurationDefaultRequest
func (b *CrossTenantAccessPolicyConfigurationDefaultRequestBuilder) Request() *CrossTenantAccessPolicyConfigurationDefaultRequest {
	return &CrossTenantAccessPolicyConfigurationDefaultRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CrossTenantAccessPolicyConfigurationDefaultRequest is request for CrossTenantAccessPolicyConfigurationDefault
type CrossTenantAccessPolicyConfigurationDefaultRequest struct{ BaseRequest }

// Get performs GET request for CrossTenantAccessPolicyConfigurationDefault
func (r *CrossTenantAccessPolicyConfigurationDefaultRequest) Get(ctx context.Context) (resObj *CrossTenantAccessPolicyConfigurationDefault, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CrossTenantAccessPolicyConfigurationDefault
func (r *CrossTenantAccessPolicyConfigurationDefaultRequest) Update(ctx context.Context, reqObj *CrossTenantAccessPolicyConfigurationDefault) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CrossTenantAccessPolicyConfigurationDefault
func (r *CrossTenantAccessPolicyConfigurationDefaultRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CrossTenantAccessPolicyConfigurationPartnerRequestBuilder is request builder for CrossTenantAccessPolicyConfigurationPartner
type CrossTenantAccessPolicyConfigurationPartnerRequestBuilder struct{ BaseRequestBuilder }

// Request returns CrossTenantAccessPolicyConfigurationPartnerRequest
func (b *CrossTenantAccessPolicyConfigurationPartnerRequestBuilder) Request() *CrossTenantAccessPolicyConfigurationPartnerRequest {
	return &CrossTenantAccessPolicyConfigurationPartnerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CrossTenantAccessPolicyConfigurationPartnerRequest is request for CrossTenantAccessPolicyConfigurationPartner
type CrossTenantAccessPolicyConfigurationPartnerRequest struct{ BaseRequest }

// Get performs GET request for CrossTenantAccessPolicyConfigurationPartner
func (r *CrossTenantAccessPolicyConfigurationPartnerRequest) Get(ctx context.Context) (resObj *CrossTenantAccessPolicyConfigurationPartner, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CrossTenantAccessPolicyConfigurationPartner
func (r *CrossTenantAccessPolicyConfigurationPartnerRequest) Update(ctx context.Context, reqObj *CrossTenantAccessPolicyConfigurationPartner) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CrossTenantAccessPolicyConfigurationPartner
func (r *CrossTenantAccessPolicyConfigurationPartnerRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
