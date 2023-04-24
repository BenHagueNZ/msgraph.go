// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// TaskFileAttachmentRequestBuilder is request builder for TaskFileAttachment
type TaskFileAttachmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns TaskFileAttachmentRequest
func (b *TaskFileAttachmentRequestBuilder) Request() *TaskFileAttachmentRequest {
	return &TaskFileAttachmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TaskFileAttachmentRequest is request for TaskFileAttachment
type TaskFileAttachmentRequest struct{ BaseRequest }

// Get performs GET request for TaskFileAttachment
func (r *TaskFileAttachmentRequest) Get(ctx context.Context) (resObj *TaskFileAttachment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TaskFileAttachment
func (r *TaskFileAttachmentRequest) Update(ctx context.Context, reqObj *TaskFileAttachment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TaskFileAttachment
func (r *TaskFileAttachmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}