// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// SamlOrWsFedExternalDomainFederationRequestBuilder is request builder for SamlOrWsFedExternalDomainFederation
type SamlOrWsFedExternalDomainFederationRequestBuilder struct{ BaseRequestBuilder }

// Request returns SamlOrWsFedExternalDomainFederationRequest
func (b *SamlOrWsFedExternalDomainFederationRequestBuilder) Request() *SamlOrWsFedExternalDomainFederationRequest {
	return &SamlOrWsFedExternalDomainFederationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SamlOrWsFedExternalDomainFederationRequest is request for SamlOrWsFedExternalDomainFederation
type SamlOrWsFedExternalDomainFederationRequest struct{ BaseRequest }

// Get performs GET request for SamlOrWsFedExternalDomainFederation
func (r *SamlOrWsFedExternalDomainFederationRequest) Get(ctx context.Context) (resObj *SamlOrWsFedExternalDomainFederation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SamlOrWsFedExternalDomainFederation
func (r *SamlOrWsFedExternalDomainFederationRequest) Update(ctx context.Context, reqObj *SamlOrWsFedExternalDomainFederation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SamlOrWsFedExternalDomainFederation
func (r *SamlOrWsFedExternalDomainFederationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SamlOrWsFedProviderRequestBuilder is request builder for SamlOrWsFedProvider
type SamlOrWsFedProviderRequestBuilder struct{ BaseRequestBuilder }

// Request returns SamlOrWsFedProviderRequest
func (b *SamlOrWsFedProviderRequestBuilder) Request() *SamlOrWsFedProviderRequest {
	return &SamlOrWsFedProviderRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SamlOrWsFedProviderRequest is request for SamlOrWsFedProvider
type SamlOrWsFedProviderRequest struct{ BaseRequest }

// Get performs GET request for SamlOrWsFedProvider
func (r *SamlOrWsFedProviderRequest) Get(ctx context.Context) (resObj *SamlOrWsFedProvider, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SamlOrWsFedProvider
func (r *SamlOrWsFedProviderRequest) Update(ctx context.Context, reqObj *SamlOrWsFedProvider) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SamlOrWsFedProvider
func (r *SamlOrWsFedProviderRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SamlSingleSignOnSettingsRequestBuilder is request builder for SamlSingleSignOnSettings
type SamlSingleSignOnSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns SamlSingleSignOnSettingsRequest
func (b *SamlSingleSignOnSettingsRequestBuilder) Request() *SamlSingleSignOnSettingsRequest {
	return &SamlSingleSignOnSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SamlSingleSignOnSettingsRequest is request for SamlSingleSignOnSettings
type SamlSingleSignOnSettingsRequest struct{ BaseRequest }

// Get performs GET request for SamlSingleSignOnSettings
func (r *SamlSingleSignOnSettingsRequest) Get(ctx context.Context) (resObj *SamlSingleSignOnSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SamlSingleSignOnSettings
func (r *SamlSingleSignOnSettingsRequest) Update(ctx context.Context, reqObj *SamlSingleSignOnSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SamlSingleSignOnSettings
func (r *SamlSingleSignOnSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
