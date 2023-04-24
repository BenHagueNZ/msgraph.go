// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ActionResultPart undocumented
type ActionResultPart struct {
	// Object is the base model of ActionResultPart
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Error undocumented
	Error *PublicError `json:"error,omitempty"`
}

func NewActionResultPart() (*ActionResultPart, error) {
	newActionResultPart := &ActionResultPart{
		ODataType: "#microsoft.graph.ActionResultPart",
	}
	return newActionResultPart, nil
}

// ActionStep undocumented
type ActionStep struct {
	// Object is the base model of ActionStep
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ActionURL undocumented
	ActionURL *ActionURL `json:"actionUrl,omitempty"`
	// StepNumber undocumented
	StepNumber *int `json:"stepNumber,omitempty"`
	// Text undocumented
	Text *string `json:"text,omitempty"`
}

func NewActionStep() (*ActionStep, error) {
	newActionStep := &ActionStep{
		ODataType: "#microsoft.graph.ActionStep",
	}
	return newActionStep, nil
}

// ActionURL undocumented
type ActionURL struct {
	// Object is the base model of ActionURL
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// URL undocumented
	URL *string `json:"url,omitempty"`
}

func NewActionURL() (*ActionURL, error) {
	newActionURL := &ActionURL{
		ODataType: "#microsoft.graph.ActionUrl",
	}
	return newActionURL, nil
}
