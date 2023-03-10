// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// TemporaryAccessPassAuthenticationMethodRequestBuilder is request builder for TemporaryAccessPassAuthenticationMethod
type TemporaryAccessPassAuthenticationMethodRequestBuilder struct{ BaseRequestBuilder }

// Request returns TemporaryAccessPassAuthenticationMethodRequest
func (b *TemporaryAccessPassAuthenticationMethodRequestBuilder) Request() *TemporaryAccessPassAuthenticationMethodRequest {
	return &TemporaryAccessPassAuthenticationMethodRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TemporaryAccessPassAuthenticationMethodRequest is request for TemporaryAccessPassAuthenticationMethod
type TemporaryAccessPassAuthenticationMethodRequest struct{ BaseRequest }

// Get performs GET request for TemporaryAccessPassAuthenticationMethod
func (r *TemporaryAccessPassAuthenticationMethodRequest) Get(ctx context.Context) (resObj *TemporaryAccessPassAuthenticationMethod, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TemporaryAccessPassAuthenticationMethod
func (r *TemporaryAccessPassAuthenticationMethodRequest) Update(ctx context.Context, reqObj *TemporaryAccessPassAuthenticationMethod) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TemporaryAccessPassAuthenticationMethod
func (r *TemporaryAccessPassAuthenticationMethodRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TemporaryAccessPassAuthenticationMethodConfigurationRequestBuilder is request builder for TemporaryAccessPassAuthenticationMethodConfiguration
type TemporaryAccessPassAuthenticationMethodConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns TemporaryAccessPassAuthenticationMethodConfigurationRequest
func (b *TemporaryAccessPassAuthenticationMethodConfigurationRequestBuilder) Request() *TemporaryAccessPassAuthenticationMethodConfigurationRequest {
	return &TemporaryAccessPassAuthenticationMethodConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TemporaryAccessPassAuthenticationMethodConfigurationRequest is request for TemporaryAccessPassAuthenticationMethodConfiguration
type TemporaryAccessPassAuthenticationMethodConfigurationRequest struct{ BaseRequest }

// Get performs GET request for TemporaryAccessPassAuthenticationMethodConfiguration
func (r *TemporaryAccessPassAuthenticationMethodConfigurationRequest) Get(ctx context.Context) (resObj *TemporaryAccessPassAuthenticationMethodConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TemporaryAccessPassAuthenticationMethodConfiguration
func (r *TemporaryAccessPassAuthenticationMethodConfigurationRequest) Update(ctx context.Context, reqObj *TemporaryAccessPassAuthenticationMethodConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TemporaryAccessPassAuthenticationMethodConfiguration
func (r *TemporaryAccessPassAuthenticationMethodConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
