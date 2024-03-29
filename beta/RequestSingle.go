// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// SingleSignOnExtensionRequestBuilder is request builder for SingleSignOnExtension
type SingleSignOnExtensionRequestBuilder struct{ BaseRequestBuilder }

// Request returns SingleSignOnExtensionRequest
func (b *SingleSignOnExtensionRequestBuilder) Request() *SingleSignOnExtensionRequest {
	return &SingleSignOnExtensionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SingleSignOnExtensionRequest is request for SingleSignOnExtension
type SingleSignOnExtensionRequest struct{ BaseRequest }

// Get performs GET request for SingleSignOnExtension
func (r *SingleSignOnExtensionRequest) Get(ctx context.Context) (resObj *SingleSignOnExtension, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SingleSignOnExtension
func (r *SingleSignOnExtensionRequest) Update(ctx context.Context, reqObj *SingleSignOnExtension) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SingleSignOnExtension
func (r *SingleSignOnExtensionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SingleUserRequestBuilder is request builder for SingleUser
type SingleUserRequestBuilder struct{ BaseRequestBuilder }

// Request returns SingleUserRequest
func (b *SingleUserRequestBuilder) Request() *SingleUserRequest {
	return &SingleUserRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SingleUserRequest is request for SingleUser
type SingleUserRequest struct{ BaseRequest }

// Get performs GET request for SingleUser
func (r *SingleUserRequest) Get(ctx context.Context) (resObj *SingleUser, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SingleUser
func (r *SingleUserRequest) Update(ctx context.Context, reqObj *SingleUser) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SingleUser
func (r *SingleUserRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SingleValueLegacyExtendedPropertyRequestBuilder is request builder for SingleValueLegacyExtendedProperty
type SingleValueLegacyExtendedPropertyRequestBuilder struct{ BaseRequestBuilder }

// Request returns SingleValueLegacyExtendedPropertyRequest
func (b *SingleValueLegacyExtendedPropertyRequestBuilder) Request() *SingleValueLegacyExtendedPropertyRequest {
	return &SingleValueLegacyExtendedPropertyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SingleValueLegacyExtendedPropertyRequest is request for SingleValueLegacyExtendedProperty
type SingleValueLegacyExtendedPropertyRequest struct{ BaseRequest }

// Get performs GET request for SingleValueLegacyExtendedProperty
func (r *SingleValueLegacyExtendedPropertyRequest) Get(ctx context.Context) (resObj *SingleValueLegacyExtendedProperty, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SingleValueLegacyExtendedProperty
func (r *SingleValueLegacyExtendedPropertyRequest) Update(ctx context.Context, reqObj *SingleValueLegacyExtendedProperty) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SingleValueLegacyExtendedProperty
func (r *SingleValueLegacyExtendedPropertyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
