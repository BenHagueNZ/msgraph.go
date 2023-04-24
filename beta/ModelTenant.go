// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// TenantAppManagementPolicy undocumented
type TenantAppManagementPolicy struct {
	// PolicyBase is the base model of TenantAppManagementPolicy
	PolicyBase

	ODataType string `json:"@odata.type,omitempty"`
	// ApplicationRestrictions undocumented
	ApplicationRestrictions *AppManagementConfiguration `json:"applicationRestrictions,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// ServicePrincipalRestrictions undocumented
	ServicePrincipalRestrictions *AppManagementConfiguration `json:"servicePrincipalRestrictions,omitempty"`
}

func NewTenantAppManagementPolicy() (*TenantAppManagementPolicy, error) {
	newTenantAppManagementPolicy := &TenantAppManagementPolicy{
		ODataType: "#microsoft.graph.TenantAppManagementPolicy",
	}
	return newTenantAppManagementPolicy, nil
}

// TenantAttachRBAC undocumented
type TenantAttachRBAC struct {
	// Entity is the base model of TenantAttachRBAC
	Entity

	ODataType string `json:"@odata.type,omitempty"`
}

func NewTenantAttachRBAC() (*TenantAttachRBAC, error) {
	newTenantAttachRBAC := &TenantAttachRBAC{
		ODataType: "#microsoft.graph.TenantAttachRBAC",
	}
	return newTenantAttachRBAC, nil
}

// TenantAttachRBACState undocumented
type TenantAttachRBACState struct {
	// Object is the base model of TenantAttachRBACState
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Enabled undocumented
	Enabled *bool `json:"enabled,omitempty"`
}

func NewTenantAttachRBACState() (*TenantAttachRBACState, error) {
	newTenantAttachRBACState := &TenantAttachRBACState{
		ODataType: "#microsoft.graph.TenantAttachRBACState",
	}
	return newTenantAttachRBACState, nil
}

// TenantInformation undocumented
type TenantInformation struct {
	// Object is the base model of TenantInformation
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// DefaultDomainName undocumented
	DefaultDomainName *string `json:"defaultDomainName,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// FederationBrandName undocumented
	FederationBrandName *string `json:"federationBrandName,omitempty"`
	// TenantID undocumented
	TenantID *string `json:"tenantId,omitempty"`
}

func NewTenantInformation() (*TenantInformation, error) {
	newTenantInformation := &TenantInformation{
		ODataType: "#microsoft.graph.TenantInformation",
	}
	return newTenantInformation, nil
}

// TenantReference undocumented
type TenantReference struct {
	// Object is the base model of TenantReference
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// TenantID undocumented
	TenantID *string `json:"tenantId,omitempty"`
}

func NewTenantReference() (*TenantReference, error) {
	newTenantReference := &TenantReference{
		ODataType: "#microsoft.graph.TenantReference",
	}
	return newTenantReference, nil
}

// TenantRelationship undocumented
type TenantRelationship struct {
	// Object is the base model of TenantRelationship
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ManagedTenants undocumented
	ManagedTenants *ManagedTenantsManagedTenant `json:"managedTenants,omitempty"`
	// DelegatedAdminCustomers undocumented
	DelegatedAdminCustomers []DelegatedAdminCustomer `json:"delegatedAdminCustomers,omitempty"`
	// DelegatedAdminRelationships undocumented
	DelegatedAdminRelationships []DelegatedAdminRelationship `json:"delegatedAdminRelationships,omitempty"`
}

func NewTenantRelationship() (*TenantRelationship, error) {
	newTenantRelationship := &TenantRelationship{
		ODataType: "#microsoft.graph.TenantRelationship",
	}
	return newTenantRelationship, nil
}

// TenantRelationshipAccessPolicyBase undocumented
type TenantRelationshipAccessPolicyBase struct {
	// PolicyBase is the base model of TenantRelationshipAccessPolicyBase
	PolicyBase

	ODataType string `json:"@odata.type,omitempty"`
	// Definition undocumented
	Definition []string `json:"definition,omitempty"`
}

func NewTenantRelationshipAccessPolicyBase() (*TenantRelationshipAccessPolicyBase, error) {
	newTenantRelationshipAccessPolicyBase := &TenantRelationshipAccessPolicyBase{
		ODataType: "#microsoft.graph.TenantRelationshipAccessPolicyBase",
	}
	return newTenantRelationshipAccessPolicyBase, nil
}

// TenantSetupInfo undocumented
type TenantSetupInfo struct {
	// Entity is the base model of TenantSetupInfo
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// FirstTimeSetup undocumented
	FirstTimeSetup *bool `json:"firstTimeSetup,omitempty"`
	// RelevantRolesSettings undocumented
	RelevantRolesSettings []string `json:"relevantRolesSettings,omitempty"`
	// SetupStatus undocumented
	SetupStatus *SetupStatus `json:"setupStatus,omitempty"`
	// SkipSetup undocumented
	SkipSetup *bool `json:"skipSetup,omitempty"`
	// UserRolesActions undocumented
	UserRolesActions *string `json:"userRolesActions,omitempty"`
	// DefaultRolesSettings undocumented
	DefaultRolesSettings *PrivilegedRoleSettings `json:"defaultRolesSettings,omitempty"`
}

func NewTenantSetupInfo() (*TenantSetupInfo, error) {
	newTenantSetupInfo := &TenantSetupInfo{
		ODataType: "#microsoft.graph.TenantSetupInfo",
	}
	return newTenantSetupInfo, nil
}