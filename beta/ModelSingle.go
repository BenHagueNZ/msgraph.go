// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// SingleSignOnExtension undocumented
type SingleSignOnExtension struct {
	// Object is the base model of SingleSignOnExtension
	Object

	ODataType string `json:"@odata.type,omitempty"`
}

func NewSingleSignOnExtension() (*SingleSignOnExtension, error) {
	newSingleSignOnExtension := &SingleSignOnExtension{
		ODataType: "#microsoft.graph.SingleSignOnExtension",
	}
	return newSingleSignOnExtension, nil
}

// SingleUser undocumented
type SingleUser struct {
	// UserSet is the base model of SingleUser
	UserSet

	ODataType string `json:"@odata.type,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
}

func NewSingleUser() (*SingleUser, error) {
	newSingleUser := &SingleUser{
		ODataType: "#microsoft.graph.SingleUser",
	}
	return newSingleUser, nil
}

// SingleValueLegacyExtendedProperty undocumented
type SingleValueLegacyExtendedProperty struct {
	// Entity is the base model of SingleValueLegacyExtendedProperty
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Value undocumented
	Value *string `json:"value,omitempty"`
}

func NewSingleValueLegacyExtendedProperty() (*SingleValueLegacyExtendedProperty, error) {
	newSingleValueLegacyExtendedProperty := &SingleValueLegacyExtendedProperty{
		ODataType: "#microsoft.graph.SingleValueLegacyExtendedProperty",
	}
	return newSingleValueLegacyExtendedProperty, nil
}