// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ContentSharingSessionRequestBuilder is request builder for ContentSharingSession
type ContentSharingSessionRequestBuilder struct{ BaseRequestBuilder }

// Request returns ContentSharingSessionRequest
func (b *ContentSharingSessionRequestBuilder) Request() *ContentSharingSessionRequest {
	return &ContentSharingSessionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ContentSharingSessionRequest is request for ContentSharingSession
type ContentSharingSessionRequest struct{ BaseRequest }

// Get performs GET request for ContentSharingSession
func (r *ContentSharingSessionRequest) Get(ctx context.Context) (resObj *ContentSharingSession, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ContentSharingSession
func (r *ContentSharingSessionRequest) Update(ctx context.Context, reqObj *ContentSharingSession) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ContentSharingSession
func (r *ContentSharingSessionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ContentTypeRequestBuilder is request builder for ContentType
type ContentTypeRequestBuilder struct{ BaseRequestBuilder }

// Request returns ContentTypeRequest
func (b *ContentTypeRequestBuilder) Request() *ContentTypeRequest {
	return &ContentTypeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ContentTypeRequest is request for ContentType
type ContentTypeRequest struct{ BaseRequest }

// Get performs GET request for ContentType
func (r *ContentTypeRequest) Get(ctx context.Context) (resObj *ContentType, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ContentType
func (r *ContentTypeRequest) Update(ctx context.Context, reqObj *ContentType) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ContentType
func (r *ContentTypeRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type ContentTypeCollectionAddCopyRequestBuilder struct{ BaseRequestBuilder }

// AddCopy action undocumented
func (b *ContentTypeBaseTypesCollectionRequestBuilder) AddCopy(reqObj *ContentTypeCollectionAddCopyRequestParameter) *ContentTypeCollectionAddCopyRequestBuilder {
	bb := &ContentTypeCollectionAddCopyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/AddCopy"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// AddCopy action undocumented
func (b *ListContentTypesCollectionRequestBuilder) AddCopy(reqObj *ContentTypeCollectionAddCopyRequestParameter) *ContentTypeCollectionAddCopyRequestBuilder {
	bb := &ContentTypeCollectionAddCopyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/AddCopy"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// AddCopy action undocumented
func (b *SiteContentTypesCollectionRequestBuilder) AddCopy(reqObj *ContentTypeCollectionAddCopyRequestParameter) *ContentTypeCollectionAddCopyRequestBuilder {
	bb := &ContentTypeCollectionAddCopyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/AddCopy"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ContentTypeCollectionAddCopyRequest struct{ BaseRequest }

func (b *ContentTypeCollectionAddCopyRequestBuilder) Request() *ContentTypeCollectionAddCopyRequest {
	return &ContentTypeCollectionAddCopyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ContentTypeCollectionAddCopyRequest) Post(ctx context.Context) (resObj *ContentType, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type ContentTypeCollectionAddCopyFromContentTypeHubRequestBuilder struct{ BaseRequestBuilder }

// AddCopyFromContentTypeHub action undocumented
func (b *ContentTypeBaseTypesCollectionRequestBuilder) AddCopyFromContentTypeHub(reqObj *ContentTypeCollectionAddCopyFromContentTypeHubRequestParameter) *ContentTypeCollectionAddCopyFromContentTypeHubRequestBuilder {
	bb := &ContentTypeCollectionAddCopyFromContentTypeHubRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/AddCopyFromContentTypeHub"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// AddCopyFromContentTypeHub action undocumented
func (b *ListContentTypesCollectionRequestBuilder) AddCopyFromContentTypeHub(reqObj *ContentTypeCollectionAddCopyFromContentTypeHubRequestParameter) *ContentTypeCollectionAddCopyFromContentTypeHubRequestBuilder {
	bb := &ContentTypeCollectionAddCopyFromContentTypeHubRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/AddCopyFromContentTypeHub"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// AddCopyFromContentTypeHub action undocumented
func (b *SiteContentTypesCollectionRequestBuilder) AddCopyFromContentTypeHub(reqObj *ContentTypeCollectionAddCopyFromContentTypeHubRequestParameter) *ContentTypeCollectionAddCopyFromContentTypeHubRequestBuilder {
	bb := &ContentTypeCollectionAddCopyFromContentTypeHubRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/AddCopyFromContentTypeHub"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ContentTypeCollectionAddCopyFromContentTypeHubRequest struct{ BaseRequest }

func (b *ContentTypeCollectionAddCopyFromContentTypeHubRequestBuilder) Request() *ContentTypeCollectionAddCopyFromContentTypeHubRequest {
	return &ContentTypeCollectionAddCopyFromContentTypeHubRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ContentTypeCollectionAddCopyFromContentTypeHubRequest) Post(ctx context.Context) (resObj *ContentType, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
