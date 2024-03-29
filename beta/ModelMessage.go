// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Message undocumented
type Message struct {
	// OutlookItem is the base model of Message
	OutlookItem

	ODataType string `json:"@odata.type,omitempty"`
	// BccRecipients undocumented
	BccRecipients []Recipient `json:"bccRecipients,omitempty"`
	// Body undocumented
	Body *ItemBody `json:"body,omitempty"`
	// BodyPreview undocumented
	BodyPreview *string `json:"bodyPreview,omitempty"`
	// CcRecipients undocumented
	CcRecipients []Recipient `json:"ccRecipients,omitempty"`
	// ConversationID undocumented
	ConversationID *string `json:"conversationId,omitempty"`
	// ConversationIndex undocumented
	ConversationIndex *Binary `json:"conversationIndex,omitempty"`
	// Flag undocumented
	Flag *FollowupFlag `json:"flag,omitempty"`
	// From undocumented
	From *Recipient `json:"from,omitempty"`
	// HasAttachments undocumented
	HasAttachments *bool `json:"hasAttachments,omitempty"`
	// Importance undocumented
	Importance *Importance `json:"importance,omitempty"`
	// InferenceClassification undocumented
	InferenceClassification *InferenceClassificationType `json:"inferenceClassification,omitempty"`
	// InternetMessageHeaders undocumented
	InternetMessageHeaders []InternetMessageHeader `json:"internetMessageHeaders,omitempty"`
	// InternetMessageID undocumented
	InternetMessageID *string `json:"internetMessageId,omitempty"`
	// IsDeliveryReceiptRequested undocumented
	IsDeliveryReceiptRequested *bool `json:"isDeliveryReceiptRequested,omitempty"`
	// IsDraft undocumented
	IsDraft *bool `json:"isDraft,omitempty"`
	// IsRead undocumented
	IsRead *bool `json:"isRead,omitempty"`
	// IsReadReceiptRequested undocumented
	IsReadReceiptRequested *bool `json:"isReadReceiptRequested,omitempty"`
	// MentionsPreview undocumented
	MentionsPreview *MentionsPreview `json:"mentionsPreview,omitempty"`
	// ParentFolderID undocumented
	ParentFolderID *string `json:"parentFolderId,omitempty"`
	// ReceivedDateTime undocumented
	ReceivedDateTime *time.Time `json:"receivedDateTime,omitempty"`
	// ReplyTo undocumented
	ReplyTo []Recipient `json:"replyTo,omitempty"`
	// Sender undocumented
	Sender *Recipient `json:"sender,omitempty"`
	// SentDateTime undocumented
	SentDateTime *time.Time `json:"sentDateTime,omitempty"`
	// Subject undocumented
	Subject *string `json:"subject,omitempty"`
	// ToRecipients undocumented
	ToRecipients []Recipient `json:"toRecipients,omitempty"`
	// UniqueBody undocumented
	UniqueBody *ItemBody `json:"uniqueBody,omitempty"`
	// UnsubscribeData undocumented
	UnsubscribeData []string `json:"unsubscribeData,omitempty"`
	// UnsubscribeEnabled undocumented
	UnsubscribeEnabled *bool `json:"unsubscribeEnabled,omitempty"`
	// WebLink undocumented
	WebLink *string `json:"webLink,omitempty"`
	// Attachments undocumented
	Attachments []Attachment `json:"attachments,omitempty"`
	// Extensions undocumented
	Extensions []Extension `json:"extensions,omitempty"`
	// Mentions undocumented
	Mentions []Mention `json:"mentions,omitempty"`
	// MultiValueExtendedProperties undocumented
	MultiValueExtendedProperties []MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`
	// SingleValueExtendedProperties undocumented
	SingleValueExtendedProperties []SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
}

func NewMessage() (*Message, error) {
	newMessage := &Message{
		ODataType: "#microsoft.graph.Message",
	}
	return newMessage, nil
}

// MessageEvent undocumented
type MessageEvent struct {
	// Entity is the base model of MessageEvent
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// DateTime undocumented
	DateTime *time.Time `json:"dateTime,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// EventType undocumented
	EventType *MessageEventType `json:"eventType,omitempty"`
}

func NewMessageEvent() (*MessageEvent, error) {
	newMessageEvent := &MessageEvent{
		ODataType: "#microsoft.graph.MessageEvent",
	}
	return newMessageEvent, nil
}

