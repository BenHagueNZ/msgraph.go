// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// InvalidLicenseAlertConfiguration undocumented
type InvalidLicenseAlertConfiguration struct {
	// UnifiedRoleManagementAlertConfiguration is the base model of InvalidLicenseAlertConfiguration
	UnifiedRoleManagementAlertConfiguration

	ODataType string `json:"@odata.type,omitempty"`
}

func NewInvalidLicenseAlertConfiguration() (*InvalidLicenseAlertConfiguration, error) {
	newInvalidLicenseAlertConfiguration := &InvalidLicenseAlertConfiguration{
		ODataType: "#microsoft.graph.InvalidLicenseAlertConfiguration",
	}
	return newInvalidLicenseAlertConfiguration, nil
}

// InvalidLicenseAlertIncident undocumented
type InvalidLicenseAlertIncident struct {
	// UnifiedRoleManagementAlertIncident is the base model of InvalidLicenseAlertIncident
	UnifiedRoleManagementAlertIncident

	ODataType string `json:"@odata.type,omitempty"`
	// TenantLicenseStatus undocumented
	TenantLicenseStatus *string `json:"tenantLicenseStatus,omitempty"`
}

func NewInvalidLicenseAlertIncident() (*InvalidLicenseAlertIncident, error) {
	newInvalidLicenseAlertIncident := &InvalidLicenseAlertIncident{
		ODataType: "#microsoft.graph.InvalidLicenseAlertIncident",
	}
	return newInvalidLicenseAlertIncident, nil
}
