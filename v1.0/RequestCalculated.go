// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// CalculatedColumnRequestBuilder is request builder for CalculatedColumn
type CalculatedColumnRequestBuilder struct{ BaseRequestBuilder }

// Request returns CalculatedColumnRequest
func (b *CalculatedColumnRequestBuilder) Request() *CalculatedColumnRequest {
	return &CalculatedColumnRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CalculatedColumnRequest is request for CalculatedColumn
type CalculatedColumnRequest struct{ BaseRequest }

// Get performs GET request for CalculatedColumn
func (r *CalculatedColumnRequest) Get(ctx context.Context) (resObj *CalculatedColumn, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CalculatedColumn
func (r *CalculatedColumnRequest) Update(ctx context.Context, reqObj *CalculatedColumn) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CalculatedColumn
func (r *CalculatedColumnRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
