
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// RemoveAccessApplyAction undocumented
type RemoveAccessApplyAction struct {
    // AccessReviewApplyAction is the base model of RemoveAccessApplyAction
    AccessReviewApplyAction
}

// RemoveContentFooterAction undocumented
type RemoveContentFooterAction struct {
    // InformationProtectionAction is the base model of RemoveContentFooterAction
    InformationProtectionAction
    // UIElementNames undocumented
    UIElementNames []string `json:"uiElementNames,omitempty"`
}

// RemoveContentHeaderAction undocumented
type RemoveContentHeaderAction struct {
    // InformationProtectionAction is the base model of RemoveContentHeaderAction
    InformationProtectionAction
    // UIElementNames undocumented
    UIElementNames []string `json:"uiElementNames,omitempty"`
}

// RemoveProtectionAction undocumented
type RemoveProtectionAction struct {
    // InformationProtectionAction is the base model of RemoveProtectionAction
    InformationProtectionAction
}

// RemoveWatermarkAction undocumented
type RemoveWatermarkAction struct {
    // InformationProtectionAction is the base model of RemoveWatermarkAction
    InformationProtectionAction
    // UIElementNames undocumented
    UIElementNames []string `json:"uiElementNames,omitempty"`
}
