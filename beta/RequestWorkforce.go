// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// WorkforceIntegrationRequestBuilder is request builder for WorkforceIntegration
type WorkforceIntegrationRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkforceIntegrationRequest
func (b *WorkforceIntegrationRequestBuilder) Request() *WorkforceIntegrationRequest {
	return &WorkforceIntegrationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkforceIntegrationRequest is request for WorkforceIntegration
type WorkforceIntegrationRequest struct{ BaseRequest }

// Get performs GET request for WorkforceIntegration
func (r *WorkforceIntegrationRequest) Get(ctx context.Context) (resObj *WorkforceIntegration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkforceIntegration
func (r *WorkforceIntegrationRequest) Update(ctx context.Context, reqObj *WorkforceIntegration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkforceIntegration
func (r *WorkforceIntegrationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// WorkforceIntegrationEncryptionRequestBuilder is request builder for WorkforceIntegrationEncryption
type WorkforceIntegrationEncryptionRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkforceIntegrationEncryptionRequest
func (b *WorkforceIntegrationEncryptionRequestBuilder) Request() *WorkforceIntegrationEncryptionRequest {
	return &WorkforceIntegrationEncryptionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkforceIntegrationEncryptionRequest is request for WorkforceIntegrationEncryption
type WorkforceIntegrationEncryptionRequest struct{ BaseRequest }

// Get performs GET request for WorkforceIntegrationEncryption
func (r *WorkforceIntegrationEncryptionRequest) Get(ctx context.Context) (resObj *WorkforceIntegrationEncryption, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkforceIntegrationEncryption
func (r *WorkforceIntegrationEncryptionRequest) Update(ctx context.Context, reqObj *WorkforceIntegrationEncryption) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkforceIntegrationEncryption
func (r *WorkforceIntegrationEncryptionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}