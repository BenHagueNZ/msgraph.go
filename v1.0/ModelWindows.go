// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// WindowsAppXAppAssignmentSettings undocumented
type WindowsAppXAppAssignmentSettings struct {
	// MobileAppAssignmentSettings is the base model of WindowsAppXAppAssignmentSettings
	MobileAppAssignmentSettings
	// UseDeviceContext undocumented
	UseDeviceContext *bool `json:"useDeviceContext,omitempty"`
}

// WindowsAutopilotDeviceIdentity undocumented
type WindowsAutopilotDeviceIdentity struct {
	// Entity is the base model of WindowsAutopilotDeviceIdentity
	Entity
	// AddressableUserName undocumented
	AddressableUserName *string `json:"addressableUserName,omitempty"`
	// AzureActiveDirectoryDeviceID undocumented
	AzureActiveDirectoryDeviceID *string `json:"azureActiveDirectoryDeviceId,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// EnrollmentState undocumented
	EnrollmentState *EnrollmentState `json:"enrollmentState,omitempty"`
	// GroupTag undocumented
	GroupTag *string `json:"groupTag,omitempty"`
	// LastContactedDateTime undocumented
	LastContactedDateTime *time.Time `json:"lastContactedDateTime,omitempty"`
	// ManagedDeviceID undocumented
	ManagedDeviceID *string `json:"managedDeviceId,omitempty"`
	// Manufacturer undocumented
	Manufacturer *string `json:"manufacturer,omitempty"`
	// Model undocumented
	Model *string `json:"model,omitempty"`
	// ProductKey undocumented
	ProductKey *string `json:"productKey,omitempty"`
	// PurchaseOrderIdentifier undocumented
	PurchaseOrderIdentifier *string `json:"purchaseOrderIdentifier,omitempty"`
	// ResourceName undocumented
	ResourceName *string `json:"resourceName,omitempty"`
	// SerialNumber undocumented
	SerialNumber *string `json:"serialNumber,omitempty"`
	// SKUNumber undocumented
	SKUNumber *string `json:"skuNumber,omitempty"`
	// SystemFamily undocumented
	SystemFamily *string `json:"systemFamily,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}

// WindowsDefenderAdvancedThreatProtectionConfiguration undocumented
type WindowsDefenderAdvancedThreatProtectionConfiguration struct {
	// DeviceConfiguration is the base model of WindowsDefenderAdvancedThreatProtectionConfiguration
	DeviceConfiguration
	// AllowSampleSharing undocumented
	AllowSampleSharing *bool `json:"allowSampleSharing,omitempty"`
	// EnableExpeditedTelemetryReporting undocumented
	EnableExpeditedTelemetryReporting *bool `json:"enableExpeditedTelemetryReporting,omitempty"`
}

// WindowsDefenderScanActionResult undocumented
type WindowsDefenderScanActionResult struct {
	// DeviceActionResult is the base model of WindowsDefenderScanActionResult
	DeviceActionResult
	// ScanType undocumented
	ScanType *string `json:"scanType,omitempty"`
}

// WindowsDeviceADAccount undocumented
type WindowsDeviceADAccount struct {
	// WindowsDeviceAccount is the base model of WindowsDeviceADAccount
	WindowsDeviceAccount
	// DomainName undocumented
	DomainName *string `json:"domainName,omitempty"`
	// UserName undocumented
	UserName *string `json:"userName,omitempty"`
}

// WindowsDeviceAccount undocumented
type WindowsDeviceAccount struct {
	// Object is the base model of WindowsDeviceAccount
	Object
	// Password undocumented
	Password *string `json:"password,omitempty"`
}

// WindowsDeviceAzureADAccount undocumented
type WindowsDeviceAzureADAccount struct {
	// WindowsDeviceAccount is the base model of WindowsDeviceAzureADAccount
	WindowsDeviceAccount
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}

// WindowsFirewallNetworkProfile undocumented
type WindowsFirewallNetworkProfile struct {
	// Object is the base model of WindowsFirewallNetworkProfile
	Object
	// AuthorizedApplicationRulesFromGroupPolicyMerged undocumented
	AuthorizedApplicationRulesFromGroupPolicyMerged *bool `json:"authorizedApplicationRulesFromGroupPolicyMerged,omitempty"`
	// ConnectionSecurityRulesFromGroupPolicyMerged undocumented
	ConnectionSecurityRulesFromGroupPolicyMerged *bool `json:"connectionSecurityRulesFromGroupPolicyMerged,omitempty"`
	// FirewallEnabled undocumented
	FirewallEnabled *StateManagementSetting `json:"firewallEnabled,omitempty"`
	// GlobalPortRulesFromGroupPolicyMerged undocumented
	GlobalPortRulesFromGroupPolicyMerged *bool `json:"globalPortRulesFromGroupPolicyMerged,omitempty"`
	// InboundConnectionsBlocked undocumented
	InboundConnectionsBlocked *bool `json:"inboundConnectionsBlocked,omitempty"`
	// InboundNotificationsBlocked undocumented
	InboundNotificationsBlocked *bool `json:"inboundNotificationsBlocked,omitempty"`
	// IncomingTrafficBlocked undocumented
	IncomingTrafficBlocked *bool `json:"incomingTrafficBlocked,omitempty"`
	// OutboundConnectionsBlocked undocumented
	OutboundConnectionsBlocked *bool `json:"outboundConnectionsBlocked,omitempty"`
	// PolicyRulesFromGroupPolicyMerged undocumented
	PolicyRulesFromGroupPolicyMerged *bool `json:"policyRulesFromGroupPolicyMerged,omitempty"`
	// SecuredPacketExemptionAllowed undocumented
	SecuredPacketExemptionAllowed *bool `json:"securedPacketExemptionAllowed,omitempty"`
	// StealthModeBlocked undocumented
	StealthModeBlocked *bool `json:"stealthModeBlocked,omitempty"`
	// UnicastResponsesToMulticastBroadcastsBlocked undocumented
	UnicastResponsesToMulticastBroadcastsBlocked *bool `json:"unicastResponsesToMulticastBroadcastsBlocked,omitempty"`
}

// WindowsHelloForBusinessAuthenticationMethod undocumented
type WindowsHelloForBusinessAuthenticationMethod struct {
	// AuthenticationMethod is the base model of WindowsHelloForBusinessAuthenticationMethod
	AuthenticationMethod
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// KeyStrength undocumented
	KeyStrength *AuthenticationMethodKeyStrength `json:"keyStrength,omitempty"`
	// Device undocumented
	Device *Device `json:"device,omitempty"`
}

// WindowsInformationProtection undocumented
type WindowsInformationProtection struct {
	// ManagedAppPolicy is the base model of WindowsInformationProtection
	ManagedAppPolicy
	// AzureRightsManagementServicesAllowed undocumented
	AzureRightsManagementServicesAllowed *bool `json:"azureRightsManagementServicesAllowed,omitempty"`
	// DataRecoveryCertificate undocumented
	DataRecoveryCertificate *WindowsInformationProtectionDataRecoveryCertificate `json:"dataRecoveryCertificate,omitempty"`
	// EnforcementLevel undocumented
	EnforcementLevel *WindowsInformationProtectionEnforcementLevel `json:"enforcementLevel,omitempty"`
	// EnterpriseDomain undocumented
	EnterpriseDomain *string `json:"enterpriseDomain,omitempty"`
	// EnterpriseInternalProxyServers undocumented
	EnterpriseInternalProxyServers []WindowsInformationProtectionResourceCollection `json:"enterpriseInternalProxyServers,omitempty"`
	// EnterpriseIPRanges undocumented
	EnterpriseIPRanges []WindowsInformationProtectionIPRangeCollection `json:"enterpriseIPRanges,omitempty"`
	// EnterpriseIPRangesAreAuthoritative undocumented
	EnterpriseIPRangesAreAuthoritative *bool `json:"enterpriseIPRangesAreAuthoritative,omitempty"`
	// EnterpriseNetworkDomainNames undocumented
	EnterpriseNetworkDomainNames []WindowsInformationProtectionResourceCollection `json:"enterpriseNetworkDomainNames,omitempty"`
	// EnterpriseProtectedDomainNames undocumented
	EnterpriseProtectedDomainNames []WindowsInformationProtectionResourceCollection `json:"enterpriseProtectedDomainNames,omitempty"`
	// EnterpriseProxiedDomains undocumented
	EnterpriseProxiedDomains []WindowsInformationProtectionProxiedDomainCollection `json:"enterpriseProxiedDomains,omitempty"`
	// EnterpriseProxyServers undocumented
	EnterpriseProxyServers []WindowsInformationProtectionResourceCollection `json:"enterpriseProxyServers,omitempty"`
	// EnterpriseProxyServersAreAuthoritative undocumented
	EnterpriseProxyServersAreAuthoritative *bool `json:"enterpriseProxyServersAreAuthoritative,omitempty"`
	// ExemptApps undocumented
	ExemptApps []WindowsInformationProtectionApp `json:"exemptApps,omitempty"`
	// IconsVisible undocumented
	IconsVisible *bool `json:"iconsVisible,omitempty"`
	// IndexingEncryptedStoresOrItemsBlocked undocumented
	IndexingEncryptedStoresOrItemsBlocked *bool `json:"indexingEncryptedStoresOrItemsBlocked,omitempty"`
	// IsAssigned undocumented
	IsAssigned *bool `json:"isAssigned,omitempty"`
	// NeutralDomainResources undocumented
	NeutralDomainResources []WindowsInformationProtectionResourceCollection `json:"neutralDomainResources,omitempty"`
	// ProtectedApps undocumented
	ProtectedApps []WindowsInformationProtectionApp `json:"protectedApps,omitempty"`
	// ProtectionUnderLockConfigRequired undocumented
	ProtectionUnderLockConfigRequired *bool `json:"protectionUnderLockConfigRequired,omitempty"`
	// RevokeOnUnenrollDisabled undocumented
	RevokeOnUnenrollDisabled *bool `json:"revokeOnUnenrollDisabled,omitempty"`
	// RightsManagementServicesTemplateID undocumented
	RightsManagementServicesTemplateID *UUID `json:"rightsManagementServicesTemplateId,omitempty"`
	// SmbAutoEncryptedFileExtensions undocumented
	SmbAutoEncryptedFileExtensions []WindowsInformationProtectionResourceCollection `json:"smbAutoEncryptedFileExtensions,omitempty"`
	// Assignments undocumented
	Assignments []TargetedManagedAppPolicyAssignment `json:"assignments,omitempty"`
	// ExemptAppLockerFiles undocumented
	ExemptAppLockerFiles []WindowsInformationProtectionAppLockerFile `json:"exemptAppLockerFiles,omitempty"`
	// ProtectedAppLockerFiles undocumented
	ProtectedAppLockerFiles []WindowsInformationProtectionAppLockerFile `json:"protectedAppLockerFiles,omitempty"`
}

// WindowsInformationProtectionApp undocumented
type WindowsInformationProtectionApp struct {
	// Object is the base model of WindowsInformationProtectionApp
	Object
	// Denied undocumented
	Denied *bool `json:"denied,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ProductName undocumented
	ProductName *string `json:"productName,omitempty"`
	// PublisherName undocumented
	PublisherName *string `json:"publisherName,omitempty"`
}

// WindowsInformationProtectionAppLearningSummary undocumented
type WindowsInformationProtectionAppLearningSummary struct {
	// Entity is the base model of WindowsInformationProtectionAppLearningSummary
	Entity
	// ApplicationName undocumented
	ApplicationName *string `json:"applicationName,omitempty"`
	// ApplicationType undocumented
	ApplicationType *ApplicationType `json:"applicationType,omitempty"`
	// DeviceCount undocumented
	DeviceCount *int `json:"deviceCount,omitempty"`
}

// WindowsInformationProtectionAppLockerFile undocumented
type WindowsInformationProtectionAppLockerFile struct {
	// Entity is the base model of WindowsInformationProtectionAppLockerFile
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// File undocumented
	File *Binary `json:"file,omitempty"`
	// FileHash undocumented
	FileHash *string `json:"fileHash,omitempty"`
	// Version undocumented
	Version *string `json:"version,omitempty"`
}

// WindowsInformationProtectionDataRecoveryCertificate undocumented
type WindowsInformationProtectionDataRecoveryCertificate struct {
	// Object is the base model of WindowsInformationProtectionDataRecoveryCertificate
	Object
	// Certificate undocumented
	Certificate *Binary `json:"certificate,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// ExpirationDateTime undocumented
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// SubjectName undocumented
	SubjectName *string `json:"subjectName,omitempty"`
}

// WindowsInformationProtectionDesktopApp undocumented
type WindowsInformationProtectionDesktopApp struct {
	// WindowsInformationProtectionApp is the base model of WindowsInformationProtectionDesktopApp
	WindowsInformationProtectionApp
	// BinaryName undocumented
	BinaryName *string `json:"binaryName,omitempty"`
	// BinaryVersionHigh undocumented
	BinaryVersionHigh *string `json:"binaryVersionHigh,omitempty"`
	// BinaryVersionLow undocumented
	BinaryVersionLow *string `json:"binaryVersionLow,omitempty"`
}

// WindowsInformationProtectionIPRangeCollection undocumented
type WindowsInformationProtectionIPRangeCollection struct {
	// Object is the base model of WindowsInformationProtectionIPRangeCollection
	Object
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Ranges undocumented
	Ranges []IPRange `json:"ranges,omitempty"`
}

// WindowsInformationProtectionNetworkLearningSummary undocumented
type WindowsInformationProtectionNetworkLearningSummary struct {
	// Entity is the base model of WindowsInformationProtectionNetworkLearningSummary
	Entity
	// DeviceCount undocumented
	DeviceCount *int `json:"deviceCount,omitempty"`
	// URL undocumented
	URL *string `json:"url,omitempty"`
}

// WindowsInformationProtectionPolicy undocumented
type WindowsInformationProtectionPolicy struct {
	// WindowsInformationProtection is the base model of WindowsInformationProtectionPolicy
	WindowsInformationProtection
	// DaysWithoutContactBeforeUnenroll undocumented
	DaysWithoutContactBeforeUnenroll *int `json:"daysWithoutContactBeforeUnenroll,omitempty"`
	// MDMEnrollmentURL undocumented
	MDMEnrollmentURL *string `json:"mdmEnrollmentUrl,omitempty"`
	// MinutesOfInactivityBeforeDeviceLock undocumented
	MinutesOfInactivityBeforeDeviceLock *int `json:"minutesOfInactivityBeforeDeviceLock,omitempty"`
	// NumberOfPastPinsRemembered undocumented
	NumberOfPastPinsRemembered *int `json:"numberOfPastPinsRemembered,omitempty"`
	// PasswordMaximumAttemptCount undocumented
	PasswordMaximumAttemptCount *int `json:"passwordMaximumAttemptCount,omitempty"`
	// PinExpirationDays undocumented
	PinExpirationDays *int `json:"pinExpirationDays,omitempty"`
	// PinLowercaseLetters undocumented
	PinLowercaseLetters *WindowsInformationProtectionPinCharacterRequirements `json:"pinLowercaseLetters,omitempty"`
	// PinMinimumLength undocumented
	PinMinimumLength *int `json:"pinMinimumLength,omitempty"`
	// PinSpecialCharacters undocumented
	PinSpecialCharacters *WindowsInformationProtectionPinCharacterRequirements `json:"pinSpecialCharacters,omitempty"`
	// PinUppercaseLetters undocumented
	PinUppercaseLetters *WindowsInformationProtectionPinCharacterRequirements `json:"pinUppercaseLetters,omitempty"`
	// RevokeOnMDMHandoffDisabled undocumented
	RevokeOnMDMHandoffDisabled *bool `json:"revokeOnMdmHandoffDisabled,omitempty"`
	// WindowsHelloForBusinessBlocked undocumented
	WindowsHelloForBusinessBlocked *bool `json:"windowsHelloForBusinessBlocked,omitempty"`
}

// WindowsInformationProtectionProxiedDomainCollection undocumented
type WindowsInformationProtectionProxiedDomainCollection struct {
	// Object is the base model of WindowsInformationProtectionProxiedDomainCollection
	Object
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ProxiedDomains undocumented
	ProxiedDomains []ProxiedDomain `json:"proxiedDomains,omitempty"`
}

// WindowsInformationProtectionResourceCollection undocumented
type WindowsInformationProtectionResourceCollection struct {
	// Object is the base model of WindowsInformationProtectionResourceCollection
	Object
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Resources undocumented
	Resources []string `json:"resources,omitempty"`
}

// WindowsInformationProtectionStoreApp undocumented
type WindowsInformationProtectionStoreApp struct {
	// WindowsInformationProtectionApp is the base model of WindowsInformationProtectionStoreApp
	WindowsInformationProtectionApp
}

// WindowsMicrosoftEdgeApp undocumented
type WindowsMicrosoftEdgeApp struct {
	// MobileApp is the base model of WindowsMicrosoftEdgeApp
	MobileApp
	// Channel undocumented
	Channel *MicrosoftEdgeChannel `json:"channel,omitempty"`
	// DisplayLanguageLocale undocumented
	DisplayLanguageLocale *string `json:"displayLanguageLocale,omitempty"`
}

// WindowsMinimumOperatingSystem undocumented
type WindowsMinimumOperatingSystem struct {
	// Object is the base model of WindowsMinimumOperatingSystem
	Object
	// V10_0 undocumented
	V10_0 *bool `json:"v10_0,omitempty"`
	// V8_0 undocumented
	V8_0 *bool `json:"v8_0,omitempty"`
	// V8_1 undocumented
	V8_1 *bool `json:"v8_1,omitempty"`
}

// WindowsMobileMSI undocumented
type WindowsMobileMSI struct {
	// MobileLobApp is the base model of WindowsMobileMSI
	MobileLobApp
	// CommandLine undocumented
	CommandLine *string `json:"commandLine,omitempty"`
	// IgnoreVersionDetection undocumented
	IgnoreVersionDetection *bool `json:"ignoreVersionDetection,omitempty"`
	// ProductCode undocumented
	ProductCode *string `json:"productCode,omitempty"`
	// ProductVersion undocumented
	ProductVersion *string `json:"productVersion,omitempty"`
}

// WindowsPhone81CompliancePolicy undocumented
type WindowsPhone81CompliancePolicy struct {
	// DeviceCompliancePolicy is the base model of WindowsPhone81CompliancePolicy
	DeviceCompliancePolicy
	// OsMaximumVersion undocumented
	OsMaximumVersion *string `json:"osMaximumVersion,omitempty"`
	// OsMinimumVersion undocumented
	OsMinimumVersion *string `json:"osMinimumVersion,omitempty"`
	// PasswordBlockSimple undocumented
	PasswordBlockSimple *bool `json:"passwordBlockSimple,omitempty"`
	// PasswordExpirationDays undocumented
	PasswordExpirationDays *int `json:"passwordExpirationDays,omitempty"`
	// PasswordMinimumCharacterSetCount undocumented
	PasswordMinimumCharacterSetCount *int `json:"passwordMinimumCharacterSetCount,omitempty"`
	// PasswordMinimumLength undocumented
	PasswordMinimumLength *int `json:"passwordMinimumLength,omitempty"`
	// PasswordMinutesOfInactivityBeforeLock undocumented
	PasswordMinutesOfInactivityBeforeLock *int `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`
	// PasswordPreviousPasswordBlockCount undocumented
	PasswordPreviousPasswordBlockCount *int `json:"passwordPreviousPasswordBlockCount,omitempty"`
	// PasswordRequired undocumented
	PasswordRequired *bool `json:"passwordRequired,omitempty"`
	// PasswordRequiredType undocumented
	PasswordRequiredType *RequiredPasswordType `json:"passwordRequiredType,omitempty"`
	// StorageRequireEncryption undocumented
	StorageRequireEncryption *bool `json:"storageRequireEncryption,omitempty"`
}

// WindowsPhone81CustomConfiguration undocumented
type WindowsPhone81CustomConfiguration struct {
	// DeviceConfiguration is the base model of WindowsPhone81CustomConfiguration
	DeviceConfiguration
	// OMASettings undocumented
	OMASettings []OMASetting `json:"omaSettings,omitempty"`
}

// WindowsPhone81GeneralConfiguration undocumented
type WindowsPhone81GeneralConfiguration struct {
	// DeviceConfiguration is the base model of WindowsPhone81GeneralConfiguration
	DeviceConfiguration
	// ApplyOnlyToWindowsPhone81 undocumented
	ApplyOnlyToWindowsPhone81 *bool `json:"applyOnlyToWindowsPhone81,omitempty"`
	// AppsBlockCopyPaste undocumented
	AppsBlockCopyPaste *bool `json:"appsBlockCopyPaste,omitempty"`
	// BluetoothBlocked undocumented
	BluetoothBlocked *bool `json:"bluetoothBlocked,omitempty"`
	// CameraBlocked undocumented
	CameraBlocked *bool `json:"cameraBlocked,omitempty"`
	// CellularBlockWiFiTethering undocumented
	CellularBlockWiFiTethering *bool `json:"cellularBlockWifiTethering,omitempty"`
	// CompliantAppListType undocumented
	CompliantAppListType *AppListType `json:"compliantAppListType,omitempty"`
	// CompliantAppsList undocumented
	CompliantAppsList []AppListItem `json:"compliantAppsList,omitempty"`
	// DiagnosticDataBlockSubmission undocumented
	DiagnosticDataBlockSubmission *bool `json:"diagnosticDataBlockSubmission,omitempty"`
	// EmailBlockAddingAccounts undocumented
	EmailBlockAddingAccounts *bool `json:"emailBlockAddingAccounts,omitempty"`
	// LocationServicesBlocked undocumented
	LocationServicesBlocked *bool `json:"locationServicesBlocked,omitempty"`
	// MicrosoftAccountBlocked undocumented
	MicrosoftAccountBlocked *bool `json:"microsoftAccountBlocked,omitempty"`
	// NfcBlocked undocumented
	NfcBlocked *bool `json:"nfcBlocked,omitempty"`
	// PasswordBlockSimple undocumented
	PasswordBlockSimple *bool `json:"passwordBlockSimple,omitempty"`
	// PasswordExpirationDays undocumented
	PasswordExpirationDays *int `json:"passwordExpirationDays,omitempty"`
	// PasswordMinimumCharacterSetCount undocumented
	PasswordMinimumCharacterSetCount *int `json:"passwordMinimumCharacterSetCount,omitempty"`
	// PasswordMinimumLength undocumented
	PasswordMinimumLength *int `json:"passwordMinimumLength,omitempty"`
	// PasswordMinutesOfInactivityBeforeScreenTimeout undocumented
	PasswordMinutesOfInactivityBeforeScreenTimeout *int `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
	// PasswordPreviousPasswordBlockCount undocumented
	PasswordPreviousPasswordBlockCount *int `json:"passwordPreviousPasswordBlockCount,omitempty"`
	// PasswordRequired undocumented
	PasswordRequired *bool `json:"passwordRequired,omitempty"`
	// PasswordRequiredType undocumented
	PasswordRequiredType *RequiredPasswordType `json:"passwordRequiredType,omitempty"`
	// PasswordSignInFailureCountBeforeFactoryReset undocumented
	PasswordSignInFailureCountBeforeFactoryReset *int `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`
	// ScreenCaptureBlocked undocumented
	ScreenCaptureBlocked *bool `json:"screenCaptureBlocked,omitempty"`
	// StorageBlockRemovableStorage undocumented
	StorageBlockRemovableStorage *bool `json:"storageBlockRemovableStorage,omitempty"`
	// StorageRequireEncryption undocumented
	StorageRequireEncryption *bool `json:"storageRequireEncryption,omitempty"`
	// WebBrowserBlocked undocumented
	WebBrowserBlocked *bool `json:"webBrowserBlocked,omitempty"`
	// WiFiBlockAutomaticConnectHotspots undocumented
	WiFiBlockAutomaticConnectHotspots *bool `json:"wifiBlockAutomaticConnectHotspots,omitempty"`
	// WiFiBlocked undocumented
	WiFiBlocked *bool `json:"wifiBlocked,omitempty"`
	// WiFiBlockHotspotReporting undocumented
	WiFiBlockHotspotReporting *bool `json:"wifiBlockHotspotReporting,omitempty"`
	// WindowsStoreBlocked undocumented
	WindowsStoreBlocked *bool `json:"windowsStoreBlocked,omitempty"`
}

// WindowsUniversalAppX undocumented
type WindowsUniversalAppX struct {
	// MobileLobApp is the base model of WindowsUniversalAppX
	MobileLobApp
	// ApplicableArchitectures undocumented
	ApplicableArchitectures *WindowsArchitecture `json:"applicableArchitectures,omitempty"`
	// ApplicableDeviceTypes undocumented
	ApplicableDeviceTypes *WindowsDeviceType `json:"applicableDeviceTypes,omitempty"`
	// IdentityName undocumented
	IdentityName *string `json:"identityName,omitempty"`
	// IdentityPublisherHash undocumented
	IdentityPublisherHash *string `json:"identityPublisherHash,omitempty"`
	// IdentityResourceIdentifier undocumented
	IdentityResourceIdentifier *string `json:"identityResourceIdentifier,omitempty"`
	// IdentityVersion undocumented
	IdentityVersion *string `json:"identityVersion,omitempty"`
	// IsBundle undocumented
	IsBundle *bool `json:"isBundle,omitempty"`
	// MinimumSupportedOperatingSystem undocumented
	MinimumSupportedOperatingSystem *WindowsMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`
	// CommittedContainedApps undocumented
	CommittedContainedApps []MobileContainedApp `json:"committedContainedApps,omitempty"`
}

// WindowsUniversalAppXAppAssignmentSettings undocumented
type WindowsUniversalAppXAppAssignmentSettings struct {
	// MobileAppAssignmentSettings is the base model of WindowsUniversalAppXAppAssignmentSettings
	MobileAppAssignmentSettings
	// UseDeviceContext undocumented
	UseDeviceContext *bool `json:"useDeviceContext,omitempty"`
}

// WindowsUniversalAppXContainedApp undocumented
type WindowsUniversalAppXContainedApp struct {
	// MobileContainedApp is the base model of WindowsUniversalAppXContainedApp
	MobileContainedApp
	// AppUserModelID undocumented
	AppUserModelID *string `json:"appUserModelId,omitempty"`
}

// WindowsUpdateActiveHoursInstall undocumented
type WindowsUpdateActiveHoursInstall struct {
	// WindowsUpdateInstallScheduleType is the base model of WindowsUpdateActiveHoursInstall
	WindowsUpdateInstallScheduleType
	// ActiveHoursEnd undocumented
	ActiveHoursEnd *TimeOfDay `json:"activeHoursEnd,omitempty"`
	// ActiveHoursStart undocumented
	ActiveHoursStart *TimeOfDay `json:"activeHoursStart,omitempty"`
}

// WindowsUpdateForBusinessConfiguration undocumented
type WindowsUpdateForBusinessConfiguration struct {
	// DeviceConfiguration is the base model of WindowsUpdateForBusinessConfiguration
	DeviceConfiguration
	// AllowWindows11Upgrade undocumented
	AllowWindows11Upgrade *bool `json:"allowWindows11Upgrade,omitempty"`
	// AutomaticUpdateMode undocumented
	AutomaticUpdateMode *AutomaticUpdateMode `json:"automaticUpdateMode,omitempty"`
	// AutoRestartNotificationDismissal undocumented
	AutoRestartNotificationDismissal *AutoRestartNotificationDismissalMethod `json:"autoRestartNotificationDismissal,omitempty"`
	// BusinessReadyUpdatesOnly undocumented
	BusinessReadyUpdatesOnly *WindowsUpdateType `json:"businessReadyUpdatesOnly,omitempty"`
	// DeadlineForFeatureUpdatesInDays undocumented
	DeadlineForFeatureUpdatesInDays *int `json:"deadlineForFeatureUpdatesInDays,omitempty"`
	// DeadlineForQualityUpdatesInDays undocumented
	DeadlineForQualityUpdatesInDays *int `json:"deadlineForQualityUpdatesInDays,omitempty"`
	// DeadlineGracePeriodInDays undocumented
	DeadlineGracePeriodInDays *int `json:"deadlineGracePeriodInDays,omitempty"`
	// DeliveryOptimizationMode undocumented
	DeliveryOptimizationMode *WindowsDeliveryOptimizationMode `json:"deliveryOptimizationMode,omitempty"`
	// DriversExcluded undocumented
	DriversExcluded *bool `json:"driversExcluded,omitempty"`
	// EngagedRestartDeadlineInDays undocumented
	EngagedRestartDeadlineInDays *int `json:"engagedRestartDeadlineInDays,omitempty"`
	// EngagedRestartSnoozeScheduleInDays undocumented
	EngagedRestartSnoozeScheduleInDays *int `json:"engagedRestartSnoozeScheduleInDays,omitempty"`
	// EngagedRestartTransitionScheduleInDays undocumented
	EngagedRestartTransitionScheduleInDays *int `json:"engagedRestartTransitionScheduleInDays,omitempty"`
	// FeatureUpdatesDeferralPeriodInDays undocumented
	FeatureUpdatesDeferralPeriodInDays *int `json:"featureUpdatesDeferralPeriodInDays,omitempty"`
	// FeatureUpdatesPaused undocumented
	FeatureUpdatesPaused *bool `json:"featureUpdatesPaused,omitempty"`
	// FeatureUpdatesPauseExpiryDateTime undocumented
	FeatureUpdatesPauseExpiryDateTime *time.Time `json:"featureUpdatesPauseExpiryDateTime,omitempty"`
	// FeatureUpdatesPauseStartDate undocumented
	FeatureUpdatesPauseStartDate *Date `json:"featureUpdatesPauseStartDate,omitempty"`
	// FeatureUpdatesRollbackStartDateTime undocumented
	FeatureUpdatesRollbackStartDateTime *time.Time `json:"featureUpdatesRollbackStartDateTime,omitempty"`
	// FeatureUpdatesRollbackWindowInDays undocumented
	FeatureUpdatesRollbackWindowInDays *int `json:"featureUpdatesRollbackWindowInDays,omitempty"`
	// FeatureUpdatesWillBeRolledBack undocumented
	FeatureUpdatesWillBeRolledBack *bool `json:"featureUpdatesWillBeRolledBack,omitempty"`
	// InstallationSchedule undocumented
	InstallationSchedule *WindowsUpdateInstallScheduleType `json:"installationSchedule,omitempty"`
	// MicrosoftUpdateServiceAllowed undocumented
	MicrosoftUpdateServiceAllowed *bool `json:"microsoftUpdateServiceAllowed,omitempty"`
	// PostponeRebootUntilAfterDeadline undocumented
	PostponeRebootUntilAfterDeadline *bool `json:"postponeRebootUntilAfterDeadline,omitempty"`
	// PrereleaseFeatures undocumented
	PrereleaseFeatures *PrereleaseFeatures `json:"prereleaseFeatures,omitempty"`
	// QualityUpdatesDeferralPeriodInDays undocumented
	QualityUpdatesDeferralPeriodInDays *int `json:"qualityUpdatesDeferralPeriodInDays,omitempty"`
	// QualityUpdatesPaused undocumented
	QualityUpdatesPaused *bool `json:"qualityUpdatesPaused,omitempty"`
	// QualityUpdatesPauseExpiryDateTime undocumented
	QualityUpdatesPauseExpiryDateTime *time.Time `json:"qualityUpdatesPauseExpiryDateTime,omitempty"`
	// QualityUpdatesPauseStartDate undocumented
	QualityUpdatesPauseStartDate *Date `json:"qualityUpdatesPauseStartDate,omitempty"`
	// QualityUpdatesRollbackStartDateTime undocumented
	QualityUpdatesRollbackStartDateTime *time.Time `json:"qualityUpdatesRollbackStartDateTime,omitempty"`
	// QualityUpdatesWillBeRolledBack undocumented
	QualityUpdatesWillBeRolledBack *bool `json:"qualityUpdatesWillBeRolledBack,omitempty"`
	// ScheduleImminentRestartWarningInMinutes undocumented
	ScheduleImminentRestartWarningInMinutes *int `json:"scheduleImminentRestartWarningInMinutes,omitempty"`
	// ScheduleRestartWarningInHours undocumented
	ScheduleRestartWarningInHours *int `json:"scheduleRestartWarningInHours,omitempty"`
	// SkipChecksBeforeRestart undocumented
	SkipChecksBeforeRestart *bool `json:"skipChecksBeforeRestart,omitempty"`
	// UpdateNotificationLevel undocumented
	UpdateNotificationLevel *WindowsUpdateNotificationDisplayOption `json:"updateNotificationLevel,omitempty"`
	// UpdateWeeks undocumented
	UpdateWeeks *WindowsUpdateForBusinessUpdateWeeks `json:"updateWeeks,omitempty"`
	// UserPauseAccess undocumented
	UserPauseAccess *Enablement `json:"userPauseAccess,omitempty"`
	// UserWindowsUpdateScanAccess undocumented
	UserWindowsUpdateScanAccess *Enablement `json:"userWindowsUpdateScanAccess,omitempty"`
}

// WindowsUpdateInstallScheduleType undocumented
type WindowsUpdateInstallScheduleType struct {
	// Object is the base model of WindowsUpdateInstallScheduleType
	Object
}

// WindowsUpdateScheduledInstall undocumented
type WindowsUpdateScheduledInstall struct {
	// WindowsUpdateInstallScheduleType is the base model of WindowsUpdateScheduledInstall
	WindowsUpdateInstallScheduleType
	// ScheduledInstallDay undocumented
	ScheduledInstallDay *WeeklySchedule `json:"scheduledInstallDay,omitempty"`
	// ScheduledInstallTime undocumented
	ScheduledInstallTime *TimeOfDay `json:"scheduledInstallTime,omitempty"`
}

// WindowsWebApp undocumented
type WindowsWebApp struct {
	// MobileApp is the base model of WindowsWebApp
	MobileApp
	// AppURL undocumented
	AppURL *string `json:"appUrl,omitempty"`
}
