// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// DocumentSetRequestBuilder is request builder for DocumentSet
type DocumentSetRequestBuilder struct{ BaseRequestBuilder }

// Request returns DocumentSetRequest
func (b *DocumentSetRequestBuilder) Request() *DocumentSetRequest {
	return &DocumentSetRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DocumentSetRequest is request for DocumentSet
type DocumentSetRequest struct{ BaseRequest }

// Get performs GET request for DocumentSet
func (r *DocumentSetRequest) Get(ctx context.Context) (resObj *DocumentSet, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DocumentSet
func (r *DocumentSetRequest) Update(ctx context.Context, reqObj *DocumentSet) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DocumentSet
func (r *DocumentSetRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DocumentSetVersionRequestBuilder is request builder for DocumentSetVersion
type DocumentSetVersionRequestBuilder struct{ BaseRequestBuilder }

// Request returns DocumentSetVersionRequest
func (b *DocumentSetVersionRequestBuilder) Request() *DocumentSetVersionRequest {
	return &DocumentSetVersionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DocumentSetVersionRequest is request for DocumentSetVersion
type DocumentSetVersionRequest struct{ BaseRequest }

// Get performs GET request for DocumentSetVersion
func (r *DocumentSetVersionRequest) Get(ctx context.Context) (resObj *DocumentSetVersion, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DocumentSetVersion
func (r *DocumentSetVersionRequest) Update(ctx context.Context, reqObj *DocumentSetVersion) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DocumentSetVersion
func (r *DocumentSetVersionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
