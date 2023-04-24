// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ComanagementEligibleDevice undocumented
type ComanagementEligibleDevice struct {
	// Entity is the base model of ComanagementEligibleDevice
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ClientRegistrationStatus undocumented
	ClientRegistrationStatus *DeviceRegistrationState `json:"clientRegistrationStatus,omitempty"`
	// DeviceName undocumented
	DeviceName *string `json:"deviceName,omitempty"`
	// DeviceType undocumented
	DeviceType *DeviceType `json:"deviceType,omitempty"`
	// EntitySource undocumented
	EntitySource *int `json:"entitySource,omitempty"`
	// ManagementAgents undocumented
	ManagementAgents *ManagementAgentType `json:"managementAgents,omitempty"`
	// ManagementState undocumented
	ManagementState *ManagementState `json:"managementState,omitempty"`
	// Manufacturer undocumented
	Manufacturer *string `json:"manufacturer,omitempty"`
	// MDMStatus undocumented
	MDMStatus *string `json:"mdmStatus,omitempty"`
	// Model undocumented
	Model *string `json:"model,omitempty"`
	// OsDescription undocumented
	OsDescription *string `json:"osDescription,omitempty"`
	// OsVersion undocumented
	OsVersion *string `json:"osVersion,omitempty"`
	// OwnerType undocumented
	OwnerType *OwnerType `json:"ownerType,omitempty"`
	// ReferenceID undocumented
	ReferenceID *string `json:"referenceId,omitempty"`
	// SerialNumber undocumented
	SerialNumber *string `json:"serialNumber,omitempty"`
	// Status undocumented
	Status *ComanagementEligibleType `json:"status,omitempty"`
	// Upn undocumented
	Upn *string `json:"upn,omitempty"`
	// UserEmail undocumented
	UserEmail *string `json:"userEmail,omitempty"`
	// UserID undocumented
	UserID *string `json:"userId,omitempty"`
	// UserName undocumented
	UserName *string `json:"userName,omitempty"`
}

func NewComanagementEligibleDevice() (*ComanagementEligibleDevice, error) {
	newComanagementEligibleDevice := &ComanagementEligibleDevice{
		ODataType: "#microsoft.graph.ComanagementEligibleDevice",
	}
	return newComanagementEligibleDevice, nil
}

// ComanagementEligibleDevicesSummary undocumented
type ComanagementEligibleDevicesSummary struct {
	// Object is the base model of ComanagementEligibleDevicesSummary
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ComanagedCount undocumented
	ComanagedCount *int `json:"comanagedCount,omitempty"`
	// EligibleButNotAzureAdJoinedCount undocumented
	EligibleButNotAzureAdJoinedCount *int `json:"eligibleButNotAzureAdJoinedCount,omitempty"`
	// EligibleCount undocumented
	EligibleCount *int `json:"eligibleCount,omitempty"`
	// IneligibleCount undocumented
	IneligibleCount *int `json:"ineligibleCount,omitempty"`
	// NeedsOsUpdateCount undocumented
	NeedsOsUpdateCount *int `json:"needsOsUpdateCount,omitempty"`
	// ScheduledForEnrollmentCount undocumented
	ScheduledForEnrollmentCount *int `json:"scheduledForEnrollmentCount,omitempty"`
}

func NewComanagementEligibleDevicesSummary() (*ComanagementEligibleDevicesSummary, error) {
	newComanagementEligibleDevicesSummary := &ComanagementEligibleDevicesSummary{
		ODataType: "#microsoft.graph.ComanagementEligibleDevicesSummary",
	}
	return newComanagementEligibleDevicesSummary, nil
}