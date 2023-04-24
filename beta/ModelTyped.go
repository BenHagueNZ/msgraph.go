// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// TypedEmailAddress undocumented
type TypedEmailAddress struct {
	// EmailAddress is the base model of TypedEmailAddress
	EmailAddress

	ODataType string `json:"@odata.type,omitempty"`
	// OtherLabel undocumented
	OtherLabel *string `json:"otherLabel,omitempty"`
	// Type undocumented
	Type *EmailType `json:"type,omitempty"`
}

func NewTypedEmailAddress() (*TypedEmailAddress, error) {
	newTypedEmailAddress := &TypedEmailAddress{
		ODataType: "#microsoft.graph.TypedEmailAddress",
	}
	return newTypedEmailAddress, nil
}
