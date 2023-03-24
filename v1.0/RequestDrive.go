// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// DriveRequestBuilder is request builder for Drive
type DriveRequestBuilder struct{ BaseRequestBuilder }

// Request returns DriveRequest
func (b *DriveRequestBuilder) Request() *DriveRequest {
	return &DriveRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DriveRequest is request for Drive
type DriveRequest struct{ BaseRequest }

// Get performs GET request for Drive
func (r *DriveRequest) Get(ctx context.Context) (resObj *Drive, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Drive
func (r *DriveRequest) Update(ctx context.Context, reqObj *Drive) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Drive
func (r *DriveRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DriveItemRequestBuilder is request builder for DriveItem
type DriveItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns DriveItemRequest
func (b *DriveItemRequestBuilder) Request() *DriveItemRequest {
	return &DriveItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DriveItemRequest is request for DriveItem
type DriveItemRequest struct{ BaseRequest }

// Get performs GET request for DriveItem
func (r *DriveItemRequest) Get(ctx context.Context) (resObj *DriveItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DriveItem
func (r *DriveItemRequest) Update(ctx context.Context, reqObj *DriveItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DriveItem
func (r *DriveItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DriveItemUploadablePropertiesRequestBuilder is request builder for DriveItemUploadableProperties
type DriveItemUploadablePropertiesRequestBuilder struct{ BaseRequestBuilder }

// Request returns DriveItemUploadablePropertiesRequest
func (b *DriveItemUploadablePropertiesRequestBuilder) Request() *DriveItemUploadablePropertiesRequest {
	return &DriveItemUploadablePropertiesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DriveItemUploadablePropertiesRequest is request for DriveItemUploadableProperties
type DriveItemUploadablePropertiesRequest struct{ BaseRequest }

// Get performs GET request for DriveItemUploadableProperties
func (r *DriveItemUploadablePropertiesRequest) Get(ctx context.Context) (resObj *DriveItemUploadableProperties, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DriveItemUploadableProperties
func (r *DriveItemUploadablePropertiesRequest) Update(ctx context.Context, reqObj *DriveItemUploadableProperties) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DriveItemUploadableProperties
func (r *DriveItemUploadablePropertiesRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DriveItemVersionRequestBuilder is request builder for DriveItemVersion
type DriveItemVersionRequestBuilder struct{ BaseRequestBuilder }

// Request returns DriveItemVersionRequest
func (b *DriveItemVersionRequestBuilder) Request() *DriveItemVersionRequest {
	return &DriveItemVersionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DriveItemVersionRequest is request for DriveItemVersion
type DriveItemVersionRequest struct{ BaseRequest }

// Get performs GET request for DriveItemVersion
func (r *DriveItemVersionRequest) Get(ctx context.Context) (resObj *DriveItemVersion, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DriveItemVersion
func (r *DriveItemVersionRequest) Update(ctx context.Context, reqObj *DriveItemVersion) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DriveItemVersion
func (r *DriveItemVersionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DriveRecipientRequestBuilder is request builder for DriveRecipient
type DriveRecipientRequestBuilder struct{ BaseRequestBuilder }

// Request returns DriveRecipientRequest
func (b *DriveRecipientRequestBuilder) Request() *DriveRecipientRequest {
	return &DriveRecipientRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DriveRecipientRequest is request for DriveRecipient
type DriveRecipientRequest struct{ BaseRequest }

// Get performs GET request for DriveRecipient
func (r *DriveRecipientRequest) Get(ctx context.Context) (resObj *DriveRecipient, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DriveRecipient
func (r *DriveRecipientRequest) Update(ctx context.Context, reqObj *DriveRecipient) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DriveRecipient
func (r *DriveRecipientRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
