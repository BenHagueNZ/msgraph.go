// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// SingleServicePrincipalRequestBuilder is request builder for SingleServicePrincipal
type SingleServicePrincipalRequestBuilder struct{ BaseRequestBuilder }

// Request returns SingleServicePrincipalRequest
func (b *SingleServicePrincipalRequestBuilder) Request() *SingleServicePrincipalRequest {
	return &SingleServicePrincipalRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SingleServicePrincipalRequest is request for SingleServicePrincipal
type SingleServicePrincipalRequest struct{ BaseRequest }

// Get performs GET request for SingleServicePrincipal
func (r *SingleServicePrincipalRequest) Get(ctx context.Context) (resObj *SingleServicePrincipal, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SingleServicePrincipal
func (r *SingleServicePrincipalRequest) Update(ctx context.Context, reqObj *SingleServicePrincipal) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SingleServicePrincipal
func (r *SingleServicePrincipalRequest) Delete(ctx context.Context) error {
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
