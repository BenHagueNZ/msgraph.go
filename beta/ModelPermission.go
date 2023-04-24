// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Permission undocumented
type Permission struct {
	// Entity is the base model of Permission
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ExpirationDateTime undocumented
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// GrantedTo undocumented
	GrantedTo *IdentitySet `json:"grantedTo,omitempty"`
	// GrantedToIdentities undocumented
	GrantedToIdentities []IdentitySet `json:"grantedToIdentities,omitempty"`
	// GrantedToIdentitiesV2 undocumented
	GrantedToIdentitiesV2 []SharePointIdentitySet `json:"grantedToIdentitiesV2,omitempty"`
	// GrantedToV2 undocumented
	GrantedToV2 *SharePointIdentitySet `json:"grantedToV2,omitempty"`
	// HasPassword undocumented
	HasPassword *bool `json:"hasPassword,omitempty"`
	// InheritedFrom undocumented
	InheritedFrom *ItemReference `json:"inheritedFrom,omitempty"`
	// Invitation undocumented
	Invitation *SharingInvitation `json:"invitation,omitempty"`
	// Link undocumented
	Link *SharingLink `json:"link,omitempty"`
	// Roles undocumented
	Roles []string `json:"roles,omitempty"`
	// ShareID undocumented
	ShareID *string `json:"shareId,omitempty"`
}

func NewPermission() (*Permission, error) {
	newPermission := &Permission{
		ODataType: "#microsoft.graph.Permission",
	}
	return newPermission, nil
}

// PermissionGrantConditionSet undocumented
type PermissionGrantConditionSet struct {
	// Entity is the base model of PermissionGrantConditionSet
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// CertifiedClientApplicationsOnly undocumented
	CertifiedClientApplicationsOnly *bool `json:"certifiedClientApplicationsOnly,omitempty"`
	// ClientApplicationIDs undocumented
	ClientApplicationIDs []string `json:"clientApplicationIds,omitempty"`
	// ClientApplicationPublisherIDs undocumented
	ClientApplicationPublisherIDs []string `json:"clientApplicationPublisherIds,omitempty"`
	// ClientApplicationsFromVerifiedPublisherOnly undocumented
	ClientApplicationsFromVerifiedPublisherOnly *bool `json:"clientApplicationsFromVerifiedPublisherOnly,omitempty"`
	// ClientApplicationTenantIDs undocumented
	ClientApplicationTenantIDs []string `json:"clientApplicationTenantIds,omitempty"`
	// PermissionClassification undocumented
	PermissionClassification *string `json:"permissionClassification,omitempty"`
	// Permissions undocumented
	Permissions []string `json:"permissions,omitempty"`
	// PermissionType undocumented
	PermissionType *PermissionType `json:"permissionType,omitempty"`
	// ResourceApplication undocumented
	ResourceApplication *string `json:"resourceApplication,omitempty"`
}

func NewPermissionGrantConditionSet() (*PermissionGrantConditionSet, error) {
	newPermissionGrantConditionSet := &PermissionGrantConditionSet{
		ODataType: "#microsoft.graph.PermissionGrantConditionSet",
	}
	return newPermissionGrantConditionSet, nil
}

// PermissionGrantPolicy undocumented
type PermissionGrantPolicy struct {
	// PolicyBase is the base model of PermissionGrantPolicy
	PolicyBase

	ODataType string `json:"@odata.type,omitempty"`
	// Excludes undocumented
	Excludes []PermissionGrantConditionSet `json:"excludes,omitempty"`
	// Includes undocumented
	Includes []PermissionGrantConditionSet `json:"includes,omitempty"`
}

func NewPermissionGrantPolicy() (*PermissionGrantPolicy, error) {
	newPermissionGrantPolicy := &PermissionGrantPolicy{
		ODataType: "#microsoft.graph.PermissionGrantPolicy",
	}
	return newPermissionGrantPolicy, nil
}

// PermissionScope undocumented
type PermissionScope struct {
	// Object is the base model of PermissionScope
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AdminConsentDescription undocumented
	AdminConsentDescription *string `json:"adminConsentDescription,omitempty"`
	// AdminConsentDisplayName undocumented
	AdminConsentDisplayName *string `json:"adminConsentDisplayName,omitempty"`
	// ID undocumented
	ID *UUID `json:"id,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Origin undocumented
	Origin *string `json:"origin,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
	// UserConsentDescription undocumented
	UserConsentDescription *string `json:"userConsentDescription,omitempty"`
	// UserConsentDisplayName undocumented
	UserConsentDisplayName *string `json:"userConsentDisplayName,omitempty"`
	// Value undocumented
	Value *string `json:"value,omitempty"`
}

func NewPermissionScope() (*PermissionScope, error) {
	newPermissionScope := &PermissionScope{
		ODataType: "#microsoft.graph.PermissionScope",
	}
	return newPermissionScope, nil
}