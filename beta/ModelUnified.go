// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// UnifiedRbacApplication undocumented
type UnifiedRbacApplication struct {
	// Entity is the base model of UnifiedRbacApplication
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ResourceNamespaces undocumented
	ResourceNamespaces []UnifiedRbacResourceNamespace `json:"resourceNamespaces,omitempty"`
	// RoleAssignments undocumented
	RoleAssignments []UnifiedRoleAssignment `json:"roleAssignments,omitempty"`
	// RoleDefinitions undocumented
	RoleDefinitions []UnifiedRoleDefinition `json:"roleDefinitions,omitempty"`
	// TransitiveRoleAssignments undocumented
	TransitiveRoleAssignments []UnifiedRoleAssignment `json:"transitiveRoleAssignments,omitempty"`
}

func NewUnifiedRbacApplication() (*UnifiedRbacApplication, error) {
	newUnifiedRbacApplication := &UnifiedRbacApplication{
		ODataType: "#microsoft.graph.UnifiedRbacApplication",
	}
	return newUnifiedRbacApplication, nil
}

// UnifiedRbacResourceAction undocumented
type UnifiedRbacResourceAction struct {
	// Entity is the base model of UnifiedRbacResourceAction
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ActionVerb undocumented
	ActionVerb *string `json:"actionVerb,omitempty"`
	// AuthenticationContextID undocumented
	AuthenticationContextID *string `json:"authenticationContextId,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// IsAuthenticationContextSettable undocumented
	IsAuthenticationContextSettable *bool `json:"isAuthenticationContextSettable,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// ResourceScopeID undocumented
	ResourceScopeID *string `json:"resourceScopeId,omitempty"`
	// ResourceScope undocumented
	ResourceScope *UnifiedRbacResourceScope `json:"resourceScope,omitempty"`
}

func NewUnifiedRbacResourceAction() (*UnifiedRbacResourceAction, error) {
	newUnifiedRbacResourceAction := &UnifiedRbacResourceAction{
		ODataType: "#microsoft.graph.UnifiedRbacResourceAction",
	}
	return newUnifiedRbacResourceAction, nil
}

// UnifiedRbacResourceNamespace undocumented
type UnifiedRbacResourceNamespace struct {
	// Entity is the base model of UnifiedRbacResourceNamespace
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// ResourceActions undocumented
	ResourceActions []UnifiedRbacResourceAction `json:"resourceActions,omitempty"`
}

func NewUnifiedRbacResourceNamespace() (*UnifiedRbacResourceNamespace, error) {
	newUnifiedRbacResourceNamespace := &UnifiedRbacResourceNamespace{
		ODataType: "#microsoft.graph.UnifiedRbacResourceNamespace",
	}
	return newUnifiedRbacResourceNamespace, nil
}

// UnifiedRbacResourceScope undocumented
type UnifiedRbacResourceScope struct {
	// Entity is the base model of UnifiedRbacResourceScope
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Scope undocumented
	Scope *string `json:"scope,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
}

func NewUnifiedRbacResourceScope() (*UnifiedRbacResourceScope, error) {
	newUnifiedRbacResourceScope := &UnifiedRbacResourceScope{
		ODataType: "#microsoft.graph.UnifiedRbacResourceScope",
	}
	return newUnifiedRbacResourceScope, nil
}

// UnifiedRole undocumented
type UnifiedRole struct {
	// Object is the base model of UnifiedRole
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// RoleDefinitionID undocumented
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty"`
}

func NewUnifiedRole() (*UnifiedRole, error) {
	newUnifiedRole := &UnifiedRole{
		ODataType: "#microsoft.graph.UnifiedRole",
	}
	return newUnifiedRole, nil
}

// UnifiedRoleAssignment undocumented
type UnifiedRoleAssignment struct {
	// Entity is the base model of UnifiedRoleAssignment
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AppScopeID undocumented
	AppScopeID *string `json:"appScopeId,omitempty"`
	// Condition undocumented
	Condition *string `json:"condition,omitempty"`
	// DirectoryScopeID undocumented
	DirectoryScopeID *string `json:"directoryScopeId,omitempty"`
	// PrincipalID undocumented
	PrincipalID *string `json:"principalId,omitempty"`
	// PrincipalOrganizationID undocumented
	PrincipalOrganizationID *string `json:"principalOrganizationId,omitempty"`
	// ResourceScope undocumented
	ResourceScope *string `json:"resourceScope,omitempty"`
	// RoleDefinitionID undocumented
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty"`
	// AppScope undocumented
	AppScope *AppScope `json:"appScope,omitempty"`
	// DirectoryScope undocumented
	DirectoryScope *DirectoryObject `json:"directoryScope,omitempty"`
	// Principal undocumented
	Principal *DirectoryObject `json:"principal,omitempty"`
	// RoleDefinition undocumented
	RoleDefinition *UnifiedRoleDefinition `json:"roleDefinition,omitempty"`
}

