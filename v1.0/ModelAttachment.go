// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Attachment undocumented
type Attachment struct {
	// Entity is the base model of Attachment
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ContentType undocumented
	ContentType *string `json:"contentType,omitempty"`
	// IsInline undocumented
	IsInline *bool `json:"isInline,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Size undocumented
	Size *int `json:"size,omitempty"`
}

func NewAttachment() (*Attachment, error) {
	newAttachment := &Attachment{
		ODataType: "#microsoft.graph.Attachment",
	}
	return newAttachment, nil
}

// AttachmentBase undocumented
type AttachmentBase struct {
	// Entity is the base model of AttachmentBase
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ContentType undocumented
	ContentType *string `json:"contentType,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Size undocumented
	Size *int `json:"size,omitempty"`
}

func NewAttachmentBase() (*AttachmentBase, error) {
	newAttachmentBase := &AttachmentBase{
		ODataType: "#microsoft.graph.AttachmentBase",
	}
	return newAttachmentBase, nil
}

// AttachmentInfo undocumented
type AttachmentInfo struct {
	// Object is the base model of AttachmentInfo
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AttachmentType undocumented
	AttachmentType *AttachmentType `json:"attachmentType,omitempty"`
	// ContentType undocumented
	ContentType *string `json:"contentType,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Size undocumented
	Size *int `json:"size,omitempty"`
}

func NewAttachmentInfo() (*AttachmentInfo, error) {
	newAttachmentInfo := &AttachmentInfo{
		ODataType: "#microsoft.graph.AttachmentInfo",
	}
	return newAttachmentInfo, nil
}

// AttachmentItem undocumented
type AttachmentItem struct {
	// Object is the base model of AttachmentItem
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AttachmentType undocumented
	AttachmentType *AttachmentType `json:"attachmentType,omitempty"`
	// ContentID undocumented
	ContentID *string `json:"contentId,omitempty"`
	// ContentType undocumented
	ContentType *string `json:"contentType,omitempty"`
	// IsInline undocumented
	IsInline *bool `json:"isInline,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Size undocumented
	Size *int `json:"size,omitempty"`
}

func NewAttachmentItem() (*AttachmentItem, error) {
	newAttachmentItem := &AttachmentItem{
		ODataType: "#microsoft.graph.AttachmentItem",
	}
	return newAttachmentItem, nil
}

// AttachmentSession undocumented
type AttachmentSession struct {
	// Entity is the base model of AttachmentSession
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Content undocumented
	Content *Stream `json:"content,omitempty"`
	// ExpirationDateTime undocumented
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// NextExpectedRanges undocumented
	NextExpectedRanges []string `json:"nextExpectedRanges,omitempty"`
}

func NewAttachmentSession() (*AttachmentSession, error) {
	newAttachmentSession := &AttachmentSession{
		ODataType: "#microsoft.graph.AttachmentSession",
	}
	return newAttachmentSession, nil
}
