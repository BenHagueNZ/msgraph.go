// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// StartHoldMusicOperation undocumented
type StartHoldMusicOperation struct {
	// CommsOperation is the base model of StartHoldMusicOperation
	CommsOperation

	ODataType string `json:"@odata.type,omitempty"`
}

func NewStartHoldMusicOperation() (*StartHoldMusicOperation, error) {
	newStartHoldMusicOperation := &StartHoldMusicOperation{
		ODataType: "#microsoft.graph.StartHoldMusicOperation",
	}
	return newStartHoldMusicOperation, nil
}