func NewUnifiedRoleAssignment() (*UnifiedRoleAssignment, error) {
	newUnifiedRoleAssignment := &UnifiedRoleAssignment{
		ODataType: "#microsoft.graph.UnifiedRoleAssignment",
	}
	return newUnifiedRoleAssignment, nil
}

// UnifiedRoleAssignmentMultiple undocumented
type UnifiedRoleAssignmentMultiple struct {
	// Entity is the base model of UnifiedRoleAssignmentMultiple
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AppScopeIDs undocumented
	AppScopeIDs []string `json:"appScopeIds,omitempty"`
	// Condition undocumented
	Condition *string `json:"condition,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DirectoryScopeIDs undocumented
	DirectoryScopeIDs []string `json:"directoryScopeIds,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// PrincipalIDs undocumented
	PrincipalIDs []string `json:"principalIds,omitempty"`
	// RoleDefinitionID undocumented
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty"`
	// AppScopes undocumented
	AppScopes []AppScope `json:"appScopes,omitempty"`
	// DirectoryScopes undocumented
	DirectoryScopes []DirectoryObject `json:"directoryScopes,omitempty"`
	// Principals undocumented
	Principals []DirectoryObject `json:"principals,omitempty"`
	// RoleDefinition undocumented
	RoleDefinition *UnifiedRoleDefinition `json:"roleDefinition,omitempty"`
}

func NewUnifiedRoleAssignmentMultiple() (*UnifiedRoleAssignmentMultiple, error) {
	newUnifiedRoleAssignmentMultiple := &UnifiedRoleAssignmentMultiple{
		ODataType: "#microsoft.graph.UnifiedRoleAssignmentMultiple",
	}
	return newUnifiedRoleAssignmentMultiple, nil
}

// UnifiedRoleAssignmentSchedule undocumented
type UnifiedRoleAssignmentSchedule struct {
	// UnifiedRoleScheduleBase is the base model of UnifiedRoleAssignmentSchedule
	UnifiedRoleScheduleBase

	ODataType string `json:"@odata.type,omitempty"`
	// AssignmentType undocumented
	AssignmentType *string `json:"assignmentType,omitempty"`
	// MemberType undocumented
	MemberType *string `json:"memberType,omitempty"`
	// ScheduleInfo undocumented
	ScheduleInfo *RequestSchedule `json:"scheduleInfo,omitempty"`
	// ActivatedUsing undocumented
	ActivatedUsing *UnifiedRoleEligibilitySchedule `json:"activatedUsing,omitempty"`
}

func NewUnifiedRoleAssignmentSchedule() (*UnifiedRoleAssignmentSchedule, error) {
	newUnifiedRoleAssignmentSchedule := &UnifiedRoleAssignmentSchedule{
		ODataType: "#microsoft.graph.UnifiedRoleAssignmentSchedule",
	}
	return newUnifiedRoleAssignmentSchedule, nil
}

// UnifiedRoleAssignmentScheduleInstance undocumented
type UnifiedRoleAssignmentScheduleInstance struct {
	// UnifiedRoleScheduleInstanceBase is the base model of UnifiedRoleAssignmentScheduleInstance
	UnifiedRoleScheduleInstanceBase

	ODataType string `json:"@odata.type,omitempty"`
	// AssignmentType undocumented
	AssignmentType *string `json:"assignmentType,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// MemberType undocumented
	MemberType *string `json:"memberType,omitempty"`
	// RoleAssignmentOriginID undocumented
	RoleAssignmentOriginID *string `json:"roleAssignmentOriginId,omitempty"`
	// RoleAssignmentScheduleID undocumented
	RoleAssignmentScheduleID *string `json:"roleAssignmentScheduleId,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// ActivatedUsing undocumented
	ActivatedUsing *UnifiedRoleEligibilityScheduleInstance `json:"activatedUsing,omitempty"`
}

