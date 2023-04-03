// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// MicrosoftAccountUserConversationMember undocumented
type MicrosoftAccountUserConversationMember struct {
	// ConversationMember is the base model of MicrosoftAccountUserConversationMember
	ConversationMember

	ODataType string `json:"@odata.type"`
	// UserID undocumented
	UserID *string `json:"userId,omitempty"`
}

func NewMicrosoftAccountUserConversationMember() (*MicrosoftAccountUserConversationMember, error) {
	newMicrosoftAccountUserConversationMember := &MicrosoftAccountUserConversationMember{
		ODataType: "#microsoft.graph.MicrosoftAccountUserConversationMember",
	}
	return newMicrosoftAccountUserConversationMember, nil
}

// MicrosoftAuthenticatorAuthenticationMethod undocumented
type MicrosoftAuthenticatorAuthenticationMethod struct {
	// AuthenticationMethod is the base model of MicrosoftAuthenticatorAuthenticationMethod
	AuthenticationMethod

	ODataType string `json:"@odata.type"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DeviceTag undocumented
	DeviceTag *string `json:"deviceTag,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// PhoneAppVersion undocumented
	PhoneAppVersion *string `json:"phoneAppVersion,omitempty"`
	// Device undocumented
	Device *Device `json:"device,omitempty"`
}

func NewMicrosoftAuthenticatorAuthenticationMethod() (*MicrosoftAuthenticatorAuthenticationMethod, error) {
	newMicrosoftAuthenticatorAuthenticationMethod := &MicrosoftAuthenticatorAuthenticationMethod{
		ODataType: "#microsoft.graph.MicrosoftAuthenticatorAuthenticationMethod",
	}
	return newMicrosoftAuthenticatorAuthenticationMethod, nil
}

// MicrosoftAuthenticatorAuthenticationMethodConfiguration undocumented
type MicrosoftAuthenticatorAuthenticationMethodConfiguration struct {
	// AuthenticationMethodConfiguration is the base model of MicrosoftAuthenticatorAuthenticationMethodConfiguration
	AuthenticationMethodConfiguration

	ODataType string `json:"@odata.type"`
	// FeatureSettings undocumented
	FeatureSettings *MicrosoftAuthenticatorFeatureSettings `json:"featureSettings,omitempty"`
	// IncludeTargets undocumented
	IncludeTargets []MicrosoftAuthenticatorAuthenticationMethodTarget `json:"includeTargets,omitempty"`
}

func NewMicrosoftAuthenticatorAuthenticationMethodConfiguration() (*MicrosoftAuthenticatorAuthenticationMethodConfiguration, error) {
	newMicrosoftAuthenticatorAuthenticationMethodConfiguration := &MicrosoftAuthenticatorAuthenticationMethodConfiguration{
		ODataType: "#microsoft.graph.MicrosoftAuthenticatorAuthenticationMethodConfiguration",
	}
	return newMicrosoftAuthenticatorAuthenticationMethodConfiguration, nil
}

// MicrosoftAuthenticatorAuthenticationMethodTarget undocumented
type MicrosoftAuthenticatorAuthenticationMethodTarget struct {
	// AuthenticationMethodTarget is the base model of MicrosoftAuthenticatorAuthenticationMethodTarget
	AuthenticationMethodTarget

	ODataType string `json:"@odata.type"`
	// AuthenticationMode undocumented
	AuthenticationMode *MicrosoftAuthenticatorAuthenticationMode `json:"authenticationMode,omitempty"`
}

func NewMicrosoftAuthenticatorAuthenticationMethodTarget() (*MicrosoftAuthenticatorAuthenticationMethodTarget, error) {
	newMicrosoftAuthenticatorAuthenticationMethodTarget := &MicrosoftAuthenticatorAuthenticationMethodTarget{
		ODataType: "#microsoft.graph.MicrosoftAuthenticatorAuthenticationMethodTarget",
	}
	return newMicrosoftAuthenticatorAuthenticationMethodTarget, nil
}

// MicrosoftAuthenticatorFeatureSettings undocumented
type MicrosoftAuthenticatorFeatureSettings struct {
	// Object is the base model of MicrosoftAuthenticatorFeatureSettings
	Object

	ODataType string `json:"@odata.type"`
	// DisplayAppInformationRequiredState undocumented
	DisplayAppInformationRequiredState *AuthenticationMethodFeatureConfiguration `json:"displayAppInformationRequiredState,omitempty"`
	// DisplayLocationInformationRequiredState undocumented
	DisplayLocationInformationRequiredState *AuthenticationMethodFeatureConfiguration `json:"displayLocationInformationRequiredState,omitempty"`
}

func NewMicrosoftAuthenticatorFeatureSettings() (*MicrosoftAuthenticatorFeatureSettings, error) {
	newMicrosoftAuthenticatorFeatureSettings := &MicrosoftAuthenticatorFeatureSettings{
		ODataType: "#microsoft.graph.MicrosoftAuthenticatorFeatureSettings",
	}
	return newMicrosoftAuthenticatorFeatureSettings, nil
}

// MicrosoftStoreForBusinessApp undocumented
type MicrosoftStoreForBusinessApp struct {
	// MobileApp is the base model of MicrosoftStoreForBusinessApp
	MobileApp

	ODataType string `json:"@odata.type"`
	// LicenseType undocumented
	LicenseType *MicrosoftStoreForBusinessLicenseType `json:"licenseType,omitempty"`
	// PackageIdentityName undocumented
	PackageIdentityName *string `json:"packageIdentityName,omitempty"`
	// ProductKey undocumented
	ProductKey *string `json:"productKey,omitempty"`
	// TotalLicenseCount undocumented
	TotalLicenseCount *int `json:"totalLicenseCount,omitempty"`
	// UsedLicenseCount undocumented
	UsedLicenseCount *int `json:"usedLicenseCount,omitempty"`
}

func NewMicrosoftStoreForBusinessApp() (*MicrosoftStoreForBusinessApp, error) {
	newMicrosoftStoreForBusinessApp := &MicrosoftStoreForBusinessApp{
		ODataType: "#microsoft.graph.MicrosoftStoreForBusinessApp",
	}
	return newMicrosoftStoreForBusinessApp, nil
}

// MicrosoftStoreForBusinessAppAssignmentSettings undocumented
type MicrosoftStoreForBusinessAppAssignmentSettings struct {
	// MobileAppAssignmentSettings is the base model of MicrosoftStoreForBusinessAppAssignmentSettings
	MobileAppAssignmentSettings

	ODataType string `json:"@odata.type"`
	// UseDeviceContext undocumented
	UseDeviceContext *bool `json:"useDeviceContext,omitempty"`
}

func NewMicrosoftStoreForBusinessAppAssignmentSettings() (*MicrosoftStoreForBusinessAppAssignmentSettings, error) {
	newMicrosoftStoreForBusinessAppAssignmentSettings := &MicrosoftStoreForBusinessAppAssignmentSettings{
		ODataType: "#microsoft.graph.MicrosoftStoreForBusinessAppAssignmentSettings",
	}
	return newMicrosoftStoreForBusinessAppAssignmentSettings, nil
}
