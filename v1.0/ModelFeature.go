// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// FeatureRolloutPolicy undocumented
type FeatureRolloutPolicy struct {
	// Entity is the base model of FeatureRolloutPolicy
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Feature undocumented
	Feature *StagedFeatureName `json:"feature,omitempty"`
	// IsAppliedToOrganization undocumented
	IsAppliedToOrganization *bool `json:"isAppliedToOrganization,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// AppliesTo undocumented
	AppliesTo []DirectoryObject `json:"appliesTo,omitempty"`
}

func NewFeatureRolloutPolicy() (*FeatureRolloutPolicy, error) {
	newFeatureRolloutPolicy := &FeatureRolloutPolicy{
		ODataType: "#microsoft.graph.FeatureRolloutPolicy",
	}
	return newFeatureRolloutPolicy, nil
}

// FeatureTarget undocumented
type FeatureTarget struct {
	// Object is the base model of FeatureTarget
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// TargetType undocumented
	TargetType *FeatureTargetType `json:"targetType,omitempty"`
}

func NewFeatureTarget() (*FeatureTarget, error) {
	newFeatureTarget := &FeatureTarget{
		ODataType: "#microsoft.graph.FeatureTarget",
	}
	return newFeatureTarget, nil
}