func NewUnifiedRoleAssignmentScheduleInstance() (*UnifiedRoleAssignmentScheduleInstance, error) {
	newUnifiedRoleAssignmentScheduleInstance := &UnifiedRoleAssignmentScheduleInstance{
		ODataType: "#microsoft.graph.UnifiedRoleAssignmentScheduleInstance",
	}
	return newUnifiedRoleAssignmentScheduleInstance, nil
}

// UnifiedRoleAssignmentScheduleRequestObject undocumented
type UnifiedRoleAssignmentScheduleRequestObject struct {
	// RequestObject is the base model of UnifiedRoleAssignmentScheduleRequestObject
	RequestObject

	ODataType string `json:"@odata.type,omitempty"`
	// Action undocumented
	Action *string `json:"action,omitempty"`
	// AppScopeID undocumented
	AppScopeID *string `json:"appScopeId,omitempty"`
	// DirectoryScopeID undocumented
	DirectoryScopeID *string `json:"directoryScopeId,omitempty"`
	// IsValidationOnly undocumented
	IsValidationOnly *bool `json:"isValidationOnly,omitempty"`
	// Justification undocumented
	Justification *string `json:"justification,omitempty"`
	// PrincipalID undocumented
	PrincipalID *string `json:"principalId,omitempty"`
	// RoleDefinitionID undocumented
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty"`
	// ScheduleInfo undocumented
	ScheduleInfo *RequestSchedule `json:"scheduleInfo,omitempty"`
	// TargetScheduleID undocumented
	TargetScheduleID *string `json:"targetScheduleId,omitempty"`
	// TicketInfo undocumented
	TicketInfo *TicketInfo `json:"ticketInfo,omitempty"`
	// ActivatedUsing undocumented
	ActivatedUsing *UnifiedRoleEligibilitySchedule `json:"activatedUsing,omitempty"`
	// AppScope undocumented
	AppScope *AppScope `json:"appScope,omitempty"`
	// DirectoryScope undocumented
	DirectoryScope *DirectoryObject `json:"directoryScope,omitempty"`
	// Principal undocumented
	Principal *DirectoryObject `json:"principal,omitempty"`
	// RoleDefinition undocumented
	RoleDefinition *UnifiedRoleDefinition `json:"roleDefinition,omitempty"`
	// TargetSchedule undocumented
	TargetSchedule *UnifiedRoleAssignmentSchedule `json:"targetSchedule,omitempty"`
}

func NewUnifiedRoleAssignmentScheduleRequestObject() (*UnifiedRoleAssignmentScheduleRequestObject, error) {
	newUnifiedRoleAssignmentScheduleRequestObject := &UnifiedRoleAssignmentScheduleRequestObject{
		ODataType: "#microsoft.graph.UnifiedRoleAssignmentScheduleRequestObject",
	}
	return newUnifiedRoleAssignmentScheduleRequestObject, nil
}

// UnifiedRoleDefinition undocumented
type UnifiedRoleDefinition struct {
	// Entity is the base model of UnifiedRoleDefinition
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AllowedPrincipalTypes undocumented
	AllowedPrincipalTypes *AllowedRolePrincipalTypes `json:"allowedPrincipalTypes,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IsBuiltIn undocumented
	IsBuiltIn *bool `json:"isBuiltIn,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// ResourceScopes undocumented
	ResourceScopes []string `json:"resourceScopes,omitempty"`
	// RolePermissions undocumented
	RolePermissions []UnifiedRolePermission `json:"rolePermissions,omitempty"`
	// TemplateID undocumented
	TemplateID *string `json:"templateId,omitempty"`
	// Version undocumented
	Version *string `json:"version,omitempty"`
	// InheritsPermissionsFrom undocumented
	InheritsPermissionsFrom []UnifiedRoleDefinition `json:"inheritsPermissionsFrom,omitempty"`
}

func NewUnifiedRoleDefinition() (*UnifiedRoleDefinition, error) {
	newUnifiedRoleDefinition := &UnifiedRoleDefinition{
		ODataType: "#microsoft.graph.UnifiedRoleDefinition",
	}
	return newUnifiedRoleDefinition, nil
}

