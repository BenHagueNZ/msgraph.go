// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// IPv6CidrRangeRequestBuilder is request builder for IPv6CidrRange
type IPv6CidrRangeRequestBuilder struct{ BaseRequestBuilder }

// Request returns IPv6CidrRangeRequest
func (b *IPv6CidrRangeRequestBuilder) Request() *IPv6CidrRangeRequest {
	return &IPv6CidrRangeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IPv6CidrRangeRequest is request for IPv6CidrRange
type IPv6CidrRangeRequest struct{ BaseRequest }

// Get performs GET request for IPv6CidrRange
func (r *IPv6CidrRangeRequest) Get(ctx context.Context) (resObj *IPv6CidrRange, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IPv6CidrRange
func (r *IPv6CidrRangeRequest) Update(ctx context.Context, reqObj *IPv6CidrRange) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IPv6CidrRange
func (r *IPv6CidrRangeRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IPv6RangeRequestBuilder is request builder for IPv6Range
type IPv6RangeRequestBuilder struct{ BaseRequestBuilder }

// Request returns IPv6RangeRequest
func (b *IPv6RangeRequestBuilder) Request() *IPv6RangeRequest {
	return &IPv6RangeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IPv6RangeRequest is request for IPv6Range
type IPv6RangeRequest struct{ BaseRequest }

// Get performs GET request for IPv6Range
func (r *IPv6RangeRequest) Get(ctx context.Context) (resObj *IPv6Range, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IPv6Range
func (r *IPv6RangeRequest) Update(ctx context.Context, reqObj *IPv6Range) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IPv6Range
func (r *IPv6RangeRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
