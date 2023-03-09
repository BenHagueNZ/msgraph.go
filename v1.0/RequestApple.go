// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// AppleDeviceFeaturesConfigurationBaseRequestBuilder is request builder for AppleDeviceFeaturesConfigurationBase
type AppleDeviceFeaturesConfigurationBaseRequestBuilder struct{ BaseRequestBuilder }

// Request returns AppleDeviceFeaturesConfigurationBaseRequest
func (b *AppleDeviceFeaturesConfigurationBaseRequestBuilder) Request() *AppleDeviceFeaturesConfigurationBaseRequest {
	return &AppleDeviceFeaturesConfigurationBaseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AppleDeviceFeaturesConfigurationBaseRequest is request for AppleDeviceFeaturesConfigurationBase
type AppleDeviceFeaturesConfigurationBaseRequest struct{ BaseRequest }

// Get performs GET request for AppleDeviceFeaturesConfigurationBase
func (r *AppleDeviceFeaturesConfigurationBaseRequest) Get(ctx context.Context) (resObj *AppleDeviceFeaturesConfigurationBase, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AppleDeviceFeaturesConfigurationBase
func (r *AppleDeviceFeaturesConfigurationBaseRequest) Update(ctx context.Context, reqObj *AppleDeviceFeaturesConfigurationBase) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AppleDeviceFeaturesConfigurationBase
func (r *AppleDeviceFeaturesConfigurationBaseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AppleManagedIdentityProviderRequestBuilder is request builder for AppleManagedIdentityProvider
type AppleManagedIdentityProviderRequestBuilder struct{ BaseRequestBuilder }

// Request returns AppleManagedIdentityProviderRequest
func (b *AppleManagedIdentityProviderRequestBuilder) Request() *AppleManagedIdentityProviderRequest {
	return &AppleManagedIdentityProviderRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AppleManagedIdentityProviderRequest is request for AppleManagedIdentityProvider
type AppleManagedIdentityProviderRequest struct{ BaseRequest }

// Get performs GET request for AppleManagedIdentityProvider
func (r *AppleManagedIdentityProviderRequest) Get(ctx context.Context) (resObj *AppleManagedIdentityProvider, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AppleManagedIdentityProvider
func (r *AppleManagedIdentityProviderRequest) Update(ctx context.Context, reqObj *AppleManagedIdentityProvider) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AppleManagedIdentityProvider
func (r *AppleManagedIdentityProviderRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ApplePushNotificationCertificateRequestBuilder is request builder for ApplePushNotificationCertificate
type ApplePushNotificationCertificateRequestBuilder struct{ BaseRequestBuilder }

// Request returns ApplePushNotificationCertificateRequest
func (b *ApplePushNotificationCertificateRequestBuilder) Request() *ApplePushNotificationCertificateRequest {
	return &ApplePushNotificationCertificateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ApplePushNotificationCertificateRequest is request for ApplePushNotificationCertificate
type ApplePushNotificationCertificateRequest struct{ BaseRequest }

// Get performs GET request for ApplePushNotificationCertificate
func (r *ApplePushNotificationCertificateRequest) Get(ctx context.Context) (resObj *ApplePushNotificationCertificate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ApplePushNotificationCertificate
func (r *ApplePushNotificationCertificateRequest) Update(ctx context.Context, reqObj *ApplePushNotificationCertificate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ApplePushNotificationCertificate
func (r *ApplePushNotificationCertificateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
