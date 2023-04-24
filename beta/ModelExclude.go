// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ExcludeTarget undocumented
type ExcludeTarget struct {
	// Object is the base model of ExcludeTarget
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// TargetType undocumented
	TargetType *AuthenticationMethodTargetType `json:"targetType,omitempty"`
}

func NewExcludeTarget() (*ExcludeTarget, error) {
	newExcludeTarget := &ExcludeTarget{
		ODataType: "#microsoft.graph.ExcludeTarget",
	}
	return newExcludeTarget, nil
}
