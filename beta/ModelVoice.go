// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// VoiceAuthenticationMethodConfiguration undocumented
type VoiceAuthenticationMethodConfiguration struct {
	// AuthenticationMethodConfiguration is the base model of VoiceAuthenticationMethodConfiguration
	AuthenticationMethodConfiguration

	ODataType string `json:"@odata.type,omitempty"`
	// IsOfficePhoneAllowed undocumented
	IsOfficePhoneAllowed *bool `json:"isOfficePhoneAllowed,omitempty"`
	// IncludeTargets undocumented
	IncludeTargets []VoiceAuthenticationMethodTarget `json:"includeTargets,omitempty"`
}

func NewVoiceAuthenticationMethodConfiguration() (*VoiceAuthenticationMethodConfiguration, error) {
	newVoiceAuthenticationMethodConfiguration := &VoiceAuthenticationMethodConfiguration{
		ODataType: "#microsoft.graph.VoiceAuthenticationMethodConfiguration",
	}
	return newVoiceAuthenticationMethodConfiguration, nil
}

// VoiceAuthenticationMethodTarget undocumented
type VoiceAuthenticationMethodTarget struct {
	// AuthenticationMethodTarget is the base model of VoiceAuthenticationMethodTarget
	AuthenticationMethodTarget

	ODataType string `json:"@odata.type,omitempty"`
}

func NewVoiceAuthenticationMethodTarget() (*VoiceAuthenticationMethodTarget, error) {
	newVoiceAuthenticationMethodTarget := &VoiceAuthenticationMethodTarget{
		ODataType: "#microsoft.graph.VoiceAuthenticationMethodTarget",
	}
	return newVoiceAuthenticationMethodTarget, nil
}