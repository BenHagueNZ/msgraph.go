// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// TermsAndConditions undocumented
type TermsAndConditions struct {
	// Entity is the base model of TermsAndConditions
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AcceptanceStatement undocumented
	AcceptanceStatement *string `json:"acceptanceStatement,omitempty"`
	// BodyText undocumented
	BodyText *string `json:"bodyText,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Title undocumented
	Title *string `json:"title,omitempty"`
	// Version undocumented
	Version *int `json:"version,omitempty"`
	// AcceptanceStatuses undocumented
	AcceptanceStatuses []TermsAndConditionsAcceptanceStatus `json:"acceptanceStatuses,omitempty"`
	// Assignments undocumented
	Assignments []TermsAndConditionsAssignment `json:"assignments,omitempty"`
}

func NewTermsAndConditions() (*TermsAndConditions, error) {
	newTermsAndConditions := &TermsAndConditions{
		ODataType: "#microsoft.graph.TermsAndConditions",
	}
	return newTermsAndConditions, nil
}

// TermsAndConditionsAcceptanceStatus undocumented
type TermsAndConditionsAcceptanceStatus struct {
	// Entity is the base model of TermsAndConditionsAcceptanceStatus
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AcceptedDateTime undocumented
	AcceptedDateTime *time.Time `json:"acceptedDateTime,omitempty"`
	// AcceptedVersion undocumented
	AcceptedVersion *int `json:"acceptedVersion,omitempty"`
	// UserDisplayName undocumented
	UserDisplayName *string `json:"userDisplayName,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// TermsAndConditions undocumented
	TermsAndConditions *TermsAndConditions `json:"termsAndConditions,omitempty"`
}

func NewTermsAndConditionsAcceptanceStatus() (*TermsAndConditionsAcceptanceStatus, error) {
	newTermsAndConditionsAcceptanceStatus := &TermsAndConditionsAcceptanceStatus{
		ODataType: "#microsoft.graph.TermsAndConditionsAcceptanceStatus",
	}
	return newTermsAndConditionsAcceptanceStatus, nil
}

// TermsAndConditionsAssignment undocumented
type TermsAndConditionsAssignment struct {
	// Entity is the base model of TermsAndConditionsAssignment
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Target undocumented
	Target *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
}

func NewTermsAndConditionsAssignment() (*TermsAndConditionsAssignment, error) {
	newTermsAndConditionsAssignment := &TermsAndConditionsAssignment{
		ODataType: "#microsoft.graph.TermsAndConditionsAssignment",
	}
	return newTermsAndConditionsAssignment, nil
}

// TermsExpiration undocumented
type TermsExpiration struct {
	// Object is the base model of TermsExpiration
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Frequency undocumented
	Frequency *Duration `json:"frequency,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
}

func NewTermsExpiration() (*TermsExpiration, error) {
	newTermsExpiration := &TermsExpiration{
		ODataType: "#microsoft.graph.TermsExpiration",
	}
	return newTermsExpiration, nil
}

// TermsOfUseContainer undocumented
type TermsOfUseContainer struct {
	// Entity is the base model of TermsOfUseContainer
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AgreementAcceptances undocumented
	AgreementAcceptances []AgreementAcceptance `json:"agreementAcceptances,omitempty"`
	// Agreements undocumented
	Agreements []Agreement `json:"agreements,omitempty"`
}

func NewTermsOfUseContainer() (*TermsOfUseContainer, error) {
	newTermsOfUseContainer := &TermsOfUseContainer{
		ODataType: "#microsoft.graph.TermsOfUseContainer",
	}
	return newTermsOfUseContainer, nil
}
