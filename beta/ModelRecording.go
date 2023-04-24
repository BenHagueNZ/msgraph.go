// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// RecordingInfo undocumented
type RecordingInfo struct {
	// Object is the base model of RecordingInfo
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// InitiatedBy undocumented
	InitiatedBy *ParticipantInfo `json:"initiatedBy,omitempty"`
	// Initiator undocumented
	Initiator *IdentitySet `json:"initiator,omitempty"`
	// RecordingStatus undocumented
	RecordingStatus *RecordingStatus `json:"recordingStatus,omitempty"`
}

func NewRecordingInfo() (*RecordingInfo, error) {
	newRecordingInfo := &RecordingInfo{
		ODataType: "#microsoft.graph.RecordingInfo",
	}
	return newRecordingInfo, nil
}