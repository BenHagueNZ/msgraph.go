// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// StopHoldMusicOperation undocumented
type StopHoldMusicOperation struct {
	// CommsOperation is the base model of StopHoldMusicOperation
	CommsOperation

	ODataType string `json:"@odata.type,omitempty"`
}

func NewStopHoldMusicOperation() (*StopHoldMusicOperation, error) {
	newStopHoldMusicOperation := &StopHoldMusicOperation{
		ODataType: "#microsoft.graph.StopHoldMusicOperation",
	}
	return newStopHoldMusicOperation, nil
}
