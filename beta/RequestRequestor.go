// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// RequestorManagerRequestBuilder is request builder for RequestorManager
type RequestorManagerRequestBuilder struct{ BaseRequestBuilder }

// Request returns RequestorManagerRequest
func (b *RequestorManagerRequestBuilder) Request() *RequestorManagerRequest {
	return &RequestorManagerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RequestorManagerRequest is request for RequestorManager
type RequestorManagerRequest struct{ BaseRequest }

// Get performs GET request for RequestorManager
func (r *RequestorManagerRequest) Get(ctx context.Context) (resObj *RequestorManager, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RequestorManager
func (r *RequestorManagerRequest) Update(ctx context.Context, reqObj *RequestorManager) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RequestorManager
func (r *RequestorManagerRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RequestorSettingsRequestBuilder is request builder for RequestorSettings
type RequestorSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns RequestorSettingsRequest
func (b *RequestorSettingsRequestBuilder) Request() *RequestorSettingsRequest {
	return &RequestorSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RequestorSettingsRequest is request for RequestorSettings
type RequestorSettingsRequest struct{ BaseRequest }

// Get performs GET request for RequestorSettings
func (r *RequestorSettingsRequest) Get(ctx context.Context) (resObj *RequestorSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RequestorSettings
func (r *RequestorSettingsRequest) Update(ctx context.Context, reqObj *RequestorSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RequestorSettings
func (r *RequestorSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
