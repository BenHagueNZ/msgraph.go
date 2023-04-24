// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// PropertyRequestBuilder is request builder for Property
type PropertyRequestBuilder struct{ BaseRequestBuilder }

// Request returns PropertyRequest
func (b *PropertyRequestBuilder) Request() *PropertyRequest {
	return &PropertyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PropertyRequest is request for Property
type PropertyRequest struct{ BaseRequest }

// Get performs GET request for Property
func (r *PropertyRequest) Get(ctx context.Context) (resObj *Property, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Property
func (r *PropertyRequest) Update(ctx context.Context, reqObj *Property) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Property
func (r *PropertyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PropertyToEvaluateRequestBuilder is request builder for PropertyToEvaluate
type PropertyToEvaluateRequestBuilder struct{ BaseRequestBuilder }

// Request returns PropertyToEvaluateRequest
func (b *PropertyToEvaluateRequestBuilder) Request() *PropertyToEvaluateRequest {
	return &PropertyToEvaluateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PropertyToEvaluateRequest is request for PropertyToEvaluate
type PropertyToEvaluateRequest struct{ BaseRequest }

// Get performs GET request for PropertyToEvaluate
func (r *PropertyToEvaluateRequest) Get(ctx context.Context) (resObj *PropertyToEvaluate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PropertyToEvaluate
func (r *PropertyToEvaluateRequest) Update(ctx context.Context, reqObj *PropertyToEvaluate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PropertyToEvaluate
func (r *PropertyToEvaluateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
