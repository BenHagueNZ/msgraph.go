// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// WorkPosition undocumented
type WorkPosition struct {
	// ItemFacet is the base model of WorkPosition
	ItemFacet

	ODataType string `json:"@odata.type,omitempty"`
	// Categories undocumented
	Categories []string `json:"categories,omitempty"`
	// Colleagues undocumented
	Colleagues []RelatedPerson `json:"colleagues,omitempty"`
	// Detail undocumented
	Detail *PositionDetail `json:"detail,omitempty"`
	// IsCurrent undocumented
	IsCurrent *bool `json:"isCurrent,omitempty"`
	// Manager undocumented
	Manager *RelatedPerson `json:"manager,omitempty"`
}

func NewWorkPosition() (*WorkPosition, error) {
	newWorkPosition := &WorkPosition{
		ODataType: "#microsoft.graph.WorkPosition",
	}
	return newWorkPosition, nil
}
