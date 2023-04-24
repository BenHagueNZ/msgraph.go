// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ClassifcationErrorBase undocumented
type ClassifcationErrorBase struct {
	// Object is the base model of ClassifcationErrorBase
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// InnerError undocumented
	InnerError *ClassificationInnerError `json:"innerError,omitempty"`
	// Message undocumented
	Message *string `json:"message,omitempty"`
	// Target undocumented
	Target *string `json:"target,omitempty"`
}

func NewClassifcationErrorBase() (*ClassifcationErrorBase, error) {
	newClassifcationErrorBase := &ClassifcationErrorBase{
		ODataType: "#microsoft.graph.ClassifcationErrorBase",
	}
	return newClassifcationErrorBase, nil
}