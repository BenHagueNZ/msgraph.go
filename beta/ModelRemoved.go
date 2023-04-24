// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// RemovedState undocumented
type RemovedState struct {
	// Object is the base model of RemovedState
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Reason undocumented
	Reason *string `json:"reason,omitempty"`
}

func NewRemovedState() (*RemovedState, error) {
	newRemovedState := &RemovedState{
		ODataType: "#microsoft.graph.RemovedState",
	}
	return newRemovedState, nil
}