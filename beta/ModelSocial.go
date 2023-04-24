// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// SocialIdentityProvider undocumented
type SocialIdentityProvider struct {
	// IdentityProviderBase is the base model of SocialIdentityProvider
	IdentityProviderBase

	ODataType string `json:"@odata.type,omitempty"`
	// ClientID undocumented
	ClientID *string `json:"clientId,omitempty"`
	// ClientSecret undocumented
	ClientSecret *string `json:"clientSecret,omitempty"`
	// IdentityProviderType undocumented
	IdentityProviderType *string `json:"identityProviderType,omitempty"`
}

func NewSocialIdentityProvider() (*SocialIdentityProvider, error) {
	newSocialIdentityProvider := &SocialIdentityProvider{
		ODataType: "#microsoft.graph.SocialIdentityProvider",
	}
	return newSocialIdentityProvider, nil
}
