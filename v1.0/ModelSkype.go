// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// SkypeForBusinessUserConversationMember undocumented
type SkypeForBusinessUserConversationMember struct {
	// ConversationMember is the base model of SkypeForBusinessUserConversationMember
	ConversationMember

	ODataType string `json:"@odata.type"`
	// TenantID undocumented
	TenantID *string `json:"tenantId,omitempty"`
	// UserID undocumented
	UserID *string `json:"userId,omitempty"`
}

func NewSkypeForBusinessUserConversationMember() (*SkypeForBusinessUserConversationMember, error) {
	newSkypeForBusinessUserConversationMember := &SkypeForBusinessUserConversationMember{
		ODataType: "#microsoft.graph.SkypeForBusinessUserConversationMember",
	}
	return newSkypeForBusinessUserConversationMember, nil
}

// SkypeUserConversationMember undocumented
type SkypeUserConversationMember struct {
	// ConversationMember is the base model of SkypeUserConversationMember
	ConversationMember

	ODataType string `json:"@odata.type"`
	// SkypeID undocumented
	SkypeID *string `json:"skypeId,omitempty"`
}

func NewSkypeUserConversationMember() (*SkypeUserConversationMember, error) {
	newSkypeUserConversationMember := &SkypeUserConversationMember{
		ODataType: "#microsoft.graph.SkypeUserConversationMember",
	}
	return newSkypeUserConversationMember, nil
}
