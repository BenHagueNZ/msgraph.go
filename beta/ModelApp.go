
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// AppCatalogs undocumented
type AppCatalogs struct {
    // Object is the base model of AppCatalogs
    Object
    // TeamsApps undocumented
    TeamsApps []TeamsApp `json:"teamsApps,omitempty"`
}

// AppConfigurationSettingItem undocumented
type AppConfigurationSettingItem struct {
    // Object is the base model of AppConfigurationSettingItem
    Object
    // AppConfigKey undocumented
    AppConfigKey *string `json:"appConfigKey,omitempty"`
    // AppConfigKeyType undocumented
    AppConfigKeyType *MDMAppConfigKeyType `json:"appConfigKeyType,omitempty"`
    // AppConfigKeyValue undocumented
    AppConfigKeyValue *string `json:"appConfigKeyValue,omitempty"`
}

// AppConsentApprovalRoute undocumented
type AppConsentApprovalRoute struct {
    // Entity is the base model of AppConsentApprovalRoute
    Entity
    // AppConsentRequests undocumented
    AppConsentRequests []AppConsentRequest `json:"appConsentRequests,omitempty"`
}

// AppConsentRequestObject undocumented
type AppConsentRequestObject struct {
    // Entity is the base model of AppConsentRequestObject
    Entity
    // AppDisplayName undocumented
    AppDisplayName *string `json:"appDisplayName,omitempty"`
    // AppID undocumented
    AppID *string `json:"appId,omitempty"`
    // ConsentType undocumented
    ConsentType *string `json:"consentType,omitempty"`
    // PendingScopes undocumented
    PendingScopes []AppConsentRequestScope `json:"pendingScopes,omitempty"`
    // UserConsentRequests undocumented
    UserConsentRequests []UserConsentRequest `json:"userConsentRequests,omitempty"`
}

// AppConsentRequestScope undocumented
type AppConsentRequestScope struct {
    // Object is the base model of AppConsentRequestScope
    Object
    // DisplayName undocumented
    DisplayName *string `json:"displayName,omitempty"`
}

// AppHostedMediaConfig undocumented
type AppHostedMediaConfig struct {
    // MediaConfig is the base model of AppHostedMediaConfig
    MediaConfig
    // Blob undocumented
    Blob *string `json:"blob,omitempty"`
}

// AppIdentity undocumented
type AppIdentity struct {
    // Object is the base model of AppIdentity
    Object
    // AppID undocumented
    AppID *string `json:"appId,omitempty"`
    // DisplayName undocumented
    DisplayName *string `json:"displayName,omitempty"`
    // ServicePrincipalID undocumented
    ServicePrincipalID *string `json:"servicePrincipalId,omitempty"`
    // ServicePrincipalName undocumented
    ServicePrincipalName *string `json:"servicePrincipalName,omitempty"`
}

// AppListItem undocumented
type AppListItem struct {
    // Object is the base model of AppListItem
    Object
    // AppID undocumented
    AppID *string `json:"appId,omitempty"`
    // AppStoreURL undocumented
    AppStoreURL *string `json:"appStoreUrl,omitempty"`
    // Name undocumented
    Name *string `json:"name,omitempty"`
    // Publisher undocumented
    Publisher *string `json:"publisher,omitempty"`
}

// AppLogCollectionDownloadDetails undocumented
type AppLogCollectionDownloadDetails struct {
    // Object is the base model of AppLogCollectionDownloadDetails
    Object
    // AppLogDecryptionAlgorithm undocumented
    AppLogDecryptionAlgorithm *AppLogDecryptionAlgorithm `json:"appLogDecryptionAlgorithm,omitempty"`
    // DecryptionKey undocumented
    DecryptionKey *string `json:"decryptionKey,omitempty"`
    // DownloadURL undocumented
    DownloadURL *string `json:"downloadUrl,omitempty"`
}

// AppLogCollectionRequestObject undocumented
type AppLogCollectionRequestObject struct {
    // Entity is the base model of AppLogCollectionRequestObject
    Entity
    // CompletedDateTime undocumented
    CompletedDateTime *time.Time `json:"completedDateTime,omitempty"`
    // CustomLogFolders undocumented
    CustomLogFolders []string `json:"customLogFolders,omitempty"`
    // ErrorMessage undocumented
    ErrorMessage *string `json:"errorMessage,omitempty"`
    // Status undocumented
    Status *AppLogUploadState `json:"status,omitempty"`
}

// AppManagementConfiguration undocumented
type AppManagementConfiguration struct {
    // Object is the base model of AppManagementConfiguration
    Object
    // KeyCredentials undocumented
    KeyCredentials []KeyCredentialConfiguration `json:"keyCredentials,omitempty"`
    // PasswordCredentials undocumented
    PasswordCredentials []PasswordCredentialConfiguration `json:"passwordCredentials,omitempty"`
}