// UnifiedRoleEligibilitySchedule undocumented
type UnifiedRoleEligibilitySchedule struct {
	// UnifiedRoleScheduleBase is the base model of UnifiedRoleEligibilitySchedule
	UnifiedRoleScheduleBase

	ODataType string `json:"@odata.type,omitempty"`
	// MemberType undocumented
	MemberType *string `json:"memberType,omitempty"`
	// ScheduleInfo undocumented
	ScheduleInfo *RequestSchedule `json:"scheduleInfo,omitempty"`
}

func NewUnifiedRoleEligibilitySchedule() (*UnifiedRoleEligibilitySchedule, error) {
	newUnifiedRoleEligibilitySchedule := &UnifiedRoleEligibilitySchedule{
		ODataType: "#microsoft.graph.UnifiedRoleEligibilitySchedule",
	}
	return newUnifiedRoleEligibilitySchedule, nil
}

// UnifiedRoleEligibilityScheduleInstance undocumented
type UnifiedRoleEligibilityScheduleInstance struct {
	// UnifiedRoleScheduleInstanceBase is the base model of UnifiedRoleEligibilityScheduleInstance
	UnifiedRoleScheduleInstanceBase

	ODataType string `json:"@odata.type,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// MemberType undocumented
	MemberType *string `json:"memberType,omitempty"`
	// RoleEligibilityScheduleID undocumented
	RoleEligibilityScheduleID *string `json:"roleEligibilityScheduleId,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
}

func NewUnifiedRoleEligibilityScheduleInstance() (*UnifiedRoleEligibilityScheduleInstance, error) {
	newUnifiedRoleEligibilityScheduleInstance := &UnifiedRoleEligibilityScheduleInstance{
		ODataType: "#microsoft.graph.UnifiedRoleEligibilityScheduleInstance",
	}
	return newUnifiedRoleEligibilityScheduleInstance, nil
}

// UnifiedRoleEligibilityScheduleRequestObject undocumented
type UnifiedRoleEligibilityScheduleRequestObject struct {
	// RequestObject is the base model of UnifiedRoleEligibilityScheduleRequestObject
	RequestObject

	ODataType string `json:"@odata.type,omitempty"`
	// Action undocumented
	Action *string `json:"action,omitempty"`
	// AppScopeID undocumented
	AppScopeID *string `json:"appScopeId,omitempty"`
	// DirectoryScopeID undocumented
	DirectoryScopeID *string `json:"directoryScopeId,omitempty"`
	// IsValidationOnly undocumented
	IsValidationOnly *bool `json:"isValidationOnly,omitempty"`
	// Justification undocumented
	Justification *string `json:"justification,omitempty"`
	// PrincipalID undocumented
	PrincipalID *string `json:"principalId,omitempty"`
	// RoleDefinitionID undocumented
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty"`
	// ScheduleInfo undocumented
	ScheduleInfo *RequestSchedule `json:"scheduleInfo,omitempty"`
	// TargetScheduleID undocumented
	TargetScheduleID *string `json:"targetScheduleId,omitempty"`
	// TicketInfo undocumented
	TicketInfo *TicketInfo `json:"ticketInfo,omitempty"`
	// AppScope undocumented
	AppScope *AppScope `json:"appScope,omitempty"`
	// DirectoryScope undocumented
	DirectoryScope *DirectoryObject `json:"directoryScope,omitempty"`
	// Principal undocumented
	Principal *DirectoryObject `json:"principal,omitempty"`
	// RoleDefinition undocumented
	RoleDefinition *UnifiedRoleDefinition `json:"roleDefinition,omitempty"`
	// TargetSchedule undocumented
	TargetSchedule *UnifiedRoleEligibilitySchedule `json:"targetSchedule,omitempty"`
}

func NewUnifiedRoleEligibilityScheduleRequestObject() (*UnifiedRoleEligibilityScheduleRequestObject, error) {
	newUnifiedRoleEligibilityScheduleRequestObject := &UnifiedRoleEligibilityScheduleRequestObject{
		ODataType: "#microsoft.graph.UnifiedRoleEligibilityScheduleRequestObject",
	}
	return newUnifiedRoleEligibilityScheduleRequestObject, nil
}

