// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Chat undocumented
type Chat struct {
	// Entity is the base model of Chat
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ChatType undocumented
	ChatType *ChatType `json:"chatType,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// LastUpdatedDateTime undocumented
	LastUpdatedDateTime *time.Time `json:"lastUpdatedDateTime,omitempty"`
	// OnlineMeetingInfo undocumented
	OnlineMeetingInfo *TeamworkOnlineMeetingInfo `json:"onlineMeetingInfo,omitempty"`
	// TenantID undocumented
	TenantID *string `json:"tenantId,omitempty"`
	// Topic undocumented
	Topic *string `json:"topic,omitempty"`
	// Viewpoint undocumented
	Viewpoint *ChatViewpoint `json:"viewpoint,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
	// InstalledApps undocumented
	InstalledApps []TeamsAppInstallation `json:"installedApps,omitempty"`
	// LastMessagePreview undocumented
	LastMessagePreview *ChatMessageInfo `json:"lastMessagePreview,omitempty"`
	// Members undocumented
	Members []ConversationMember `json:"members,omitempty"`
	// Messages undocumented
	Messages []ChatMessage `json:"messages,omitempty"`
	// PinnedMessages undocumented
	PinnedMessages []PinnedChatMessageInfo `json:"pinnedMessages,omitempty"`
	// Tabs undocumented
	Tabs []TeamsTab `json:"tabs,omitempty"`
}

func NewChat() (*Chat, error) {
	newChat := &Chat{
		ODataType: "#microsoft.graph.Chat",
	}
	return newChat, nil
}

// ChatInfo undocumented
type ChatInfo struct {
	// Object is the base model of ChatInfo
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// MessageID undocumented
	MessageID *string `json:"messageId,omitempty"`
	// ReplyChainMessageID undocumented
	ReplyChainMessageID *string `json:"replyChainMessageId,omitempty"`
	// ThreadID undocumented
	ThreadID *string `json:"threadId,omitempty"`
}

func NewChatInfo() (*ChatInfo, error) {
	newChatInfo := &ChatInfo{
		ODataType: "#microsoft.graph.ChatInfo",
	}
	return newChatInfo, nil
}

// ChatMembersNotificationRecipient undocumented
type ChatMembersNotificationRecipient struct {
	// TeamworkNotificationRecipient is the base model of ChatMembersNotificationRecipient
	TeamworkNotificationRecipient

	ODataType string `json:"@odata.type,omitempty"`
	// ChatID undocumented
	ChatID *string `json:"chatId,omitempty"`
}

func NewChatMembersNotificationRecipient() (*ChatMembersNotificationRecipient, error) {
	newChatMembersNotificationRecipient := &ChatMembersNotificationRecipient{
		ODataType: "#microsoft.graph.ChatMembersNotificationRecipient",
	}
	return newChatMembersNotificationRecipient, nil
}

// ChatMessage undocumented
type ChatMessage struct {
	// Entity is the base model of ChatMessage
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Attachments undocumented
	Attachments []ChatMessageAttachment `json:"attachments,omitempty"`
	// Body undocumented
	Body *ItemBody `json:"body,omitempty"`
	// ChannelIdentity undocumented
	ChannelIdentity *ChannelIdentity `json:"channelIdentity,omitempty"`
	// ChatID undocumented
	ChatID *string `json:"chatId,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DeletedDateTime undocumented
	DeletedDateTime *time.Time `json:"deletedDateTime,omitempty"`
	// Etag undocumented
	Etag *string `json:"etag,omitempty"`
	// EventDetail undocumented
	EventDetail *EventMessageDetail `json:"eventDetail,omitempty"`
	// From undocumented
	From *ChatMessageFromIdentitySet `json:"from,omitempty"`
	// Importance undocumented
	Importance *ChatMessageImportance `json:"importance,omitempty"`
	// LastEditedDateTime undocumented
	LastEditedDateTime *time.Time `json:"lastEditedDateTime,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Locale undocumented
	Locale *string `json:"locale,omitempty"`
	// Mentions undocumented
	Mentions []ChatMessageMention `json:"mentions,omitempty"`
	// MessageType undocumented
	MessageType *ChatMessageType `json:"messageType,omitempty"`
	// PolicyViolation undocumented
	PolicyViolation *ChatMessagePolicyViolation `json:"policyViolation,omitempty"`
	// Reactions undocumented
	Reactions []ChatMessageReaction `json:"reactions,omitempty"`
	// ReplyToID undocumented
	ReplyToID *string `json:"replyToId,omitempty"`
	// Subject undocumented
	Subject *string `json:"subject,omitempty"`
	// Summary undocumented
	Summary *string `json:"summary,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
	// HostedContents undocumented
	HostedContents []ChatMessageHostedContent `json:"hostedContents,omitempty"`
	// Replies undocumented
	Replies []ChatMessage `json:"replies,omitempty"`
}

func NewChatMessage() (*ChatMessage, error) {
	newChatMessage := &ChatMessage{
		ODataType: "#microsoft.graph.ChatMessage",
	}
	return newChatMessage, nil
}

// ChatMessageAttachment undocumented
type ChatMessageAttachment struct {
	// Object is the base model of ChatMessageAttachment
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Content undocumented
	Content *string `json:"content,omitempty"`
	// ContentType undocumented
	ContentType *string `json:"contentType,omitempty"`
	// ContentURL undocumented
	ContentURL *string `json:"contentUrl,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// TeamsAppID undocumented
	TeamsAppID *string `json:"teamsAppId,omitempty"`
	// ThumbnailURL undocumented
	ThumbnailURL *string `json:"thumbnailUrl,omitempty"`
}

