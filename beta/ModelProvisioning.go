// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// ProvisioningErrorInfo undocumented
type ProvisioningErrorInfo struct {
	// Object is the base model of ProvisioningErrorInfo
	Object

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

func NewProvisioningErrorInfo() (*ProvisioningErrorInfo, error) {
	newProvisioningErrorInfo := &ProvisioningErrorInfo{
		ODataType: "#microsoft.graph.ProvisioningErrorInfo",
	}
	return newProvisioningErrorInfo, nil
}

// ProvisioningObjectSummary undocumented
type ProvisioningObjectSummary struct {
	// Entity is the base model of ProvisioningObjectSummary
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Action undocumented
	Action *string `json:"action,omitempty"`
	// ActivityDateTime undocumented
	ActivityDateTime *time.Time `json:"activityDateTime,omitempty"`
	// ChangeID undocumented
	ChangeID *string `json:"changeId,omitempty"`
	// CycleID undocumented
	CycleID *string `json:"cycleId,omitempty"`
	// DurationInMilliseconds undocumented
	DurationInMilliseconds *int `json:"durationInMilliseconds,omitempty"`
	// InitiatedBy undocumented
	InitiatedBy *Initiator `json:"initiatedBy,omitempty"`
	// JobID undocumented
	JobID *string `json:"jobId,omitempty"`
	// ModifiedProperties undocumented
	ModifiedProperties []ModifiedProperty `json:"modifiedProperties,omitempty"`
	// ProvisioningAction undocumented
	ProvisioningAction *ProvisioningAction `json:"provisioningAction,omitempty"`
	// ProvisioningStatusInfo undocumented
	ProvisioningStatusInfo *ProvisioningStatusInfo `json:"provisioningStatusInfo,omitempty"`
	// ProvisioningSteps undocumented
	ProvisioningSteps []ProvisioningStep `json:"provisioningSteps,omitempty"`
	// ServicePrincipal undocumented
	ServicePrincipal *ProvisioningServicePrincipal `json:"servicePrincipal,omitempty"`
	// SourceIdentity undocumented
	SourceIdentity *ProvisionedIdentity `json:"sourceIdentity,omitempty"`
	// SourceSystem undocumented
	SourceSystem *ProvisioningSystem `json:"sourceSystem,omitempty"`
	// StatusInfo undocumented
	StatusInfo *StatusBase `json:"statusInfo,omitempty"`
	// TargetIdentity undocumented
	TargetIdentity *ProvisionedIdentity `json:"targetIdentity,omitempty"`
	// TargetSystem undocumented
	TargetSystem *ProvisioningSystem `json:"targetSystem,omitempty"`
	// TenantID undocumented
	TenantID *string `json:"tenantId,omitempty"`
}

func NewProvisioningObjectSummary() (*ProvisioningObjectSummary, error) {
	newProvisioningObjectSummary := &ProvisioningObjectSummary{
		ODataType: "#microsoft.graph.ProvisioningObjectSummary",
	}
	return newProvisioningObjectSummary, nil
}

// ProvisioningServicePrincipal undocumented
type ProvisioningServicePrincipal struct {
	// Identity is the base model of ProvisioningServicePrincipal
	Identity

	ODataType string `json:"@odata.type,omitempty"`
}

func NewProvisioningServicePrincipal() (*ProvisioningServicePrincipal, error) {
	newProvisioningServicePrincipal := &ProvisioningServicePrincipal{
		ODataType: "#microsoft.graph.ProvisioningServicePrincipal",
	}
	return newProvisioningServicePrincipal, nil
}

// ProvisioningStatusInfo undocumented
type ProvisioningStatusInfo struct {
	// Object is the base model of ProvisioningStatusInfo
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ErrorInformation undocumented
	ErrorInformation *ProvisioningErrorInfo `json:"errorInformation,omitempty"`
	// Status undocumented
	Status *ProvisioningResult `json:"status,omitempty"`
}

func NewProvisioningStatusInfo() (*ProvisioningStatusInfo, error) {
	newProvisioningStatusInfo := &ProvisioningStatusInfo{
		ODataType: "#microsoft.graph.ProvisioningStatusInfo",
	}
	return newProvisioningStatusInfo, nil
}

// ProvisioningStep undocumented
type ProvisioningStep struct {
	// Object is the base model of ProvisioningStep
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// Details undocumented
	Details *DetailsInfo `json:"details,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// ProvisioningStepType undocumented
	ProvisioningStepType *ProvisioningStepType `json:"provisioningStepType,omitempty"`
	// Status undocumented
	Status *ProvisioningResult `json:"status,omitempty"`
}

func NewProvisioningStep() (*ProvisioningStep, error) {
	newProvisioningStep := &ProvisioningStep{
		ODataType: "#microsoft.graph.ProvisioningStep",
	}
	return newProvisioningStep, nil
}

// ProvisioningSystem undocumented
type ProvisioningSystem struct {
	// Identity is the base model of ProvisioningSystem
	Identity

	ODataType string `json:"@odata.type,omitempty"`
	// Details undocumented
	Details *DetailsInfo `json:"details,omitempty"`
}

func NewProvisioningSystem() (*ProvisioningSystem, error) {
	newProvisioningSystem := &ProvisioningSystem{
		ODataType: "#microsoft.graph.ProvisioningSystem",
	}
	return newProvisioningSystem, nil
}
