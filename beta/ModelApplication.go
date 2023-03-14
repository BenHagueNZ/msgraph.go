
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// Application undocumented
type Application struct {
    // DirectoryObject is the base model of Application
    DirectoryObject
    // API undocumented
    API *APIApplication `json:"api,omitempty"`
    // AppID undocumented
    AppID *string `json:"appId,omitempty"`
    // AppRoles undocumented
    AppRoles []AppRole `json:"appRoles,omitempty"`
    // Certification undocumented
    Certification *Certification `json:"certification,omitempty"`
    // CreatedDateTime undocumented
    CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
    // DefaultRedirectURI undocumented
    DefaultRedirectURI *string `json:"defaultRedirectUri,omitempty"`
    // Description undocumented
    Description *string `json:"description,omitempty"`
    // DisabledByMicrosoftStatus undocumented
    DisabledByMicrosoftStatus *string `json:"disabledByMicrosoftStatus,omitempty"`
    // DisplayName undocumented
    DisplayName *string `json:"displayName,omitempty"`
    // GroupMembershipClaims undocumented
    GroupMembershipClaims *string `json:"groupMembershipClaims,omitempty"`
    // IdentifierUris undocumented
    IdentifierUris []string `json:"identifierUris,omitempty"`
    // Info undocumented
    Info *InformationalURL `json:"info,omitempty"`
    // IsDeviceOnlyAuthSupported undocumented
    IsDeviceOnlyAuthSupported *bool `json:"isDeviceOnlyAuthSupported,omitempty"`
    // IsFallbackPublicClient undocumented
    IsFallbackPublicClient *bool `json:"isFallbackPublicClient,omitempty"`
    // KeyCredentials undocumented
    KeyCredentials []KeyCredential `json:"keyCredentials,omitempty"`
    // Logo undocumented
    Logo *Stream `json:"logo,omitempty"`
    // Notes undocumented
    Notes *string `json:"notes,omitempty"`
    // OptionalClaims undocumented
    OptionalClaims *OptionalClaims `json:"optionalClaims,omitempty"`
    // ParentalControlSettings undocumented
    ParentalControlSettings *ParentalControlSettings `json:"parentalControlSettings,omitempty"`
    // PasswordCredentials undocumented
    PasswordCredentials []PasswordCredential `json:"passwordCredentials,omitempty"`
    // PublicClient undocumented
    PublicClient *PublicClientApplication `json:"publicClient,omitempty"`
    // PublisherDomain undocumented
    PublisherDomain *string `json:"publisherDomain,omitempty"`
    // RequestSignatureVerification undocumented
    RequestSignatureVerification *RequestSignatureVerification `json:"requestSignatureVerification,omitempty"`
    // RequiredResourceAccess undocumented
    RequiredResourceAccess []RequiredResourceAccess `json:"requiredResourceAccess,omitempty"`
    // SamlMetadataURL undocumented
    SamlMetadataURL *string `json:"samlMetadataUrl,omitempty"`
    // ServiceManagementReference undocumented
    ServiceManagementReference *string `json:"serviceManagementReference,omitempty"`
    // ServicePrincipalLockConfiguration undocumented
    ServicePrincipalLockConfiguration *ServicePrincipalLockConfiguration `json:"servicePrincipalLockConfiguration,omitempty"`
    // SignInAudience undocumented
    SignInAudience *string `json:"signInAudience,omitempty"`
    // Spa undocumented
    Spa *SpaApplication `json:"spa,omitempty"`
    // Tags undocumented
    Tags []string `json:"tags,omitempty"`
    // TokenEncryptionKeyID undocumented
    TokenEncryptionKeyID *UUID `json:"tokenEncryptionKeyId,omitempty"`
    // UniqueName undocumented
    UniqueName *string `json:"uniqueName,omitempty"`
    // VerifiedPublisher undocumented
    VerifiedPublisher *VerifiedPublisher `json:"verifiedPublisher,omitempty"`
    // Web undocumented
    Web *WebApplication `json:"web,omitempty"`
    // Windows undocumented
    Windows *WindowsApplication `json:"windows,omitempty"`
    // OnPremisesPublishing undocumented
    OnPremisesPublishing *OnPremisesPublishing `json:"onPremisesPublishing,omitempty"`
    // AppManagementPolicies undocumented
    AppManagementPolicies []AppManagementPolicy `json:"appManagementPolicies,omitempty"`
    // CreatedOnBehalfOf undocumented
    CreatedOnBehalfOf *DirectoryObject `json:"createdOnBehalfOf,omitempty"`
    // ExtensionProperties undocumented
    ExtensionProperties []ExtensionProperty `json:"extensionProperties,omitempty"`
    // FederatedIdentityCredentials undocumented
    FederatedIdentityCredentials []FederatedIdentityCredential `json:"federatedIdentityCredentials,omitempty"`
    // HomeRealmDiscoveryPolicies undocumented
    HomeRealmDiscoveryPolicies []HomeRealmDiscoveryPolicy `json:"homeRealmDiscoveryPolicies,omitempty"`
    // Owners undocumented
    Owners []DirectoryObject `json:"owners,omitempty"`
    // TokenIssuancePolicies undocumented
    TokenIssuancePolicies []TokenIssuancePolicy `json:"tokenIssuancePolicies,omitempty"`
    // TokenLifetimePolicies undocumented
    TokenLifetimePolicies []TokenLifetimePolicy `json:"tokenLifetimePolicies,omitempty"`
    // ConnectorGroup undocumented
    ConnectorGroup *ConnectorGroup `json:"connectorGroup,omitempty"`
    // Synchronization undocumented
    Synchronization *Synchronization `json:"synchronization,omitempty"`
}