func NewChatMessageAttachment() (*ChatMessageAttachment, error) {
	newChatMessageAttachment := &ChatMessageAttachment{
		ODataType: "#microsoft.graph.ChatMessageAttachment",
	}
	return newChatMessageAttachment, nil
}

// ChatMessageFromIdentitySet undocumented
type ChatMessageFromIdentitySet struct {
	// IdentitySet is the base model of ChatMessageFromIdentitySet
	IdentitySet

	ODataType string `json:"@odata.type,omitempty"`
}

func NewChatMessageFromIdentitySet() (*ChatMessageFromIdentitySet, error) {
	newChatMessageFromIdentitySet := &ChatMessageFromIdentitySet{
		ODataType: "#microsoft.graph.ChatMessageFromIdentitySet",
	}
	return newChatMessageFromIdentitySet, nil
}

// ChatMessageHostedContent undocumented
type ChatMessageHostedContent struct {
	// TeamworkHostedContent is the base model of ChatMessageHostedContent
	TeamworkHostedContent

	ODataType string `json:"@odata.type,omitempty"`
}

func NewChatMessageHostedContent() (*ChatMessageHostedContent, error) {
	newChatMessageHostedContent := &ChatMessageHostedContent{
		ODataType: "#microsoft.graph.ChatMessageHostedContent",
	}
	return newChatMessageHostedContent, nil
}

