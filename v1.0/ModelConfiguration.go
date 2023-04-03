// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ConfigurationManagerClientEnabledFeatures undocumented
type ConfigurationManagerClientEnabledFeatures struct {
	// Object is the base model of ConfigurationManagerClientEnabledFeatures
	Object

	ODataType string `json:"@odata.type"`
	// CompliancePolicy undocumented
	CompliancePolicy *bool `json:"compliancePolicy,omitempty"`
	// DeviceConfiguration undocumented
	DeviceConfiguration *bool `json:"deviceConfiguration,omitempty"`
	// Inventory undocumented
	Inventory *bool `json:"inventory,omitempty"`
	// ModernApps undocumented
	ModernApps *bool `json:"modernApps,omitempty"`
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

// ConfigurationManagerCollectionAssignmentTarget undocumented
type ConfigurationManagerCollectionAssignmentTarget struct {
	// DeviceAndAppManagementAssignmentTarget is the base model of ConfigurationManagerCollectionAssignmentTarget
	DeviceAndAppManagementAssignmentTarget

	ODataType string `json:"@odata.type"`
	// CollectionID undocumented
	CollectionID *string `json:"collectionId,omitempty"`
}

func NewConfigurationManagerCollectionAssignmentTarget() (*ConfigurationManagerCollectionAssignmentTarget, error) {
	newConfigurationManagerCollectionAssignmentTarget := &ConfigurationManagerCollectionAssignmentTarget{
		ODataType: "#microsoft.graph.ConfigurationManagerCollectionAssignmentTarget",
	}
	return newConfigurationManagerCollectionAssignmentTarget, nil
}
