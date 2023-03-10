// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// TermsAndConditions A termsAndConditions entity represents the metadata and contents of a given Terms and Conditions (T&C) policy. T&C policies’ contents are presented to users upon their first attempt to enroll into Intune and subsequently upon edits where an administrator has required re-acceptance. They enable administrators to communicate the provisions to which a user must agree in order to have devices enrolled into Intune.
type TermsAndConditions struct {
	// Entity is the base model of TermsAndConditions
	Entity
	// AcceptanceStatement Administrator-supplied explanation of the terms and conditions, typically describing what it means to accept the terms and conditions set out in the T&C policy. This is shown to the user on prompts to accept the T&C policy.
	AcceptanceStatement *string `json:"acceptanceStatement,omitempty"`
	// BodyText Administrator-supplied body text of the terms and conditions, typically the terms themselves. This is shown to the user on prompts to accept the T&C policy.
	BodyText *string `json:"bodyText,omitempty"`
	// CreatedDateTime DateTime the object was created.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Description Administrator-supplied description of the T&C policy.
	Description *string `json:"description,omitempty"`
	// DisplayName Administrator-supplied name for the T&C policy.
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedDateTime DateTime the object was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Title Administrator-supplied title of the terms and conditions. This is shown to the user on prompts to accept the T&C policy.
	Title *string `json:"title,omitempty"`
	// Version Integer indicating the current version of the terms. Incremented when an administrator makes a change to the terms and wishes to require users to re-accept the modified T&C policy.
	Version *int `json:"version,omitempty"`
	// AcceptanceStatuses undocumented
	AcceptanceStatuses []TermsAndConditionsAcceptanceStatus `json:"acceptanceStatuses,omitempty"`
	// Assignments undocumented
	Assignments []TermsAndConditionsAssignment `json:"assignments,omitempty"`
}

// TermsAndConditionsAcceptanceStatus A termsAndConditionsAcceptanceStatus entity represents the acceptance status of a given Terms and Conditions (T&C) policy by a given user. Users must accept the most up-to-date version of the terms in order to retain access to the Company Portal.
type TermsAndConditionsAcceptanceStatus struct {
	// Entity is the base model of TermsAndConditionsAcceptanceStatus
	Entity
	// AcceptedDateTime DateTime when the terms were last accepted by the user.
	AcceptedDateTime *time.Time `json:"acceptedDateTime,omitempty"`
	// AcceptedVersion Most recent version number of the T&C accepted by the user.
	AcceptedVersion *int `json:"acceptedVersion,omitempty"`
	// UserDisplayName Display name of the user whose acceptance the entity represents.
	UserDisplayName *string `json:"userDisplayName,omitempty"`
	// UserPrincipalName The userPrincipalName of the User that accepted the term.
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// TermsAndConditions undocumented
	TermsAndConditions *TermsAndConditions `json:"termsAndConditions,omitempty"`
}

// TermsAndConditionsAssignment A termsAndConditionsAssignment entity represents the assignment of a given Terms and Conditions (T&C) policy to a given group. Users in the group will be required to accept the terms in order to have devices enrolled into Intune.
type TermsAndConditionsAssignment struct {
	// Entity is the base model of TermsAndConditionsAssignment
	Entity
	// Target Assignment target that the T&C policy is assigned to.
	Target *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
}

// TermsExpiration undocumented
type TermsExpiration struct {
	// Object is the base model of TermsExpiration
	Object
	// Frequency undocumented
	Frequency *Duration `json:"frequency,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
}

// TermsOfUseContainer undocumented
type TermsOfUseContainer struct {
	// Entity is the base model of TermsOfUseContainer
	Entity
	// AgreementAcceptances undocumented
	AgreementAcceptances []AgreementAcceptance `json:"agreementAcceptances,omitempty"`
	// Agreements undocumented
	Agreements []Agreement `json:"agreements,omitempty"`
}
