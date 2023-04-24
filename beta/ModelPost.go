// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Post undocumented
type Post struct {
	// OutlookItem is the base model of Post
	OutlookItem

	ODataType string `json:"@odata.type,omitempty"`
	// Body undocumented
	Body *ItemBody `json:"body,omitempty"`
	// ConversationID undocumented
	ConversationID *string `json:"conversationId,omitempty"`
	// ConversationThreadID undocumented
	ConversationThreadID *string `json:"conversationThreadId,omitempty"`
	// From undocumented
	From *Recipient `json:"from,omitempty"`
	// HasAttachments undocumented
	HasAttachments *bool `json:"hasAttachments,omitempty"`
	// Importance undocumented
	Importance *Importance `json:"importance,omitempty"`
	// NewParticipants undocumented
	NewParticipants []Recipient `json:"newParticipants,omitempty"`
	// ReceivedDateTime undocumented
	ReceivedDateTime *time.Time `json:"receivedDateTime,omitempty"`
	// Sender undocumented
	Sender *Recipient `json:"sender,omitempty"`
	// Attachments undocumented
	Attachments []Attachment `json:"attachments,omitempty"`
	// Extensions undocumented
	Extensions []Extension `json:"extensions,omitempty"`
	// InReplyTo undocumented
	InReplyTo *Post `json:"inReplyTo,omitempty"`
	// Mentions undocumented
	Mentions []Mention `json:"mentions,omitempty"`
	// MultiValueExtendedProperties undocumented
	MultiValueExtendedProperties []MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`
	// SingleValueExtendedProperties undocumented
	SingleValueExtendedProperties []SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
}

func NewPost() (*Post, error) {
	newPost := &Post{
		ODataType: "#microsoft.graph.Post",
	}
	return newPost, nil
}
