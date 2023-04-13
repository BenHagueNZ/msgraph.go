// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// TemporaryAccessPassAuthenticationMethod undocumented
type TemporaryAccessPassAuthenticationMethod struct {
	// AuthenticationMethod is the base model of TemporaryAccessPassAuthenticationMethod
	AuthenticationMethod

	ODataType string `json:"@odata.type,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// IsUsable undocumented
	IsUsable *bool `json:"isUsable,omitempty"`
	// IsUsableOnce undocumented
	IsUsableOnce *bool `json:"isUsableOnce,omitempty"`
	// LifetimeInMinutes undocumented
	LifetimeInMinutes *int `json:"lifetimeInMinutes,omitempty"`
	// MethodUsabilityReason undocumented
	MethodUsabilityReason *string `json:"methodUsabilityReason,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// TemporaryAccessPass undocumented
	TemporaryAccessPass *string `json:"temporaryAccessPass,omitempty"`
}

func NewTemporaryAccessPassAuthenticationMethod() (*TemporaryAccessPassAuthenticationMethod, error) {
	newTemporaryAccessPassAuthenticationMethod := &TemporaryAccessPassAuthenticationMethod{
		ODataType: "#microsoft.graph.TemporaryAccessPassAuthenticationMethod",
	}
	return newTemporaryAccessPassAuthenticationMethod, nil
}

// TemporaryAccessPassAuthenticationMethodConfiguration undocumented
type TemporaryAccessPassAuthenticationMethodConfiguration struct {
	// AuthenticationMethodConfiguration is the base model of TemporaryAccessPassAuthenticationMethodConfiguration
	AuthenticationMethodConfiguration

	ODataType string `json:"@odata.type,omitempty"`
	// DefaultLength undocumented
	DefaultLength *int `json:"defaultLength,omitempty"`
	// DefaultLifetimeInMinutes undocumented
	DefaultLifetimeInMinutes *int `json:"defaultLifetimeInMinutes,omitempty"`
	// IsUsableOnce undocumented
	IsUsableOnce *bool `json:"isUsableOnce,omitempty"`
	// MaximumLifetimeInMinutes undocumented
	MaximumLifetimeInMinutes *int `json:"maximumLifetimeInMinutes,omitempty"`
	// MinimumLifetimeInMinutes undocumented
	MinimumLifetimeInMinutes *int `json:"minimumLifetimeInMinutes,omitempty"`
	// IncludeTargets undocumented
	IncludeTargets []AuthenticationMethodTarget `json:"includeTargets,omitempty"`
}

func NewTemporaryAccessPassAuthenticationMethodConfiguration() (*TemporaryAccessPassAuthenticationMethodConfiguration, error) {
	newTemporaryAccessPassAuthenticationMethodConfiguration := &TemporaryAccessPassAuthenticationMethodConfiguration{
		ODataType: "#microsoft.graph.TemporaryAccessPassAuthenticationMethodConfiguration",
	}
	return newTemporaryAccessPassAuthenticationMethodConfiguration, nil
}
