// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ProviderTenantSettingRequestBuilder is request builder for ProviderTenantSetting
type ProviderTenantSettingRequestBuilder struct{ BaseRequestBuilder }

// Request returns ProviderTenantSettingRequest
func (b *ProviderTenantSettingRequestBuilder) Request() *ProviderTenantSettingRequest {
	return &ProviderTenantSettingRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ProviderTenantSettingRequest is request for ProviderTenantSetting
type ProviderTenantSettingRequest struct{ BaseRequest }

// Get performs GET request for ProviderTenantSetting
func (r *ProviderTenantSettingRequest) Get(ctx context.Context) (resObj *ProviderTenantSetting, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ProviderTenantSetting
func (r *ProviderTenantSettingRequest) Update(ctx context.Context, reqObj *ProviderTenantSetting) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ProviderTenantSetting
func (r *ProviderTenantSettingRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}