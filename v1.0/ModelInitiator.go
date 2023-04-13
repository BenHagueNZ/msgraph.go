// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Initiator undocumented
type Initiator struct {
	// Identity is the base model of Initiator
	Identity

	ODataType string `json:"@odata.type,omitempty"`
	// InitiatorType undocumented
	InitiatorType *InitiatorType `json:"initiatorType,omitempty"`
}

func NewInitiator() (*Initiator, error) {
	newInitiator := &Initiator{
		ODataType: "#microsoft.graph.Initiator",
	}
	return newInitiator, nil
}
