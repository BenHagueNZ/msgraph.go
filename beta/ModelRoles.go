// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// RolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration undocumented
type RolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration struct {
	// UnifiedRoleManagementAlertConfiguration is the base model of RolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration
	UnifiedRoleManagementAlertConfiguration

	ODataType string `json:"@odata.type,omitempty"`
}

func NewRolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration() (*RolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration, error) {
	newRolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration := &RolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration{
		ODataType: "#microsoft.graph.RolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration",
	}
	return newRolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration, nil
}

// RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident undocumented
type RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident struct {
	// UnifiedRoleManagementAlertIncident is the base model of RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident
	UnifiedRoleManagementAlertIncident

	ODataType string `json:"@odata.type,omitempty"`
	// AssigneeDisplayName undocumented
	AssigneeDisplayName *string `json:"assigneeDisplayName,omitempty"`
	// AssigneeID undocumented
	AssigneeID *string `json:"assigneeId,omitempty"`
	// AssigneeUserPrincipalName undocumented
	AssigneeUserPrincipalName *string `json:"assigneeUserPrincipalName,omitempty"`
	// AssignmentCreatedDateTime undocumented
	AssignmentCreatedDateTime *time.Time `json:"assignmentCreatedDateTime,omitempty"`
	// RoleDefinitionID undocumented
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty"`
	// RoleDisplayName undocumented
	RoleDisplayName *string `json:"roleDisplayName,omitempty"`
	// RoleTemplateID undocumented
	RoleTemplateID *string `json:"roleTemplateId,omitempty"`
}

func NewRolesAssignedOutsidePrivilegedIdentityManagementAlertIncident() (*RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident, error) {
	newRolesAssignedOutsidePrivilegedIdentityManagementAlertIncident := &RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident{
		ODataType: "#microsoft.graph.RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident",
	}
	return newRolesAssignedOutsidePrivilegedIdentityManagementAlertIncident, nil
}