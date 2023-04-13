// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// TrainingEventsContent undocumented
type TrainingEventsContent struct {
	// Object is the base model of TrainingEventsContent
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AssignedTrainingsInfos undocumented
	AssignedTrainingsInfos []AssignedTrainingInfo `json:"assignedTrainingsInfos,omitempty"`
	// TrainingsAssignedUserCount undocumented
	TrainingsAssignedUserCount *int `json:"trainingsAssignedUserCount,omitempty"`
}

func NewTrainingEventsContent() (*TrainingEventsContent, error) {
	newTrainingEventsContent := &TrainingEventsContent{
		ODataType: "#microsoft.graph.TrainingEventsContent",
	}
	return newTrainingEventsContent, nil
}
