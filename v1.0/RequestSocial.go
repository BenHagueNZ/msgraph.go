// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// SocialIdentityProviderRequestBuilder is request builder for SocialIdentityProvider
type SocialIdentityProviderRequestBuilder struct{ BaseRequestBuilder }

// Request returns SocialIdentityProviderRequest
func (b *SocialIdentityProviderRequestBuilder) Request() *SocialIdentityProviderRequest {
	return &SocialIdentityProviderRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SocialIdentityProviderRequest is request for SocialIdentityProvider
type SocialIdentityProviderRequest struct{ BaseRequest }

// Get performs GET request for SocialIdentityProvider
func (r *SocialIdentityProviderRequest) Get(ctx context.Context) (resObj *SocialIdentityProvider, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SocialIdentityProvider
func (r *SocialIdentityProviderRequest) Update(ctx context.Context, reqObj *SocialIdentityProvider) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SocialIdentityProvider
func (r *SocialIdentityProviderRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}