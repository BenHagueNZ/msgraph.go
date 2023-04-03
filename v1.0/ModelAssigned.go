// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// AssignedLabel undocumented
type AssignedLabel struct {
	// Object is the base model of AssignedLabel
	Object

	ODataType string `json:"@odata.type"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LabelID undocumented
	LabelID *string `json:"labelId,omitempty"`
}

func NewAssignedLabel() (*AssignedLabel, error) {
	newAssignedLabel := &AssignedLabel{
		ODataType: "#microsoft.graph.AssignedLabel",
	}
	return newAssignedLabel, nil
}

// AssignedLicense undocumented
type AssignedLicense struct {
	// Object is the base model of AssignedLicense
	Object

	ODataType string `json:"@odata.type"`
	// DisabledPlans undocumented
	DisabledPlans []UUID `json:"disabledPlans,omitempty"`
	// SKUID undocumented
	SKUID *UUID `json:"skuId,omitempty"`
}

func NewAssignedLicense() (*AssignedLicense, error) {
	newAssignedLicense := &AssignedLicense{
		ODataType: "#microsoft.graph.AssignedLicense",
	}
	return newAssignedLicense, nil
}

// AssignedPlan undocumented
type AssignedPlan struct {
	// Object is the base model of AssignedPlan
	Object

	ODataType string `json:"@odata.type"`
	// AssignedDateTime undocumented
	AssignedDateTime *time.Time `json:"assignedDateTime,omitempty"`
	// CapabilityStatus undocumented
	CapabilityStatus *string `json:"capabilityStatus,omitempty"`
	// Service undocumented
	Service *string `json:"service,omitempty"`
	// ServicePlanID undocumented
	ServicePlanID *UUID `json:"servicePlanId,omitempty"`
}

func NewAssignedPlan() (*AssignedPlan, error) {
	newAssignedPlan := &AssignedPlan{
		ODataType: "#microsoft.graph.AssignedPlan",
	}
	return newAssignedPlan, nil
}

// AssignedTrainingInfo undocumented
type AssignedTrainingInfo struct {
	// Object is the base model of AssignedTrainingInfo
	Object

	ODataType string `json:"@odata.type"`
	// AssignedUserCount undocumented
	AssignedUserCount *int `json:"assignedUserCount,omitempty"`
	// CompletedUserCount undocumented
	CompletedUserCount *int `json:"completedUserCount,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
}

func NewAssignedTrainingInfo() (*AssignedTrainingInfo, error) {
	newAssignedTrainingInfo := &AssignedTrainingInfo{
		ODataType: "#microsoft.graph.AssignedTrainingInfo",
	}
	return newAssignedTrainingInfo, nil
}
