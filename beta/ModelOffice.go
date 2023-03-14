
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// OfficeClientCheckinStatus undocumented
type OfficeClientCheckinStatus struct {
    // Object is the base model of OfficeClientCheckinStatus
    Object
    // AppliedPolicies undocumented
    AppliedPolicies []string `json:"appliedPolicies,omitempty"`
    // CheckinDateTime undocumented
    CheckinDateTime *time.Time `json:"checkinDateTime,omitempty"`
    // DeviceName undocumented
    DeviceName *string `json:"deviceName,omitempty"`
    // DevicePlatform undocumented
    DevicePlatform *string `json:"devicePlatform,omitempty"`
    // DevicePlatformVersion undocumented
    DevicePlatformVersion *string `json:"devicePlatformVersion,omitempty"`
    // ErrorMessage undocumented
    ErrorMessage *string `json:"errorMessage,omitempty"`
    // UserID undocumented
    UserID *string `json:"userId,omitempty"`
    // UserPrincipalName undocumented
    UserPrincipalName *string `json:"userPrincipalName,omitempty"`
    // WasSuccessful undocumented
    WasSuccessful *bool `json:"wasSuccessful,omitempty"`
}

// OfficeClientConfiguration undocumented
type OfficeClientConfiguration struct {
    // Entity is the base model of OfficeClientConfiguration
    Entity
    // CheckinStatuses undocumented
    CheckinStatuses []OfficeClientCheckinStatus `json:"checkinStatuses,omitempty"`
    // Description undocumented
    Description *string `json:"description,omitempty"`
    // DisplayName undocumented
    DisplayName *string `json:"displayName,omitempty"`
    // PolicyPayload undocumented
    PolicyPayload *Stream `json:"policyPayload,omitempty"`
    // Priority undocumented
    Priority *int `json:"priority,omitempty"`
    // UserCheckinSummary undocumented
    UserCheckinSummary *OfficeUserCheckinSummary `json:"userCheckinSummary,omitempty"`
    // UserPreferencePayload undocumented
    UserPreferencePayload *Stream `json:"userPreferencePayload,omitempty"`
    // Assignments undocumented
    Assignments []OfficeClientConfigurationAssignment `json:"assignments,omitempty"`
}

// OfficeClientConfigurationAssignment undocumented
type OfficeClientConfigurationAssignment struct {
    // Entity is the base model of OfficeClientConfigurationAssignment
    Entity
    // Target undocumented
    Target *OfficeConfigurationAssignmentTarget `json:"target,omitempty"`
}

// OfficeConfiguration undocumented
type OfficeConfiguration struct {
    // Object is the base model of OfficeConfiguration
    Object
    // TenantCheckinStatuses undocumented
    TenantCheckinStatuses []OfficeClientCheckinStatus `json:"tenantCheckinStatuses,omitempty"`
    // TenantUserCheckinSummary undocumented
    TenantUserCheckinSummary *OfficeUserCheckinSummary `json:"tenantUserCheckinSummary,omitempty"`
    // ClientConfigurations undocumented
    ClientConfigurations []OfficeClientConfiguration `json:"clientConfigurations,omitempty"`
}

// OfficeConfigurationAssignmentTarget undocumented
type OfficeConfigurationAssignmentTarget struct {
    // Object is the base model of OfficeConfigurationAssignmentTarget
    Object
}

// OfficeConfigurationGroupAssignmentTarget undocumented
type OfficeConfigurationGroupAssignmentTarget struct {
    // OfficeConfigurationAssignmentTarget is the base model of OfficeConfigurationGroupAssignmentTarget
    OfficeConfigurationAssignmentTarget
    // GroupID undocumented
    GroupID *string `json:"groupId,omitempty"`
}

// OfficeGraphInsights undocumented
type OfficeGraphInsights struct {
    // Entity is the base model of OfficeGraphInsights
    Entity
    // Shared undocumented
    Shared []SharedInsight `json:"shared,omitempty"`
    // Trending undocumented
    Trending []Trending `json:"trending,omitempty"`
    // Used undocumented
    Used []UsedInsight `json:"used,omitempty"`
}

// OfficeSuiteApp undocumented
type OfficeSuiteApp struct {
    // MobileApp is the base model of OfficeSuiteApp
    MobileApp
    // AutoAcceptEula undocumented
    AutoAcceptEula *bool `json:"autoAcceptEula,omitempty"`
    // ExcludedApps undocumented
    ExcludedApps *ExcludedApps `json:"excludedApps,omitempty"`
    // InstallProgressDisplayLevel undocumented
    InstallProgressDisplayLevel *OfficeSuiteInstallProgressDisplayLevel `json:"installProgressDisplayLevel,omitempty"`
    // LocalesToInstall undocumented
    LocalesToInstall []string `json:"localesToInstall,omitempty"`
    // OfficeConfigurationXML undocumented
    OfficeConfigurationXML *Binary `json:"officeConfigurationXml,omitempty"`
    // OfficePlatformArchitecture undocumented
    OfficePlatformArchitecture *WindowsArchitecture `json:"officePlatformArchitecture,omitempty"`
    // OfficeSuiteAppDefaultFileFormat undocumented
    OfficeSuiteAppDefaultFileFormat *OfficeSuiteDefaultFileFormatType `json:"officeSuiteAppDefaultFileFormat,omitempty"`
    // ProductIDs undocumented
    ProductIDs []OfficeProductID `json:"productIds,omitempty"`
    // ShouldUninstallOlderVersionsOfOffice undocumented
    ShouldUninstallOlderVersionsOfOffice *bool `json:"shouldUninstallOlderVersionsOfOffice,omitempty"`
    // TargetVersion undocumented
    TargetVersion *string `json:"targetVersion,omitempty"`
    // UpdateChannel undocumented
    UpdateChannel *OfficeUpdateChannel `json:"updateChannel,omitempty"`
    // UpdateVersion undocumented
    UpdateVersion *string `json:"updateVersion,omitempty"`
    // UseSharedComputerActivation undocumented
    UseSharedComputerActivation *bool `json:"useSharedComputerActivation,omitempty"`
}

// OfficeUserCheckinSummary undocumented
type OfficeUserCheckinSummary struct {
    // Object is the base model of OfficeUserCheckinSummary
    Object
    // FailedUserCount undocumented
    FailedUserCount *int `json:"failedUserCount,omitempty"`
    // SucceededUserCount undocumented
    SucceededUserCount *int `json:"succeededUserCount,omitempty"`
}
