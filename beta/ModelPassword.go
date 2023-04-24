// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// PasswordAuthenticationMethod undocumented
type PasswordAuthenticationMethod struct {
	// AuthenticationMethod is the base model of PasswordAuthenticationMethod
	AuthenticationMethod

	ODataType string `json:"@odata.type,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Password undocumented
	Password *string `json:"password,omitempty"`
}

func NewPasswordAuthenticationMethod() (*PasswordAuthenticationMethod, error) {
	newPasswordAuthenticationMethod := &PasswordAuthenticationMethod{
		ODataType: "#microsoft.graph.PasswordAuthenticationMethod",
	}
	return newPasswordAuthenticationMethod, nil
}

// PasswordCredential undocumented
type PasswordCredential struct {
	// Object is the base model of PasswordCredential
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// CustomKeyIdentifier undocumented
	CustomKeyIdentifier *Binary `json:"customKeyIdentifier,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// Hint undocumented
	Hint *string `json:"hint,omitempty"`
	// KeyID undocumented
	KeyID *UUID `json:"keyId,omitempty"`
	// SecretText undocumented
	SecretText *string `json:"secretText,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
}

func NewPasswordCredential() (*PasswordCredential, error) {
	newPasswordCredential := &PasswordCredential{
		ODataType: "#microsoft.graph.PasswordCredential",
	}
	return newPasswordCredential, nil
}

// PasswordCredentialConfiguration undocumented
type PasswordCredentialConfiguration struct {
	// Object is the base model of PasswordCredentialConfiguration
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// MaxLifetime undocumented
	MaxLifetime *Duration `json:"maxLifetime,omitempty"`
	// RestrictForAppsCreatedAfterDateTime undocumented
	RestrictForAppsCreatedAfterDateTime *time.Time `json:"restrictForAppsCreatedAfterDateTime,omitempty"`
	// RestrictionType undocumented
	RestrictionType *AppCredentialRestrictionType `json:"restrictionType,omitempty"`
}

func NewPasswordCredentialConfiguration() (*PasswordCredentialConfiguration, error) {
	newPasswordCredentialConfiguration := &PasswordCredentialConfiguration{
		ODataType: "#microsoft.graph.PasswordCredentialConfiguration",
	}
	return newPasswordCredentialConfiguration, nil
}

// PasswordProfile undocumented
type PasswordProfile struct {
	// Object is the base model of PasswordProfile
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ForceChangePasswordNextSignIn undocumented
	ForceChangePasswordNextSignIn *bool `json:"forceChangePasswordNextSignIn,omitempty"`
	// ForceChangePasswordNextSignInWithMFA undocumented
	ForceChangePasswordNextSignInWithMFA *bool `json:"forceChangePasswordNextSignInWithMfa,omitempty"`
	// Password undocumented
	Password *string `json:"password,omitempty"`
}

func NewPasswordProfile() (*PasswordProfile, error) {
	newPasswordProfile := &PasswordProfile{
		ODataType: "#microsoft.graph.PasswordProfile",
	}
	return newPasswordProfile, nil
}

// PasswordResetResponse undocumented
type PasswordResetResponse struct {
	// Object is the base model of PasswordResetResponse
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// NewPassword undocumented
	NewPassword *string `json:"newPassword,omitempty"`
}

func NewPasswordResetResponse() (*PasswordResetResponse, error) {
	newPasswordResetResponse := &PasswordResetResponse{
		ODataType: "#microsoft.graph.PasswordResetResponse",
	}
	return newPasswordResetResponse, nil
}

// PasswordSingleSignOnCredentialSet undocumented
type PasswordSingleSignOnCredentialSet struct {
	// Object is the base model of PasswordSingleSignOnCredentialSet
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Credentials undocumented
	Credentials []Credential `json:"credentials,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
}

func NewPasswordSingleSignOnCredentialSet() (*PasswordSingleSignOnCredentialSet, error) {
	newPasswordSingleSignOnCredentialSet := &PasswordSingleSignOnCredentialSet{
		ODataType: "#microsoft.graph.PasswordSingleSignOnCredentialSet",
	}
	return newPasswordSingleSignOnCredentialSet, nil
}

// PasswordSingleSignOnField undocumented
type PasswordSingleSignOnField struct {
	// Object is the base model of PasswordSingleSignOnField
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// CustomizedLabel undocumented
	CustomizedLabel *string `json:"customizedLabel,omitempty"`
	// DefaultLabel undocumented
	DefaultLabel *string `json:"defaultLabel,omitempty"`
	// FieldID undocumented
	FieldID *string `json:"fieldId,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
}

func NewPasswordSingleSignOnField() (*PasswordSingleSignOnField, error) {
	newPasswordSingleSignOnField := &PasswordSingleSignOnField{
		ODataType: "#microsoft.graph.PasswordSingleSignOnField",
	}
	return newPasswordSingleSignOnField, nil
}

// PasswordSingleSignOnSettings undocumented
type PasswordSingleSignOnSettings struct {
	// Object is the base model of PasswordSingleSignOnSettings
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Fields undocumented
	Fields []PasswordSingleSignOnField `json:"fields,omitempty"`
}

func NewPasswordSingleSignOnSettings() (*PasswordSingleSignOnSettings, error) {
	newPasswordSingleSignOnSettings := &PasswordSingleSignOnSettings{
		ODataType: "#microsoft.graph.PasswordSingleSignOnSettings",
	}
	return newPasswordSingleSignOnSettings, nil
}

// PasswordValidationInformation undocumented
type PasswordValidationInformation struct {
	// Object is the base model of PasswordValidationInformation
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// IsValid undocumented
	IsValid *bool `json:"isValid,omitempty"`
	// ValidationResults undocumented
	ValidationResults []ValidationResult `json:"validationResults,omitempty"`
}

func NewPasswordValidationInformation() (*PasswordValidationInformation, error) {
	newPasswordValidationInformation := &PasswordValidationInformation{
		ODataType: "#microsoft.graph.PasswordValidationInformation",
	}
	return newPasswordValidationInformation, nil
}