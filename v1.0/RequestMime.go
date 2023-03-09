// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// MimeContentRequestBuilder is request builder for MimeContent
type MimeContentRequestBuilder struct{ BaseRequestBuilder }

// Request returns MimeContentRequest
func (b *MimeContentRequestBuilder) Request() *MimeContentRequest {
	return &MimeContentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MimeContentRequest is request for MimeContent
type MimeContentRequest struct{ BaseRequest }

// Get performs GET request for MimeContent
func (r *MimeContentRequest) Get(ctx context.Context) (resObj *MimeContent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MimeContent
func (r *MimeContentRequest) Update(ctx context.Context, reqObj *MimeContent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MimeContent
func (r *MimeContentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
