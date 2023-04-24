// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ChannelRequestBuilder is request builder for Channel
type ChannelRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChannelRequest
func (b *ChannelRequestBuilder) Request() *ChannelRequest {
	return &ChannelRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChannelRequest is request for Channel
type ChannelRequest struct{ BaseRequest }

// Get performs GET request for Channel
func (r *ChannelRequest) Get(ctx context.Context) (resObj *Channel, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Channel
func (r *ChannelRequest) Update(ctx context.Context, reqObj *Channel) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Channel
func (r *ChannelRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChannelAddedEventMessageDetailRequestBuilder is request builder for ChannelAddedEventMessageDetail
type ChannelAddedEventMessageDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChannelAddedEventMessageDetailRequest
func (b *ChannelAddedEventMessageDetailRequestBuilder) Request() *ChannelAddedEventMessageDetailRequest {
	return &ChannelAddedEventMessageDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChannelAddedEventMessageDetailRequest is request for ChannelAddedEventMessageDetail
type ChannelAddedEventMessageDetailRequest struct{ BaseRequest }

// Get performs GET request for ChannelAddedEventMessageDetail
func (r *ChannelAddedEventMessageDetailRequest) Get(ctx context.Context) (resObj *ChannelAddedEventMessageDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChannelAddedEventMessageDetail
func (r *ChannelAddedEventMessageDetailRequest) Update(ctx context.Context, reqObj *ChannelAddedEventMessageDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChannelAddedEventMessageDetail
func (r *ChannelAddedEventMessageDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChannelDeletedEventMessageDetailRequestBuilder is request builder for ChannelDeletedEventMessageDetail
type ChannelDeletedEventMessageDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChannelDeletedEventMessageDetailRequest
func (b *ChannelDeletedEventMessageDetailRequestBuilder) Request() *ChannelDeletedEventMessageDetailRequest {
	return &ChannelDeletedEventMessageDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChannelDeletedEventMessageDetailRequest is request for ChannelDeletedEventMessageDetail
type ChannelDeletedEventMessageDetailRequest struct{ BaseRequest }

// Get performs GET request for ChannelDeletedEventMessageDetail
func (r *ChannelDeletedEventMessageDetailRequest) Get(ctx context.Context) (resObj *ChannelDeletedEventMessageDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChannelDeletedEventMessageDetail
func (r *ChannelDeletedEventMessageDetailRequest) Update(ctx context.Context, reqObj *ChannelDeletedEventMessageDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChannelDeletedEventMessageDetail
func (r *ChannelDeletedEventMessageDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChannelDescriptionUpdatedEventMessageDetailRequestBuilder is request builder for ChannelDescriptionUpdatedEventMessageDetail
type ChannelDescriptionUpdatedEventMessageDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChannelDescriptionUpdatedEventMessageDetailRequest
func (b *ChannelDescriptionUpdatedEventMessageDetailRequestBuilder) Request() *ChannelDescriptionUpdatedEventMessageDetailRequest {
	return &ChannelDescriptionUpdatedEventMessageDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChannelDescriptionUpdatedEventMessageDetailRequest is request for ChannelDescriptionUpdatedEventMessageDetail
type ChannelDescriptionUpdatedEventMessageDetailRequest struct{ BaseRequest }

// Get performs GET request for ChannelDescriptionUpdatedEventMessageDetail
func (r *ChannelDescriptionUpdatedEventMessageDetailRequest) Get(ctx context.Context) (resObj *ChannelDescriptionUpdatedEventMessageDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChannelDescriptionUpdatedEventMessageDetail
func (r *ChannelDescriptionUpdatedEventMessageDetailRequest) Update(ctx context.Context, reqObj *ChannelDescriptionUpdatedEventMessageDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChannelDescriptionUpdatedEventMessageDetail
func (r *ChannelDescriptionUpdatedEventMessageDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChannelIdentityRequestBuilder is request builder for ChannelIdentity
type ChannelIdentityRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChannelIdentityRequest
func (b *ChannelIdentityRequestBuilder) Request() *ChannelIdentityRequest {
	return &ChannelIdentityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChannelIdentityRequest is request for ChannelIdentity
type ChannelIdentityRequest struct{ BaseRequest }

// Get performs GET request for ChannelIdentity
func (r *ChannelIdentityRequest) Get(ctx context.Context) (resObj *ChannelIdentity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChannelIdentity
func (r *ChannelIdentityRequest) Update(ctx context.Context, reqObj *ChannelIdentity) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChannelIdentity
func (r *ChannelIdentityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChannelMembersNotificationRecipientRequestBuilder is request builder for ChannelMembersNotificationRecipient
type ChannelMembersNotificationRecipientRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChannelMembersNotificationRecipientRequest
func (b *ChannelMembersNotificationRecipientRequestBuilder) Request() *ChannelMembersNotificationRecipientRequest {
	return &ChannelMembersNotificationRecipientRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChannelMembersNotificationRecipientRequest is request for ChannelMembersNotificationRecipient
type ChannelMembersNotificationRecipientRequest struct{ BaseRequest }

// Get performs GET request for ChannelMembersNotificationRecipient
func (r *ChannelMembersNotificationRecipientRequest) Get(ctx context.Context) (resObj *ChannelMembersNotificationRecipient, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChannelMembersNotificationRecipient
func (r *ChannelMembersNotificationRecipientRequest) Update(ctx context.Context, reqObj *ChannelMembersNotificationRecipient) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChannelMembersNotificationRecipient
func (r *ChannelMembersNotificationRecipientRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChannelModerationSettingsRequestBuilder is request builder for ChannelModerationSettings
type ChannelModerationSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChannelModerationSettingsRequest
func (b *ChannelModerationSettingsRequestBuilder) Request() *ChannelModerationSettingsRequest {
	return &ChannelModerationSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChannelModerationSettingsRequest is request for ChannelModerationSettings
type ChannelModerationSettingsRequest struct{ BaseRequest }

// Get performs GET request for ChannelModerationSettings
func (r *ChannelModerationSettingsRequest) Get(ctx context.Context) (resObj *ChannelModerationSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChannelModerationSettings
func (r *ChannelModerationSettingsRequest) Update(ctx context.Context, reqObj *ChannelModerationSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChannelModerationSettings
func (r *ChannelModerationSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChannelRenamedEventMessageDetailRequestBuilder is request builder for ChannelRenamedEventMessageDetail
type ChannelRenamedEventMessageDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChannelRenamedEventMessageDetailRequest
func (b *ChannelRenamedEventMessageDetailRequestBuilder) Request() *ChannelRenamedEventMessageDetailRequest {
	return &ChannelRenamedEventMessageDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChannelRenamedEventMessageDetailRequest is request for ChannelRenamedEventMessageDetail
type ChannelRenamedEventMessageDetailRequest struct{ BaseRequest }

// Get performs GET request for ChannelRenamedEventMessageDetail
func (r *ChannelRenamedEventMessageDetailRequest) Get(ctx context.Context) (resObj *ChannelRenamedEventMessageDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChannelRenamedEventMessageDetail
func (r *ChannelRenamedEventMessageDetailRequest) Update(ctx context.Context, reqObj *ChannelRenamedEventMessageDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChannelRenamedEventMessageDetail
func (r *ChannelRenamedEventMessageDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChannelSetAsFavoriteByDefaultEventMessageDetailRequestBuilder is request builder for ChannelSetAsFavoriteByDefaultEventMessageDetail
type ChannelSetAsFavoriteByDefaultEventMessageDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChannelSetAsFavoriteByDefaultEventMessageDetailRequest
func (b *ChannelSetAsFavoriteByDefaultEventMessageDetailRequestBuilder) Request() *ChannelSetAsFavoriteByDefaultEventMessageDetailRequest {
	return &ChannelSetAsFavoriteByDefaultEventMessageDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChannelSetAsFavoriteByDefaultEventMessageDetailRequest is request for ChannelSetAsFavoriteByDefaultEventMessageDetail
type ChannelSetAsFavoriteByDefaultEventMessageDetailRequest struct{ BaseRequest }

// Get performs GET request for ChannelSetAsFavoriteByDefaultEventMessageDetail
func (r *ChannelSetAsFavoriteByDefaultEventMessageDetailRequest) Get(ctx context.Context) (resObj *ChannelSetAsFavoriteByDefaultEventMessageDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChannelSetAsFavoriteByDefaultEventMessageDetail
func (r *ChannelSetAsFavoriteByDefaultEventMessageDetailRequest) Update(ctx context.Context, reqObj *ChannelSetAsFavoriteByDefaultEventMessageDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChannelSetAsFavoriteByDefaultEventMessageDetail
func (r *ChannelSetAsFavoriteByDefaultEventMessageDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChannelSummaryRequestBuilder is request builder for ChannelSummary
type ChannelSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChannelSummaryRequest
func (b *ChannelSummaryRequestBuilder) Request() *ChannelSummaryRequest {
	return &ChannelSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChannelSummaryRequest is request for ChannelSummary
type ChannelSummaryRequest struct{ BaseRequest }

// Get performs GET request for ChannelSummary
func (r *ChannelSummaryRequest) Get(ctx context.Context) (resObj *ChannelSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChannelSummary
func (r *ChannelSummaryRequest) Update(ctx context.Context, reqObj *ChannelSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChannelSummary
func (r *ChannelSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChannelUnsetAsFavoriteByDefaultEventMessageDetailRequestBuilder is request builder for ChannelUnsetAsFavoriteByDefaultEventMessageDetail
type ChannelUnsetAsFavoriteByDefaultEventMessageDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChannelUnsetAsFavoriteByDefaultEventMessageDetailRequest
func (b *ChannelUnsetAsFavoriteByDefaultEventMessageDetailRequestBuilder) Request() *ChannelUnsetAsFavoriteByDefaultEventMessageDetailRequest {
	return &ChannelUnsetAsFavoriteByDefaultEventMessageDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChannelUnsetAsFavoriteByDefaultEventMessageDetailRequest is request for ChannelUnsetAsFavoriteByDefaultEventMessageDetail
type ChannelUnsetAsFavoriteByDefaultEventMessageDetailRequest struct{ BaseRequest }

// Get performs GET request for ChannelUnsetAsFavoriteByDefaultEventMessageDetail
func (r *ChannelUnsetAsFavoriteByDefaultEventMessageDetailRequest) Get(ctx context.Context) (resObj *ChannelUnsetAsFavoriteByDefaultEventMessageDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChannelUnsetAsFavoriteByDefaultEventMessageDetail
func (r *ChannelUnsetAsFavoriteByDefaultEventMessageDetailRequest) Update(ctx context.Context, reqObj *ChannelUnsetAsFavoriteByDefaultEventMessageDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChannelUnsetAsFavoriteByDefaultEventMessageDetail
func (r *ChannelUnsetAsFavoriteByDefaultEventMessageDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type ChannelCompleteMigrationRequestBuilder struct{ BaseRequestBuilder }

// CompleteMigration action undocumentedrav
func (b *ChannelRequestBuilder) CompleteMigration(reqObj *ChannelCompleteMigrationRequestParameter) *ChannelCompleteMigrationRequestBuilder {
	bb := &ChannelCompleteMigrationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/CompleteMigration"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ChannelCompleteMigrationRequest struct{ BaseRequest }

func (b *ChannelCompleteMigrationRequestBuilder) Request() *ChannelCompleteMigrationRequest {
	return &ChannelCompleteMigrationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ChannelCompleteMigrationRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type ChannelProvisionEmailRequestBuilder struct{ BaseRequestBuilder }

// ProvisionEmail action undocumentedras
func (b *ChannelRequestBuilder) ProvisionEmail(reqObj *ChannelProvisionEmailRequestParameter) *ChannelProvisionEmailRequestBuilder {
	bb := &ChannelProvisionEmailRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ProvisionEmail"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ChannelProvisionEmailRequest struct{ BaseRequest }

func (b *ChannelProvisionEmailRequestBuilder) Request() *ChannelProvisionEmailRequest {
	return &ChannelProvisionEmailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ChannelProvisionEmailRequest) Post(ctx context.Context) (resObj *ProvisionChannelEmailResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type ChannelRemoveEmailRequestBuilder struct{ BaseRequestBuilder }

// RemoveEmail action undocumentedrav
func (b *ChannelRequestBuilder) RemoveEmail(reqObj *ChannelRemoveEmailRequestParameter) *ChannelRemoveEmailRequestBuilder {
	bb := &ChannelRemoveEmailRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/RemoveEmail"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ChannelRemoveEmailRequest struct{ BaseRequest }

func (b *ChannelRemoveEmailRequestBuilder) Request() *ChannelRemoveEmailRequest {
	return &ChannelRemoveEmailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ChannelRemoveEmailRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}