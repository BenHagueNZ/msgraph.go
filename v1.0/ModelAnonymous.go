// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// AnonymousGuestConversationMember undocumented
type AnonymousGuestConversationMember struct {
	// ConversationMember is the base model of AnonymousGuestConversationMember
	ConversationMember

	ODataType string `json:"@odata.type,omitempty"`
	// AnonymousGuestID undocumented
	AnonymousGuestID *string `json:"anonymousGuestId,omitempty"`
}

func NewAnonymousGuestConversationMember() (*AnonymousGuestConversationMember, error) {
	newAnonymousGuestConversationMember := &AnonymousGuestConversationMember{
		ODataType: "#microsoft.graph.AnonymousGuestConversationMember",
	}
	return newAnonymousGuestConversationMember, nil
}
