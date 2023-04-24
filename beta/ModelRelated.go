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
	// ID undocumented
	ID *string `json:"id,omitempty"`
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

// RelatedPerson undocumented
type RelatedPerson struct {
	// Object is the base model of RelatedPerson
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Relationship undocumented
	Relationship *PersonRelationship `json:"relationship,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}

func NewRelatedPerson() (*RelatedPerson, error) {
	newRelatedPerson := &RelatedPerson{
		ODataType: "#microsoft.graph.RelatedPerson",
	}
	return newRelatedPerson, nil
}