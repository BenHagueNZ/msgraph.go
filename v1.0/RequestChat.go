// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ChatRequestBuilder is request builder for Chat
type ChatRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChatRequest
func (b *ChatRequestBuilder) Request() *ChatRequest {
	return &ChatRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChatRequest is request for Chat
type ChatRequest struct{ BaseRequest }

// Get performs GET request for Chat
func (r *ChatRequest) Get(ctx context.Context) (resObj *Chat, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Chat
func (r *ChatRequest) Update(ctx context.Context, reqObj *Chat) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Chat
func (r *ChatRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChatInfoRequestBuilder is request builder for ChatInfo
type ChatInfoRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChatInfoRequest
func (b *ChatInfoRequestBuilder) Request() *ChatInfoRequest {
	return &ChatInfoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChatInfoRequest is request for ChatInfo
type ChatInfoRequest struct{ BaseRequest }

// Get performs GET request for ChatInfo
func (r *ChatInfoRequest) Get(ctx context.Context) (resObj *ChatInfo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChatInfo
func (r *ChatInfoRequest) Update(ctx context.Context, reqObj *ChatInfo) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChatInfo
func (r *ChatInfoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChatMembersNotificationRecipientRequestBuilder is request builder for ChatMembersNotificationRecipient
type ChatMembersNotificationRecipientRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChatMembersNotificationRecipientRequest
func (b *ChatMembersNotificationRecipientRequestBuilder) Request() *ChatMembersNotificationRecipientRequest {
	return &ChatMembersNotificationRecipientRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChatMembersNotificationRecipientRequest is request for ChatMembersNotificationRecipient
type ChatMembersNotificationRecipientRequest struct{ BaseRequest }

// Get performs GET request for ChatMembersNotificationRecipient
func (r *ChatMembersNotificationRecipientRequest) Get(ctx context.Context) (resObj *ChatMembersNotificationRecipient, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChatMembersNotificationRecipient
func (r *ChatMembersNotificationRecipientRequest) Update(ctx context.Context, reqObj *ChatMembersNotificationRecipient) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChatMembersNotificationRecipient
func (r *ChatMembersNotificationRecipientRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChatMessageRequestBuilder is request builder for ChatMessage
type ChatMessageRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChatMessageRequest
func (b *ChatMessageRequestBuilder) Request() *ChatMessageRequest {
	return &ChatMessageRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChatMessageRequest is request for ChatMessage
type ChatMessageRequest struct{ BaseRequest }

// Get performs GET request for ChatMessage
func (r *ChatMessageRequest) Get(ctx context.Context) (resObj *ChatMessage, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChatMessage
func (r *ChatMessageRequest) Update(ctx context.Context, reqObj *ChatMessage) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChatMessage
func (r *ChatMessageRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChatMessageAttachmentRequestBuilder is request builder for ChatMessageAttachment
type ChatMessageAttachmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChatMessageAttachmentRequest
func (b *ChatMessageAttachmentRequestBuilder) Request() *ChatMessageAttachmentRequest {
	return &ChatMessageAttachmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChatMessageAttachmentRequest is request for ChatMessageAttachment
type ChatMessageAttachmentRequest struct{ BaseRequest }

// Get performs GET request for ChatMessageAttachment
func (r *ChatMessageAttachmentRequest) Get(ctx context.Context) (resObj *ChatMessageAttachment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChatMessageAttachment
func (r *ChatMessageAttachmentRequest) Update(ctx context.Context, reqObj *ChatMessageAttachment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChatMessageAttachment
func (r *ChatMessageAttachmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChatMessageFromIdentitySetRequestBuilder is request builder for ChatMessageFromIdentitySet
type ChatMessageFromIdentitySetRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChatMessageFromIdentitySetRequest
func (b *ChatMessageFromIdentitySetRequestBuilder) Request() *ChatMessageFromIdentitySetRequest {
	return &ChatMessageFromIdentitySetRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChatMessageFromIdentitySetRequest is request for ChatMessageFromIdentitySet
type ChatMessageFromIdentitySetRequest struct{ BaseRequest }

// Get performs GET request for ChatMessageFromIdentitySet
func (r *ChatMessageFromIdentitySetRequest) Get(ctx context.Context) (resObj *ChatMessageFromIdentitySet, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChatMessageFromIdentitySet
func (r *ChatMessageFromIdentitySetRequest) Update(ctx context.Context, reqObj *ChatMessageFromIdentitySet) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChatMessageFromIdentitySet
func (r *ChatMessageFromIdentitySetRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChatMessageHostedContentRequestBuilder is request builder for ChatMessageHostedContent
type ChatMessageHostedContentRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChatMessageHostedContentRequest
func (b *ChatMessageHostedContentRequestBuilder) Request() *ChatMessageHostedContentRequest {
	return &ChatMessageHostedContentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChatMessageHostedContentRequest is request for ChatMessageHostedContent
type ChatMessageHostedContentRequest struct{ BaseRequest }

// Get performs GET request for ChatMessageHostedContent
func (r *ChatMessageHostedContentRequest) Get(ctx context.Context) (resObj *ChatMessageHostedContent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChatMessageHostedContent
func (r *ChatMessageHostedContentRequest) Update(ctx context.Context, reqObj *ChatMessageHostedContent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChatMessageHostedContent
func (r *ChatMessageHostedContentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChatMessageInfoRequestBuilder is request builder for ChatMessageInfo
type ChatMessageInfoRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChatMessageInfoRequest
func (b *ChatMessageInfoRequestBuilder) Request() *ChatMessageInfoRequest {
	return &ChatMessageInfoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChatMessageInfoRequest is request for ChatMessageInfo
type ChatMessageInfoRequest struct{ BaseRequest }

// Get performs GET request for ChatMessageInfo
func (r *ChatMessageInfoRequest) Get(ctx context.Context) (resObj *ChatMessageInfo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChatMessageInfo
func (r *ChatMessageInfoRequest) Update(ctx context.Context, reqObj *ChatMessageInfo) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChatMessageInfo
func (r *ChatMessageInfoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChatMessageMentionRequestBuilder is request builder for ChatMessageMention
type ChatMessageMentionRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChatMessageMentionRequest
func (b *ChatMessageMentionRequestBuilder) Request() *ChatMessageMentionRequest {
	return &ChatMessageMentionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChatMessageMentionRequest is request for ChatMessageMention
type ChatMessageMentionRequest struct{ BaseRequest }

// Get performs GET request for ChatMessageMention
func (r *ChatMessageMentionRequest) Get(ctx context.Context) (resObj *ChatMessageMention, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChatMessageMention
func (r *ChatMessageMentionRequest) Update(ctx context.Context, reqObj *ChatMessageMention) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChatMessageMention
func (r *ChatMessageMentionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChatMessageMentionedIdentitySetRequestBuilder is request builder for ChatMessageMentionedIdentitySet
type ChatMessageMentionedIdentitySetRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChatMessageMentionedIdentitySetRequest
func (b *ChatMessageMentionedIdentitySetRequestBuilder) Request() *ChatMessageMentionedIdentitySetRequest {
	return &ChatMessageMentionedIdentitySetRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChatMessageMentionedIdentitySetRequest is request for ChatMessageMentionedIdentitySet
type ChatMessageMentionedIdentitySetRequest struct{ BaseRequest }

// Get performs GET request for ChatMessageMentionedIdentitySet
func (r *ChatMessageMentionedIdentitySetRequest) Get(ctx context.Context) (resObj *ChatMessageMentionedIdentitySet, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChatMessageMentionedIdentitySet
func (r *ChatMessageMentionedIdentitySetRequest) Update(ctx context.Context, reqObj *ChatMessageMentionedIdentitySet) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChatMessageMentionedIdentitySet
func (r *ChatMessageMentionedIdentitySetRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChatMessagePolicyViolationRequestBuilder is request builder for ChatMessagePolicyViolation
type ChatMessagePolicyViolationRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChatMessagePolicyViolationRequest
func (b *ChatMessagePolicyViolationRequestBuilder) Request() *ChatMessagePolicyViolationRequest {
	return &ChatMessagePolicyViolationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChatMessagePolicyViolationRequest is request for ChatMessagePolicyViolation
type ChatMessagePolicyViolationRequest struct{ BaseRequest }

// Get performs GET request for ChatMessagePolicyViolation
func (r *ChatMessagePolicyViolationRequest) Get(ctx context.Context) (resObj *ChatMessagePolicyViolation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChatMessagePolicyViolation
func (r *ChatMessagePolicyViolationRequest) Update(ctx context.Context, reqObj *ChatMessagePolicyViolation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChatMessagePolicyViolation
func (r *ChatMessagePolicyViolationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChatMessagePolicyViolationPolicyTipRequestBuilder is request builder for ChatMessagePolicyViolationPolicyTip
type ChatMessagePolicyViolationPolicyTipRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChatMessagePolicyViolationPolicyTipRequest
func (b *ChatMessagePolicyViolationPolicyTipRequestBuilder) Request() *ChatMessagePolicyViolationPolicyTipRequest {
	return &ChatMessagePolicyViolationPolicyTipRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChatMessagePolicyViolationPolicyTipRequest is request for ChatMessagePolicyViolationPolicyTip
type ChatMessagePolicyViolationPolicyTipRequest struct{ BaseRequest }

// Get performs GET request for ChatMessagePolicyViolationPolicyTip
func (r *ChatMessagePolicyViolationPolicyTipRequest) Get(ctx context.Context) (resObj *ChatMessagePolicyViolationPolicyTip, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChatMessagePolicyViolationPolicyTip
func (r *ChatMessagePolicyViolationPolicyTipRequest) Update(ctx context.Context, reqObj *ChatMessagePolicyViolationPolicyTip) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChatMessagePolicyViolationPolicyTip
func (r *ChatMessagePolicyViolationPolicyTipRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChatMessageReactionRequestBuilder is request builder for ChatMessageReaction
type ChatMessageReactionRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChatMessageReactionRequest
func (b *ChatMessageReactionRequestBuilder) Request() *ChatMessageReactionRequest {
	return &ChatMessageReactionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChatMessageReactionRequest is request for ChatMessageReaction
type ChatMessageReactionRequest struct{ BaseRequest }

// Get performs GET request for ChatMessageReaction
func (r *ChatMessageReactionRequest) Get(ctx context.Context) (resObj *ChatMessageReaction, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChatMessageReaction
func (r *ChatMessageReactionRequest) Update(ctx context.Context, reqObj *ChatMessageReaction) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChatMessageReaction
func (r *ChatMessageReactionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChatMessageReactionIdentitySetRequestBuilder is request builder for ChatMessageReactionIdentitySet
type ChatMessageReactionIdentitySetRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChatMessageReactionIdentitySetRequest
func (b *ChatMessageReactionIdentitySetRequestBuilder) Request() *ChatMessageReactionIdentitySetRequest {
	return &ChatMessageReactionIdentitySetRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChatMessageReactionIdentitySetRequest is request for ChatMessageReactionIdentitySet
type ChatMessageReactionIdentitySetRequest struct{ BaseRequest }

// Get performs GET request for ChatMessageReactionIdentitySet
func (r *ChatMessageReactionIdentitySetRequest) Get(ctx context.Context) (resObj *ChatMessageReactionIdentitySet, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChatMessageReactionIdentitySet
func (r *ChatMessageReactionIdentitySetRequest) Update(ctx context.Context, reqObj *ChatMessageReactionIdentitySet) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChatMessageReactionIdentitySet
func (r *ChatMessageReactionIdentitySetRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChatRenamedEventMessageDetailRequestBuilder is request builder for ChatRenamedEventMessageDetail
type ChatRenamedEventMessageDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChatRenamedEventMessageDetailRequest
func (b *ChatRenamedEventMessageDetailRequestBuilder) Request() *ChatRenamedEventMessageDetailRequest {
	return &ChatRenamedEventMessageDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChatRenamedEventMessageDetailRequest is request for ChatRenamedEventMessageDetail
type ChatRenamedEventMessageDetailRequest struct{ BaseRequest }

// Get performs GET request for ChatRenamedEventMessageDetail
func (r *ChatRenamedEventMessageDetailRequest) Get(ctx context.Context) (resObj *ChatRenamedEventMessageDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChatRenamedEventMessageDetail
func (r *ChatRenamedEventMessageDetailRequest) Update(ctx context.Context, reqObj *ChatRenamedEventMessageDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChatRenamedEventMessageDetail
func (r *ChatRenamedEventMessageDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChatViewpointRequestBuilder is request builder for ChatViewpoint
type ChatViewpointRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChatViewpointRequest
func (b *ChatViewpointRequestBuilder) Request() *ChatViewpointRequest {
	return &ChatViewpointRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChatViewpointRequest is request for ChatViewpoint
type ChatViewpointRequest struct{ BaseRequest }

// Get performs GET request for ChatViewpoint
func (r *ChatViewpointRequest) Get(ctx context.Context) (resObj *ChatViewpoint, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChatViewpoint
func (r *ChatViewpointRequest) Update(ctx context.Context, reqObj *ChatViewpoint) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChatViewpoint
func (r *ChatViewpointRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type ChatSendActivityNotificationRequestBuilder struct{ BaseRequestBuilder }

// SendActivityNotification action undocumented
func (b *ChatRequestBuilder) SendActivityNotification(reqObj *ChatSendActivityNotificationRequestParameter) *ChatSendActivityNotificationRequestBuilder {
	bb := &ChatSendActivityNotificationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/SendActivityNotification"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ChatSendActivityNotificationRequest struct{ BaseRequest }

func (b *ChatSendActivityNotificationRequestBuilder) Request() *ChatSendActivityNotificationRequest {
	return &ChatSendActivityNotificationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ChatSendActivityNotificationRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type ChatHideForUserRequestBuilder struct{ BaseRequestBuilder }

// HideForUser action undocumented
func (b *ChatRequestBuilder) HideForUser(reqObj *ChatHideForUserRequestParameter) *ChatHideForUserRequestBuilder {
	bb := &ChatHideForUserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/HideForUser"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ChatHideForUserRequest struct{ BaseRequest }

func (b *ChatHideForUserRequestBuilder) Request() *ChatHideForUserRequest {
	return &ChatHideForUserRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ChatHideForUserRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type ChatMarkChatReadForUserRequestBuilder struct{ BaseRequestBuilder }

// MarkChatReadForUser action undocumented
func (b *ChatRequestBuilder) MarkChatReadForUser(reqObj *ChatMarkChatReadForUserRequestParameter) *ChatMarkChatReadForUserRequestBuilder {
	bb := &ChatMarkChatReadForUserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/MarkChatReadForUser"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ChatMarkChatReadForUserRequest struct{ BaseRequest }

func (b *ChatMarkChatReadForUserRequestBuilder) Request() *ChatMarkChatReadForUserRequest {
	return &ChatMarkChatReadForUserRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ChatMarkChatReadForUserRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type ChatMarkChatUnreadForUserRequestBuilder struct{ BaseRequestBuilder }

// MarkChatUnreadForUser action undocumented
func (b *ChatRequestBuilder) MarkChatUnreadForUser(reqObj *ChatMarkChatUnreadForUserRequestParameter) *ChatMarkChatUnreadForUserRequestBuilder {
	bb := &ChatMarkChatUnreadForUserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/MarkChatUnreadForUser"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ChatMarkChatUnreadForUserRequest struct{ BaseRequest }

func (b *ChatMarkChatUnreadForUserRequestBuilder) Request() *ChatMarkChatUnreadForUserRequest {
	return &ChatMarkChatUnreadForUserRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ChatMarkChatUnreadForUserRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type ChatUnhideForUserRequestBuilder struct{ BaseRequestBuilder }

// UnhideForUser action undocumented
func (b *ChatRequestBuilder) UnhideForUser(reqObj *ChatUnhideForUserRequestParameter) *ChatUnhideForUserRequestBuilder {
	bb := &ChatUnhideForUserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/UnhideForUser"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ChatUnhideForUserRequest struct{ BaseRequest }

func (b *ChatUnhideForUserRequestBuilder) Request() *ChatUnhideForUserRequest {
	return &ChatUnhideForUserRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ChatUnhideForUserRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type ChatMessageSoftDeleteRequestBuilder struct{ BaseRequestBuilder }

// SoftDelete action undocumented
func (b *ChatMessageRequestBuilder) SoftDelete(reqObj *ChatMessageSoftDeleteRequestParameter) *ChatMessageSoftDeleteRequestBuilder {
	bb := &ChatMessageSoftDeleteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/SoftDelete"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ChatMessageSoftDeleteRequest struct{ BaseRequest }

func (b *ChatMessageSoftDeleteRequestBuilder) Request() *ChatMessageSoftDeleteRequest {
	return &ChatMessageSoftDeleteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ChatMessageSoftDeleteRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type ChatMessageUndoSoftDeleteRequestBuilder struct{ BaseRequestBuilder }

// UndoSoftDelete action undocumented
func (b *ChatMessageRequestBuilder) UndoSoftDelete(reqObj *ChatMessageUndoSoftDeleteRequestParameter) *ChatMessageUndoSoftDeleteRequestBuilder {
	bb := &ChatMessageUndoSoftDeleteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/UndoSoftDelete"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ChatMessageUndoSoftDeleteRequest struct{ BaseRequest }

func (b *ChatMessageUndoSoftDeleteRequestBuilder) Request() *ChatMessageUndoSoftDeleteRequest {
	return &ChatMessageUndoSoftDeleteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ChatMessageUndoSoftDeleteRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
