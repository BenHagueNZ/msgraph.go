// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// RestrictedAppsViolation undocumented
type RestrictedAppsViolation struct {
	// Entity is the base model of RestrictedAppsViolation
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// DeviceConfigurationID undocumented
	DeviceConfigurationID *string `json:"deviceConfigurationId,omitempty"`
	// DeviceConfigurationName undocumented
	DeviceConfigurationName *string `json:"deviceConfigurationName,omitempty"`
	// DeviceName undocumented
	DeviceName *string `json:"deviceName,omitempty"`
	// ManagedDeviceID undocumented
	ManagedDeviceID *string `json:"managedDeviceId,omitempty"`
	// PlatformType undocumented
	PlatformType *PolicyPlatformType `json:"platformType,omitempty"`
	// RestrictedApps undocumented
	RestrictedApps []ManagedDeviceReportedApp `json:"restrictedApps,omitempty"`
	// RestrictedAppsState undocumented
	RestrictedAppsState *RestrictedAppsState `json:"restrictedAppsState,omitempty"`
	// UserID undocumented
	UserID *string `json:"userId,omitempty"`
	// UserName undocumented
	UserName *string `json:"userName,omitempty"`
}

func NewRestrictedAppsViolation() (*RestrictedAppsViolation, error) {
	newRestrictedAppsViolation := &RestrictedAppsViolation{
		ODataType: "#microsoft.graph.RestrictedAppsViolation",
	}
	return newRestrictedAppsViolation, nil
}