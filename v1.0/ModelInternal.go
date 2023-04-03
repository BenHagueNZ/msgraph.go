// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// InternalDomainFederation undocumented
type InternalDomainFederation struct {
	// SamlOrWsFedProvider is the base model of InternalDomainFederation
	SamlOrWsFedProvider

	ODataType string `json:"@odata.type"`
	// ActiveSignInURI undocumented
	ActiveSignInURI *string `json:"activeSignInUri,omitempty"`
	// FederatedIdpMFABehavior undocumented
	FederatedIdpMFABehavior *FederatedIdpMFABehavior `json:"federatedIdpMfaBehavior,omitempty"`
	// IsSignedAuthenticationRequestRequired undocumented
	IsSignedAuthenticationRequestRequired *bool `json:"isSignedAuthenticationRequestRequired,omitempty"`
	// NextSigningCertificate undocumented
	NextSigningCertificate *string `json:"nextSigningCertificate,omitempty"`
	// PromptLoginBehavior undocumented
	PromptLoginBehavior *PromptLoginBehavior `json:"promptLoginBehavior,omitempty"`
	// SigningCertificateUpdateStatus undocumented
	SigningCertificateUpdateStatus *SigningCertificateUpdateStatus `json:"signingCertificateUpdateStatus,omitempty"`
	// SignOutURI undocumented
	SignOutURI *string `json:"signOutUri,omitempty"`
}

func NewInternalDomainFederation() (*InternalDomainFederation, error) {
	newInternalDomainFederation := &InternalDomainFederation{
		ODataType: "#microsoft.graph.InternalDomainFederation",
	}
	return newInternalDomainFederation, nil
}

// InternalSponsors undocumented
type InternalSponsors struct {
	// SubjectSet is the base model of InternalSponsors
	SubjectSet

	ODataType string `json:"@odata.type"`
}

func NewInternalSponsors() (*InternalSponsors, error) {
	newInternalSponsors := &InternalSponsors{
		ODataType: "#microsoft.graph.InternalSponsors",
	}
	return newInternalSponsors, nil
}