// ApplicationEnforcedRestrictionsSessionControl undocumented
type ApplicationEnforcedRestrictionsSessionControl struct {
    // ConditionalAccessSessionControl is the base model of ApplicationEnforcedRestrictionsSessionControl
    ConditionalAccessSessionControl
}

// ApplicationSegment undocumented
type ApplicationSegment struct {
    // Entity is the base model of ApplicationSegment
    Entity
}

// ApplicationServicePrincipal undocumented
type ApplicationServicePrincipal struct {
    // Object is the base model of ApplicationServicePrincipal
    Object
    // Application undocumented
    Application *Application `json:"application,omitempty"`
    // ServicePrincipal undocumented
    ServicePrincipal *ServicePrincipal `json:"servicePrincipal,omitempty"`
}

// ApplicationSignInDetailedSummary undocumented
type ApplicationSignInDetailedSummary struct {
    // Entity is the base model of ApplicationSignInDetailedSummary
    Entity
    // AggregatedEventDateTime undocumented
    AggregatedEventDateTime *time.Time `json:"aggregatedEventDateTime,omitempty"`
    // AppDisplayName undocumented
    AppDisplayName *string `json:"appDisplayName,omitempty"`
    // AppID undocumented
    AppID *string `json:"appId,omitempty"`
    // SignInCount undocumented
    SignInCount *int `json:"signInCount,omitempty"`
    // Status undocumented
    Status *SignInStatus `json:"status,omitempty"`
}

// ApplicationSignInSummary undocumented
type ApplicationSignInSummary struct {
    // Entity is the base model of ApplicationSignInSummary
    Entity
    // AppDisplayName undocumented
    AppDisplayName *string `json:"appDisplayName,omitempty"`
    // FailedSignInCount undocumented
    FailedSignInCount *int `json:"failedSignInCount,omitempty"`
    // SuccessfulSignInCount undocumented
    SuccessfulSignInCount *int `json:"successfulSignInCount,omitempty"`
    // SuccessPercentage undocumented
    SuccessPercentage *float64 `json:"successPercentage,omitempty"`
}

// ApplicationTemplate undocumented
type ApplicationTemplate struct {
    // Entity is the base model of ApplicationTemplate
    Entity
    // Categories undocumented
    Categories []string `json:"categories,omitempty"`
    // Description undocumented
    Description *string `json:"description,omitempty"`
    // DisplayName undocumented
    DisplayName *string `json:"displayName,omitempty"`
    // HomePageURL undocumented
    HomePageURL *string `json:"homePageUrl,omitempty"`
    // InformationalUrls undocumented
    InformationalUrls *InformationalUrls `json:"informationalUrls,omitempty"`
    // LogoURL undocumented
    LogoURL *string `json:"logoUrl,omitempty"`
    // Publisher undocumented
    Publisher *string `json:"publisher,omitempty"`
    // SupportedClaimConfiguration undocumented
    SupportedClaimConfiguration *SupportedClaimConfiguration `json:"supportedClaimConfiguration,omitempty"`
    // SupportedProvisioningTypes undocumented
    SupportedProvisioningTypes []string `json:"supportedProvisioningTypes,omitempty"`
    // SupportedSingleSignOnModes undocumented
    SupportedSingleSignOnModes []string `json:"supportedSingleSignOnModes,omitempty"`
}
