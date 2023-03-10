// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// RedirectURISettingsRequestBuilder is request builder for RedirectURISettings
type RedirectURISettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns RedirectURISettingsRequest
func (b *RedirectURISettingsRequestBuilder) Request() *RedirectURISettingsRequest {
	return &RedirectURISettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RedirectURISettingsRequest is request for RedirectURISettings
type RedirectURISettingsRequest struct{ BaseRequest }

// Get performs GET request for RedirectURISettings
func (r *RedirectURISettingsRequest) Get(ctx context.Context) (resObj *RedirectURISettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RedirectURISettings
func (r *RedirectURISettingsRequest) Update(ctx context.Context, reqObj *RedirectURISettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RedirectURISettings
func (r *RedirectURISettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
