// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// MentionsPreviewRequestBuilder is request builder for MentionsPreview
type MentionsPreviewRequestBuilder struct{ BaseRequestBuilder }

// Request returns MentionsPreviewRequest
func (b *MentionsPreviewRequestBuilder) Request() *MentionsPreviewRequest {
	return &MentionsPreviewRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MentionsPreviewRequest is request for MentionsPreview
type MentionsPreviewRequest struct{ BaseRequest }

// Get performs GET request for MentionsPreview
func (r *MentionsPreviewRequest) Get(ctx context.Context) (resObj *MentionsPreview, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MentionsPreview
func (r *MentionsPreviewRequest) Update(ctx context.Context, reqObj *MentionsPreview) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MentionsPreview
func (r *MentionsPreviewRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
