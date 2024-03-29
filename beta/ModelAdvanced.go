// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// AdvancedThreatProtectionOnboardingDeviceSettingState undocumented
type AdvancedThreatProtectionOnboardingDeviceSettingState struct {
	// Entity is the base model of AdvancedThreatProtectionOnboardingDeviceSettingState
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ComplianceGracePeriodExpirationDateTime undocumented
	ComplianceGracePeriodExpirationDateTime *time.Time `json:"complianceGracePeriodExpirationDateTime,omitempty"`
	// DeviceID undocumented
	DeviceID *string `json:"deviceId,omitempty"`
	// DeviceModel undocumented
	DeviceModel *string `json:"deviceModel,omitempty"`
	// DeviceName undocumented
	DeviceName *string `json:"deviceName,omitempty"`
	// PlatformType undocumented
	PlatformType *DeviceType `json:"platformType,omitempty"`
	// Setting undocumented
	Setting *string `json:"setting,omitempty"`
	// SettingName undocumented
	SettingName *string `json:"settingName,omitempty"`
	// State undocumented
	State *ComplianceStatus `json:"state,omitempty"`
	// UserEmail undocumented
	UserEmail *string `json:"userEmail,omitempty"`
	// UserID undocumented
	UserID *string `json:"userId,omitempty"`
	// UserName undocumented
	UserName *string `json:"userName,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}

func NewAdvancedThreatProtectionOnboardingDeviceSettingState() (*AdvancedThreatProtectionOnboardingDeviceSettingState, error) {
	newAdvancedThreatProtectionOnboardingDeviceSettingState := &AdvancedThreatProtectionOnboardingDeviceSettingState{
		ODataType: "#microsoft.graph.AdvancedThreatProtectionOnboardingDeviceSettingState",
	}
	return newAdvancedThreatProtectionOnboardingDeviceSettingState, nil
}

// AdvancedThreatProtectionOnboardingStateSummary undocumented
type AdvancedThreatProtectionOnboardingStateSummary struct {
	// Entity is the base model of AdvancedThreatProtectionOnboardingStateSummary
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// CompliantDeviceCount undocumented
	CompliantDeviceCount *int `json:"compliantDeviceCount,omitempty"`
	// ConflictDeviceCount undocumented
	ConflictDeviceCount *int `json:"conflictDeviceCount,omitempty"`
	// ErrorDeviceCount undocumented
	ErrorDeviceCount *int `json:"errorDeviceCount,omitempty"`
	// NonCompliantDeviceCount undocumented
	NonCompliantDeviceCount *int `json:"nonCompliantDeviceCount,omitempty"`
	// NotApplicableDeviceCount undocumented
	NotApplicableDeviceCount *int `json:"notApplicableDeviceCount,omitempty"`
	// NotAssignedDeviceCount undocumented
	NotAssignedDeviceCount *int `json:"notAssignedDeviceCount,omitempty"`
	// RemediatedDeviceCount undocumented
	RemediatedDeviceCount *int `json:"remediatedDeviceCount,omitempty"`
	// UnknownDeviceCount undocumented
	UnknownDeviceCount *int `json:"unknownDeviceCount,omitempty"`
	// AdvancedThreatProtectionOnboardingDeviceSettingStates undocumented
	AdvancedThreatProtectionOnboardingDeviceSettingStates []AdvancedThreatProtectionOnboardingDeviceSettingState `json:"advancedThreatProtectionOnboardingDeviceSettingStates,omitempty"`
}

func NewAdvancedThreatProtectionOnboardingStateSummary() (*AdvancedThreatProtectionOnboardingStateSummary, error) {
	newAdvancedThreatProtectionOnboardingStateSummary := &AdvancedThreatProtectionOnboardingStateSummary{
		ODataType: "#microsoft.graph.AdvancedThreatProtectionOnboardingStateSummary",
	}
	return newAdvancedThreatProtectionOnboardingStateSummary, nil
}
