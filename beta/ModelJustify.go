// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// JustifyAction undocumented
type JustifyAction struct {
	// InformationProtectionAction is the base model of JustifyAction
	InformationProtectionAction

	ODataType string `json:"@odata.type,omitempty"`
}

func NewJustifyAction() (*JustifyAction, error) {
	newJustifyAction := &JustifyAction{
		ODataType: "#microsoft.graph.JustifyAction",
	}
	return newJustifyAction, nil
}
