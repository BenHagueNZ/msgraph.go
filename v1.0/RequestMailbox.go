// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// MailboxSettingsRequestBuilder is request builder for MailboxSettings
type MailboxSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns MailboxSettingsRequest
func (b *MailboxSettingsRequestBuilder) Request() *MailboxSettingsRequest {
	return &MailboxSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MailboxSettingsRequest is request for MailboxSettings
type MailboxSettingsRequest struct{ BaseRequest }

// Get performs GET request for MailboxSettings
func (r *MailboxSettingsRequest) Get(ctx context.Context) (resObj *MailboxSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MailboxSettings
func (r *MailboxSettingsRequest) Update(ctx context.Context, reqObj *MailboxSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MailboxSettings
func (r *MailboxSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}