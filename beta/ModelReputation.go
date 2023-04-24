// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ReputationCategory undocumented
type ReputationCategory struct {
	// Object is the base model of ReputationCategory
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Vendor undocumented
	Vendor *string `json:"vendor,omitempty"`
}

func NewReputationCategory() (*ReputationCategory, error) {
	newReputationCategory := &ReputationCategory{
		ODataType: "#microsoft.graph.ReputationCategory",
	}
	return newReputationCategory, nil
}