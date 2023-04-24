// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// CommentActionRequestBuilder is request builder for CommentAction
type CommentActionRequestBuilder struct{ BaseRequestBuilder }

// Request returns CommentActionRequest
func (b *CommentActionRequestBuilder) Request() *CommentActionRequest {
	return &CommentActionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CommentActionRequest is request for CommentAction
type CommentActionRequest struct{ BaseRequest }

// Get performs GET request for CommentAction
func (r *CommentActionRequest) Get(ctx context.Context) (resObj *CommentAction, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CommentAction
func (r *CommentActionRequest) Update(ctx context.Context, reqObj *CommentAction) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CommentAction
func (r *CommentActionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}