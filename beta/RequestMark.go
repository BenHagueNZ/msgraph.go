// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// MarkContentRequestBuilder is request builder for MarkContent
type MarkContentRequestBuilder struct{ BaseRequestBuilder }

// Request returns MarkContentRequest
func (b *MarkContentRequestBuilder) Request() *MarkContentRequest {
	return &MarkContentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MarkContentRequest is request for MarkContent
type MarkContentRequest struct{ BaseRequest }

// Get performs GET request for MarkContent
func (r *MarkContentRequest) Get(ctx context.Context) (resObj *MarkContent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MarkContent
func (r *MarkContentRequest) Update(ctx context.Context, reqObj *MarkContent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MarkContent
func (r *MarkContentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}