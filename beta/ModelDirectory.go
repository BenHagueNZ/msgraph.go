// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Directory undocumented
type Directory struct {
	// Entity is the base model of Directory
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ImpactedResources undocumented
	ImpactedResources []ImpactedResource `json:"impactedResources,omitempty"`
	// Recommendations undocumented
	Recommendations []Recommendation `json:"recommendations,omitempty"`
	// AdministrativeUnits undocumented
	AdministrativeUnits []AdministrativeUnit `json:"administrativeUnits,omitempty"`
	// AttributeSets undocumented
	AttributeSets []AttributeSet `json:"attributeSets,omitempty"`
	// CustomSecurityAttributeDefinitions undocumented
	CustomSecurityAttributeDefinitions []CustomSecurityAttributeDefinition `json:"customSecurityAttributeDefinitions,omitempty"`
	// DeletedItems undocumented
	DeletedItems []DirectoryObject `json:"deletedItems,omitempty"`
	// FederationConfigurations undocumented
	FederationConfigurations []IdentityProviderBase `json:"federationConfigurations,omitempty"`
	// InboundSharedUserProfiles undocumented
	InboundSharedUserProfiles []InboundSharedUserProfile `json:"inboundSharedUserProfiles,omitempty"`
	// OnPremisesSynchronization undocumented
	OnPremisesSynchronization []OnPremisesDirectorySynchronization `json:"onPremisesSynchronization,omitempty"`
	// OutboundSharedUserProfiles undocumented
	OutboundSharedUserProfiles []OutboundSharedUserProfile `json:"outboundSharedUserProfiles,omitempty"`
	// SharedEmailDomains undocumented
	SharedEmailDomains []SharedEmailDomain `json:"sharedEmailDomains,omitempty"`
	// FeatureRolloutPolicies undocumented
	FeatureRolloutPolicies []FeatureRolloutPolicy `json:"featureRolloutPolicies,omitempty"`
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
	// UserAgent undocumented
	UserAgent *string `json:"userAgent,omitempty"`
}

func NewDirectoryAudit() (*DirectoryAudit, error) {
	newDirectoryAudit := &DirectoryAudit{
		ODataType: "#microsoft.graph.DirectoryAudit",
	}
	return newDirectoryAudit, nil
}

// DirectoryDefinition undocumented
type DirectoryDefinition struct {
	// Entity is the base model of DirectoryDefinition
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Discoverabilities undocumented
	Discoverabilities *DirectoryDefinitionDiscoverabilities `json:"discoverabilities,omitempty"`
	// DiscoveryDateTime undocumented
	DiscoveryDateTime *time.Time `json:"discoveryDateTime,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Objects undocumented
	Objects []ObjectDefinition `json:"objects,omitempty"`
	// ReadOnly undocumented
	ReadOnly *bool `json:"readOnly,omitempty"`
	// Version undocumented
	Version *string `json:"version,omitempty"`
}

func NewDirectoryDefinition() (*DirectoryDefinition, error) {
	newDirectoryDefinition := &DirectoryDefinition{
		ODataType: "#microsoft.graph.DirectoryDefinition",
	}
	return newDirectoryDefinition, nil
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

// DirectoryRoleAccessReviewPolicy undocumented
type DirectoryRoleAccessReviewPolicy struct {
	// Entity is the base model of DirectoryRoleAccessReviewPolicy
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Settings undocumented
	Settings *AccessReviewScheduleSettings `json:"settings,omitempty"`
}

func NewDirectoryRoleAccessReviewPolicy() (*DirectoryRoleAccessReviewPolicy, error) {
	newDirectoryRoleAccessReviewPolicy := &DirectoryRoleAccessReviewPolicy{
		ODataType: "#microsoft.graph.DirectoryRoleAccessReviewPolicy",
	}
	return newDirectoryRoleAccessReviewPolicy, nil
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

// DirectorySetting undocumented
type DirectorySetting struct {
	// Entity is the base model of DirectorySetting
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// TemplateID undocumented
	TemplateID *string `json:"templateId,omitempty"`
	// Values undocumented
	Values []SettingValue `json:"values,omitempty"`
}

func NewDirectorySetting() (*DirectorySetting, error) {
	newDirectorySetting := &DirectorySetting{
		ODataType: "#microsoft.graph.DirectorySetting",
	}
	return newDirectorySetting, nil
}

// DirectorySettingTemplate undocumented
type DirectorySettingTemplate struct {
	// DirectoryObject is the base model of DirectorySettingTemplate
	DirectoryObject

	ODataType string `json:"@odata.type,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Values undocumented
	Values []SettingTemplateValue `json:"values,omitempty"`
}

func NewDirectorySettingTemplate() (*DirectorySettingTemplate, error) {
	newDirectorySettingTemplate := &DirectorySettingTemplate{
		ODataType: "#microsoft.graph.DirectorySettingTemplate",
	}
	return newDirectorySettingTemplate, nil
}

// DirectorySizeQuota undocumented
type DirectorySizeQuota struct {
	// Object is the base model of DirectorySizeQuota
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Total undocumented
	Total *int `json:"total,omitempty"`
	// Used undocumented
	Used *int `json:"used,omitempty"`
}

func NewDirectorySizeQuota() (*DirectorySizeQuota, error) {
	newDirectorySizeQuota := &DirectorySizeQuota{
		ODataType: "#microsoft.graph.DirectorySizeQuota",
	}
	return newDirectorySizeQuota, nil
}