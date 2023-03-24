// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// TermColumnRequestBuilder is request builder for TermColumn
type TermColumnRequestBuilder struct{ BaseRequestBuilder }

// Request returns TermColumnRequest
func (b *TermColumnRequestBuilder) Request() *TermColumnRequest {
	return &TermColumnRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TermColumnRequest is request for TermColumn
type TermColumnRequest struct{ BaseRequest }

// Get performs GET request for TermColumn
func (r *TermColumnRequest) Get(ctx context.Context) (resObj *TermColumn, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TermColumn
func (r *TermColumnRequest) Update(ctx context.Context, reqObj *TermColumn) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TermColumn
func (r *TermColumnRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TermStoreSetRequestBuilder is request builder for TermStoreSet
type TermStoreSetRequestBuilder struct{ BaseRequestBuilder }

// Request returns TermStoreSetRequest
func (b *TermStoreSetRequestBuilder) Request() *TermStoreSetRequest {
	return &TermStoreSetRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TermStoreSetRequest is request for TermStoreSet
type TermStoreSetRequest struct{ BaseRequest }

// Get performs GET request for TermStoreSet
func (r *TermStoreSetRequest) Get(ctx context.Context) (resObj *TermStoreSet, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TermStoreSet
func (r *TermStoreSetRequest) Update(ctx context.Context, reqObj *TermStoreSet) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TermStoreSet
func (r *TermStoreSetRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TermStoreStoreRequestBuilder is request builder for TermStoreStore
type TermStoreStoreRequestBuilder struct{ BaseRequestBuilder }

// Request returns TermStoreStoreRequest
func (b *TermStoreStoreRequestBuilder) Request() *TermStoreStoreRequest {
	return &TermStoreStoreRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TermStoreStoreRequest is request for TermStoreStore
type TermStoreStoreRequest struct{ BaseRequest }

// Get performs GET request for TermStoreStore
func (r *TermStoreStoreRequest) Get(ctx context.Context) (resObj *TermStoreStore, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TermStoreStore
func (r *TermStoreStoreRequest) Update(ctx context.Context, reqObj *TermStoreStore) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TermStoreStore
func (r *TermStoreStoreRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TermStoreTermRequestBuilder is request builder for TermStoreTerm
type TermStoreTermRequestBuilder struct{ BaseRequestBuilder }

// Request returns TermStoreTermRequest
func (b *TermStoreTermRequestBuilder) Request() *TermStoreTermRequest {
	return &TermStoreTermRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TermStoreTermRequest is request for TermStoreTerm
type TermStoreTermRequest struct{ BaseRequest }

// Get performs GET request for TermStoreTerm
func (r *TermStoreTermRequest) Get(ctx context.Context) (resObj *TermStoreTerm, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TermStoreTerm
func (r *TermStoreTermRequest) Update(ctx context.Context, reqObj *TermStoreTerm) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TermStoreTerm
func (r *TermStoreTermRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TermStoreGroupRequestBuilder is request builder for TermStoreGroup
type TermStoreGroupRequestBuilder struct{ BaseRequestBuilder }

// Request returns TermStoreGroupRequest
func (b *TermStoreGroupRequestBuilder) Request() *TermStoreGroupRequest {
	return &TermStoreGroupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TermStoreGroupRequest is request for TermStoreGroup
type TermStoreGroupRequest struct{ BaseRequest }

// Get performs GET request for TermStoreGroup
func (r *TermStoreGroupRequest) Get(ctx context.Context) (resObj *TermStoreGroup, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TermStoreGroup
func (r *TermStoreGroupRequest) Update(ctx context.Context, reqObj *TermStoreGroup) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TermStoreGroup
func (r *TermStoreGroupRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TermStoreRelationRequestBuilder is request builder for TermStoreRelation
type TermStoreRelationRequestBuilder struct{ BaseRequestBuilder }

// Request returns TermStoreRelationRequest
func (b *TermStoreRelationRequestBuilder) Request() *TermStoreRelationRequest {
	return &TermStoreRelationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TermStoreRelationRequest is request for TermStoreRelation
type TermStoreRelationRequest struct{ BaseRequest }

// Get performs GET request for TermStoreRelation
func (r *TermStoreRelationRequest) Get(ctx context.Context) (resObj *TermStoreRelation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TermStoreRelation
func (r *TermStoreRelationRequest) Update(ctx context.Context, reqObj *TermStoreRelation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TermStoreRelation
func (r *TermStoreRelationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
