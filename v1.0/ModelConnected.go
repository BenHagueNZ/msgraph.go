// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// ConnectedOrganization undocumented
type ConnectedOrganization struct {
	// Entity is the base model of ConnectedOrganization
	Entity
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IdentitySources undocumented
	IdentitySources []IdentitySource `json:"identitySources,omitempty"`
	// ModifiedDateTime undocumented
	ModifiedDateTime *time.Time `json:"modifiedDateTime,omitempty"`
	// State undocumented
	State *ConnectedOrganizationState `json:"state,omitempty"`
	// ExternalSponsors undocumented
	ExternalSponsors []DirectoryObject `json:"externalSponsors,omitempty"`
	// InternalSponsors undocumented
	InternalSponsors []DirectoryObject `json:"internalSponsors,omitempty"`
}

// ConnectedOrganizationMembers undocumented
type ConnectedOrganizationMembers struct {
	// SubjectSet is the base model of ConnectedOrganizationMembers
	SubjectSet
	// ConnectedOrganizationID undocumented
	ConnectedOrganizationID *string `json:"connectedOrganizationId,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
}
