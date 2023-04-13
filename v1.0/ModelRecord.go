// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// RecordOperation undocumented
type RecordOperation struct {
	// CommsOperation is the base model of RecordOperation
	CommsOperation

	ODataType string `json:"@odata.type,omitempty"`
	// RecordingAccessToken undocumented
	RecordingAccessToken *string `json:"recordingAccessToken,omitempty"`
	// RecordingLocation undocumented
	RecordingLocation *string `json:"recordingLocation,omitempty"`
}

func NewRecordOperation() (*RecordOperation, error) {
	newRecordOperation := &RecordOperation{
		ODataType: "#microsoft.graph.RecordOperation",
	}
	return newRecordOperation, nil
}
