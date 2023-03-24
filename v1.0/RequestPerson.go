// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// PersonRequestBuilder is request builder for Person
type PersonRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonRequest
func (b *PersonRequestBuilder) Request() *PersonRequest {
	return &PersonRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PersonRequest is request for Person
type PersonRequest struct{ BaseRequest }

// Get performs GET request for Person
func (r *PersonRequest) Get(ctx context.Context) (resObj *Person, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Person
func (r *PersonRequest) Update(ctx context.Context, reqObj *Person) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Person
func (r *PersonRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PersonOrGroupColumnRequestBuilder is request builder for PersonOrGroupColumn
type PersonOrGroupColumnRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonOrGroupColumnRequest
func (b *PersonOrGroupColumnRequestBuilder) Request() *PersonOrGroupColumnRequest {
	return &PersonOrGroupColumnRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PersonOrGroupColumnRequest is request for PersonOrGroupColumn
type PersonOrGroupColumnRequest struct{ BaseRequest }

// Get performs GET request for PersonOrGroupColumn
func (r *PersonOrGroupColumnRequest) Get(ctx context.Context) (resObj *PersonOrGroupColumn, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PersonOrGroupColumn
func (r *PersonOrGroupColumnRequest) Update(ctx context.Context, reqObj *PersonOrGroupColumn) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PersonOrGroupColumn
func (r *PersonOrGroupColumnRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PersonTypeRequestBuilder is request builder for PersonType
type PersonTypeRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonTypeRequest
func (b *PersonTypeRequestBuilder) Request() *PersonTypeRequest {
	return &PersonTypeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PersonTypeRequest is request for PersonType
type PersonTypeRequest struct{ BaseRequest }

// Get performs GET request for PersonType
func (r *PersonTypeRequest) Get(ctx context.Context) (resObj *PersonType, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PersonType
func (r *PersonTypeRequest) Update(ctx context.Context, reqObj *PersonType) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PersonType
func (r *PersonTypeRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
