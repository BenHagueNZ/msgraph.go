// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// VoiceAuthenticationMethodConfigurationRequestBuilder is request builder for VoiceAuthenticationMethodConfiguration
type VoiceAuthenticationMethodConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns VoiceAuthenticationMethodConfigurationRequest
func (b *VoiceAuthenticationMethodConfigurationRequestBuilder) Request() *VoiceAuthenticationMethodConfigurationRequest {
	return &VoiceAuthenticationMethodConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// VoiceAuthenticationMethodConfigurationRequest is request for VoiceAuthenticationMethodConfiguration
type VoiceAuthenticationMethodConfigurationRequest struct{ BaseRequest }

// Get performs GET request for VoiceAuthenticationMethodConfiguration
func (r *VoiceAuthenticationMethodConfigurationRequest) Get(ctx context.Context) (resObj *VoiceAuthenticationMethodConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for VoiceAuthenticationMethodConfiguration
func (r *VoiceAuthenticationMethodConfigurationRequest) Update(ctx context.Context, reqObj *VoiceAuthenticationMethodConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for VoiceAuthenticationMethodConfiguration
func (r *VoiceAuthenticationMethodConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// VoiceAuthenticationMethodTargetRequestBuilder is request builder for VoiceAuthenticationMethodTarget
type VoiceAuthenticationMethodTargetRequestBuilder struct{ BaseRequestBuilder }

// Request returns VoiceAuthenticationMethodTargetRequest
func (b *VoiceAuthenticationMethodTargetRequestBuilder) Request() *VoiceAuthenticationMethodTargetRequest {
	return &VoiceAuthenticationMethodTargetRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// VoiceAuthenticationMethodTargetRequest is request for VoiceAuthenticationMethodTarget
type VoiceAuthenticationMethodTargetRequest struct{ BaseRequest }

// Get performs GET request for VoiceAuthenticationMethodTarget
func (r *VoiceAuthenticationMethodTargetRequest) Get(ctx context.Context) (resObj *VoiceAuthenticationMethodTarget, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for VoiceAuthenticationMethodTarget
func (r *VoiceAuthenticationMethodTargetRequest) Update(ctx context.Context, reqObj *VoiceAuthenticationMethodTarget) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for VoiceAuthenticationMethodTarget
func (r *VoiceAuthenticationMethodTargetRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
