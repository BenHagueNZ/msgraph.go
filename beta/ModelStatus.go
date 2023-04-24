// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// StatusBase undocumented
type StatusBase struct {
	// Object is the base model of StatusBase
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Status undocumented
	Status *ProvisioningResult `json:"status,omitempty"`
}

func NewStatusBase() (*StatusBase, error) {
	newStatusBase := &StatusBase{
		ODataType: "#microsoft.graph.StatusBase",
	}
	return newStatusBase, nil
}

// StatusDetails undocumented
type StatusDetails struct {
	// StatusBase is the base model of StatusDetails
	StatusBase

	ODataType string `json:"@odata.type,omitempty"`
	// AdditionalDetails undocumented
	AdditionalDetails *string `json:"additionalDetails,omitempty"`
	// ErrorCategory undocumented
	ErrorCategory *ProvisioningStatusErrorCategory `json:"errorCategory,omitempty"`
	// ErrorCode undocumented
	ErrorCode *string `json:"errorCode,omitempty"`
	// Reason undocumented
	Reason *string `json:"reason,omitempty"`
	// RecommendedAction undocumented
	RecommendedAction *string `json:"recommendedAction,omitempty"`
}

func NewStatusDetails() (*StatusDetails, error) {
	newStatusDetails := &StatusDetails{
		ODataType: "#microsoft.graph.StatusDetails",
	}
	return newStatusDetails, nil
}