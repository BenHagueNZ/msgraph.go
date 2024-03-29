// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// CorsConfigurationRequestBuilder is request builder for CorsConfiguration
type CorsConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns CorsConfigurationRequest
func (b *CorsConfigurationRequestBuilder) Request() *CorsConfigurationRequest {
	return &CorsConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CorsConfigurationRequest is request for CorsConfiguration
type CorsConfigurationRequest struct{ BaseRequest }

// Get performs GET request for CorsConfiguration
func (r *CorsConfigurationRequest) Get(ctx context.Context) (resObj *CorsConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CorsConfiguration
func (r *CorsConfigurationRequest) Update(ctx context.Context, reqObj *CorsConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CorsConfiguration
func (r *CorsConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CorsConfiguration_v2RequestBuilder is request builder for CorsConfiguration_v2
type CorsConfiguration_v2RequestBuilder struct{ BaseRequestBuilder }

// Request returns CorsConfiguration_v2Request
func (b *CorsConfiguration_v2RequestBuilder) Request() *CorsConfiguration_v2Request {
	return &CorsConfiguration_v2Request{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CorsConfiguration_v2Request is request for CorsConfiguration_v2
type CorsConfiguration_v2Request struct{ BaseRequest }

// Get performs GET request for CorsConfiguration_v2
func (r *CorsConfiguration_v2Request) Get(ctx context.Context) (resObj *CorsConfiguration_v2, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CorsConfiguration_v2
func (r *CorsConfiguration_v2Request) Update(ctx context.Context, reqObj *CorsConfiguration_v2) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CorsConfiguration_v2
func (r *CorsConfiguration_v2Request) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
