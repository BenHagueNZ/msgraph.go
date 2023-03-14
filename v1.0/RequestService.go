// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ServiceAnnouncementRequestBuilder is request builder for ServiceAnnouncement
type ServiceAnnouncementRequestBuilder struct{ BaseRequestBuilder }

// Request returns ServiceAnnouncementRequest
func (b *ServiceAnnouncementRequestBuilder) Request() *ServiceAnnouncementRequest {
	return &ServiceAnnouncementRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ServiceAnnouncementRequest is request for ServiceAnnouncement
type ServiceAnnouncementRequest struct{ BaseRequest }

// Get performs GET request for ServiceAnnouncement
func (r *ServiceAnnouncementRequest) Get(ctx context.Context) (resObj *ServiceAnnouncement, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ServiceAnnouncement
func (r *ServiceAnnouncementRequest) Update(ctx context.Context, reqObj *ServiceAnnouncement) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ServiceAnnouncement
func (r *ServiceAnnouncementRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ServiceAnnouncementAttachmentRequestBuilder is request builder for ServiceAnnouncementAttachment
type ServiceAnnouncementAttachmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns ServiceAnnouncementAttachmentRequest
func (b *ServiceAnnouncementAttachmentRequestBuilder) Request() *ServiceAnnouncementAttachmentRequest {
	return &ServiceAnnouncementAttachmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ServiceAnnouncementAttachmentRequest is request for ServiceAnnouncementAttachment
type ServiceAnnouncementAttachmentRequest struct{ BaseRequest }

// Get performs GET request for ServiceAnnouncementAttachment
func (r *ServiceAnnouncementAttachmentRequest) Get(ctx context.Context) (resObj *ServiceAnnouncementAttachment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ServiceAnnouncementAttachment
func (r *ServiceAnnouncementAttachmentRequest) Update(ctx context.Context, reqObj *ServiceAnnouncementAttachment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ServiceAnnouncementAttachment
func (r *ServiceAnnouncementAttachmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ServiceAnnouncementBaseRequestBuilder is request builder for ServiceAnnouncementBase
type ServiceAnnouncementBaseRequestBuilder struct{ BaseRequestBuilder }

// Request returns ServiceAnnouncementBaseRequest
func (b *ServiceAnnouncementBaseRequestBuilder) Request() *ServiceAnnouncementBaseRequest {
	return &ServiceAnnouncementBaseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ServiceAnnouncementBaseRequest is request for ServiceAnnouncementBase
type ServiceAnnouncementBaseRequest struct{ BaseRequest }

// Get performs GET request for ServiceAnnouncementBase
func (r *ServiceAnnouncementBaseRequest) Get(ctx context.Context) (resObj *ServiceAnnouncementBase, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ServiceAnnouncementBase
func (r *ServiceAnnouncementBaseRequest) Update(ctx context.Context, reqObj *ServiceAnnouncementBase) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ServiceAnnouncementBase
func (r *ServiceAnnouncementBaseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ServiceHealthRequestBuilder is request builder for ServiceHealth
type ServiceHealthRequestBuilder struct{ BaseRequestBuilder }

// Request returns ServiceHealthRequest
func (b *ServiceHealthRequestBuilder) Request() *ServiceHealthRequest {
	return &ServiceHealthRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ServiceHealthRequest is request for ServiceHealth
type ServiceHealthRequest struct{ BaseRequest }

// Get performs GET request for ServiceHealth
func (r *ServiceHealthRequest) Get(ctx context.Context) (resObj *ServiceHealth, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ServiceHealth
func (r *ServiceHealthRequest) Update(ctx context.Context, reqObj *ServiceHealth) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ServiceHealth
func (r *ServiceHealthRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ServiceHealthIssueRequestBuilder is request builder for ServiceHealthIssue
type ServiceHealthIssueRequestBuilder struct{ BaseRequestBuilder }

// Request returns ServiceHealthIssueRequest
func (b *ServiceHealthIssueRequestBuilder) Request() *ServiceHealthIssueRequest {
	return &ServiceHealthIssueRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ServiceHealthIssueRequest is request for ServiceHealthIssue
type ServiceHealthIssueRequest struct{ BaseRequest }

// Get performs GET request for ServiceHealthIssue
func (r *ServiceHealthIssueRequest) Get(ctx context.Context) (resObj *ServiceHealthIssue, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ServiceHealthIssue
func (r *ServiceHealthIssueRequest) Update(ctx context.Context, reqObj *ServiceHealthIssue) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ServiceHealthIssue
func (r *ServiceHealthIssueRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ServiceHealthIssuePostRequestBuilder is request builder for ServiceHealthIssuePost
type ServiceHealthIssuePostRequestBuilder struct{ BaseRequestBuilder }

// Request returns ServiceHealthIssuePostRequest
func (b *ServiceHealthIssuePostRequestBuilder) Request() *ServiceHealthIssuePostRequest {
	return &ServiceHealthIssuePostRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ServiceHealthIssuePostRequest is request for ServiceHealthIssuePost
type ServiceHealthIssuePostRequest struct{ BaseRequest }

// Get performs GET request for ServiceHealthIssuePost
func (r *ServiceHealthIssuePostRequest) Get(ctx context.Context) (resObj *ServiceHealthIssuePost, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ServiceHealthIssuePost
func (r *ServiceHealthIssuePostRequest) Update(ctx context.Context, reqObj *ServiceHealthIssuePost) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ServiceHealthIssuePost
func (r *ServiceHealthIssuePostRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ServiceHostedMediaConfigRequestBuilder is request builder for ServiceHostedMediaConfig
type ServiceHostedMediaConfigRequestBuilder struct{ BaseRequestBuilder }

// Request returns ServiceHostedMediaConfigRequest
func (b *ServiceHostedMediaConfigRequestBuilder) Request() *ServiceHostedMediaConfigRequest {
	return &ServiceHostedMediaConfigRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ServiceHostedMediaConfigRequest is request for ServiceHostedMediaConfig
type ServiceHostedMediaConfigRequest struct{ BaseRequest }

// Get performs GET request for ServiceHostedMediaConfig
func (r *ServiceHostedMediaConfigRequest) Get(ctx context.Context) (resObj *ServiceHostedMediaConfig, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ServiceHostedMediaConfig
func (r *ServiceHostedMediaConfigRequest) Update(ctx context.Context, reqObj *ServiceHostedMediaConfig) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ServiceHostedMediaConfig
func (r *ServiceHostedMediaConfigRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ServicePlanInfoRequestBuilder is request builder for ServicePlanInfo
type ServicePlanInfoRequestBuilder struct{ BaseRequestBuilder }

// Request returns ServicePlanInfoRequest
func (b *ServicePlanInfoRequestBuilder) Request() *ServicePlanInfoRequest {
	return &ServicePlanInfoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ServicePlanInfoRequest is request for ServicePlanInfo
type ServicePlanInfoRequest struct{ BaseRequest }

// Get performs GET request for ServicePlanInfo
func (r *ServicePlanInfoRequest) Get(ctx context.Context) (resObj *ServicePlanInfo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ServicePlanInfo
func (r *ServicePlanInfoRequest) Update(ctx context.Context, reqObj *ServicePlanInfo) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ServicePlanInfo
func (r *ServicePlanInfoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ServicePrincipalRequestBuilder is request builder for ServicePrincipal
type ServicePrincipalRequestBuilder struct{ BaseRequestBuilder }

// Request returns ServicePrincipalRequest
func (b *ServicePrincipalRequestBuilder) Request() *ServicePrincipalRequest {
	return &ServicePrincipalRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ServicePrincipalRequest is request for ServicePrincipal
type ServicePrincipalRequest struct{ BaseRequest }

// Get performs GET request for ServicePrincipal
func (r *ServicePrincipalRequest) Get(ctx context.Context) (resObj *ServicePrincipal, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ServicePrincipal
func (r *ServicePrincipalRequest) Update(ctx context.Context, reqObj *ServicePrincipal) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ServicePrincipal
func (r *ServicePrincipalRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ServicePrincipalIdentityRequestBuilder is request builder for ServicePrincipalIdentity
type ServicePrincipalIdentityRequestBuilder struct{ BaseRequestBuilder }

// Request returns ServicePrincipalIdentityRequest
func (b *ServicePrincipalIdentityRequestBuilder) Request() *ServicePrincipalIdentityRequest {
	return &ServicePrincipalIdentityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ServicePrincipalIdentityRequest is request for ServicePrincipalIdentity
type ServicePrincipalIdentityRequest struct{ BaseRequest }

// Get performs GET request for ServicePrincipalIdentity
func (r *ServicePrincipalIdentityRequest) Get(ctx context.Context) (resObj *ServicePrincipalIdentity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ServicePrincipalIdentity
func (r *ServicePrincipalIdentityRequest) Update(ctx context.Context, reqObj *ServicePrincipalIdentity) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ServicePrincipalIdentity
func (r *ServicePrincipalIdentityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ServicePrincipalRiskDetectionRequestBuilder is request builder for ServicePrincipalRiskDetection
type ServicePrincipalRiskDetectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns ServicePrincipalRiskDetectionRequest
func (b *ServicePrincipalRiskDetectionRequestBuilder) Request() *ServicePrincipalRiskDetectionRequest {
	return &ServicePrincipalRiskDetectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ServicePrincipalRiskDetectionRequest is request for ServicePrincipalRiskDetection
type ServicePrincipalRiskDetectionRequest struct{ BaseRequest }

// Get performs GET request for ServicePrincipalRiskDetection
func (r *ServicePrincipalRiskDetectionRequest) Get(ctx context.Context) (resObj *ServicePrincipalRiskDetection, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ServicePrincipalRiskDetection
func (r *ServicePrincipalRiskDetectionRequest) Update(ctx context.Context, reqObj *ServicePrincipalRiskDetection) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ServicePrincipalRiskDetection
func (r *ServicePrincipalRiskDetectionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ServiceUpdateMessageRequestBuilder is request builder for ServiceUpdateMessage
type ServiceUpdateMessageRequestBuilder struct{ BaseRequestBuilder }

// Request returns ServiceUpdateMessageRequest
func (b *ServiceUpdateMessageRequestBuilder) Request() *ServiceUpdateMessageRequest {
	return &ServiceUpdateMessageRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ServiceUpdateMessageRequest is request for ServiceUpdateMessage
type ServiceUpdateMessageRequest struct{ BaseRequest }

// Get performs GET request for ServiceUpdateMessage
func (r *ServiceUpdateMessageRequest) Get(ctx context.Context) (resObj *ServiceUpdateMessage, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ServiceUpdateMessage
func (r *ServiceUpdateMessageRequest) Update(ctx context.Context, reqObj *ServiceUpdateMessage) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ServiceUpdateMessage
func (r *ServiceUpdateMessageRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ServiceUpdateMessageViewpointRequestBuilder is request builder for ServiceUpdateMessageViewpoint
type ServiceUpdateMessageViewpointRequestBuilder struct{ BaseRequestBuilder }

// Request returns ServiceUpdateMessageViewpointRequest
func (b *ServiceUpdateMessageViewpointRequestBuilder) Request() *ServiceUpdateMessageViewpointRequest {
	return &ServiceUpdateMessageViewpointRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ServiceUpdateMessageViewpointRequest is request for ServiceUpdateMessageViewpoint
type ServiceUpdateMessageViewpointRequest struct{ BaseRequest }

// Get performs GET request for ServiceUpdateMessageViewpoint
func (r *ServiceUpdateMessageViewpointRequest) Get(ctx context.Context) (resObj *ServiceUpdateMessageViewpoint, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ServiceUpdateMessageViewpoint
func (r *ServiceUpdateMessageViewpointRequest) Update(ctx context.Context, reqObj *ServiceUpdateMessageViewpoint) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ServiceUpdateMessageViewpoint
func (r *ServiceUpdateMessageViewpointRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type ServiceUpdateMessageCollectionArchiveRequestBuilder struct{ BaseRequestBuilder }

// Archive action undocumented
func (b *ServiceAnnouncementMessagesCollectionRequestBuilder) Archive(reqObj *ServiceUpdateMessageCollectionArchiveRequestParameter) *ServiceUpdateMessageCollectionArchiveRequestBuilder {
	bb := &ServiceUpdateMessageCollectionArchiveRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Archive"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ServiceUpdateMessageCollectionArchiveRequest struct{ BaseRequest }

func (b *ServiceUpdateMessageCollectionArchiveRequestBuilder) Request() *ServiceUpdateMessageCollectionArchiveRequest {
	return &ServiceUpdateMessageCollectionArchiveRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ServiceUpdateMessageCollectionArchiveRequest) Post(ctx context.Context) (resObj *bool, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type ServiceUpdateMessageCollectionFavoriteRequestBuilder struct{ BaseRequestBuilder }

// Favorite action undocumented
func (b *ServiceAnnouncementMessagesCollectionRequestBuilder) Favorite(reqObj *ServiceUpdateMessageCollectionFavoriteRequestParameter) *ServiceUpdateMessageCollectionFavoriteRequestBuilder {
	bb := &ServiceUpdateMessageCollectionFavoriteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Favorite"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ServiceUpdateMessageCollectionFavoriteRequest struct{ BaseRequest }

func (b *ServiceUpdateMessageCollectionFavoriteRequestBuilder) Request() *ServiceUpdateMessageCollectionFavoriteRequest {
	return &ServiceUpdateMessageCollectionFavoriteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ServiceUpdateMessageCollectionFavoriteRequest) Post(ctx context.Context) (resObj *bool, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type ServiceUpdateMessageCollectionMarkReadRequestBuilder struct{ BaseRequestBuilder }

// MarkRead action undocumented
func (b *ServiceAnnouncementMessagesCollectionRequestBuilder) MarkRead(reqObj *ServiceUpdateMessageCollectionMarkReadRequestParameter) *ServiceUpdateMessageCollectionMarkReadRequestBuilder {
	bb := &ServiceUpdateMessageCollectionMarkReadRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/MarkRead"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ServiceUpdateMessageCollectionMarkReadRequest struct{ BaseRequest }

func (b *ServiceUpdateMessageCollectionMarkReadRequestBuilder) Request() *ServiceUpdateMessageCollectionMarkReadRequest {
	return &ServiceUpdateMessageCollectionMarkReadRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ServiceUpdateMessageCollectionMarkReadRequest) Post(ctx context.Context) (resObj *bool, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type ServiceUpdateMessageCollectionMarkUnreadRequestBuilder struct{ BaseRequestBuilder }

// MarkUnread action undocumented
func (b *ServiceAnnouncementMessagesCollectionRequestBuilder) MarkUnread(reqObj *ServiceUpdateMessageCollectionMarkUnreadRequestParameter) *ServiceUpdateMessageCollectionMarkUnreadRequestBuilder {
	bb := &ServiceUpdateMessageCollectionMarkUnreadRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/MarkUnread"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ServiceUpdateMessageCollectionMarkUnreadRequest struct{ BaseRequest }

func (b *ServiceUpdateMessageCollectionMarkUnreadRequestBuilder) Request() *ServiceUpdateMessageCollectionMarkUnreadRequest {
	return &ServiceUpdateMessageCollectionMarkUnreadRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ServiceUpdateMessageCollectionMarkUnreadRequest) Post(ctx context.Context) (resObj *bool, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type ServiceUpdateMessageCollectionUnarchiveRequestBuilder struct{ BaseRequestBuilder }

// Unarchive action undocumented
func (b *ServiceAnnouncementMessagesCollectionRequestBuilder) Unarchive(reqObj *ServiceUpdateMessageCollectionUnarchiveRequestParameter) *ServiceUpdateMessageCollectionUnarchiveRequestBuilder {
	bb := &ServiceUpdateMessageCollectionUnarchiveRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Unarchive"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ServiceUpdateMessageCollectionUnarchiveRequest struct{ BaseRequest }

func (b *ServiceUpdateMessageCollectionUnarchiveRequestBuilder) Request() *ServiceUpdateMessageCollectionUnarchiveRequest {
	return &ServiceUpdateMessageCollectionUnarchiveRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ServiceUpdateMessageCollectionUnarchiveRequest) Post(ctx context.Context) (resObj *bool, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type ServiceUpdateMessageCollectionUnfavoriteRequestBuilder struct{ BaseRequestBuilder }

// Unfavorite action undocumented
func (b *ServiceAnnouncementMessagesCollectionRequestBuilder) Unfavorite(reqObj *ServiceUpdateMessageCollectionUnfavoriteRequestParameter) *ServiceUpdateMessageCollectionUnfavoriteRequestBuilder {
	bb := &ServiceUpdateMessageCollectionUnfavoriteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Unfavorite"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ServiceUpdateMessageCollectionUnfavoriteRequest struct{ BaseRequest }

func (b *ServiceUpdateMessageCollectionUnfavoriteRequestBuilder) Request() *ServiceUpdateMessageCollectionUnfavoriteRequest {
	return &ServiceUpdateMessageCollectionUnfavoriteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ServiceUpdateMessageCollectionUnfavoriteRequest) Post(ctx context.Context) (resObj *bool, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
