// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Configuration undocumented
type Configuration struct {
	// Object is the base model of Configuration
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AuthorizedAppIDs undocumented
	AuthorizedAppIDs []string `json:"authorizedAppIds,omitempty"`
	// AuthorizedApps undocumented
	AuthorizedApps []string `json:"authorizedApps,omitempty"`
}

func NewConfiguration() (*Configuration, error) {
	newConfiguration := &Configuration{
		ODataType: "#microsoft.graph.Configuration",
	}
	return newConfiguration, nil
}

// ConfigurationManagerAction undocumented
type ConfigurationManagerAction struct {
	// Object is the base model of ConfigurationManagerAction
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Action undocumented
	Action *ConfigurationManagerActionType `json:"action,omitempty"`
}

func NewConfigurationManagerAction() (*ConfigurationManagerAction, error) {
	newConfigurationManagerAction := &ConfigurationManagerAction{
		ODataType: "#microsoft.graph.ConfigurationManagerAction",
	}
	return newConfigurationManagerAction, nil
}

// ConfigurationManagerActionResult undocumented
type ConfigurationManagerActionResult struct {
	// DeviceActionResult is the base model of ConfigurationManagerActionResult
	DeviceActionResult

	ODataType string `json:"@odata.type,omitempty"`
	// ActionDeliveryStatus undocumented
	ActionDeliveryStatus *ConfigurationManagerActionDeliveryStatus `json:"actionDeliveryStatus,omitempty"`
	// ErrorCode undocumented
	ErrorCode *int `json:"errorCode,omitempty"`
}

func NewConfigurationManagerActionResult() (*ConfigurationManagerActionResult, error) {
	newConfigurationManagerActionResult := &ConfigurationManagerActionResult{
		ODataType: "#microsoft.graph.ConfigurationManagerActionResult",
	}
	return newConfigurationManagerActionResult, nil
}

// ConfigurationManagerClientEnabledFeatures undocumented
type ConfigurationManagerClientEnabledFeatures struct {
	// Object is the base model of ConfigurationManagerClientEnabledFeatures
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// CompliancePolicy undocumented
	CompliancePolicy *bool `json:"compliancePolicy,omitempty"`
	// DeviceConfiguration undocumented
	DeviceConfiguration *bool `json:"deviceConfiguration,omitempty"`
	// EndpointProtection undocumented
	EndpointProtection *bool `json:"endpointProtection,omitempty"`
	// Inventory undocumented
	Inventory *bool `json:"inventory,omitempty"`
	// ModernApps undocumented
	ModernApps *bool `json:"modernApps,omitempty"`
	// OfficeApps undocumented
	OfficeApps *bool `json:"officeApps,omitempty"`
	// ResourceAccess undocumented
	ResourceAccess *bool `json:"resourceAccess,omitempty"`
	// WindowsUpdateForBusiness undocumented
	WindowsUpdateForBusiness *bool `json:"windowsUpdateForBusiness,omitempty"`
}

func NewConfigurationManagerClientEnabledFeatures() (*ConfigurationManagerClientEnabledFeatures, error) {
	newConfigurationManagerClientEnabledFeatures := &ConfigurationManagerClientEnabledFeatures{
		ODataType: "#microsoft.graph.ConfigurationManagerClientEnabledFeatures",
	}
	return newConfigurationManagerClientEnabledFeatures, nil
}

// ConfigurationManagerClientHealthState undocumented
type ConfigurationManagerClientHealthState struct {
	// Object is the base model of ConfigurationManagerClientHealthState
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ErrorCode undocumented
	ErrorCode *int `json:"errorCode,omitempty"`
	// LastSyncDateTime undocumented
	LastSyncDateTime *time.Time `json:"lastSyncDateTime,omitempty"`
	// State undocumented
	State *ConfigurationManagerClientState `json:"state,omitempty"`
}

func NewConfigurationManagerClientHealthState() (*ConfigurationManagerClientHealthState, error) {
	newConfigurationManagerClientHealthState := &ConfigurationManagerClientHealthState{
		ODataType: "#microsoft.graph.ConfigurationManagerClientHealthState",
	}
	return newConfigurationManagerClientHealthState, nil
}

// ConfigurationManagerClientInformation undocumented
type ConfigurationManagerClientInformation struct {
	// Object is the base model of ConfigurationManagerClientInformation
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ClientIdentifier undocumented
	ClientIdentifier *string `json:"clientIdentifier,omitempty"`
	// ClientVersion undocumented
	ClientVersion *string `json:"clientVersion,omitempty"`
	// IsBlocked undocumented
	IsBlocked *bool `json:"isBlocked,omitempty"`
}

func NewConfigurationManagerClientInformation() (*ConfigurationManagerClientInformation, error) {
	newConfigurationManagerClientInformation := &ConfigurationManagerClientInformation{
		ODataType: "#microsoft.graph.ConfigurationManagerClientInformation",
	}
	return newConfigurationManagerClientInformation, nil
}

// ConfigurationManagerCollectionAssignmentTarget undocumented
type ConfigurationManagerCollectionAssignmentTarget struct {
	// DeviceAndAppManagementAssignmentTarget is the base model of ConfigurationManagerCollectionAssignmentTarget
	DeviceAndAppManagementAssignmentTarget

	ODataType string `json:"@odata.type,omitempty"`
	// CollectionID undocumented
	CollectionID *string `json:"collectionId,omitempty"`
}

func NewConfigurationManagerCollectionAssignmentTarget() (*ConfigurationManagerCollectionAssignmentTarget, error) {
	newConfigurationManagerCollectionAssignmentTarget := &ConfigurationManagerCollectionAssignmentTarget{
		ODataType: "#microsoft.graph.ConfigurationManagerCollectionAssignmentTarget",
	}
	return newConfigurationManagerCollectionAssignmentTarget, nil
}
