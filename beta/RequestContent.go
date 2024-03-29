// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ContentApprovalStatusColumnRequestBuilder is request builder for ContentApprovalStatusColumn
type ContentApprovalStatusColumnRequestBuilder struct{ BaseRequestBuilder }

// Request returns ContentApprovalStatusColumnRequest
func (b *ContentApprovalStatusColumnRequestBuilder) Request() *ContentApprovalStatusColumnRequest {
	return &ContentApprovalStatusColumnRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ContentApprovalStatusColumnRequest is request for ContentApprovalStatusColumn
type ContentApprovalStatusColumnRequest struct{ BaseRequest }

// Get performs GET request for ContentApprovalStatusColumn
func (r *ContentApprovalStatusColumnRequest) Get(ctx context.Context) (resObj *ContentApprovalStatusColumn, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ContentApprovalStatusColumn
func (r *ContentApprovalStatusColumnRequest) Update(ctx context.Context, reqObj *ContentApprovalStatusColumn) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ContentApprovalStatusColumn
func (r *ContentApprovalStatusColumnRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ContentClassificationRequestBuilder is request builder for ContentClassification
type ContentClassificationRequestBuilder struct{ BaseRequestBuilder }

// Request returns ContentClassificationRequest
func (b *ContentClassificationRequestBuilder) Request() *ContentClassificationRequest {
	return &ContentClassificationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ContentClassificationRequest is request for ContentClassification
type ContentClassificationRequest struct{ BaseRequest }

// Get performs GET request for ContentClassification
func (r *ContentClassificationRequest) Get(ctx context.Context) (resObj *ContentClassification, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ContentClassification
func (r *ContentClassificationRequest) Update(ctx context.Context, reqObj *ContentClassification) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ContentClassification
func (r *ContentClassificationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ContentInfoRequestBuilder is request builder for ContentInfo
type ContentInfoRequestBuilder struct{ BaseRequestBuilder }

// Request returns ContentInfoRequest
func (b *ContentInfoRequestBuilder) Request() *ContentInfoRequest {
	return &ContentInfoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ContentInfoRequest is request for ContentInfo
type ContentInfoRequest struct{ BaseRequest }

// Get performs GET request for ContentInfo
func (r *ContentInfoRequest) Get(ctx context.Context) (resObj *ContentInfo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ContentInfo
func (r *ContentInfoRequest) Update(ctx context.Context, reqObj *ContentInfo) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ContentInfo
func (r *ContentInfoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ContentMetadataRequestBuilder is request builder for ContentMetadata
type ContentMetadataRequestBuilder struct{ BaseRequestBuilder }

// Request returns ContentMetadataRequest
func (b *ContentMetadataRequestBuilder) Request() *ContentMetadataRequest {
	return &ContentMetadataRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ContentMetadataRequest is request for ContentMetadata
type ContentMetadataRequest struct{ BaseRequest }

// Get performs GET request for ContentMetadata
func (r *ContentMetadataRequest) Get(ctx context.Context) (resObj *ContentMetadata, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ContentMetadata
func (r *ContentMetadataRequest) Update(ctx context.Context, reqObj *ContentMetadata) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ContentMetadata
func (r *ContentMetadataRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ContentPropertiesRequestBuilder is request builder for ContentProperties
type ContentPropertiesRequestBuilder struct{ BaseRequestBuilder }

// Request returns ContentPropertiesRequest
func (b *ContentPropertiesRequestBuilder) Request() *ContentPropertiesRequest {
	return &ContentPropertiesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ContentPropertiesRequest is request for ContentProperties
type ContentPropertiesRequest struct{ BaseRequest }

// Get performs GET request for ContentProperties
func (r *ContentPropertiesRequest) Get(ctx context.Context) (resObj *ContentProperties, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ContentProperties
func (r *ContentPropertiesRequest) Update(ctx context.Context, reqObj *ContentProperties) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ContentProperties
func (r *ContentPropertiesRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

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

// ContentTypeInfoRequestBuilder is request builder for ContentTypeInfo
type ContentTypeInfoRequestBuilder struct{ BaseRequestBuilder }

// Request returns ContentTypeInfoRequest
func (b *ContentTypeInfoRequestBuilder) Request() *ContentTypeInfoRequest {
	return &ContentTypeInfoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ContentTypeInfoRequest is request for ContentTypeInfo
type ContentTypeInfoRequest struct{ BaseRequest }

// Get performs GET request for ContentTypeInfo
func (r *ContentTypeInfoRequest) Get(ctx context.Context) (resObj *ContentTypeInfo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ContentTypeInfo
func (r *ContentTypeInfoRequest) Update(ctx context.Context, reqObj *ContentTypeInfo) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ContentTypeInfo
func (r *ContentTypeInfoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ContentTypeOrderRequestBuilder is request builder for ContentTypeOrder
type ContentTypeOrderRequestBuilder struct{ BaseRequestBuilder }

// Request returns ContentTypeOrderRequest
func (b *ContentTypeOrderRequestBuilder) Request() *ContentTypeOrderRequest {
	return &ContentTypeOrderRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ContentTypeOrderRequest is request for ContentTypeOrder
type ContentTypeOrderRequest struct{ BaseRequest }

// Get performs GET request for ContentTypeOrder
func (r *ContentTypeOrderRequest) Get(ctx context.Context) (resObj *ContentTypeOrder, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ContentTypeOrder
func (r *ContentTypeOrderRequest) Update(ctx context.Context, reqObj *ContentTypeOrder) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ContentTypeOrder
func (r *ContentTypeOrderRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type ContentTypeCollectionAddCopyRequestBuilder struct{ BaseRequestBuilder }

// AddCopy action undocumentedras
func (b *ContentTypeBaseTypesCollectionRequestBuilder) AddCopy(reqObj *ContentTypeCollectionAddCopyRequestParameter) *ContentTypeCollectionAddCopyRequestBuilder {
	bb := &ContentTypeCollectionAddCopyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/AddCopy"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// AddCopy action undocumentedras
func (b *ListContentTypesCollectionRequestBuilder) AddCopy(reqObj *ContentTypeCollectionAddCopyRequestParameter) *ContentTypeCollectionAddCopyRequestBuilder {
	bb := &ContentTypeCollectionAddCopyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/AddCopy"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// AddCopy action undocumentedras
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

// AddCopyFromContentTypeHub action undocumentedras
func (b *ContentTypeBaseTypesCollectionRequestBuilder) AddCopyFromContentTypeHub(reqObj *ContentTypeCollectionAddCopyFromContentTypeHubRequestParameter) *ContentTypeCollectionAddCopyFromContentTypeHubRequestBuilder {
	bb := &ContentTypeCollectionAddCopyFromContentTypeHubRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/AddCopyFromContentTypeHub"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// AddCopyFromContentTypeHub action undocumentedras
func (b *ListContentTypesCollectionRequestBuilder) AddCopyFromContentTypeHub(reqObj *ContentTypeCollectionAddCopyFromContentTypeHubRequestParameter) *ContentTypeCollectionAddCopyFromContentTypeHubRequestBuilder {
	bb := &ContentTypeCollectionAddCopyFromContentTypeHubRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/AddCopyFromContentTypeHub"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// AddCopyFromContentTypeHub action undocumentedras
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

type ContentTypePublishRequestBuilder struct{ BaseRequestBuilder }

// Publish action undocumentedrav
func (b *ContentTypeRequestBuilder) Publish(reqObj *ContentTypePublishRequestParameter) *ContentTypePublishRequestBuilder {
	bb := &ContentTypePublishRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Publish"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ContentTypePublishRequest struct{ BaseRequest }

func (b *ContentTypePublishRequestBuilder) Request() *ContentTypePublishRequest {
	return &ContentTypePublishRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ContentTypePublishRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type ContentTypeUnpublishRequestBuilder struct{ BaseRequestBuilder }

// Unpublish action undocumentedrav
func (b *ContentTypeRequestBuilder) Unpublish(reqObj *ContentTypeUnpublishRequestParameter) *ContentTypeUnpublishRequestBuilder {
	bb := &ContentTypeUnpublishRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Unpublish"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ContentTypeUnpublishRequest struct{ BaseRequest }

func (b *ContentTypeUnpublishRequestBuilder) Request() *ContentTypeUnpublishRequest {
	return &ContentTypeUnpublishRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ContentTypeUnpublishRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type ContentTypeAssociateWithHubSitesRequestBuilder struct{ BaseRequestBuilder }

// AssociateWithHubSites action undocumentedrav
func (b *ContentTypeRequestBuilder) AssociateWithHubSites(reqObj *ContentTypeAssociateWithHubSitesRequestParameter) *ContentTypeAssociateWithHubSitesRequestBuilder {
	bb := &ContentTypeAssociateWithHubSitesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/AssociateWithHubSites"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ContentTypeAssociateWithHubSitesRequest struct{ BaseRequest }

func (b *ContentTypeAssociateWithHubSitesRequestBuilder) Request() *ContentTypeAssociateWithHubSitesRequest {
	return &ContentTypeAssociateWithHubSitesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ContentTypeAssociateWithHubSitesRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type ContentTypeCopyToDefaultContentLocationRequestBuilder struct{ BaseRequestBuilder }

// CopyToDefaultContentLocation action undocumentedrav
func (b *ContentTypeRequestBuilder) CopyToDefaultContentLocation(reqObj *ContentTypeCopyToDefaultContentLocationRequestParameter) *ContentTypeCopyToDefaultContentLocationRequestBuilder {
	bb := &ContentTypeCopyToDefaultContentLocationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/CopyToDefaultContentLocation"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ContentTypeCopyToDefaultContentLocationRequest struct{ BaseRequest }

func (b *ContentTypeCopyToDefaultContentLocationRequestBuilder) Request() *ContentTypeCopyToDefaultContentLocationRequest {
	return &ContentTypeCopyToDefaultContentLocationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ContentTypeCopyToDefaultContentLocationRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