// UnifiedRoleManagementAlert undocumented
type UnifiedRoleManagementAlert struct {
	// Entity is the base model of UnifiedRoleManagementAlert
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AlertDefinitionID undocumented
	AlertDefinitionID *string `json:"alertDefinitionId,omitempty"`
	// IncidentCount undocumented
	IncidentCount *int `json:"incidentCount,omitempty"`
	// IsActive undocumented
	IsActive *bool `json:"isActive,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// LastScannedDateTime undocumented
	LastScannedDateTime *time.Time `json:"lastScannedDateTime,omitempty"`
	// ScopeID undocumented
	ScopeID *string `json:"scopeId,omitempty"`
	// ScopeType undocumented
	ScopeType *string `json:"scopeType,omitempty"`
	// AlertConfiguration undocumented
	AlertConfiguration *UnifiedRoleManagementAlertConfiguration `json:"alertConfiguration,omitempty"`
	// AlertDefinition undocumented
	AlertDefinition *UnifiedRoleManagementAlertDefinition `json:"alertDefinition,omitempty"`
	// AlertIncidents undocumented
	AlertIncidents []UnifiedRoleManagementAlertIncident `json:"alertIncidents,omitempty"`
}

func NewUnifiedRoleManagementAlert() (*UnifiedRoleManagementAlert, error) {
	newUnifiedRoleManagementAlert := &UnifiedRoleManagementAlert{
		ODataType: "#microsoft.graph.UnifiedRoleManagementAlert",
	}
	return newUnifiedRoleManagementAlert, nil
}

// UnifiedRoleManagementAlertConfiguration undocumented
type UnifiedRoleManagementAlertConfiguration struct {
	// Entity is the base model of UnifiedRoleManagementAlertConfiguration
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AlertDefinitionID undocumented
	AlertDefinitionID *string `json:"alertDefinitionId,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// ScopeID undocumented
	ScopeID *string `json:"scopeId,omitempty"`
	// ScopeType undocumented
	ScopeType *string `json:"scopeType,omitempty"`
	// AlertDefinition undocumented
	AlertDefinition *UnifiedRoleManagementAlertDefinition `json:"alertDefinition,omitempty"`
}

func NewUnifiedRoleManagementAlertConfiguration() (*UnifiedRoleManagementAlertConfiguration, error) {
	newUnifiedRoleManagementAlertConfiguration := &UnifiedRoleManagementAlertConfiguration{
		ODataType: "#microsoft.graph.UnifiedRoleManagementAlertConfiguration",
	}
	return newUnifiedRoleManagementAlertConfiguration, nil
}

// UnifiedRoleManagementAlertDefinition undocumented
type UnifiedRoleManagementAlertDefinition struct {
	// Entity is the base model of UnifiedRoleManagementAlertDefinition
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// HowToPrevent undocumented
	HowToPrevent *string `json:"howToPrevent,omitempty"`
	// IsConfigurable undocumented
	IsConfigurable *bool `json:"isConfigurable,omitempty"`
	// IsRemediatable undocumented
	IsRemediatable *bool `json:"isRemediatable,omitempty"`
	// MitigationSteps undocumented
	MitigationSteps *string `json:"mitigationSteps,omitempty"`
	// ScopeID undocumented
	ScopeID *string `json:"scopeId,omitempty"`
	// ScopeType undocumented
	ScopeType *string `json:"scopeType,omitempty"`
	// SecurityImpact undocumented
	SecurityImpact *string `json:"securityImpact,omitempty"`
	// SeverityLevel undocumented
	SeverityLevel *AlertSeverity `json:"severityLevel,omitempty"`
}

func NewUnifiedRoleManagementAlertDefinition() (*UnifiedRoleManagementAlertDefinition, error) {
	newUnifiedRoleManagementAlertDefinition := &UnifiedRoleManagementAlertDefinition{
		ODataType: "#microsoft.graph.UnifiedRoleManagementAlertDefinition",
	}
	return newUnifiedRoleManagementAlertDefinition, nil
}

// UnifiedRoleManagementAlertIncident undocumented
type UnifiedRoleManagementAlertIncident struct {
	// Entity is the base model of UnifiedRoleManagementAlertIncident
	Entity

	ODataType string `json:"@odata.type,omitempty"`
}

func NewUnifiedRoleManagementAlertIncident() (*UnifiedRoleManagementAlertIncident, error) {
	newUnifiedRoleManagementAlertIncident := &UnifiedRoleManagementAlertIncident{
		ODataType: "#microsoft.graph.UnifiedRoleManagementAlertIncident",
	}
	return newUnifiedRoleManagementAlertIncident, nil
}

