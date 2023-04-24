// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// AutoLabeling undocumented
type AutoLabeling struct {
	// Object is the base model of AutoLabeling
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Message undocumented
	Message *string `json:"message,omitempty"`
	// SensitiveTypeIDs undocumented
	SensitiveTypeIDs []string `json:"sensitiveTypeIds,omitempty"`
}

func NewAutoLabeling() (*AutoLabeling, error) {
	newAutoLabeling := &AutoLabeling{
		ODataType: "#microsoft.graph.AutoLabeling",
	}
	return newAutoLabeling, nil
}

// AutoReviewSettings undocumented
type AutoReviewSettings struct {
	// Object is the base model of AutoReviewSettings
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// NotReviewedResult undocumented
	NotReviewedResult *string `json:"notReviewedResult,omitempty"`
}

func NewAutoReviewSettings() (*AutoReviewSettings, error) {
	newAutoReviewSettings := &AutoReviewSettings{
		ODataType: "#microsoft.graph.AutoReviewSettings",
	}
	return newAutoReviewSettings, nil
}
