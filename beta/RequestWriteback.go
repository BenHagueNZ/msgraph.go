// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// WritebackConfigurationRequestBuilder is request builder for WritebackConfiguration
type WritebackConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns WritebackConfigurationRequest
func (b *WritebackConfigurationRequestBuilder) Request() *WritebackConfigurationRequest {
	return &WritebackConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WritebackConfigurationRequest is request for WritebackConfiguration
type WritebackConfigurationRequest struct{ BaseRequest }

// Get performs GET request for WritebackConfiguration
func (r *WritebackConfigurationRequest) Get(ctx context.Context) (resObj *WritebackConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WritebackConfiguration
func (r *WritebackConfigurationRequest) Update(ctx context.Context, reqObj *WritebackConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WritebackConfiguration
func (r *WritebackConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
