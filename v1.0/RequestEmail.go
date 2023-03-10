// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// EmailAddressRequestBuilder is request builder for EmailAddress
type EmailAddressRequestBuilder struct{ BaseRequestBuilder }

// Request returns EmailAddressRequest
func (b *EmailAddressRequestBuilder) Request() *EmailAddressRequest {
	return &EmailAddressRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EmailAddressRequest is request for EmailAddress
type EmailAddressRequest struct{ BaseRequest }

// Get performs GET request for EmailAddress
func (r *EmailAddressRequest) Get(ctx context.Context) (resObj *EmailAddress, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EmailAddress
func (r *EmailAddressRequest) Update(ctx context.Context, reqObj *EmailAddress) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EmailAddress
func (r *EmailAddressRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// EmailAuthenticationMethodRequestBuilder is request builder for EmailAuthenticationMethod
type EmailAuthenticationMethodRequestBuilder struct{ BaseRequestBuilder }

// Request returns EmailAuthenticationMethodRequest
func (b *EmailAuthenticationMethodRequestBuilder) Request() *EmailAuthenticationMethodRequest {
	return &EmailAuthenticationMethodRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EmailAuthenticationMethodRequest is request for EmailAuthenticationMethod
type EmailAuthenticationMethodRequest struct{ BaseRequest }

// Get performs GET request for EmailAuthenticationMethod
func (r *EmailAuthenticationMethodRequest) Get(ctx context.Context) (resObj *EmailAuthenticationMethod, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EmailAuthenticationMethod
func (r *EmailAuthenticationMethodRequest) Update(ctx context.Context, reqObj *EmailAuthenticationMethod) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EmailAuthenticationMethod
func (r *EmailAuthenticationMethodRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// EmailAuthenticationMethodConfigurationRequestBuilder is request builder for EmailAuthenticationMethodConfiguration
type EmailAuthenticationMethodConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns EmailAuthenticationMethodConfigurationRequest
func (b *EmailAuthenticationMethodConfigurationRequestBuilder) Request() *EmailAuthenticationMethodConfigurationRequest {
	return &EmailAuthenticationMethodConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EmailAuthenticationMethodConfigurationRequest is request for EmailAuthenticationMethodConfiguration
type EmailAuthenticationMethodConfigurationRequest struct{ BaseRequest }

// Get performs GET request for EmailAuthenticationMethodConfiguration
func (r *EmailAuthenticationMethodConfigurationRequest) Get(ctx context.Context) (resObj *EmailAuthenticationMethodConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EmailAuthenticationMethodConfiguration
func (r *EmailAuthenticationMethodConfigurationRequest) Update(ctx context.Context, reqObj *EmailAuthenticationMethodConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EmailAuthenticationMethodConfiguration
func (r *EmailAuthenticationMethodConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// EmailFileAssessmentRequestObjectRequestBuilder is request builder for EmailFileAssessmentRequestObject
type EmailFileAssessmentRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns EmailFileAssessmentRequestObjectRequest
func (b *EmailFileAssessmentRequestObjectRequestBuilder) Request() *EmailFileAssessmentRequestObjectRequest {
	return &EmailFileAssessmentRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EmailFileAssessmentRequestObjectRequest is request for EmailFileAssessmentRequestObject
type EmailFileAssessmentRequestObjectRequest struct{ BaseRequest }

// Get performs GET request for EmailFileAssessmentRequestObject
func (r *EmailFileAssessmentRequestObjectRequest) Get(ctx context.Context) (resObj *EmailFileAssessmentRequestObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EmailFileAssessmentRequestObject
func (r *EmailFileAssessmentRequestObjectRequest) Update(ctx context.Context, reqObj *EmailFileAssessmentRequestObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EmailFileAssessmentRequestObject
func (r *EmailFileAssessmentRequestObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// EmailIdentityRequestBuilder is request builder for EmailIdentity
type EmailIdentityRequestBuilder struct{ BaseRequestBuilder }

// Request returns EmailIdentityRequest
func (b *EmailIdentityRequestBuilder) Request() *EmailIdentityRequest {
	return &EmailIdentityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EmailIdentityRequest is request for EmailIdentity
type EmailIdentityRequest struct{ BaseRequest }

// Get performs GET request for EmailIdentity
func (r *EmailIdentityRequest) Get(ctx context.Context) (resObj *EmailIdentity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EmailIdentity
func (r *EmailIdentityRequest) Update(ctx context.Context, reqObj *EmailIdentity) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EmailIdentity
func (r *EmailIdentityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
