// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// CrossCloudAzureActiveDirectoryTenantRequestBuilder is request builder for CrossCloudAzureActiveDirectoryTenant
type CrossCloudAzureActiveDirectoryTenantRequestBuilder struct{ BaseRequestBuilder }

// Request returns CrossCloudAzureActiveDirectoryTenantRequest
func (b *CrossCloudAzureActiveDirectoryTenantRequestBuilder) Request() *CrossCloudAzureActiveDirectoryTenantRequest {
	return &CrossCloudAzureActiveDirectoryTenantRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CrossCloudAzureActiveDirectoryTenantRequest is request for CrossCloudAzureActiveDirectoryTenant
type CrossCloudAzureActiveDirectoryTenantRequest struct{ BaseRequest }

// Get performs GET request for CrossCloudAzureActiveDirectoryTenant
func (r *CrossCloudAzureActiveDirectoryTenantRequest) Get(ctx context.Context) (resObj *CrossCloudAzureActiveDirectoryTenant, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CrossCloudAzureActiveDirectoryTenant
func (r *CrossCloudAzureActiveDirectoryTenantRequest) Update(ctx context.Context, reqObj *CrossCloudAzureActiveDirectoryTenant) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CrossCloudAzureActiveDirectoryTenant
func (r *CrossCloudAzureActiveDirectoryTenantRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

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

// CrossTenantAccessPolicyB2BSettingRequestBuilder is request builder for CrossTenantAccessPolicyB2BSetting
type CrossTenantAccessPolicyB2BSettingRequestBuilder struct{ BaseRequestBuilder }

// Request returns CrossTenantAccessPolicyB2BSettingRequest
func (b *CrossTenantAccessPolicyB2BSettingRequestBuilder) Request() *CrossTenantAccessPolicyB2BSettingRequest {
	return &CrossTenantAccessPolicyB2BSettingRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CrossTenantAccessPolicyB2BSettingRequest is request for CrossTenantAccessPolicyB2BSetting
type CrossTenantAccessPolicyB2BSettingRequest struct{ BaseRequest }

// Get performs GET request for CrossTenantAccessPolicyB2BSetting
func (r *CrossTenantAccessPolicyB2BSettingRequest) Get(ctx context.Context) (resObj *CrossTenantAccessPolicyB2BSetting, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CrossTenantAccessPolicyB2BSetting
func (r *CrossTenantAccessPolicyB2BSettingRequest) Update(ctx context.Context, reqObj *CrossTenantAccessPolicyB2BSetting) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CrossTenantAccessPolicyB2BSetting
func (r *CrossTenantAccessPolicyB2BSettingRequest) Delete(ctx context.Context) error {
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

// CrossTenantAccessPolicyInboundTrustRequestBuilder is request builder for CrossTenantAccessPolicyInboundTrust
type CrossTenantAccessPolicyInboundTrustRequestBuilder struct{ BaseRequestBuilder }

// Request returns CrossTenantAccessPolicyInboundTrustRequest
func (b *CrossTenantAccessPolicyInboundTrustRequestBuilder) Request() *CrossTenantAccessPolicyInboundTrustRequest {
	return &CrossTenantAccessPolicyInboundTrustRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CrossTenantAccessPolicyInboundTrustRequest is request for CrossTenantAccessPolicyInboundTrust
type CrossTenantAccessPolicyInboundTrustRequest struct{ BaseRequest }

// Get performs GET request for CrossTenantAccessPolicyInboundTrust
func (r *CrossTenantAccessPolicyInboundTrustRequest) Get(ctx context.Context) (resObj *CrossTenantAccessPolicyInboundTrust, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CrossTenantAccessPolicyInboundTrust
func (r *CrossTenantAccessPolicyInboundTrustRequest) Update(ctx context.Context, reqObj *CrossTenantAccessPolicyInboundTrust) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CrossTenantAccessPolicyInboundTrust
func (r *CrossTenantAccessPolicyInboundTrustRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CrossTenantAccessPolicyTargetRequestBuilder is request builder for CrossTenantAccessPolicyTarget
type CrossTenantAccessPolicyTargetRequestBuilder struct{ BaseRequestBuilder }

// Request returns CrossTenantAccessPolicyTargetRequest
func (b *CrossTenantAccessPolicyTargetRequestBuilder) Request() *CrossTenantAccessPolicyTargetRequest {
	return &CrossTenantAccessPolicyTargetRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CrossTenantAccessPolicyTargetRequest is request for CrossTenantAccessPolicyTarget
type CrossTenantAccessPolicyTargetRequest struct{ BaseRequest }

// Get performs GET request for CrossTenantAccessPolicyTarget
func (r *CrossTenantAccessPolicyTargetRequest) Get(ctx context.Context) (resObj *CrossTenantAccessPolicyTarget, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CrossTenantAccessPolicyTarget
func (r *CrossTenantAccessPolicyTargetRequest) Update(ctx context.Context, reqObj *CrossTenantAccessPolicyTarget) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CrossTenantAccessPolicyTarget
func (r *CrossTenantAccessPolicyTargetRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CrossTenantAccessPolicyTargetConfigurationRequestBuilder is request builder for CrossTenantAccessPolicyTargetConfiguration
type CrossTenantAccessPolicyTargetConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns CrossTenantAccessPolicyTargetConfigurationRequest
func (b *CrossTenantAccessPolicyTargetConfigurationRequestBuilder) Request() *CrossTenantAccessPolicyTargetConfigurationRequest {
	return &CrossTenantAccessPolicyTargetConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CrossTenantAccessPolicyTargetConfigurationRequest is request for CrossTenantAccessPolicyTargetConfiguration
type CrossTenantAccessPolicyTargetConfigurationRequest struct{ BaseRequest }

// Get performs GET request for CrossTenantAccessPolicyTargetConfiguration
func (r *CrossTenantAccessPolicyTargetConfigurationRequest) Get(ctx context.Context) (resObj *CrossTenantAccessPolicyTargetConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CrossTenantAccessPolicyTargetConfiguration
func (r *CrossTenantAccessPolicyTargetConfigurationRequest) Update(ctx context.Context, reqObj *CrossTenantAccessPolicyTargetConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CrossTenantAccessPolicyTargetConfiguration
func (r *CrossTenantAccessPolicyTargetConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type CrossTenantAccessPolicyConfigurationDefaultResetToSystemDefaultRequestBuilder struct{ BaseRequestBuilder }

// ResetToSystemDefault action undocumented
func (b *CrossTenantAccessPolicyConfigurationDefaultRequestBuilder) ResetToSystemDefault(reqObj *CrossTenantAccessPolicyConfigurationDefaultResetToSystemDefaultRequestParameter) *CrossTenantAccessPolicyConfigurationDefaultResetToSystemDefaultRequestBuilder {
	bb := &CrossTenantAccessPolicyConfigurationDefaultResetToSystemDefaultRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/resetToSystemDefault"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type CrossTenantAccessPolicyConfigurationDefaultResetToSystemDefaultRequest struct{ BaseRequest }

func (b *CrossTenantAccessPolicyConfigurationDefaultResetToSystemDefaultRequestBuilder) Request() *CrossTenantAccessPolicyConfigurationDefaultResetToSystemDefaultRequest {
	return &CrossTenantAccessPolicyConfigurationDefaultResetToSystemDefaultRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *CrossTenantAccessPolicyConfigurationDefaultResetToSystemDefaultRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