// UnifiedRoleManagementPolicy undocumented
type UnifiedRoleManagementPolicy struct {
	// Entity is the base model of UnifiedRoleManagementPolicy
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IsOrganizationDefault undocumented
	IsOrganizationDefault *bool `json:"isOrganizationDefault,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *Identity `json:"lastModifiedBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// ScopeID undocumented
	ScopeID *string `json:"scopeId,omitempty"`
	// ScopeType undocumented
	ScopeType *string `json:"scopeType,omitempty"`
	// EffectiveRules undocumented
	EffectiveRules []UnifiedRoleManagementPolicyRule `json:"effectiveRules,omitempty"`
	// Rules undocumented
	Rules []UnifiedRoleManagementPolicyRule `json:"rules,omitempty"`
}

func NewUnifiedRoleManagementPolicy() (*UnifiedRoleManagementPolicy, error) {
	newUnifiedRoleManagementPolicy := &UnifiedRoleManagementPolicy{
		ODataType: "#microsoft.graph.UnifiedRoleManagementPolicy",
	}
	return newUnifiedRoleManagementPolicy, nil
}

// UnifiedRoleManagementPolicyApprovalRule undocumented
type UnifiedRoleManagementPolicyApprovalRule struct {
	// UnifiedRoleManagementPolicyRule is the base model of UnifiedRoleManagementPolicyApprovalRule
	UnifiedRoleManagementPolicyRule

	ODataType string `json:"@odata.type,omitempty"`
	// Setting undocumented
	Setting *ApprovalSettings `json:"setting,omitempty"`
}

func NewUnifiedRoleManagementPolicyApprovalRule() (*UnifiedRoleManagementPolicyApprovalRule, error) {
	newUnifiedRoleManagementPolicyApprovalRule := &UnifiedRoleManagementPolicyApprovalRule{
		ODataType: "#microsoft.graph.UnifiedRoleManagementPolicyApprovalRule",
	}
	return newUnifiedRoleManagementPolicyApprovalRule, nil
}

// UnifiedRoleManagementPolicyAssignment undocumented
type UnifiedRoleManagementPolicyAssignment struct {
	// Entity is the base model of UnifiedRoleManagementPolicyAssignment
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// PolicyID undocumented
	PolicyID *string `json:"policyId,omitempty"`
	// RoleDefinitionID undocumented
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty"`
	// ScopeID undocumented
	ScopeID *string `json:"scopeId,omitempty"`
	// ScopeType undocumented
	ScopeType *string `json:"scopeType,omitempty"`
	// Policy undocumented
	Policy *UnifiedRoleManagementPolicy `json:"policy,omitempty"`
}

func NewUnifiedRoleManagementPolicyAssignment() (*UnifiedRoleManagementPolicyAssignment, error) {
	newUnifiedRoleManagementPolicyAssignment := &UnifiedRoleManagementPolicyAssignment{
		ODataType: "#microsoft.graph.UnifiedRoleManagementPolicyAssignment",
	}
	return newUnifiedRoleManagementPolicyAssignment, nil
}

