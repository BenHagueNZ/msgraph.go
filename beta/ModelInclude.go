// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// IncludeAllAccountTargetContent undocumented
type IncludeAllAccountTargetContent struct {
	// AccountTargetContent is the base model of IncludeAllAccountTargetContent
	AccountTargetContent

	ODataType string `json:"@odata.type,omitempty"`
}

func NewIncludeAllAccountTargetContent() (*IncludeAllAccountTargetContent, error) {
	newIncludeAllAccountTargetContent := &IncludeAllAccountTargetContent{
		ODataType: "#microsoft.graph.IncludeAllAccountTargetContent",
	}
	return newIncludeAllAccountTargetContent, nil
}

// IncludeTarget undocumented
type IncludeTarget struct {
	// Object is the base model of IncludeTarget
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// TargetType undocumented
	TargetType *AuthenticationMethodTargetType `json:"targetType,omitempty"`
}

func NewIncludeTarget() (*IncludeTarget, error) {
	newIncludeTarget := &IncludeTarget{
		ODataType: "#microsoft.graph.IncludeTarget",
	}
	return newIncludeTarget, nil
}