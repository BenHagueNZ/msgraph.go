// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Directory undocumented
type Directory struct {
	// Entity is the base model of Directory
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AdministrativeUnits undocumented
	AdministrativeUnits []AdministrativeUnit `json:"administrativeUnits,omitempty"`
	// DeletedItems undocumented
	DeletedItems []DirectoryObject `json:"deletedItems,omitempty"`
	// FederationConfigurations undocumented
	FederationConfigurations []IdentityProviderBase `json:"federationConfigurations,omitempty"`
	// OnPremisesSynchronization undocumented
	OnPremisesSynchronization []OnPremisesDirectorySynchronization `json:"onPremisesSynchronization,omitempty"`
}

func NewDirectory() (*Directory, error) {
	newDirectory := &Directory{
		ODataType: "#microsoft.graph.Directory",
	}
	return newDirectory, nil
}

// DirectoryAudit undocumented
type DirectoryAudit struct {
	// Entity is the base model of DirectoryAudit
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ActivityDateTime undocumented
	ActivityDateTime *time.Time `json:"activityDateTime,omitempty"`
	// ActivityDisplayName undocumented
	ActivityDisplayName *string `json:"activityDisplayName,omitempty"`
	// AdditionalDetails undocumented
	AdditionalDetails []KeyValue `json:"additionalDetails,omitempty"`
	// Category undocumented
	Category *string `json:"category,omitempty"`
	// CorrelationID undocumented
	CorrelationID *string `json:"correlationId,omitempty"`
	// InitiatedBy undocumented
	InitiatedBy *AuditActivityInitiator `json:"initiatedBy,omitempty"`
	// LoggedByService undocumented
	LoggedByService *string `json:"loggedByService,omitempty"`
	// OperationType undocumented
	OperationType *string `json:"operationType,omitempty"`
	// Result undocumented
	Result *OperationResult `json:"result,omitempty"`
	// ResultReason undocumented
	ResultReason *string `json:"resultReason,omitempty"`
	// TargetResources undocumented
	TargetResources []TargetResource `json:"targetResources,omitempty"`
}

func NewDirectoryAudit() (*DirectoryAudit, error) {
	newDirectoryAudit := &DirectoryAudit{
		ODataType: "#microsoft.graph.DirectoryAudit",
	}
	return newDirectoryAudit, nil
}

// DirectoryObject undocumented
type DirectoryObject struct {
	// Entity is the base model of DirectoryObject
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// DeletedDateTime undocumented
	DeletedDateTime *time.Time `json:"deletedDateTime,omitempty"`
}

func NewDirectoryObject() (*DirectoryObject, error) {
	newDirectoryObject := &DirectoryObject{
		ODataType: "#microsoft.graph.DirectoryObject",
	}
	return newDirectoryObject, nil
}

// DirectoryObjectPartnerReference undocumented
type DirectoryObjectPartnerReference struct {
	// DirectoryObject is the base model of DirectoryObjectPartnerReference
	DirectoryObject

	ODataType string `json:"@odata.type,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ExternalPartnerTenantID undocumented
	ExternalPartnerTenantID *UUID `json:"externalPartnerTenantId,omitempty"`
	// ObjectType undocumented
	ObjectType *string `json:"objectType,omitempty"`
}

func NewDirectoryObjectPartnerReference() (*DirectoryObjectPartnerReference, error) {
	newDirectoryObjectPartnerReference := &DirectoryObjectPartnerReference{
		ODataType: "#microsoft.graph.DirectoryObjectPartnerReference",
	}
	return newDirectoryObjectPartnerReference, nil
}

// DirectoryRole undocumented
type DirectoryRole struct {
	// DirectoryObject is the base model of DirectoryRole
	DirectoryObject

	ODataType string `json:"@odata.type,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// RoleTemplateID undocumented
	RoleTemplateID *string `json:"roleTemplateId,omitempty"`
	// Members undocumented
	Members []DirectoryObject `json:"members,omitempty"`
	// ScopedMembers undocumented
	ScopedMembers []ScopedRoleMembership `json:"scopedMembers,omitempty"`
}

func NewDirectoryRole() (*DirectoryRole, error) {
	newDirectoryRole := &DirectoryRole{
		ODataType: "#microsoft.graph.DirectoryRole",
	}
	return newDirectoryRole, nil
}

// DirectoryRoleTemplate undocumented
type DirectoryRoleTemplate struct {
	// DirectoryObject is the base model of DirectoryRoleTemplate
	DirectoryObject

	ODataType string `json:"@odata.type,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
}

func NewDirectoryRoleTemplate() (*DirectoryRoleTemplate, error) {
	newDirectoryRoleTemplate := &DirectoryRoleTemplate{
		ODataType: "#microsoft.graph.DirectoryRoleTemplate",
	}
	return newDirectoryRoleTemplate, nil
}
