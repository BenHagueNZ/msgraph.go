// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// RelatedContact undocumented
type RelatedContact struct {
	// Object is the base model of RelatedContact
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AccessConsent undocumented
	AccessConsent *bool `json:"accessConsent,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// EmailAddress undocumented
	EmailAddress *string `json:"emailAddress,omitempty"`
	// MobilePhone undocumented
	MobilePhone *string `json:"mobilePhone,omitempty"`
	// Relationship undocumented
	Relationship *ContactRelationship `json:"relationship,omitempty"`
}

func NewRelatedContact() (*RelatedContact, error) {
	newRelatedContact := &RelatedContact{
		ODataType: "#microsoft.graph.RelatedContact",
	}
	return newRelatedContact, nil
}
