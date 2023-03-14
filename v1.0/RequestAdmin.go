// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// AdminRequestBuilder is request builder for Admin
type AdminRequestBuilder struct{ BaseRequestBuilder }

// Request returns AdminRequest
func (b *AdminRequestBuilder) Request() *AdminRequest {
	return &AdminRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AdminRequest is request for Admin
type AdminRequest struct{ BaseRequest }

// Get performs GET request for Admin
func (r *AdminRequest) Get(ctx context.Context) (resObj *Admin, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Admin
func (r *AdminRequest) Update(ctx context.Context, reqObj *Admin) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Admin
func (r *AdminRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AdminConsentRequestPolicyRequestBuilder is request builder for AdminConsentRequestPolicy
type AdminConsentRequestPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns AdminConsentRequestPolicyRequest
func (b *AdminConsentRequestPolicyRequestBuilder) Request() *AdminConsentRequestPolicyRequest {
	return &AdminConsentRequestPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AdminConsentRequestPolicyRequest is request for AdminConsentRequestPolicy
type AdminConsentRequestPolicyRequest struct{ BaseRequest }

// Get performs GET request for AdminConsentRequestPolicy
func (r *AdminConsentRequestPolicyRequest) Get(ctx context.Context) (resObj *AdminConsentRequestPolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AdminConsentRequestPolicy
func (r *AdminConsentRequestPolicyRequest) Update(ctx context.Context, reqObj *AdminConsentRequestPolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AdminConsentRequestPolicy
func (r *AdminConsentRequestPolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}