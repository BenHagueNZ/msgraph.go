// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// AttachmentRequestBuilder is request builder for Attachment
type AttachmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns AttachmentRequest
func (b *AttachmentRequestBuilder) Request() *AttachmentRequest {
	return &AttachmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AttachmentRequest is request for Attachment
type AttachmentRequest struct{ BaseRequest }

// Get performs GET request for Attachment
func (r *AttachmentRequest) Get(ctx context.Context) (resObj *Attachment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Attachment
func (r *AttachmentRequest) Update(ctx context.Context, reqObj *Attachment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Attachment
func (r *AttachmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AttachmentBaseRequestBuilder is request builder for AttachmentBase
type AttachmentBaseRequestBuilder struct{ BaseRequestBuilder }

// Request returns AttachmentBaseRequest
func (b *AttachmentBaseRequestBuilder) Request() *AttachmentBaseRequest {
	return &AttachmentBaseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AttachmentBaseRequest is request for AttachmentBase
type AttachmentBaseRequest struct{ BaseRequest }

// Get performs GET request for AttachmentBase
func (r *AttachmentBaseRequest) Get(ctx context.Context) (resObj *AttachmentBase, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AttachmentBase
func (r *AttachmentBaseRequest) Update(ctx context.Context, reqObj *AttachmentBase) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AttachmentBase
func (r *AttachmentBaseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AttachmentInfoRequestBuilder is request builder for AttachmentInfo
type AttachmentInfoRequestBuilder struct{ BaseRequestBuilder }

// Request returns AttachmentInfoRequest
func (b *AttachmentInfoRequestBuilder) Request() *AttachmentInfoRequest {
	return &AttachmentInfoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AttachmentInfoRequest is request for AttachmentInfo
type AttachmentInfoRequest struct{ BaseRequest }

// Get performs GET request for AttachmentInfo
func (r *AttachmentInfoRequest) Get(ctx context.Context) (resObj *AttachmentInfo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AttachmentInfo
func (r *AttachmentInfoRequest) Update(ctx context.Context, reqObj *AttachmentInfo) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AttachmentInfo
func (r *AttachmentInfoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AttachmentItemRequestBuilder is request builder for AttachmentItem
type AttachmentItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns AttachmentItemRequest
func (b *AttachmentItemRequestBuilder) Request() *AttachmentItemRequest {
	return &AttachmentItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AttachmentItemRequest is request for AttachmentItem
type AttachmentItemRequest struct{ BaseRequest }

// Get performs GET request for AttachmentItem
func (r *AttachmentItemRequest) Get(ctx context.Context) (resObj *AttachmentItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AttachmentItem
func (r *AttachmentItemRequest) Update(ctx context.Context, reqObj *AttachmentItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AttachmentItem
func (r *AttachmentItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AttachmentSessionRequestBuilder is request builder for AttachmentSession
type AttachmentSessionRequestBuilder struct{ BaseRequestBuilder }

// Request returns AttachmentSessionRequest
func (b *AttachmentSessionRequestBuilder) Request() *AttachmentSessionRequest {
	return &AttachmentSessionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AttachmentSessionRequest is request for AttachmentSession
type AttachmentSessionRequest struct{ BaseRequest }

// Get performs GET request for AttachmentSession
func (r *AttachmentSessionRequest) Get(ctx context.Context) (resObj *AttachmentSession, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AttachmentSession
func (r *AttachmentSessionRequest) Update(ctx context.Context, reqObj *AttachmentSession) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AttachmentSession
func (r *AttachmentSessionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type AttachmentCollectionCreateUploadSessionRequestBuilder struct{ BaseRequestBuilder }

// CreateUploadSession action undocumentedras
func (b *EventAttachmentsCollectionRequestBuilder) CreateUploadSession(reqObj *AttachmentCollectionCreateUploadSessionRequestParameter) *AttachmentCollectionCreateUploadSessionRequestBuilder {
	bb := &AttachmentCollectionCreateUploadSessionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/CreateUploadSession"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// CreateUploadSession action undocumentedras
func (b *MessageAttachmentsCollectionRequestBuilder) CreateUploadSession(reqObj *AttachmentCollectionCreateUploadSessionRequestParameter) *AttachmentCollectionCreateUploadSessionRequestBuilder {
	bb := &AttachmentCollectionCreateUploadSessionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/CreateUploadSession"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// CreateUploadSession action undocumentedras
func (b *PostAttachmentsCollectionRequestBuilder) CreateUploadSession(reqObj *AttachmentCollectionCreateUploadSessionRequestParameter) *AttachmentCollectionCreateUploadSessionRequestBuilder {
	bb := &AttachmentCollectionCreateUploadSessionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/CreateUploadSession"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type AttachmentCollectionCreateUploadSessionRequest struct{ BaseRequest }

func (b *AttachmentCollectionCreateUploadSessionRequestBuilder) Request() *AttachmentCollectionCreateUploadSessionRequest {
	return &AttachmentCollectionCreateUploadSessionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *AttachmentCollectionCreateUploadSessionRequest) Post(ctx context.Context) (resObj *UploadSession, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type AttachmentBaseCollectionCreateUploadSessionRequestBuilder struct{ BaseRequestBuilder }

// CreateUploadSession action undocumentedras
func (b *TodoTaskAttachmentsCollectionRequestBuilder) CreateUploadSession(reqObj *AttachmentBaseCollectionCreateUploadSessionRequestParameter) *AttachmentBaseCollectionCreateUploadSessionRequestBuilder {
	bb := &AttachmentBaseCollectionCreateUploadSessionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/CreateUploadSession"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type AttachmentBaseCollectionCreateUploadSessionRequest struct{ BaseRequest }

func (b *AttachmentBaseCollectionCreateUploadSessionRequestBuilder) Request() *AttachmentBaseCollectionCreateUploadSessionRequest {
	return &AttachmentBaseCollectionCreateUploadSessionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *AttachmentBaseCollectionCreateUploadSessionRequest) Post(ctx context.Context) (resObj *UploadSession, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
