// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ApplyLabelAction undocumented
type ApplyLabelAction struct {
	// InformationProtectionAction is the base model of ApplyLabelAction
	InformationProtectionAction

	ODataType string `json:"@odata.type,omitempty"`
	// Actions undocumented
	Actions []InformationProtectionAction `json:"actions,omitempty"`
	// ActionSource undocumented
	ActionSource *ActionSource `json:"actionSource,omitempty"`
	// Label undocumented
	Label *LabelDetails `json:"label,omitempty"`
	// ResponsibleSensitiveTypeIDs undocumented
	ResponsibleSensitiveTypeIDs []UUID `json:"responsibleSensitiveTypeIds,omitempty"`
}

func NewApplyLabelAction() (*ApplyLabelAction, error) {
	newApplyLabelAction := &ApplyLabelAction{
		ODataType: "#microsoft.graph.ApplyLabelAction",
	}
	return newApplyLabelAction, nil
}