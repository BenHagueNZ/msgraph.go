// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// SettingsRequestBuilder is request builder for Settings
type SettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns SettingsRequest
func (b *SettingsRequestBuilder) Request() *SettingsRequest {
	return &SettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SettingsRequest is request for Settings
type SettingsRequest struct{ BaseRequest }

// Get performs GET request for Settings
func (r *SettingsRequest) Get(ctx context.Context) (resObj *Settings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Settings
func (r *SettingsRequest) Update(ctx context.Context, reqObj *Settings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Settings
func (r *SettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
