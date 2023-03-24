// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// DelegatedAdminAccessAssignment undocumented
type DelegatedAdminAccessAssignment struct {
	// Entity is the base model of DelegatedAdminAccessAssignment
	Entity
	// AccessContainer undocumented
	AccessContainer *DelegatedAdminAccessContainer `json:"accessContainer,omitempty"`
	// AccessDetails undocumented
	AccessDetails *DelegatedAdminAccessDetails `json:"accessDetails,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Status undocumented
	Status *DelegatedAdminAccessAssignmentStatus `json:"status,omitempty"`
}

// DelegatedAdminAccessContainer undocumented
type DelegatedAdminAccessContainer struct {
	// Object is the base model of DelegatedAdminAccessContainer
	Object
	// AccessContainerID undocumented
	AccessContainerID *string `json:"accessContainerId,omitempty"`
	// AccessContainerType undocumented
	AccessContainerType *DelegatedAdminAccessContainerType `json:"accessContainerType,omitempty"`
}

// DelegatedAdminAccessDetails undocumented
type DelegatedAdminAccessDetails struct {
	// Object is the base model of DelegatedAdminAccessDetails
	Object
	// UnifiedRoles undocumented
	UnifiedRoles []UnifiedRole `json:"unifiedRoles,omitempty"`
}

// DelegatedAdminCustomer undocumented
type DelegatedAdminCustomer struct {
	// Entity is the base model of DelegatedAdminCustomer
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// TenantID undocumented
	TenantID *string `json:"tenantId,omitempty"`
	// ServiceManagementDetails undocumented
	ServiceManagementDetails []DelegatedAdminServiceManagementDetail `json:"serviceManagementDetails,omitempty"`
}

// DelegatedAdminRelationship undocumented
type DelegatedAdminRelationship struct {
	// Entity is the base model of DelegatedAdminRelationship
	Entity
	// AccessDetails undocumented
	AccessDetails *DelegatedAdminAccessDetails `json:"accessDetails,omitempty"`
	// ActivatedDateTime undocumented
	ActivatedDateTime *time.Time `json:"activatedDateTime,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Customer undocumented
	Customer *DelegatedAdminRelationshipCustomerParticipant `json:"customer,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Duration undocumented
	Duration *Duration `json:"duration,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Status undocumented
	Status *DelegatedAdminRelationshipStatus `json:"status,omitempty"`
	// AccessAssignments undocumented
	AccessAssignments []DelegatedAdminAccessAssignment `json:"accessAssignments,omitempty"`
	// Operations undocumented
	Operations []DelegatedAdminRelationshipOperation `json:"operations,omitempty"`
	// Requests undocumented
	Requests []DelegatedAdminRelationshipRequestObject `json:"requests,omitempty"`
}

// DelegatedAdminRelationshipCustomerParticipant undocumented
type DelegatedAdminRelationshipCustomerParticipant struct {
	// Object is the base model of DelegatedAdminRelationshipCustomerParticipant
	Object
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// TenantID undocumented
	TenantID *string `json:"tenantId,omitempty"`
}

// DelegatedAdminRelationshipOperation undocumented
type DelegatedAdminRelationshipOperation struct {
	// Entity is the base model of DelegatedAdminRelationshipOperation
	Entity
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Data undocumented
	Data *string `json:"data,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// OperationType undocumented
	OperationType *DelegatedAdminRelationshipOperationType `json:"operationType,omitempty"`
	// Status undocumented
	Status *LongRunningOperationStatus `json:"status,omitempty"`
}

// DelegatedAdminRelationshipRequestObject undocumented
type DelegatedAdminRelationshipRequestObject struct {
	// Entity is the base model of DelegatedAdminRelationshipRequestObject
	Entity
	// Action undocumented
	Action *DelegatedAdminRelationshipRequestAction `json:"action,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Status undocumented
	Status *DelegatedAdminRelationshipRequestStatus `json:"status,omitempty"`
}

// DelegatedAdminServiceManagementDetail undocumented
type DelegatedAdminServiceManagementDetail struct {
	// Entity is the base model of DelegatedAdminServiceManagementDetail
	Entity
	// ServiceManagementURL undocumented
	ServiceManagementURL *string `json:"serviceManagementUrl,omitempty"`
	// ServiceName undocumented
	ServiceName *string `json:"serviceName,omitempty"`
}

// DelegatedPermissionClassification undocumented
type DelegatedPermissionClassification struct {
	// Entity is the base model of DelegatedPermissionClassification
	Entity
	// Classification undocumented
	Classification *PermissionClassificationType `json:"classification,omitempty"`
	// PermissionID undocumented
	PermissionID *string `json:"permissionId,omitempty"`
	// PermissionName undocumented
	PermissionName *string `json:"permissionName,omitempty"`
}
