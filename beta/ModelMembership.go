// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// MembershipOutlierInsight undocumented
type MembershipOutlierInsight struct {
	// GovernanceInsight is the base model of MembershipOutlierInsight
	GovernanceInsight

	ODataType string `json:"@odata.type,omitempty"`
	// ContainerID undocumented
	ContainerID *string `json:"containerId,omitempty"`
	// MemberID undocumented
	MemberID *string `json:"memberId,omitempty"`
	// OutlierContainerType undocumented
	OutlierContainerType *OutlierContainerType `json:"outlierContainerType,omitempty"`
	// OutlierMemberType undocumented
	OutlierMemberType *OutlierMemberType `json:"outlierMemberType,omitempty"`
	// Container undocumented
	Container *DirectoryObject `json:"container,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *User `json:"lastModifiedBy,omitempty"`
	// Member undocumented
	Member *DirectoryObject `json:"member,omitempty"`
}

func NewMembershipOutlierInsight() (*MembershipOutlierInsight, error) {
	newMembershipOutlierInsight := &MembershipOutlierInsight{
		ODataType: "#microsoft.graph.MembershipOutlierInsight",
	}
	return newMembershipOutlierInsight, nil
}

// MembershipRuleEvaluationDetails undocumented
type MembershipRuleEvaluationDetails struct {
	// Object is the base model of MembershipRuleEvaluationDetails
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// MembershipRuleEvaluationDetails undocumented
	MembershipRuleEvaluationDetails *ExpressionEvaluationDetails `json:"membershipRuleEvaluationDetails,omitempty"`
}

func NewMembershipRuleEvaluationDetails() (*MembershipRuleEvaluationDetails, error) {
	newMembershipRuleEvaluationDetails := &MembershipRuleEvaluationDetails{
		ODataType: "#microsoft.graph.MembershipRuleEvaluationDetails",
	}
	return newMembershipRuleEvaluationDetails, nil
}

// MembershipRuleProcessingStatus undocumented
type MembershipRuleProcessingStatus struct {
	// Object is the base model of MembershipRuleProcessingStatus
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ErrorMessage undocumented
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// LastMembershipUpdated undocumented
	LastMembershipUpdated *time.Time `json:"lastMembershipUpdated,omitempty"`
	// Status undocumented
	Status *MembershipRuleProcessingStatusDetails `json:"status,omitempty"`
}

func NewMembershipRuleProcessingStatus() (*MembershipRuleProcessingStatus, error) {
	newMembershipRuleProcessingStatus := &MembershipRuleProcessingStatus{
		ODataType: "#microsoft.graph.MembershipRuleProcessingStatus",
	}
	return newMembershipRuleProcessingStatus, nil
}
