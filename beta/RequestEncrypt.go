// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// EncryptContentRequestBuilder is request builder for EncryptContent
type EncryptContentRequestBuilder struct{ BaseRequestBuilder }

// Request returns EncryptContentRequest
func (b *EncryptContentRequestBuilder) Request() *EncryptContentRequest {
	return &EncryptContentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EncryptContentRequest is request for EncryptContent
type EncryptContentRequest struct{ BaseRequest }

// Get performs GET request for EncryptContent
func (r *EncryptContentRequest) Get(ctx context.Context) (resObj *EncryptContent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EncryptContent
func (r *EncryptContentRequest) Update(ctx context.Context, reqObj *EncryptContent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EncryptContent
func (r *EncryptContentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// EncryptWithTemplateRequestBuilder is request builder for EncryptWithTemplate
type EncryptWithTemplateRequestBuilder struct{ BaseRequestBuilder }

// Request returns EncryptWithTemplateRequest
func (b *EncryptWithTemplateRequestBuilder) Request() *EncryptWithTemplateRequest {
	return &EncryptWithTemplateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EncryptWithTemplateRequest is request for EncryptWithTemplate
type EncryptWithTemplateRequest struct{ BaseRequest }

// Get performs GET request for EncryptWithTemplate
func (r *EncryptWithTemplateRequest) Get(ctx context.Context) (resObj *EncryptWithTemplate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EncryptWithTemplate
func (r *EncryptWithTemplateRequest) Update(ctx context.Context, reqObj *EncryptWithTemplate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EncryptWithTemplate
func (r *EncryptWithTemplateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// EncryptWithUserDefinedRightsRequestBuilder is request builder for EncryptWithUserDefinedRights
type EncryptWithUserDefinedRightsRequestBuilder struct{ BaseRequestBuilder }

// Request returns EncryptWithUserDefinedRightsRequest
func (b *EncryptWithUserDefinedRightsRequestBuilder) Request() *EncryptWithUserDefinedRightsRequest {
	return &EncryptWithUserDefinedRightsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EncryptWithUserDefinedRightsRequest is request for EncryptWithUserDefinedRights
type EncryptWithUserDefinedRightsRequest struct{ BaseRequest }

// Get performs GET request for EncryptWithUserDefinedRights
func (r *EncryptWithUserDefinedRightsRequest) Get(ctx context.Context) (resObj *EncryptWithUserDefinedRights, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EncryptWithUserDefinedRights
func (r *EncryptWithUserDefinedRightsRequest) Update(ctx context.Context, reqObj *EncryptWithUserDefinedRights) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EncryptWithUserDefinedRights
func (r *EncryptWithUserDefinedRightsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}