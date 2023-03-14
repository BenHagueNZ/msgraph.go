
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// PasswordAuthenticationMethod undocumented
type PasswordAuthenticationMethod struct {
    // AuthenticationMethod is the base model of PasswordAuthenticationMethod
    AuthenticationMethod
    // CreatedDateTime undocumented
    CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
    // Password undocumented
    Password *string `json:"password,omitempty"`
}

// PasswordCredential undocumented
type PasswordCredential struct {
    // Object is the base model of PasswordCredential
    Object
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

// PasswordCredentialConfiguration undocumented
type PasswordCredentialConfiguration struct {
    // Object is the base model of PasswordCredentialConfiguration
    Object
    // MaxLifetime undocumented
    MaxLifetime *Duration `json:"maxLifetime,omitempty"`
    // RestrictForAppsCreatedAfterDateTime undocumented
    RestrictForAppsCreatedAfterDateTime *time.Time `json:"restrictForAppsCreatedAfterDateTime,omitempty"`
    // RestrictionType undocumented
    RestrictionType *AppCredentialRestrictionType `json:"restrictionType,omitempty"`
}

// PasswordProfile undocumented
type PasswordProfile struct {
    // Object is the base model of PasswordProfile
    Object
    // ForceChangePasswordNextSignIn undocumented
    ForceChangePasswordNextSignIn *bool `json:"forceChangePasswordNextSignIn,omitempty"`
    // ForceChangePasswordNextSignInWithMFA undocumented
    ForceChangePasswordNextSignInWithMFA *bool `json:"forceChangePasswordNextSignInWithMfa,omitempty"`
    // Password undocumented
    Password *string `json:"password,omitempty"`
}

// PasswordResetResponse undocumented
type PasswordResetResponse struct {
    // Object is the base model of PasswordResetResponse
    Object
    // NewPassword undocumented
    NewPassword *string `json:"newPassword,omitempty"`
}

// PasswordSingleSignOnCredentialSet undocumented
type PasswordSingleSignOnCredentialSet struct {
    // Object is the base model of PasswordSingleSignOnCredentialSet
    Object
    // Credentials undocumented
    Credentials []Credential `json:"credentials,omitempty"`
    // ID undocumented
    ID *string `json:"id,omitempty"`
}

// PasswordSingleSignOnField undocumented
type PasswordSingleSignOnField struct {
    // Object is the base model of PasswordSingleSignOnField
    Object
    // CustomizedLabel undocumented
    CustomizedLabel *string `json:"customizedLabel,omitempty"`
    // DefaultLabel undocumented
    DefaultLabel *string `json:"defaultLabel,omitempty"`
    // FieldID undocumented
    FieldID *string `json:"fieldId,omitempty"`
    // Type undocumented
    Type *string `json:"type,omitempty"`
}

// PasswordSingleSignOnSettings undocumented
type PasswordSingleSignOnSettings struct {
    // Object is the base model of PasswordSingleSignOnSettings
    Object
    // Fields undocumented
    Fields []PasswordSingleSignOnField `json:"fields,omitempty"`
}

// PasswordValidationInformation undocumented
type PasswordValidationInformation struct {
    // Object is the base model of PasswordValidationInformation
    Object
    // IsValid undocumented
    IsValid *bool `json:"isValid,omitempty"`
    // ValidationResults undocumented
    ValidationResults []ValidationResult `json:"validationResults,omitempty"`
}
