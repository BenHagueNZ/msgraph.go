// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// SkypeForBusinessUserConversationMember undocumented
type SkypeForBusinessUserConversationMember struct {
	// ConversationMember is the base model of SkypeForBusinessUserConversationMember
	ConversationMember
	// TenantID undocumented
	TenantID *string `json:"tenantId,omitempty"`
	// UserID undocumented
	UserID *string `json:"userId,omitempty"`
}

// SkypeUserConversationMember undocumented
type SkypeUserConversationMember struct {
	// ConversationMember is the base model of SkypeUserConversationMember
	ConversationMember
	// SkypeID undocumented
	SkypeID *string `json:"skypeId,omitempty"`
}
