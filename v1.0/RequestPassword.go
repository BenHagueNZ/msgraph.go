// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// PasswordAuthenticationMethodRequestBuilder is request builder for PasswordAuthenticationMethod
type PasswordAuthenticationMethodRequestBuilder struct{ BaseRequestBuilder }

// Request returns PasswordAuthenticationMethodRequest
func (b *PasswordAuthenticationMethodRequestBuilder) Request() *PasswordAuthenticationMethodRequest {
	return &PasswordAuthenticationMethodRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PasswordAuthenticationMethodRequest is request for PasswordAuthenticationMethod
type PasswordAuthenticationMethodRequest struct{ BaseRequest }

// Get performs GET request for PasswordAuthenticationMethod
func (r *PasswordAuthenticationMethodRequest) Get(ctx context.Context) (resObj *PasswordAuthenticationMethod, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PasswordAuthenticationMethod
func (r *PasswordAuthenticationMethodRequest) Update(ctx context.Context, reqObj *PasswordAuthenticationMethod) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PasswordAuthenticationMethod
func (r *PasswordAuthenticationMethodRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PasswordCredentialRequestBuilder is request builder for PasswordCredential
type PasswordCredentialRequestBuilder struct{ BaseRequestBuilder }

// Request returns PasswordCredentialRequest
func (b *PasswordCredentialRequestBuilder) Request() *PasswordCredentialRequest {
	return &PasswordCredentialRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PasswordCredentialRequest is request for PasswordCredential
type PasswordCredentialRequest struct{ BaseRequest }

// Get performs GET request for PasswordCredential
func (r *PasswordCredentialRequest) Get(ctx context.Context) (resObj *PasswordCredential, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PasswordCredential
func (r *PasswordCredentialRequest) Update(ctx context.Context, reqObj *PasswordCredential) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PasswordCredential
func (r *PasswordCredentialRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PasswordCredentialConfigurationRequestBuilder is request builder for PasswordCredentialConfiguration
type PasswordCredentialConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns PasswordCredentialConfigurationRequest
func (b *PasswordCredentialConfigurationRequestBuilder) Request() *PasswordCredentialConfigurationRequest {
	return &PasswordCredentialConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PasswordCredentialConfigurationRequest is request for PasswordCredentialConfiguration
type PasswordCredentialConfigurationRequest struct{ BaseRequest }

// Get performs GET request for PasswordCredentialConfiguration
func (r *PasswordCredentialConfigurationRequest) Get(ctx context.Context) (resObj *PasswordCredentialConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PasswordCredentialConfiguration
func (r *PasswordCredentialConfigurationRequest) Update(ctx context.Context, reqObj *PasswordCredentialConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PasswordCredentialConfiguration
func (r *PasswordCredentialConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PasswordProfileRequestBuilder is request builder for PasswordProfile
type PasswordProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns PasswordProfileRequest
func (b *PasswordProfileRequestBuilder) Request() *PasswordProfileRequest {
	return &PasswordProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PasswordProfileRequest is request for PasswordProfile
type PasswordProfileRequest struct{ BaseRequest }

// Get performs GET request for PasswordProfile
func (r *PasswordProfileRequest) Get(ctx context.Context) (resObj *PasswordProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PasswordProfile
func (r *PasswordProfileRequest) Update(ctx context.Context, reqObj *PasswordProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PasswordProfile
func (r *PasswordProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PasswordResetResponseRequestBuilder is request builder for PasswordResetResponse
type PasswordResetResponseRequestBuilder struct{ BaseRequestBuilder }

// Request returns PasswordResetResponseRequest
func (b *PasswordResetResponseRequestBuilder) Request() *PasswordResetResponseRequest {
	return &PasswordResetResponseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PasswordResetResponseRequest is request for PasswordResetResponse
type PasswordResetResponseRequest struct{ BaseRequest }

// Get performs GET request for PasswordResetResponse
func (r *PasswordResetResponseRequest) Get(ctx context.Context) (resObj *PasswordResetResponse, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PasswordResetResponse
func (r *PasswordResetResponseRequest) Update(ctx context.Context, reqObj *PasswordResetResponse) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PasswordResetResponse
func (r *PasswordResetResponseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
