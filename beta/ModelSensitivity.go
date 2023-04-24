// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// SensitivityLabel undocumented
type SensitivityLabel struct {
	// Entity is the base model of SensitivityLabel
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ApplicableTo undocumented
	ApplicableTo *SensitivityLabelTarget `json:"applicableTo,omitempty"`
	// ApplicationMode undocumented
	ApplicationMode *ApplicationMode `json:"applicationMode,omitempty"`
	// AssignedPolicies undocumented
	AssignedPolicies []LabelPolicy `json:"assignedPolicies,omitempty"`
	// AutoLabeling undocumented
	AutoLabeling *AutoLabeling `json:"autoLabeling,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IsDefault undocumented
	IsDefault *bool `json:"isDefault,omitempty"`
	// IsEndpointProtectionEnabled undocumented
	IsEndpointProtectionEnabled *bool `json:"isEndpointProtectionEnabled,omitempty"`
	// LabelActions undocumented
	LabelActions []LabelActionBase `json:"labelActions,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Priority undocumented
	Priority *int `json:"priority,omitempty"`
	// ToolTip undocumented
	ToolTip *string `json:"toolTip,omitempty"`
	// Sublabels undocumented
	Sublabels []SensitivityLabel `json:"sublabels,omitempty"`
}

func NewSensitivityLabel() (*SensitivityLabel, error) {
	newSensitivityLabel := &SensitivityLabel{
		ODataType: "#microsoft.graph.SensitivityLabel",
	}
	return newSensitivityLabel, nil
}

// SensitivityLabelAssignment undocumented
type SensitivityLabelAssignment struct {
	// Object is the base model of SensitivityLabelAssignment
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AssignmentMethod undocumented
	AssignmentMethod *SensitivityLabelAssignmentMethod `json:"assignmentMethod,omitempty"`
	// SensitivityLabelID undocumented
	SensitivityLabelID *string `json:"sensitivityLabelId,omitempty"`
	// TenantID undocumented
	TenantID *string `json:"tenantId,omitempty"`
}

func NewSensitivityLabelAssignment() (*SensitivityLabelAssignment, error) {
	newSensitivityLabelAssignment := &SensitivityLabelAssignment{
		ODataType: "#microsoft.graph.SensitivityLabelAssignment",
	}
	return newSensitivityLabelAssignment, nil
}

// SensitivityPolicySettings undocumented
type SensitivityPolicySettings struct {
	// Entity is the base model of SensitivityPolicySettings
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ApplicableTo undocumented
	ApplicableTo *SensitivityLabelTarget `json:"applicableTo,omitempty"`
	// DowngradeSensitivityRequiresJustification undocumented
	DowngradeSensitivityRequiresJustification *bool `json:"downgradeSensitivityRequiresJustification,omitempty"`
	// HelpWebURL undocumented
	HelpWebURL *string `json:"helpWebUrl,omitempty"`
	// IsMandatory undocumented
	IsMandatory *bool `json:"isMandatory,omitempty"`
}

func NewSensitivityPolicySettings() (*SensitivityPolicySettings, error) {
	newSensitivityPolicySettings := &SensitivityPolicySettings{
		ODataType: "#microsoft.graph.SensitivityPolicySettings",
	}
	return newSensitivityPolicySettings, nil
}
