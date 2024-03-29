// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// SizeRange undocumented
type SizeRange struct {
	// Object is the base model of SizeRange
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// MaximumSize undocumented
	MaximumSize *int `json:"maximumSize,omitempty"`
	// MinimumSize undocumented
	MinimumSize *int `json:"minimumSize,omitempty"`
}

func NewSizeRange() (*SizeRange, error) {
	newSizeRange := &SizeRange{
		ODataType: "#microsoft.graph.SizeRange",
	}
	return newSizeRange, nil
}