// AppManagementPolicy undocumented
type AppManagementPolicy struct {
    // PolicyBase is the base model of AppManagementPolicy
    PolicyBase
    // IsEnabled undocumented
    IsEnabled *bool `json:"isEnabled,omitempty"`
    // Restrictions undocumented
    Restrictions *AppManagementConfiguration `json:"restrictions,omitempty"`
    // AppliesTo undocumented
    AppliesTo []DirectoryObject `json:"appliesTo,omitempty"`
}

// AppMetadata undocumented
type AppMetadata struct {
    // Object is the base model of AppMetadata
    Object
    // Data undocumented
    Data []AppMetadataEntry `json:"data,omitempty"`
    // Version undocumented
    Version *int `json:"version,omitempty"`
}

// AppMetadataEntry undocumented
type AppMetadataEntry struct {
    // Object is the base model of AppMetadataEntry
    Object
    // Key undocumented
    Key *string `json:"key,omitempty"`
    // Value undocumented
    Value *Binary `json:"value,omitempty"`
}

// AppRole undocumented
type AppRole struct {
    // Object is the base model of AppRole
    Object
    // AllowedMemberTypes undocumented
    AllowedMemberTypes []string `json:"allowedMemberTypes,omitempty"`
    // Description undocumented
    Description *string `json:"description,omitempty"`
    // DisplayName undocumented
    DisplayName *string `json:"displayName,omitempty"`
    // ID undocumented
    ID *UUID `json:"id,omitempty"`
    // IsEnabled undocumented
    IsEnabled *bool `json:"isEnabled,omitempty"`
    // Origin undocumented
    Origin *string `json:"origin,omitempty"`
    // Value undocumented
    Value *string `json:"value,omitempty"`
}

// AppRoleAssignment undocumented
type AppRoleAssignment struct {
    // Entity is the base model of AppRoleAssignment
    Entity
    // AppRoleID undocumented
    AppRoleID *UUID `json:"appRoleId,omitempty"`
    // CreationTimestamp undocumented
    CreationTimestamp *time.Time `json:"creationTimestamp,omitempty"`
    // PrincipalDisplayName undocumented
    PrincipalDisplayName *string `json:"principalDisplayName,omitempty"`
    // PrincipalID undocumented
    PrincipalID *UUID `json:"principalId,omitempty"`
    // PrincipalType undocumented
    PrincipalType *string `json:"principalType,omitempty"`
    // ResourceDisplayName undocumented
    ResourceDisplayName *string `json:"resourceDisplayName,omitempty"`
    // ResourceID undocumented
    ResourceID *UUID `json:"resourceId,omitempty"`
}

// AppScope undocumented
type AppScope struct {
    // Entity is the base model of AppScope
    Entity
    // DisplayName undocumented
    DisplayName *string `json:"displayName,omitempty"`
    // Type undocumented
    Type *string `json:"type,omitempty"`
}

// AppVulnerabilityManagedDevice undocumented
type AppVulnerabilityManagedDevice struct {
    // Entity is the base model of AppVulnerabilityManagedDevice
    Entity
    // DisplayName undocumented
    DisplayName *string `json:"displayName,omitempty"`
    // LastSyncDateTime undocumented
    LastSyncDateTime *time.Time `json:"lastSyncDateTime,omitempty"`
    // ManagedDeviceID undocumented
    ManagedDeviceID *string `json:"managedDeviceId,omitempty"`
}

// AppVulnerabilityMobileApp undocumented
type AppVulnerabilityMobileApp struct {
    // Entity is the base model of AppVulnerabilityMobileApp
    Entity
    // CreatedDateTime undocumented
    CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
    // DisplayName undocumented
    DisplayName *string `json:"displayName,omitempty"`
    // LastModifiedDateTime undocumented
    LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
    // MobileAppID undocumented
    MobileAppID *string `json:"mobileAppId,omitempty"`
    // MobileAppType undocumented
    MobileAppType *string `json:"mobileAppType,omitempty"`
    // Version undocumented
    Version *string `json:"version,omitempty"`
}

// AppVulnerabilityTask undocumented
type AppVulnerabilityTask struct {
    // DeviceAppManagementTask is the base model of AppVulnerabilityTask
    DeviceAppManagementTask
    // AppName undocumented
    AppName *string `json:"appName,omitempty"`
    // AppPublisher undocumented
    AppPublisher *string `json:"appPublisher,omitempty"`
    // AppVersion undocumented
    AppVersion *string `json:"appVersion,omitempty"`
    // Insights undocumented
    Insights *string `json:"insights,omitempty"`
    // ManagedDeviceCount undocumented
    ManagedDeviceCount *int `json:"managedDeviceCount,omitempty"`
    // MitigationType undocumented
    MitigationType *AppVulnerabilityTaskMitigationType `json:"mitigationType,omitempty"`
    // MobileAppCount undocumented
    MobileAppCount *int `json:"mobileAppCount,omitempty"`
    // Remediation undocumented
    Remediation *string `json:"remediation,omitempty"`
    // ManagedDevices undocumented
    ManagedDevices []AppVulnerabilityManagedDevice `json:"managedDevices,omitempty"`
    // MobileApps undocumented
    MobileApps []AppVulnerabilityMobileApp `json:"mobileApps,omitempty"`
}
