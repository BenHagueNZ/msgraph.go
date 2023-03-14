// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// EnrollmentConfigurationAssignment undocumented
type EnrollmentConfigurationAssignment struct {
	// Entity is the base model of EnrollmentConfigurationAssignment
	Entity
	// Target undocumented
	Target *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
}

// EnrollmentTroubleshootingEvent undocumented
type EnrollmentTroubleshootingEvent struct {
	// DeviceManagementTroubleshootingEvent is the base model of EnrollmentTroubleshootingEvent
	DeviceManagementTroubleshootingEvent
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
