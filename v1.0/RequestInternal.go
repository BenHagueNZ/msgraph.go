// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// InternalDomainFederationRequestBuilder is request builder for InternalDomainFederation
type InternalDomainFederationRequestBuilder struct{ BaseRequestBuilder }

// Request returns InternalDomainFederationRequest
func (b *InternalDomainFederationRequestBuilder) Request() *InternalDomainFederationRequest {
	return &InternalDomainFederationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// InternalDomainFederationRequest is request for InternalDomainFederation
type InternalDomainFederationRequest struct{ BaseRequest }

// Get performs GET request for InternalDomainFederation
func (r *InternalDomainFederationRequest) Get(ctx context.Context) (resObj *InternalDomainFederation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for InternalDomainFederation
func (r *InternalDomainFederationRequest) Update(ctx context.Context, reqObj *InternalDomainFederation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for InternalDomainFederation
func (r *InternalDomainFederationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// InternalSponsorsRequestBuilder is request builder for InternalSponsors
type InternalSponsorsRequestBuilder struct{ BaseRequestBuilder }

// Request returns InternalSponsorsRequest
func (b *InternalSponsorsRequestBuilder) Request() *InternalSponsorsRequest {
	return &InternalSponsorsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// InternalSponsorsRequest is request for InternalSponsors
type InternalSponsorsRequest struct{ BaseRequest }

// Get performs GET request for InternalSponsors
func (r *InternalSponsorsRequest) Get(ctx context.Context) (resObj *InternalSponsors, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for InternalSponsors
func (r *InternalSponsorsRequest) Update(ctx context.Context, reqObj *InternalSponsors) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for InternalSponsors
func (r *InternalSponsorsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
