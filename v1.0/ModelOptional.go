// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// OptionalClaim undocumented
type OptionalClaim struct {
	// Object is the base model of OptionalClaim
	Object

	ODataType string `json:"@odata.type"`
	// AdditionalProperties undocumented
	AdditionalProperties []string `json:"additionalProperties,omitempty"`
	// Essential undocumented
	Essential *bool `json:"essential,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Source undocumented
	Source *string `json:"source,omitempty"`
}

func NewOptionalClaim() (*OptionalClaim, error) {
	newOptionalClaim := &OptionalClaim{
		ODataType: "#microsoft.graph.OptionalClaim",
	}
	return newOptionalClaim, nil
}

// OptionalClaims undocumented
type OptionalClaims struct {
	// Object is the base model of OptionalClaims
	Object

	ODataType string `json:"@odata.type"`
	// AccessToken undocumented
	AccessToken []OptionalClaim `json:"accessToken,omitempty"`
	// IDToken undocumented
	IDToken []OptionalClaim `json:"idToken,omitempty"`
	// Saml2Token undocumented
	Saml2Token []OptionalClaim `json:"saml2Token,omitempty"`
}

func NewOptionalClaims() (*OptionalClaims, error) {
	newOptionalClaims := &OptionalClaims{
		ODataType: "#microsoft.graph.OptionalClaims",
	}
	return newOptionalClaims, nil
}
