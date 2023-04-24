// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// TooManyGlobalAdminsAssignedToTenantAlertConfiguration undocumented
type TooManyGlobalAdminsAssignedToTenantAlertConfiguration struct {
	// UnifiedRoleManagementAlertConfiguration is the base model of TooManyGlobalAdminsAssignedToTenantAlertConfiguration
	UnifiedRoleManagementAlertConfiguration

	ODataType string `json:"@odata.type,omitempty"`
	// GlobalAdminCountThreshold undocumented
	GlobalAdminCountThreshold *int `json:"globalAdminCountThreshold,omitempty"`
	// PercentageOfGlobalAdminsOutOfRolesThreshold undocumented
	PercentageOfGlobalAdminsOutOfRolesThreshold *int `json:"percentageOfGlobalAdminsOutOfRolesThreshold,omitempty"`
}

func NewTooManyGlobalAdminsAssignedToTenantAlertConfiguration() (*TooManyGlobalAdminsAssignedToTenantAlertConfiguration, error) {
	newTooManyGlobalAdminsAssignedToTenantAlertConfiguration := &TooManyGlobalAdminsAssignedToTenantAlertConfiguration{
		ODataType: "#microsoft.graph.TooManyGlobalAdminsAssignedToTenantAlertConfiguration",
	}
	return newTooManyGlobalAdminsAssignedToTenantAlertConfiguration, nil
}

// TooManyGlobalAdminsAssignedToTenantAlertIncident undocumented
type TooManyGlobalAdminsAssignedToTenantAlertIncident struct {
	// UnifiedRoleManagementAlertIncident is the base model of TooManyGlobalAdminsAssignedToTenantAlertIncident
	UnifiedRoleManagementAlertIncident

	ODataType string `json:"@odata.type,omitempty"`
	// AssigneeDisplayName undocumented
	AssigneeDisplayName *string `json:"assigneeDisplayName,omitempty"`
	// AssigneeID undocumented
	AssigneeID *string `json:"assigneeId,omitempty"`
	// AssigneeUserPrincipalName undocumented
	AssigneeUserPrincipalName *string `json:"assigneeUserPrincipalName,omitempty"`
}

func NewTooManyGlobalAdminsAssignedToTenantAlertIncident() (*TooManyGlobalAdminsAssignedToTenantAlertIncident, error) {
	newTooManyGlobalAdminsAssignedToTenantAlertIncident := &TooManyGlobalAdminsAssignedToTenantAlertIncident{
		ODataType: "#microsoft.graph.TooManyGlobalAdminsAssignedToTenantAlertIncident",
	}
	return newTooManyGlobalAdminsAssignedToTenantAlertIncident, nil
}
