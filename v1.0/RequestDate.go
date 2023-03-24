// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// DateTimeColumnRequestBuilder is request builder for DateTimeColumn
type DateTimeColumnRequestBuilder struct{ BaseRequestBuilder }

// Request returns DateTimeColumnRequest
func (b *DateTimeColumnRequestBuilder) Request() *DateTimeColumnRequest {
	return &DateTimeColumnRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DateTimeColumnRequest is request for DateTimeColumn
type DateTimeColumnRequest struct{ BaseRequest }

// Get performs GET request for DateTimeColumn
func (r *DateTimeColumnRequest) Get(ctx context.Context) (resObj *DateTimeColumn, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DateTimeColumn
func (r *DateTimeColumnRequest) Update(ctx context.Context, reqObj *DateTimeColumn) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DateTimeColumn
func (r *DateTimeColumnRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DateTimeTimeZoneRequestBuilder is request builder for DateTimeTimeZone
type DateTimeTimeZoneRequestBuilder struct{ BaseRequestBuilder }

// Request returns DateTimeTimeZoneRequest
func (b *DateTimeTimeZoneRequestBuilder) Request() *DateTimeTimeZoneRequest {
	return &DateTimeTimeZoneRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DateTimeTimeZoneRequest is request for DateTimeTimeZone
type DateTimeTimeZoneRequest struct{ BaseRequest }

// Get performs GET request for DateTimeTimeZone
func (r *DateTimeTimeZoneRequest) Get(ctx context.Context) (resObj *DateTimeTimeZone, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DateTimeTimeZone
func (r *DateTimeTimeZoneRequest) Update(ctx context.Context, reqObj *DateTimeTimeZone) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DateTimeTimeZone
func (r *DateTimeTimeZoneRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}