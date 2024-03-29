// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// TargetedManagedAppConfiguration undocumented
type TargetedManagedAppConfiguration struct {
	// ManagedAppConfiguration is the base model of TargetedManagedAppConfiguration
	ManagedAppConfiguration

	ODataType string `json:"@odata.type,omitempty"`
	// AppGroupType undocumented
	AppGroupType *TargetedManagedAppGroupType `json:"appGroupType,omitempty"`
	// DeployedAppCount undocumented
	DeployedAppCount *int `json:"deployedAppCount,omitempty"`
	// IsAssigned undocumented
	IsAssigned *bool `json:"isAssigned,omitempty"`
	// TargetedAppManagementLevels undocumented
	TargetedAppManagementLevels *AppManagementLevel `json:"targetedAppManagementLevels,omitempty"`
	// Apps undocumented
	Apps []ManagedMobileApp `json:"apps,omitempty"`
	// Assignments undocumented
	Assignments []TargetedManagedAppPolicyAssignment `json:"assignments,omitempty"`
	// DeploymentSummary undocumented
	DeploymentSummary *ManagedAppPolicyDeploymentSummary `json:"deploymentSummary,omitempty"`
}

func NewTargetedManagedAppConfiguration() (*TargetedManagedAppConfiguration, error) {
	newTargetedManagedAppConfiguration := &TargetedManagedAppConfiguration{
		ODataType: "#microsoft.graph.TargetedManagedAppConfiguration",
	}
	return newTargetedManagedAppConfiguration, nil
}

// TargetedManagedAppConfigurationPolicySetItem undocumented
type TargetedManagedAppConfigurationPolicySetItem struct {
	// PolicySetItem is the base model of TargetedManagedAppConfigurationPolicySetItem
	PolicySetItem

	ODataType string `json:"@odata.type,omitempty"`
}

func NewTargetedManagedAppConfigurationPolicySetItem() (*TargetedManagedAppConfigurationPolicySetItem, error) {
	newTargetedManagedAppConfigurationPolicySetItem := &TargetedManagedAppConfigurationPolicySetItem{
		ODataType: "#microsoft.graph.TargetedManagedAppConfigurationPolicySetItem",
	}
	return newTargetedManagedAppConfigurationPolicySetItem, nil
}

// TargetedManagedAppPolicyAssignment undocumented
type TargetedManagedAppPolicyAssignment struct {
	// Entity is the base model of TargetedManagedAppPolicyAssignment
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Source undocumented
	Source *DeviceAndAppManagementAssignmentSource `json:"source,omitempty"`
	// SourceID undocumented
	SourceID *string `json:"sourceId,omitempty"`
	// Target undocumented
	Target *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
}

func NewTargetedManagedAppPolicyAssignment() (*TargetedManagedAppPolicyAssignment, error) {
	newTargetedManagedAppPolicyAssignment := &TargetedManagedAppPolicyAssignment{
		ODataType: "#microsoft.graph.TargetedManagedAppPolicyAssignment",
	}
	return newTargetedManagedAppPolicyAssignment, nil
}

// TargetedManagedAppProtection undocumented
type TargetedManagedAppProtection struct {
	// ManagedAppProtection is the base model of TargetedManagedAppProtection
	ManagedAppProtection

	ODataType string `json:"@odata.type,omitempty"`
	// AppGroupType undocumented
	AppGroupType *TargetedManagedAppGroupType `json:"appGroupType,omitempty"`
	// IsAssigned undocumented
	IsAssigned *bool `json:"isAssigned,omitempty"`
	// TargetedAppManagementLevels undocumented
	TargetedAppManagementLevels *AppManagementLevel `json:"targetedAppManagementLevels,omitempty"`
	// Assignments undocumented
	Assignments []TargetedManagedAppPolicyAssignment `json:"assignments,omitempty"`
}

func NewTargetedManagedAppProtection() (*TargetedManagedAppProtection, error) {
	newTargetedManagedAppProtection := &TargetedManagedAppProtection{
		ODataType: "#microsoft.graph.TargetedManagedAppProtection",
	}
	return newTargetedManagedAppProtection, nil
}
