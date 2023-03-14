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

// DocumentSetContentRequestBuilder is request builder for DocumentSetContent
type DocumentSetContentRequestBuilder struct{ BaseRequestBuilder }

// Request returns DocumentSetContentRequest
func (b *DocumentSetContentRequestBuilder) Request() *DocumentSetContentRequest {
	return &DocumentSetContentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DocumentSetContentRequest is request for DocumentSetContent
type DocumentSetContentRequest struct{ BaseRequest }

// Get performs GET request for DocumentSetContent
func (r *DocumentSetContentRequest) Get(ctx context.Context) (resObj *DocumentSetContent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DocumentSetContent
func (r *DocumentSetContentRequest) Update(ctx context.Context, reqObj *DocumentSetContent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DocumentSetContent
func (r *DocumentSetContentRequest) Delete(ctx context.Context) error {
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

// DocumentSetVersionItemRequestBuilder is request builder for DocumentSetVersionItem
type DocumentSetVersionItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns DocumentSetVersionItemRequest
func (b *DocumentSetVersionItemRequestBuilder) Request() *DocumentSetVersionItemRequest {
	return &DocumentSetVersionItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DocumentSetVersionItemRequest is request for DocumentSetVersionItem
type DocumentSetVersionItemRequest struct{ BaseRequest }

// Get performs GET request for DocumentSetVersionItem
func (r *DocumentSetVersionItemRequest) Get(ctx context.Context) (resObj *DocumentSetVersionItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DocumentSetVersionItem
func (r *DocumentSetVersionItemRequest) Update(ctx context.Context, reqObj *DocumentSetVersionItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DocumentSetVersionItem
func (r *DocumentSetVersionItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}