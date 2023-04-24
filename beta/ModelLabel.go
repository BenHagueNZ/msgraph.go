// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// LabelActionBase undocumented
type LabelActionBase struct {
	// Object is the base model of LabelActionBase
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
}

func NewLabelActionBase() (*LabelActionBase, error) {
	newLabelActionBase := &LabelActionBase{
		ODataType: "#microsoft.graph.LabelActionBase",
	}
	return newLabelActionBase, nil
}

// LabelDetails undocumented
type LabelDetails struct {
	// ParentLabelDetails is the base model of LabelDetails
	ParentLabelDetails

	ODataType string `json:"@odata.type,omitempty"`
}

func NewLabelDetails() (*LabelDetails, error) {
	newLabelDetails := &LabelDetails{
		ODataType: "#microsoft.graph.LabelDetails",
	}
	return newLabelDetails, nil
}

// LabelPolicy undocumented
type LabelPolicy struct {
	// Object is the base model of LabelPolicy
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
}

func NewLabelPolicy() (*LabelPolicy, error) {
	newLabelPolicy := &LabelPolicy{
		ODataType: "#microsoft.graph.LabelPolicy",
	}
	return newLabelPolicy, nil
}