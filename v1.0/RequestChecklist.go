// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ChecklistItemRequestBuilder is request builder for ChecklistItem
type ChecklistItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChecklistItemRequest
func (b *ChecklistItemRequestBuilder) Request() *ChecklistItemRequest {
	return &ChecklistItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChecklistItemRequest is request for ChecklistItem
type ChecklistItemRequest struct{ BaseRequest }

// Get performs GET request for ChecklistItem
func (r *ChecklistItemRequest) Get(ctx context.Context) (resObj *ChecklistItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChecklistItem
func (r *ChecklistItemRequest) Update(ctx context.Context, reqObj *ChecklistItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChecklistItem
func (r *ChecklistItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
