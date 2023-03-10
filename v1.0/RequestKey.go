// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// KeyCredentialRequestBuilder is request builder for KeyCredential
type KeyCredentialRequestBuilder struct{ BaseRequestBuilder }

// Request returns KeyCredentialRequest
func (b *KeyCredentialRequestBuilder) Request() *KeyCredentialRequest {
	return &KeyCredentialRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// KeyCredentialRequest is request for KeyCredential
type KeyCredentialRequest struct{ BaseRequest }

// Get performs GET request for KeyCredential
func (r *KeyCredentialRequest) Get(ctx context.Context) (resObj *KeyCredential, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for KeyCredential
func (r *KeyCredentialRequest) Update(ctx context.Context, reqObj *KeyCredential) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for KeyCredential
func (r *KeyCredentialRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// KeyCredentialConfigurationRequestBuilder is request builder for KeyCredentialConfiguration
type KeyCredentialConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns KeyCredentialConfigurationRequest
func (b *KeyCredentialConfigurationRequestBuilder) Request() *KeyCredentialConfigurationRequest {
	return &KeyCredentialConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// KeyCredentialConfigurationRequest is request for KeyCredentialConfiguration
type KeyCredentialConfigurationRequest struct{ BaseRequest }

// Get performs GET request for KeyCredentialConfiguration
func (r *KeyCredentialConfigurationRequest) Get(ctx context.Context) (resObj *KeyCredentialConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for KeyCredentialConfiguration
func (r *KeyCredentialConfigurationRequest) Update(ctx context.Context, reqObj *KeyCredentialConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for KeyCredentialConfiguration
func (r *KeyCredentialConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// KeyValueRequestBuilder is request builder for KeyValue
type KeyValueRequestBuilder struct{ BaseRequestBuilder }

// Request returns KeyValueRequest
func (b *KeyValueRequestBuilder) Request() *KeyValueRequest {
	return &KeyValueRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// KeyValueRequest is request for KeyValue
type KeyValueRequest struct{ BaseRequest }

// Get performs GET request for KeyValue
func (r *KeyValueRequest) Get(ctx context.Context) (resObj *KeyValue, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for KeyValue
func (r *KeyValueRequest) Update(ctx context.Context, reqObj *KeyValue) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for KeyValue
func (r *KeyValueRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// KeyValuePairRequestBuilder is request builder for KeyValuePair
type KeyValuePairRequestBuilder struct{ BaseRequestBuilder }

// Request returns KeyValuePairRequest
func (b *KeyValuePairRequestBuilder) Request() *KeyValuePairRequest {
	return &KeyValuePairRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// KeyValuePairRequest is request for KeyValuePair
type KeyValuePairRequest struct{ BaseRequest }

// Get performs GET request for KeyValuePair
func (r *KeyValuePairRequest) Get(ctx context.Context) (resObj *KeyValuePair, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for KeyValuePair
func (r *KeyValuePairRequest) Update(ctx context.Context, reqObj *KeyValuePair) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for KeyValuePair
func (r *KeyValuePairRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