// UnifiedRoleManagementPolicyAuthenticationContextRule undocumented
type UnifiedRoleManagementPolicyAuthenticationContextRule struct {
	// UnifiedRoleManagementPolicyRule is the base model of UnifiedRoleManagementPolicyAuthenticationContextRule
	UnifiedRoleManagementPolicyRule

	ODataType string `json:"@odata.type,omitempty"`
	// ClaimValue undocumented
	ClaimValue *string `json:"claimValue,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
}

func NewUnifiedRoleManagementPolicyAuthenticationContextRule() (*UnifiedRoleManagementPolicyAuthenticationContextRule, error) {
	newUnifiedRoleManagementPolicyAuthenticationContextRule := &UnifiedRoleManagementPolicyAuthenticationContextRule{
		ODataType: "#microsoft.graph.UnifiedRoleManagementPolicyAuthenticationContextRule",
	}
	return newUnifiedRoleManagementPolicyAuthenticationContextRule, nil
}

// UnifiedRoleManagementPolicyEnablementRule undocumented
type UnifiedRoleManagementPolicyEnablementRule struct {
	// UnifiedRoleManagementPolicyRule is the base model of UnifiedRoleManagementPolicyEnablementRule
	UnifiedRoleManagementPolicyRule

	ODataType string `json:"@odata.type,omitempty"`
	// EnabledRules undocumented
	EnabledRules []string `json:"enabledRules,omitempty"`
}

func NewUnifiedRoleManagementPolicyEnablementRule() (*UnifiedRoleManagementPolicyEnablementRule, error) {
	newUnifiedRoleManagementPolicyEnablementRule := &UnifiedRoleManagementPolicyEnablementRule{
		ODataType: "#microsoft.graph.UnifiedRoleManagementPolicyEnablementRule",
	}
	return newUnifiedRoleManagementPolicyEnablementRule, nil
}

// UnifiedRoleManagementPolicyExpirationRule undocumented
type UnifiedRoleManagementPolicyExpirationRule struct {
	// UnifiedRoleManagementPolicyRule is the base model of UnifiedRoleManagementPolicyExpirationRule
	UnifiedRoleManagementPolicyRule

	ODataType string `json:"@odata.type,omitempty"`
	// IsExpirationRequired undocumented
	IsExpirationRequired *bool `json:"isExpirationRequired,omitempty"`
	// MaximumDuration undocumented
	MaximumDuration *Duration `json:"maximumDuration,omitempty"`
}

func NewUnifiedRoleManagementPolicyExpirationRule() (*UnifiedRoleManagementPolicyExpirationRule, error) {
	newUnifiedRoleManagementPolicyExpirationRule := &UnifiedRoleManagementPolicyExpirationRule{
		ODataType: "#microsoft.graph.UnifiedRoleManagementPolicyExpirationRule",
	}
	return newUnifiedRoleManagementPolicyExpirationRule, nil
}

// UnifiedRoleManagementPolicyNotificationRule undocumented
type UnifiedRoleManagementPolicyNotificationRule struct {
	// UnifiedRoleManagementPolicyRule is the base model of UnifiedRoleManagementPolicyNotificationRule
	UnifiedRoleManagementPolicyRule

	ODataType string `json:"@odata.type,omitempty"`
	// IsDefaultRecipientsEnabled undocumented
	IsDefaultRecipientsEnabled *bool `json:"isDefaultRecipientsEnabled,omitempty"`
	// NotificationLevel undocumented
	NotificationLevel *string `json:"notificationLevel,omitempty"`
	// NotificationRecipients undocumented
	NotificationRecipients []string `json:"notificationRecipients,omitempty"`
	// NotificationType undocumented
	NotificationType *string `json:"notificationType,omitempty"`
	// RecipientType undocumented
	RecipientType *string `json:"recipientType,omitempty"`
}

func NewUnifiedRoleManagementPolicyNotificationRule() (*UnifiedRoleManagementPolicyNotificationRule, error) {
	newUnifiedRoleManagementPolicyNotificationRule := &UnifiedRoleManagementPolicyNotificationRule{
		ODataType: "#microsoft.graph.UnifiedRoleManagementPolicyNotificationRule",
	}
	return newUnifiedRoleManagementPolicyNotificationRule, nil
}

// UnifiedRoleManagementPolicyRule undocumented
type UnifiedRoleManagementPolicyRule struct {
	// Entity is the base model of UnifiedRoleManagementPolicyRule
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Target undocumented
	Target *UnifiedRoleManagementPolicyRuleTarget `json:"target,omitempty"`
}

func NewUnifiedRoleManagementPolicyRule() (*UnifiedRoleManagementPolicyRule, error) {
	newUnifiedRoleManagementPolicyRule := &UnifiedRoleManagementPolicyRule{
		ODataType: "#microsoft.graph.UnifiedRoleManagementPolicyRule",
	}
	return newUnifiedRoleManagementPolicyRule, nil
}

// UnifiedRoleManagementPolicyRuleTarget undocumented
type UnifiedRoleManagementPolicyRuleTarget struct {
	// Object is the base model of UnifiedRoleManagementPolicyRuleTarget
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Caller undocumented
	Caller *string `json:"caller,omitempty"`
	// EnforcedSettings undocumented
	EnforcedSettings []string `json:"enforcedSettings,omitempty"`
	// InheritableSettings undocumented
	InheritableSettings []string `json:"inheritableSettings,omitempty"`
	// Level undocumented
	Level *string `json:"level,omitempty"`
	// Operations undocumented
	Operations []string `json:"operations,omitempty"`
	// TargetObjects undocumented
	TargetObjects []DirectoryObject `json:"targetObjects,omitempty"`
}

func NewUnifiedRoleManagementPolicyRuleTarget() (*UnifiedRoleManagementPolicyRuleTarget, error) {
	newUnifiedRoleManagementPolicyRuleTarget := &UnifiedRoleManagementPolicyRuleTarget{
		ODataType: "#microsoft.graph.UnifiedRoleManagementPolicyRuleTarget",
	}
	return newUnifiedRoleManagementPolicyRuleTarget, nil
}

// UnifiedRolePermission undocumented
type UnifiedRolePermission struct {
	// Object is the base model of UnifiedRolePermission
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AllowedResourceActions undocumented
	AllowedResourceActions []string `json:"allowedResourceActions,omitempty"`
	// Condition undocumented
	Condition *string `json:"condition,omitempty"`
	// ExcludedResourceActions undocumented
	ExcludedResourceActions []string `json:"excludedResourceActions,omitempty"`
}

func NewUnifiedRolePermission() (*UnifiedRolePermission, error) {
	newUnifiedRolePermission := &UnifiedRolePermission{
		ODataType: "#microsoft.graph.UnifiedRolePermission",
	}
	return newUnifiedRolePermission, nil
}

// UnifiedRoleScheduleBase undocumented
type UnifiedRoleScheduleBase struct {
	// Entity is the base model of UnifiedRoleScheduleBase
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AppScopeID undocumented
	AppScopeID *string `json:"appScopeId,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// CreatedUsing undocumented
	CreatedUsing *string `json:"createdUsing,omitempty"`
	// DirectoryScopeID undocumented
	DirectoryScopeID *string `json:"directoryScopeId,omitempty"`
	// ModifiedDateTime undocumented
	ModifiedDateTime *time.Time `json:"modifiedDateTime,omitempty"`
	// PrincipalID undocumented
	PrincipalID *string `json:"principalId,omitempty"`
	// RoleDefinitionID undocumented
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty"`
	// Status undocumented
	Status *string `json:"status,omitempty"`
	// AppScope undocumented
	AppScope *AppScope `json:"appScope,omitempty"`
	// DirectoryScope undocumented
	DirectoryScope *DirectoryObject `json:"directoryScope,omitempty"`
	// Principal undocumented
	Principal *DirectoryObject `json:"principal,omitempty"`
	// RoleDefinition undocumented
	RoleDefinition *UnifiedRoleDefinition `json:"roleDefinition,omitempty"`
}

