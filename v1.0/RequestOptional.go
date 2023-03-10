// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// OptionalClaimRequestBuilder is request builder for OptionalClaim
type OptionalClaimRequestBuilder struct{ BaseRequestBuilder }

// Request returns OptionalClaimRequest
func (b *OptionalClaimRequestBuilder) Request() *OptionalClaimRequest {
	return &OptionalClaimRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OptionalClaimRequest is request for OptionalClaim
type OptionalClaimRequest struct{ BaseRequest }

// Get performs GET request for OptionalClaim
func (r *OptionalClaimRequest) Get(ctx context.Context) (resObj *OptionalClaim, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OptionalClaim
func (r *OptionalClaimRequest) Update(ctx context.Context, reqObj *OptionalClaim) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OptionalClaim
func (r *OptionalClaimRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OptionalClaimsRequestBuilder is request builder for OptionalClaims
type OptionalClaimsRequestBuilder struct{ BaseRequestBuilder }

// Request returns OptionalClaimsRequest
func (b *OptionalClaimsRequestBuilder) Request() *OptionalClaimsRequest {
	return &OptionalClaimsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OptionalClaimsRequest is request for OptionalClaims
type OptionalClaimsRequest struct{ BaseRequest }

// Get performs GET request for OptionalClaims
func (r *OptionalClaimsRequest) Get(ctx context.Context) (resObj *OptionalClaims, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OptionalClaims
func (r *OptionalClaimsRequest) Update(ctx context.Context, reqObj *OptionalClaims) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OptionalClaims
func (r *OptionalClaimsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
