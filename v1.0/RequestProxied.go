// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ProxiedDomainRequestBuilder is request builder for ProxiedDomain
type ProxiedDomainRequestBuilder struct{ BaseRequestBuilder }

// Request returns ProxiedDomainRequest
func (b *ProxiedDomainRequestBuilder) Request() *ProxiedDomainRequest {
	return &ProxiedDomainRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ProxiedDomainRequest is request for ProxiedDomain
type ProxiedDomainRequest struct{ BaseRequest }

// Get performs GET request for ProxiedDomain
func (r *ProxiedDomainRequest) Get(ctx context.Context) (resObj *ProxiedDomain, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ProxiedDomain
func (r *ProxiedDomainRequest) Update(ctx context.Context, reqObj *ProxiedDomain) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ProxiedDomain
func (r *ProxiedDomainRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}