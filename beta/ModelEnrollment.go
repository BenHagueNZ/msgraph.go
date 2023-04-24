// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// EnrollmentConfigurationAssignment undocumented
type EnrollmentConfigurationAssignment struct {
	// Entity is the base model of EnrollmentConfigurationAssignment
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Source undocumented
	Source *DeviceAndAppManagementAssignmentSource `json:"source,omitempty"`
	// SourceID undocumented
	SourceID *string `json:"sourceId,omitempty"`
	// Target undocumented
	Target *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
}

func NewEnrollmentConfigurationAssignment() (*EnrollmentConfigurationAssignment, error) {
	newEnrollmentConfigurationAssignment := &EnrollmentConfigurationAssignment{
		ODataType: "#microsoft.graph.EnrollmentConfigurationAssignment",
	}
	return newEnrollmentConfigurationAssignment, nil
}

// EnrollmentProfile undocumented
type EnrollmentProfile struct {
	// Entity is the base model of EnrollmentProfile
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ConfigurationEndpointURL undocumented
	ConfigurationEndpointURL *string `json:"configurationEndpointUrl,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// EnableAuthenticationViaCompanyPortal undocumented
	EnableAuthenticationViaCompanyPortal *bool `json:"enableAuthenticationViaCompanyPortal,omitempty"`
	// RequireCompanyPortalOnSetupAssistantEnrolledDevices undocumented
	RequireCompanyPortalOnSetupAssistantEnrolledDevices *bool `json:"requireCompanyPortalOnSetupAssistantEnrolledDevices,omitempty"`
	// RequiresUserAuthentication undocumented
	RequiresUserAuthentication *bool `json:"requiresUserAuthentication,omitempty"`
}

func NewEnrollmentProfile() (*EnrollmentProfile, error) {
	newEnrollmentProfile := &EnrollmentProfile{
		ODataType: "#microsoft.graph.EnrollmentProfile",
	}
	return newEnrollmentProfile, nil
}

// EnrollmentRestrictionsConfigurationPolicySetItem undocumented
type EnrollmentRestrictionsConfigurationPolicySetItem struct {
	// PolicySetItem is the base model of EnrollmentRestrictionsConfigurationPolicySetItem
	PolicySetItem

	ODataType string `json:"@odata.type,omitempty"`
	// Limit undocumented
	Limit *int `json:"limit,omitempty"`
	// Priority undocumented
	Priority *int `json:"priority,omitempty"`
}

func NewEnrollmentRestrictionsConfigurationPolicySetItem() (*EnrollmentRestrictionsConfigurationPolicySetItem, error) {
	newEnrollmentRestrictionsConfigurationPolicySetItem := &EnrollmentRestrictionsConfigurationPolicySetItem{
		ODataType: "#microsoft.graph.EnrollmentRestrictionsConfigurationPolicySetItem",
	}
	return newEnrollmentRestrictionsConfigurationPolicySetItem, nil
}

// EnrollmentTroubleshootingEvent undocumented
type EnrollmentTroubleshootingEvent struct {
	// DeviceManagementTroubleshootingEvent is the base model of EnrollmentTroubleshootingEvent
	DeviceManagementTroubleshootingEvent

	ODataType string `json:"@odata.type,omitempty"`
	// DeviceID undocumented
	DeviceID *string `json:"deviceId,omitempty"`
	// EnrollmentType undocumented
	EnrollmentType *DeviceEnrollmentType `json:"enrollmentType,omitempty"`
	// FailureCategory undocumented
	FailureCategory *DeviceEnrollmentFailureReason `json:"failureCategory,omitempty"`
	// FailureReason undocumented
	FailureReason *string `json:"failureReason,omitempty"`
	// ManagedDeviceIdentifier undocumented
	ManagedDeviceIdentifier *string `json:"managedDeviceIdentifier,omitempty"`
	// OperatingSystem undocumented
	OperatingSystem *string `json:"operatingSystem,omitempty"`
	// OsVersion undocumented
	OsVersion *string `json:"osVersion,omitempty"`
	// UserID undocumented
	UserID *string `json:"userId,omitempty"`
}

func NewEnrollmentTroubleshootingEvent() (*EnrollmentTroubleshootingEvent, error) {
	newEnrollmentTroubleshootingEvent := &EnrollmentTroubleshootingEvent{
		ODataType: "#microsoft.graph.EnrollmentTroubleshootingEvent",
	}
	return newEnrollmentTroubleshootingEvent, nil
}