// MessagePinnedEventMessageDetail undocumented
type MessagePinnedEventMessageDetail struct {
	// EventMessageDetail is the base model of MessagePinnedEventMessageDetail
	EventMessageDetail

	ODataType string `json:"@odata.type,omitempty"`
	// EventDateTime undocumented
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`
	// Initiator undocumented
	Initiator *IdentitySet `json:"initiator,omitempty"`
}

func NewMessagePinnedEventMessageDetail() (*MessagePinnedEventMessageDetail, error) {
	newMessagePinnedEventMessageDetail := &MessagePinnedEventMessageDetail{
		ODataType: "#microsoft.graph.MessagePinnedEventMessageDetail",
	}
	return newMessagePinnedEventMessageDetail, nil
}

// MessageRecipient undocumented
type MessageRecipient struct {
	// Entity is the base model of MessageRecipient
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// DeliveryStatus undocumented
	DeliveryStatus *MessageStatus `json:"deliveryStatus,omitempty"`
	// RecipientEmail undocumented
	RecipientEmail *string `json:"recipientEmail,omitempty"`
	// Events undocumented
	Events []MessageEvent `json:"events,omitempty"`
}

func NewMessageRecipient() (*MessageRecipient, error) {
	newMessageRecipient := &MessageRecipient{
		ODataType: "#microsoft.graph.MessageRecipient",
	}
	return newMessageRecipient, nil
}

// MessageRule undocumented
type MessageRule struct {
	// Entity is the base model of MessageRule
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Actions undocumented
	Actions *MessageRuleActions `json:"actions,omitempty"`
	// Conditions undocumented
	Conditions *MessageRulePredicates `json:"conditions,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Exceptions undocumented
	Exceptions *MessageRulePredicates `json:"exceptions,omitempty"`
	// HasError undocumented
	HasError *bool `json:"hasError,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// IsReadOnly undocumented
	IsReadOnly *bool `json:"isReadOnly,omitempty"`
	// Sequence undocumented
	Sequence *int `json:"sequence,omitempty"`
}

func NewMessageRule() (*MessageRule, error) {
	newMessageRule := &MessageRule{
		ODataType: "#microsoft.graph.MessageRule",
	}
	return newMessageRule, nil
}

// MessageRuleActions undocumented
type MessageRuleActions struct {
	// Object is the base model of MessageRuleActions
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AssignCategories undocumented
	AssignCategories []string `json:"assignCategories,omitempty"`
	// CopyToFolder undocumented
	CopyToFolder *string `json:"copyToFolder,omitempty"`
	// Delete undocumented
	Delete *bool `json:"delete,omitempty"`
	// ForwardAsAttachmentTo undocumented
	ForwardAsAttachmentTo []Recipient `json:"forwardAsAttachmentTo,omitempty"`
	// ForwardTo undocumented
	ForwardTo []Recipient `json:"forwardTo,omitempty"`
	// MarkAsRead undocumented
	MarkAsRead *bool `json:"markAsRead,omitempty"`
	// MarkImportance undocumented
	MarkImportance *Importance `json:"markImportance,omitempty"`
	// MoveToFolder undocumented
	MoveToFolder *string `json:"moveToFolder,omitempty"`
	// PermanentDelete undocumented
	PermanentDelete *bool `json:"permanentDelete,omitempty"`
	// RedirectTo undocumented
	RedirectTo []Recipient `json:"redirectTo,omitempty"`
	// StopProcessingRules undocumented
	StopProcessingRules *bool `json:"stopProcessingRules,omitempty"`
}

func NewMessageRuleActions() (*MessageRuleActions, error) {
	newMessageRuleActions := &MessageRuleActions{
		ODataType: "#microsoft.graph.MessageRuleActions",
	}
	return newMessageRuleActions, nil
}

// MessageRulePredicates undocumented
type MessageRulePredicates struct {
	// Object is the base model of MessageRulePredicates
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// BodyContains undocumented
	BodyContains []string `json:"bodyContains,omitempty"`
	// BodyOrSubjectContains undocumented
	BodyOrSubjectContains []string `json:"bodyOrSubjectContains,omitempty"`
	// Categories undocumented
	Categories []string `json:"categories,omitempty"`
	// FromAddresses undocumented
	FromAddresses []Recipient `json:"fromAddresses,omitempty"`
	// HasAttachments undocumented
	HasAttachments *bool `json:"hasAttachments,omitempty"`
	// HeaderContains undocumented
	HeaderContains []string `json:"headerContains,omitempty"`
	// Importance undocumented
	Importance *Importance `json:"importance,omitempty"`
	// IsApprovalRequest undocumented
	IsApprovalRequest *bool `json:"isApprovalRequest,omitempty"`
	// IsAutomaticForward undocumented
	IsAutomaticForward *bool `json:"isAutomaticForward,omitempty"`
	// IsAutomaticReply undocumented
	IsAutomaticReply *bool `json:"isAutomaticReply,omitempty"`
	// IsEncrypted undocumented
	IsEncrypted *bool `json:"isEncrypted,omitempty"`
	// IsMeetingRequest undocumented
	IsMeetingRequest *bool `json:"isMeetingRequest,omitempty"`
	// IsMeetingResponse undocumented
	IsMeetingResponse *bool `json:"isMeetingResponse,omitempty"`
	// IsNonDeliveryReport undocumented
	IsNonDeliveryReport *bool `json:"isNonDeliveryReport,omitempty"`
	// IsPermissionControlled undocumented
	IsPermissionControlled *bool `json:"isPermissionControlled,omitempty"`
	// IsReadReceipt undocumented
	IsReadReceipt *bool `json:"isReadReceipt,omitempty"`
	// IsSigned undocumented
	IsSigned *bool `json:"isSigned,omitempty"`
	// IsVoicemail undocumented
	IsVoicemail *bool `json:"isVoicemail,omitempty"`
	// MessageActionFlag undocumented
	MessageActionFlag *MessageActionFlag `json:"messageActionFlag,omitempty"`
	// NotSentToMe undocumented
	NotSentToMe *bool `json:"notSentToMe,omitempty"`
	// RecipientContains undocumented
	RecipientContains []string `json:"recipientContains,omitempty"`
	// SenderContains undocumented
	SenderContains []string `json:"senderContains,omitempty"`
	// Sensitivity undocumented
	Sensitivity *Sensitivity `json:"sensitivity,omitempty"`
	// SentCcMe undocumented
	SentCcMe *bool `json:"sentCcMe,omitempty"`
	// SentOnlyToMe undocumented
	SentOnlyToMe *bool `json:"sentOnlyToMe,omitempty"`
	// SentToAddresses undocumented
	SentToAddresses []Recipient `json:"sentToAddresses,omitempty"`
	// SentToMe undocumented
	SentToMe *bool `json:"sentToMe,omitempty"`
	// SentToOrCcMe undocumented
	SentToOrCcMe *bool `json:"sentToOrCcMe,omitempty"`
	// SubjectContains undocumented
	SubjectContains []string `json:"subjectContains,omitempty"`
	// WithinSizeRange undocumented
	WithinSizeRange *SizeRange `json:"withinSizeRange,omitempty"`
}

func NewMessageRulePredicates() (*MessageRulePredicates, error) {
	newMessageRulePredicates := &MessageRulePredicates{
		ODataType: "#microsoft.graph.MessageRulePredicates",
	}
	return newMessageRulePredicates, nil
}

// MessageSecurityState undocumented
type MessageSecurityState struct {
	// Object is the base model of MessageSecurityState
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ConnectingIP undocumented
	ConnectingIP *string `json:"connectingIP,omitempty"`
	// DeliveryAction undocumented
	DeliveryAction *string `json:"deliveryAction,omitempty"`
	// DeliveryLocation undocumented
	DeliveryLocation *string `json:"deliveryLocation,omitempty"`
	// Directionality undocumented
	Directionality *string `json:"directionality,omitempty"`
	// InternetMessageID undocumented
	InternetMessageID *string `json:"internetMessageId,omitempty"`
	// MessageFingerprint undocumented
	MessageFingerprint *string `json:"messageFingerprint,omitempty"`
	// MessageReceivedDateTime undocumented
	MessageReceivedDateTime *time.Time `json:"messageReceivedDateTime,omitempty"`
	// MessageSubject undocumented
	MessageSubject *string `json:"messageSubject,omitempty"`
	// NetworkMessageID undocumented
	NetworkMessageID *string `json:"networkMessageId,omitempty"`
}

func NewMessageSecurityState() (*MessageSecurityState, error) {
	newMessageSecurityState := &MessageSecurityState{
		ODataType: "#microsoft.graph.MessageSecurityState",
	}
	return newMessageSecurityState, nil
}

// MessageTrace undocumented
type MessageTrace struct {
	// Entity is the base model of MessageTrace
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// DestinationIPAddress undocumented
	DestinationIPAddress *string `json:"destinationIPAddress,omitempty"`
	// MessageID undocumented
	MessageID *string `json:"messageId,omitempty"`
	// ReceivedDateTime undocumented
	ReceivedDateTime *time.Time `json:"receivedDateTime,omitempty"`
	// SenderEmail undocumented
	SenderEmail *string `json:"senderEmail,omitempty"`
	// Size undocumented
	Size *int `json:"size,omitempty"`
	// SourceIPAddress undocumented
	SourceIPAddress *string `json:"sourceIPAddress,omitempty"`
	// Subject undocumented
	Subject *string `json:"subject,omitempty"`
	// Recipients undocumented
	Recipients []MessageRecipient `json:"recipients,omitempty"`
}

func NewMessageTrace() (*MessageTrace, error) {
	newMessageTrace := &MessageTrace{
		ODataType: "#microsoft.graph.MessageTrace",
	}
	return newMessageTrace, nil
}

// MessageUnpinnedEventMessageDetail undocumented
type MessageUnpinnedEventMessageDetail struct {
	// EventMessageDetail is the base model of MessageUnpinnedEventMessageDetail
	EventMessageDetail

	ODataType string `json:"@odata.type,omitempty"`
	// EventDateTime undocumented
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`
	// Initiator undocumented
	Initiator *IdentitySet `json:"initiator,omitempty"`
}

func NewMessageUnpinnedEventMessageDetail() (*MessageUnpinnedEventMessageDetail, error) {
	newMessageUnpinnedEventMessageDetail := &MessageUnpinnedEventMessageDetail{
		ODataType: "#microsoft.graph.MessageUnpinnedEventMessageDetail",
	}
	return newMessageUnpinnedEventMessageDetail, nil
}