// ChatMessageInfo undocumented
type ChatMessageInfo struct {
	// Entity is the base model of ChatMessageInfo
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Body undocumented
	Body *ItemBody `json:"body,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// EventDetail undocumented
	EventDetail *EventMessageDetail `json:"eventDetail,omitempty"`
	// From undocumented
	From *ChatMessageFromIdentitySet `json:"from,omitempty"`
	// IsDeleted undocumented
	IsDeleted *bool `json:"isDeleted,omitempty"`
	// MessageType undocumented
	MessageType *ChatMessageType `json:"messageType,omitempty"`
}

func NewChatMessageInfo() (*ChatMessageInfo, error) {
	newChatMessageInfo := &ChatMessageInfo{
		ODataType: "#microsoft.graph.ChatMessageInfo",
	}
	return newChatMessageInfo, nil
}

// ChatMessageMention undocumented
type ChatMessageMention struct {
	// Object is the base model of ChatMessageMention
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ID undocumented
	ID *int `json:"id,omitempty"`
	// Mentioned undocumented
	Mentioned *ChatMessageMentionedIdentitySet `json:"mentioned,omitempty"`
	// MentionText undocumented
	MentionText *string `json:"mentionText,omitempty"`
}

func NewChatMessageMention() (*ChatMessageMention, error) {
	newChatMessageMention := &ChatMessageMention{
		ODataType: "#microsoft.graph.ChatMessageMention",
	}
	return newChatMessageMention, nil
}

// ChatMessageMentionedIdentitySet undocumented
type ChatMessageMentionedIdentitySet struct {
	// IdentitySet is the base model of ChatMessageMentionedIdentitySet
	IdentitySet

	ODataType string `json:"@odata.type,omitempty"`
	// Conversation undocumented
	Conversation *TeamworkConversationIdentity `json:"conversation,omitempty"`
}

func NewChatMessageMentionedIdentitySet() (*ChatMessageMentionedIdentitySet, error) {
	newChatMessageMentionedIdentitySet := &ChatMessageMentionedIdentitySet{
		ODataType: "#microsoft.graph.ChatMessageMentionedIdentitySet",
	}
	return newChatMessageMentionedIdentitySet, nil
}

// ChatMessagePolicyViolation undocumented
type ChatMessagePolicyViolation struct {
	// Object is the base model of ChatMessagePolicyViolation
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// DlpAction undocumented
	DlpAction *ChatMessagePolicyViolationDlpActionTypes `json:"dlpAction,omitempty"`
	// JustificationText undocumented
	JustificationText *string `json:"justificationText,omitempty"`
	// PolicyTip undocumented
	PolicyTip *ChatMessagePolicyViolationPolicyTip `json:"policyTip,omitempty"`
	// UserAction undocumented
	UserAction *ChatMessagePolicyViolationUserActionTypes `json:"userAction,omitempty"`
	// VerdictDetails undocumented
	VerdictDetails *ChatMessagePolicyViolationVerdictDetailsTypes `json:"verdictDetails,omitempty"`
}

func NewChatMessagePolicyViolation() (*ChatMessagePolicyViolation, error) {
	newChatMessagePolicyViolation := &ChatMessagePolicyViolation{
		ODataType: "#microsoft.graph.ChatMessagePolicyViolation",
	}
	return newChatMessagePolicyViolation, nil
}

// ChatMessagePolicyViolationPolicyTip undocumented
type ChatMessagePolicyViolationPolicyTip struct {
	// Object is the base model of ChatMessagePolicyViolationPolicyTip
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ComplianceURL undocumented
	ComplianceURL *string `json:"complianceUrl,omitempty"`
	// GeneralText undocumented
	GeneralText *string `json:"generalText,omitempty"`
	// MatchedConditionDescriptions undocumented
	MatchedConditionDescriptions []string `json:"matchedConditionDescriptions,omitempty"`
}

func NewChatMessagePolicyViolationPolicyTip() (*ChatMessagePolicyViolationPolicyTip, error) {
	newChatMessagePolicyViolationPolicyTip := &ChatMessagePolicyViolationPolicyTip{
		ODataType: "#microsoft.graph.ChatMessagePolicyViolationPolicyTip",
	}
	return newChatMessagePolicyViolationPolicyTip, nil
}

// ChatMessageReaction undocumented
type ChatMessageReaction struct {
	// Object is the base model of ChatMessageReaction
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// ReactionType undocumented
	ReactionType *string `json:"reactionType,omitempty"`
	// User undocumented
	User *ChatMessageReactionIdentitySet `json:"user,omitempty"`
}

func NewChatMessageReaction() (*ChatMessageReaction, error) {
	newChatMessageReaction := &ChatMessageReaction{
		ODataType: "#microsoft.graph.ChatMessageReaction",
	}
	return newChatMessageReaction, nil
}

// ChatMessageReactionIdentitySet undocumented
type ChatMessageReactionIdentitySet struct {
	// IdentitySet is the base model of ChatMessageReactionIdentitySet
	IdentitySet

	ODataType string `json:"@odata.type,omitempty"`
}

func NewChatMessageReactionIdentitySet() (*ChatMessageReactionIdentitySet, error) {
	newChatMessageReactionIdentitySet := &ChatMessageReactionIdentitySet{
		ODataType: "#microsoft.graph.ChatMessageReactionIdentitySet",
	}
	return newChatMessageReactionIdentitySet, nil
}

// ChatRenamedEventMessageDetail undocumented
type ChatRenamedEventMessageDetail struct {
	// EventMessageDetail is the base model of ChatRenamedEventMessageDetail
	EventMessageDetail

	ODataType string `json:"@odata.type,omitempty"`
	// ChatDisplayName undocumented
	ChatDisplayName *string `json:"chatDisplayName,omitempty"`
	// ChatID undocumented
	ChatID *string `json:"chatId,omitempty"`
	// Initiator undocumented
	Initiator *IdentitySet `json:"initiator,omitempty"`
}

func NewChatRenamedEventMessageDetail() (*ChatRenamedEventMessageDetail, error) {
	newChatRenamedEventMessageDetail := &ChatRenamedEventMessageDetail{
		ODataType: "#microsoft.graph.ChatRenamedEventMessageDetail",
	}
	return newChatRenamedEventMessageDetail, nil
}

// ChatViewpoint undocumented
type ChatViewpoint struct {
	// Object is the base model of ChatViewpoint
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// IsHidden undocumented
	IsHidden *bool `json:"isHidden,omitempty"`
	// LastMessageReadDateTime undocumented
	LastMessageReadDateTime *time.Time `json:"lastMessageReadDateTime,omitempty"`
}

func NewChatViewpoint() (*ChatViewpoint, error) {
	newChatViewpoint := &ChatViewpoint{
		ODataType: "#microsoft.graph.ChatViewpoint",
	}
	return newChatViewpoint, nil
}
