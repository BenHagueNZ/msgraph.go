// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// TrainingEventsContent undocumented
type TrainingEventsContent struct {
	// Object is the base model of TrainingEventsContent
	Object
	// AssignedTrainingsInfos undocumented
	AssignedTrainingsInfos []AssignedTrainingInfo `json:"assignedTrainingsInfos,omitempty"`
	// TrainingsAssignedUserCount undocumented
	TrainingsAssignedUserCount *int `json:"trainingsAssignedUserCount,omitempty"`
}
