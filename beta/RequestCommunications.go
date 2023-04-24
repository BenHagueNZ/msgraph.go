// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// CommunicationsApplicationIdentityRequestBuilder is request builder for CommunicationsApplicationIdentity
type CommunicationsApplicationIdentityRequestBuilder struct{ BaseRequestBuilder }

// Request returns CommunicationsApplicationIdentityRequest
func (b *CommunicationsApplicationIdentityRequestBuilder) Request() *CommunicationsApplicationIdentityRequest {
	return &CommunicationsApplicationIdentityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CommunicationsApplicationIdentityRequest is request for CommunicationsApplicationIdentity
type CommunicationsApplicationIdentityRequest struct{ BaseRequest }

// Get performs GET request for CommunicationsApplicationIdentity
func (r *CommunicationsApplicationIdentityRequest) Get(ctx context.Context) (resObj *CommunicationsApplicationIdentity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CommunicationsApplicationIdentity
func (r *CommunicationsApplicationIdentityRequest) Update(ctx context.Context, reqObj *CommunicationsApplicationIdentity) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CommunicationsApplicationIdentity
func (r *CommunicationsApplicationIdentityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CommunicationsApplicationInstanceIdentityRequestBuilder is request builder for CommunicationsApplicationInstanceIdentity
type CommunicationsApplicationInstanceIdentityRequestBuilder struct{ BaseRequestBuilder }

// Request returns CommunicationsApplicationInstanceIdentityRequest
func (b *CommunicationsApplicationInstanceIdentityRequestBuilder) Request() *CommunicationsApplicationInstanceIdentityRequest {
	return &CommunicationsApplicationInstanceIdentityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CommunicationsApplicationInstanceIdentityRequest is request for CommunicationsApplicationInstanceIdentity
type CommunicationsApplicationInstanceIdentityRequest struct{ BaseRequest }

// Get performs GET request for CommunicationsApplicationInstanceIdentity
func (r *CommunicationsApplicationInstanceIdentityRequest) Get(ctx context.Context) (resObj *CommunicationsApplicationInstanceIdentity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CommunicationsApplicationInstanceIdentity
func (r *CommunicationsApplicationInstanceIdentityRequest) Update(ctx context.Context, reqObj *CommunicationsApplicationInstanceIdentity) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CommunicationsApplicationInstanceIdentity
func (r *CommunicationsApplicationInstanceIdentityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CommunicationsEncryptedIdentityRequestBuilder is request builder for CommunicationsEncryptedIdentity
type CommunicationsEncryptedIdentityRequestBuilder struct{ BaseRequestBuilder }

// Request returns CommunicationsEncryptedIdentityRequest
func (b *CommunicationsEncryptedIdentityRequestBuilder) Request() *CommunicationsEncryptedIdentityRequest {
	return &CommunicationsEncryptedIdentityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CommunicationsEncryptedIdentityRequest is request for CommunicationsEncryptedIdentity
type CommunicationsEncryptedIdentityRequest struct{ BaseRequest }

// Get performs GET request for CommunicationsEncryptedIdentity
func (r *CommunicationsEncryptedIdentityRequest) Get(ctx context.Context) (resObj *CommunicationsEncryptedIdentity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CommunicationsEncryptedIdentity
func (r *CommunicationsEncryptedIdentityRequest) Update(ctx context.Context, reqObj *CommunicationsEncryptedIdentity) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CommunicationsEncryptedIdentity
func (r *CommunicationsEncryptedIdentityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CommunicationsGuestIdentityRequestBuilder is request builder for CommunicationsGuestIdentity
type CommunicationsGuestIdentityRequestBuilder struct{ BaseRequestBuilder }

// Request returns CommunicationsGuestIdentityRequest
func (b *CommunicationsGuestIdentityRequestBuilder) Request() *CommunicationsGuestIdentityRequest {
	return &CommunicationsGuestIdentityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CommunicationsGuestIdentityRequest is request for CommunicationsGuestIdentity
type CommunicationsGuestIdentityRequest struct{ BaseRequest }

// Get performs GET request for CommunicationsGuestIdentity
func (r *CommunicationsGuestIdentityRequest) Get(ctx context.Context) (resObj *CommunicationsGuestIdentity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CommunicationsGuestIdentity
func (r *CommunicationsGuestIdentityRequest) Update(ctx context.Context, reqObj *CommunicationsGuestIdentity) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CommunicationsGuestIdentity
func (r *CommunicationsGuestIdentityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CommunicationsIdentitySetRequestBuilder is request builder for CommunicationsIdentitySet
type CommunicationsIdentitySetRequestBuilder struct{ BaseRequestBuilder }

// Request returns CommunicationsIdentitySetRequest
func (b *CommunicationsIdentitySetRequestBuilder) Request() *CommunicationsIdentitySetRequest {
	return &CommunicationsIdentitySetRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CommunicationsIdentitySetRequest is request for CommunicationsIdentitySet
type CommunicationsIdentitySetRequest struct{ BaseRequest }

// Get performs GET request for CommunicationsIdentitySet
func (r *CommunicationsIdentitySetRequest) Get(ctx context.Context) (resObj *CommunicationsIdentitySet, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CommunicationsIdentitySet
func (r *CommunicationsIdentitySetRequest) Update(ctx context.Context, reqObj *CommunicationsIdentitySet) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CommunicationsIdentitySet
func (r *CommunicationsIdentitySetRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CommunicationsPhoneIdentityRequestBuilder is request builder for CommunicationsPhoneIdentity
type CommunicationsPhoneIdentityRequestBuilder struct{ BaseRequestBuilder }

// Request returns CommunicationsPhoneIdentityRequest
func (b *CommunicationsPhoneIdentityRequestBuilder) Request() *CommunicationsPhoneIdentityRequest {
	return &CommunicationsPhoneIdentityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CommunicationsPhoneIdentityRequest is request for CommunicationsPhoneIdentity
type CommunicationsPhoneIdentityRequest struct{ BaseRequest }

// Get performs GET request for CommunicationsPhoneIdentity
func (r *CommunicationsPhoneIdentityRequest) Get(ctx context.Context) (resObj *CommunicationsPhoneIdentity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CommunicationsPhoneIdentity
func (r *CommunicationsPhoneIdentityRequest) Update(ctx context.Context, reqObj *CommunicationsPhoneIdentity) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CommunicationsPhoneIdentity
func (r *CommunicationsPhoneIdentityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CommunicationsUserIdentityRequestBuilder is request builder for CommunicationsUserIdentity
type CommunicationsUserIdentityRequestBuilder struct{ BaseRequestBuilder }

// Request returns CommunicationsUserIdentityRequest
func (b *CommunicationsUserIdentityRequestBuilder) Request() *CommunicationsUserIdentityRequest {
	return &CommunicationsUserIdentityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CommunicationsUserIdentityRequest is request for CommunicationsUserIdentity
type CommunicationsUserIdentityRequest struct{ BaseRequest }

// Get performs GET request for CommunicationsUserIdentity
func (r *CommunicationsUserIdentityRequest) Get(ctx context.Context) (resObj *CommunicationsUserIdentity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CommunicationsUserIdentity
func (r *CommunicationsUserIdentityRequest) Update(ctx context.Context, reqObj *CommunicationsUserIdentity) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CommunicationsUserIdentity
func (r *CommunicationsUserIdentityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}