func NewUnifiedRoleScheduleBase() (*UnifiedRoleScheduleBase, error) {
	newUnifiedRoleScheduleBase := &UnifiedRoleScheduleBase{
		ODataType: "#microsoft.graph.UnifiedRoleScheduleBase",
	}
	return newUnifiedRoleScheduleBase, nil
}

// UnifiedRoleScheduleInstanceBase undocumented
type UnifiedRoleScheduleInstanceBase struct {
	// Entity is the base model of UnifiedRoleScheduleInstanceBase
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AppScopeID undocumented
	AppScopeID *string `json:"appScopeId,omitempty"`
	// DirectoryScopeID undocumented
	DirectoryScopeID *string `json:"directoryScopeId,omitempty"`
	// PrincipalID undocumented
	PrincipalID *string `json:"principalId,omitempty"`
	// RoleDefinitionID undocumented
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty"`
	// AppScope undocumented
	AppScope *AppScope `json:"appScope,omitempty"`
	// DirectoryScope undocumented
	DirectoryScope *DirectoryObject `json:"directoryScope,omitempty"`
	// Principal undocumented
	Principal *DirectoryObject `json:"principal,omitempty"`
	// RoleDefinition undocumented
	RoleDefinition *UnifiedRoleDefinition `json:"roleDefinition,omitempty"`
}

func NewUnifiedRoleScheduleInstanceBase() (*UnifiedRoleScheduleInstanceBase, error) {
	newUnifiedRoleScheduleInstanceBase := &UnifiedRoleScheduleInstanceBase{
		ODataType: "#microsoft.graph.UnifiedRoleScheduleInstanceBase",
	}
	return newUnifiedRoleScheduleInstanceBase, nil
}
