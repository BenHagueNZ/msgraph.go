// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// EmailAddress undocumented
type EmailAddress struct {
	// Object is the base model of EmailAddress
	Object
	// Address undocumented
	Address *string `json:"address,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
}

// EmailAuthenticationMethod undocumented
type EmailAuthenticationMethod struct {
	// AuthenticationMethod is the base model of EmailAuthenticationMethod
	AuthenticationMethod
	// EmailAddress undocumented
	EmailAddress *string `json:"emailAddress,omitempty"`
}

// EmailAuthenticationMethodConfiguration undocumented
type EmailAuthenticationMethodConfiguration struct {
	// AuthenticationMethodConfiguration is the base model of EmailAuthenticationMethodConfiguration
	AuthenticationMethodConfiguration
	// AllowExternalIDToUseEmailOtp undocumented
	AllowExternalIDToUseEmailOtp *ExternalEmailOtpState `json:"allowExternalIdToUseEmailOtp,omitempty"`
	// IncludeTargets undocumented
	IncludeTargets []AuthenticationMethodTarget `json:"includeTargets,omitempty"`
}

// EmailFileAssessmentRequestObject undocumented
type EmailFileAssessmentRequestObject struct {
	// ThreatAssessmentRequestObject is the base model of EmailFileAssessmentRequestObject
	ThreatAssessmentRequestObject
	// ContentData undocumented
	ContentData *string `json:"contentData,omitempty"`
	// DestinationRoutingReason undocumented
	DestinationRoutingReason *MailDestinationRoutingReason `json:"destinationRoutingReason,omitempty"`
	// RecipientEmail undocumented
	RecipientEmail *string `json:"recipientEmail,omitempty"`
}

// EmailIdentity undocumented
type EmailIdentity struct {
	// Identity is the base model of EmailIdentity
	Identity
	// Email undocumented
	Email *string `json:"email,omitempty"`
}
