// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ChoiceColumnRequestBuilder is request builder for ChoiceColumn
type ChoiceColumnRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChoiceColumnRequest
func (b *ChoiceColumnRequestBuilder) Request() *ChoiceColumnRequest {
	return &ChoiceColumnRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChoiceColumnRequest is request for ChoiceColumn
type ChoiceColumnRequest struct{ BaseRequest }

// Get performs GET request for ChoiceColumn
func (r *ChoiceColumnRequest) Get(ctx context.Context) (resObj *ChoiceColumn, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChoiceColumn
func (r *ChoiceColumnRequest) Update(ctx context.Context, reqObj *ChoiceColumn) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChoiceColumn
func (r *ChoiceColumnRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
