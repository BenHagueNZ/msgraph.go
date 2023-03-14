
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// ChromeOSDeviceProperty undocumented
type ChromeOSDeviceProperty struct {
    // Object is the base model of ChromeOSDeviceProperty
    Object
    // Name undocumented
    Name *string `json:"name,omitempty"`
    // Updatable undocumented
    Updatable *bool `json:"updatable,omitempty"`
    // Value undocumented
    Value *string `json:"value,omitempty"`
    // ValueType undocumented
    ValueType *string `json:"valueType,omitempty"`
}

// ChromeOSOnboardingSettings undocumented
type ChromeOSOnboardingSettings struct {
    // Entity is the base model of ChromeOSOnboardingSettings
    Entity
    // LastDirectorySyncDateTime undocumented
    LastDirectorySyncDateTime *time.Time `json:"lastDirectorySyncDateTime,omitempty"`
    // LastModifiedDateTime undocumented
    LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
    // OnboardingStatus undocumented
    OnboardingStatus *OnboardingStatus `json:"onboardingStatus,omitempty"`
    // OwnerUserPrincipalName undocumented
    OwnerUserPrincipalName *string `json:"ownerUserPrincipalName,omitempty"`
}
