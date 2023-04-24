// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// InvitedUserMessageInfo undocumented
type InvitedUserMessageInfo struct {
	// Object is the base model of InvitedUserMessageInfo
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// CcRecipients undocumented
	CcRecipients []Recipient `json:"ccRecipients,omitempty"`
	// CustomizedMessageBody undocumented
	CustomizedMessageBody *string `json:"customizedMessageBody,omitempty"`
	// MessageLanguage undocumented
	MessageLanguage *string `json:"messageLanguage,omitempty"`
}

func NewInvitedUserMessageInfo() (*InvitedUserMessageInfo, error) {
	newInvitedUserMessageInfo := &InvitedUserMessageInfo{
		ODataType: "#microsoft.graph.InvitedUserMessageInfo",
	}
	return newInvitedUserMessageInfo, nil
}