// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// OrganizationalBrandingRequestBuilder is request builder for OrganizationalBranding
type OrganizationalBrandingRequestBuilder struct{ BaseRequestBuilder }

// Request returns OrganizationalBrandingRequest
func (b *OrganizationalBrandingRequestBuilder) Request() *OrganizationalBrandingRequest {
	return &OrganizationalBrandingRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OrganizationalBrandingRequest is request for OrganizationalBranding
type OrganizationalBrandingRequest struct{ BaseRequest }

// Get performs GET request for OrganizationalBranding
func (r *OrganizationalBrandingRequest) Get(ctx context.Context) (resObj *OrganizationalBranding, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OrganizationalBranding
func (r *OrganizationalBrandingRequest) Update(ctx context.Context, reqObj *OrganizationalBranding) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OrganizationalBranding
func (r *OrganizationalBrandingRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OrganizationalBrandingLocalizationRequestBuilder is request builder for OrganizationalBrandingLocalization
type OrganizationalBrandingLocalizationRequestBuilder struct{ BaseRequestBuilder }

// Request returns OrganizationalBrandingLocalizationRequest
func (b *OrganizationalBrandingLocalizationRequestBuilder) Request() *OrganizationalBrandingLocalizationRequest {
	return &OrganizationalBrandingLocalizationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OrganizationalBrandingLocalizationRequest is request for OrganizationalBrandingLocalization
type OrganizationalBrandingLocalizationRequest struct{ BaseRequest }

// Get performs GET request for OrganizationalBrandingLocalization
func (r *OrganizationalBrandingLocalizationRequest) Get(ctx context.Context) (resObj *OrganizationalBrandingLocalization, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OrganizationalBrandingLocalization
func (r *OrganizationalBrandingLocalizationRequest) Update(ctx context.Context, reqObj *OrganizationalBrandingLocalization) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OrganizationalBrandingLocalization
func (r *OrganizationalBrandingLocalizationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OrganizationalBrandingPropertiesRequestBuilder is request builder for OrganizationalBrandingProperties
type OrganizationalBrandingPropertiesRequestBuilder struct{ BaseRequestBuilder }

// Request returns OrganizationalBrandingPropertiesRequest
func (b *OrganizationalBrandingPropertiesRequestBuilder) Request() *OrganizationalBrandingPropertiesRequest {
	return &OrganizationalBrandingPropertiesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OrganizationalBrandingPropertiesRequest is request for OrganizationalBrandingProperties
type OrganizationalBrandingPropertiesRequest struct{ BaseRequest }

// Get performs GET request for OrganizationalBrandingProperties
func (r *OrganizationalBrandingPropertiesRequest) Get(ctx context.Context) (resObj *OrganizationalBrandingProperties, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OrganizationalBrandingProperties
func (r *OrganizationalBrandingPropertiesRequest) Update(ctx context.Context, reqObj *OrganizationalBrandingProperties) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OrganizationalBrandingProperties
func (r *OrganizationalBrandingPropertiesRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
