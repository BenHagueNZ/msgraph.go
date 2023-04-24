// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// LearningContentRequestBuilder is request builder for LearningContent
type LearningContentRequestBuilder struct{ BaseRequestBuilder }

// Request returns LearningContentRequest
func (b *LearningContentRequestBuilder) Request() *LearningContentRequest {
	return &LearningContentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// LearningContentRequest is request for LearningContent
type LearningContentRequest struct{ BaseRequest }

// Get performs GET request for LearningContent
func (r *LearningContentRequest) Get(ctx context.Context) (resObj *LearningContent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for LearningContent
func (r *LearningContentRequest) Update(ctx context.Context, reqObj *LearningContent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for LearningContent
func (r *LearningContentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// LearningProviderRequestBuilder is request builder for LearningProvider
type LearningProviderRequestBuilder struct{ BaseRequestBuilder }

// Request returns LearningProviderRequest
func (b *LearningProviderRequestBuilder) Request() *LearningProviderRequest {
	return &LearningProviderRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// LearningProviderRequest is request for LearningProvider
type LearningProviderRequest struct{ BaseRequest }

// Get performs GET request for LearningProvider
func (r *LearningProviderRequest) Get(ctx context.Context) (resObj *LearningProvider, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for LearningProvider
func (r *LearningProviderRequest) Update(ctx context.Context, reqObj *LearningProvider) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for LearningProvider
func (r *LearningProviderRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
