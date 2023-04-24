// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Conversation undocumented
type Conversation struct {
	// Entity is the base model of Conversation
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// HasAttachments undocumented
	HasAttachments *bool `json:"hasAttachments,omitempty"`
	// LastDeliveredDateTime undocumented
	LastDeliveredDateTime *time.Time `json:"lastDeliveredDateTime,omitempty"`
	// Preview undocumented
	Preview *string `json:"preview,omitempty"`
	// Topic undocumented
	Topic *string `json:"topic,omitempty"`
	// UniqueSenders undocumented
	UniqueSenders []string `json:"uniqueSenders,omitempty"`
	// Threads undocumented
	Threads []ConversationThread `json:"threads,omitempty"`
}

func NewConversation() (*Conversation, error) {
	newConversation := &Conversation{
		ODataType: "#microsoft.graph.Conversation",
	}
	return newConversation, nil
}

// ConversationMember undocumented
type ConversationMember struct {
	// Entity is the base model of ConversationMember
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Roles undocumented
	Roles []string `json:"roles,omitempty"`
	// VisibleHistoryStartDateTime undocumented
	VisibleHistoryStartDateTime *time.Time `json:"visibleHistoryStartDateTime,omitempty"`
}

func NewConversationMember() (*ConversationMember, error) {
	newConversationMember := &ConversationMember{
		ODataType: "#microsoft.graph.ConversationMember",
	}
	return newConversationMember, nil
}

// ConversationMemberRoleUpdatedEventMessageDetail undocumented
type ConversationMemberRoleUpdatedEventMessageDetail struct {
	// EventMessageDetail is the base model of ConversationMemberRoleUpdatedEventMessageDetail
	EventMessageDetail

	ODataType string `json:"@odata.type,omitempty"`
	// ConversationMemberRoles undocumented
	ConversationMemberRoles []string `json:"conversationMemberRoles,omitempty"`
	// ConversationMemberUser undocumented
	ConversationMemberUser *TeamworkUserIdentity `json:"conversationMemberUser,omitempty"`
	// Initiator undocumented
	Initiator *IdentitySet `json:"initiator,omitempty"`
}

func NewConversationMemberRoleUpdatedEventMessageDetail() (*ConversationMemberRoleUpdatedEventMessageDetail, error) {
	newConversationMemberRoleUpdatedEventMessageDetail := &ConversationMemberRoleUpdatedEventMessageDetail{
		ODataType: "#microsoft.graph.ConversationMemberRoleUpdatedEventMessageDetail",
	}
	return newConversationMemberRoleUpdatedEventMessageDetail, nil
}

// ConversationThread undocumented
type ConversationThread struct {
	// Entity is the base model of ConversationThread
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// CcRecipients undocumented
	CcRecipients []Recipient `json:"ccRecipients,omitempty"`
	// HasAttachments undocumented
	HasAttachments *bool `json:"hasAttachments,omitempty"`
	// IsLocked undocumented
	IsLocked *bool `json:"isLocked,omitempty"`
	// LastDeliveredDateTime undocumented
	LastDeliveredDateTime *time.Time `json:"lastDeliveredDateTime,omitempty"`
	// Preview undocumented
	Preview *string `json:"preview,omitempty"`
	// Topic undocumented
	Topic *string `json:"topic,omitempty"`
	// ToRecipients undocumented
	ToRecipients []Recipient `json:"toRecipients,omitempty"`
	// UniqueSenders undocumented
	UniqueSenders []string `json:"uniqueSenders,omitempty"`
	// Posts undocumented
	Posts []Post `json:"posts,omitempty"`
}

func NewConversationThread() (*ConversationThread, error) {
	newConversationThread := &ConversationThread{
		ODataType: "#microsoft.graph.ConversationThread",
	}
	return newConversationThread, nil
}
