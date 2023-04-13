// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// AndroidCompliancePolicy undocumented
type AndroidCompliancePolicy struct {
	// DeviceCompliancePolicy is the base model of AndroidCompliancePolicy
	DeviceCompliancePolicy

	ODataType string `json:"@odata.type,omitempty"`
	// DeviceThreatProtectionEnabled undocumented
	DeviceThreatProtectionEnabled *bool `json:"deviceThreatProtectionEnabled,omitempty"`
	// DeviceThreatProtectionRequiredSecurityLevel undocumented
	DeviceThreatProtectionRequiredSecurityLevel *DeviceThreatProtectionLevel `json:"deviceThreatProtectionRequiredSecurityLevel,omitempty"`
	// MinAndroidSecurityPatchLevel undocumented
	MinAndroidSecurityPatchLevel *string `json:"minAndroidSecurityPatchLevel,omitempty"`
	// OsMaximumVersion undocumented
	OsMaximumVersion *string `json:"osMaximumVersion,omitempty"`
	// OsMinimumVersion undocumented
	OsMinimumVersion *string `json:"osMinimumVersion,omitempty"`
	// PasswordExpirationDays undocumented
	PasswordExpirationDays *int `json:"passwordExpirationDays,omitempty"`
	// PasswordMinimumLength undocumented
	PasswordMinimumLength *int `json:"passwordMinimumLength,omitempty"`
	// PasswordMinutesOfInactivityBeforeLock undocumented
	PasswordMinutesOfInactivityBeforeLock *int `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`
	// PasswordPreviousPasswordBlockCount undocumented
	PasswordPreviousPasswordBlockCount *int `json:"passwordPreviousPasswordBlockCount,omitempty"`
	// PasswordRequired undocumented
	PasswordRequired *bool `json:"passwordRequired,omitempty"`
	// PasswordRequiredType undocumented
	PasswordRequiredType *AndroidRequiredPasswordType `json:"passwordRequiredType,omitempty"`
	// SecurityBlockJailbrokenDevices undocumented
	SecurityBlockJailbrokenDevices *bool `json:"securityBlockJailbrokenDevices,omitempty"`
	// SecurityDisableUsbDebugging undocumented
	SecurityDisableUsbDebugging *bool `json:"securityDisableUsbDebugging,omitempty"`
	// SecurityPreventInstallAppsFromUnknownSources undocumented
	SecurityPreventInstallAppsFromUnknownSources *bool `json:"securityPreventInstallAppsFromUnknownSources,omitempty"`
	// SecurityRequireCompanyPortalAppIntegrity undocumented
	SecurityRequireCompanyPortalAppIntegrity *bool `json:"securityRequireCompanyPortalAppIntegrity,omitempty"`
	// SecurityRequireGooglePlayServices undocumented
	SecurityRequireGooglePlayServices *bool `json:"securityRequireGooglePlayServices,omitempty"`
	// SecurityRequireSafetyNetAttestationBasicIntegrity undocumented
	SecurityRequireSafetyNetAttestationBasicIntegrity *bool `json:"securityRequireSafetyNetAttestationBasicIntegrity,omitempty"`
	// SecurityRequireSafetyNetAttestationCertifiedDevice undocumented
	SecurityRequireSafetyNetAttestationCertifiedDevice *bool `json:"securityRequireSafetyNetAttestationCertifiedDevice,omitempty"`
	// SecurityRequireUpToDateSecurityProviders undocumented
	SecurityRequireUpToDateSecurityProviders *bool `json:"securityRequireUpToDateSecurityProviders,omitempty"`
	// SecurityRequireVerifyApps undocumented
	SecurityRequireVerifyApps *bool `json:"securityRequireVerifyApps,omitempty"`
	// StorageRequireEncryption undocumented
	StorageRequireEncryption *bool `json:"storageRequireEncryption,omitempty"`
}

func NewAndroidCompliancePolicy() (*AndroidCompliancePolicy, error) {
	newAndroidCompliancePolicy := &AndroidCompliancePolicy{
		ODataType: "#microsoft.graph.AndroidCompliancePolicy",
	}
	return newAndroidCompliancePolicy, nil
}

// AndroidCustomConfiguration undocumented
type AndroidCustomConfiguration struct {
	// DeviceConfiguration is the base model of AndroidCustomConfiguration
	DeviceConfiguration

	ODataType string `json:"@odata.type,omitempty"`
	// OMASettings undocumented
	OMASettings []OMASetting `json:"omaSettings,omitempty"`
}

func NewAndroidCustomConfiguration() (*AndroidCustomConfiguration, error) {
	newAndroidCustomConfiguration := &AndroidCustomConfiguration{
		ODataType: "#microsoft.graph.AndroidCustomConfiguration",
	}
	return newAndroidCustomConfiguration, nil
}

// AndroidGeneralDeviceConfiguration undocumented
type AndroidGeneralDeviceConfiguration struct {
	// DeviceConfiguration is the base model of AndroidGeneralDeviceConfiguration
	DeviceConfiguration

	ODataType string `json:"@odata.type,omitempty"`
	// AppsBlockClipboardSharing undocumented
	AppsBlockClipboardSharing *bool `json:"appsBlockClipboardSharing,omitempty"`
	// AppsBlockCopyPaste undocumented
	AppsBlockCopyPaste *bool `json:"appsBlockCopyPaste,omitempty"`
	// AppsBlockYouTube undocumented
	AppsBlockYouTube *bool `json:"appsBlockYouTube,omitempty"`
	// AppsHideList undocumented
	AppsHideList []AppListItem `json:"appsHideList,omitempty"`
	// AppsInstallAllowList undocumented
	AppsInstallAllowList []AppListItem `json:"appsInstallAllowList,omitempty"`
	// AppsLaunchBlockList undocumented
	AppsLaunchBlockList []AppListItem `json:"appsLaunchBlockList,omitempty"`
	// BluetoothBlocked undocumented
	BluetoothBlocked *bool `json:"bluetoothBlocked,omitempty"`
	// CameraBlocked undocumented
	CameraBlocked *bool `json:"cameraBlocked,omitempty"`
	// CellularBlockDataRoaming undocumented
	CellularBlockDataRoaming *bool `json:"cellularBlockDataRoaming,omitempty"`
	// CellularBlockMessaging undocumented
	CellularBlockMessaging *bool `json:"cellularBlockMessaging,omitempty"`
	// CellularBlockVoiceRoaming undocumented
	CellularBlockVoiceRoaming *bool `json:"cellularBlockVoiceRoaming,omitempty"`
	// CellularBlockWiFiTethering undocumented
	CellularBlockWiFiTethering *bool `json:"cellularBlockWiFiTethering,omitempty"`
	// CompliantAppListType undocumented
	CompliantAppListType *AppListType `json:"compliantAppListType,omitempty"`
	// CompliantAppsList undocumented
	CompliantAppsList []AppListItem `json:"compliantAppsList,omitempty"`
	// DeviceSharingAllowed undocumented
	DeviceSharingAllowed *bool `json:"deviceSharingAllowed,omitempty"`
	// DiagnosticDataBlockSubmission undocumented
	DiagnosticDataBlockSubmission *bool `json:"diagnosticDataBlockSubmission,omitempty"`
	// FactoryResetBlocked undocumented
	FactoryResetBlocked *bool `json:"factoryResetBlocked,omitempty"`
	// GoogleAccountBlockAutoSync undocumented
	GoogleAccountBlockAutoSync *bool `json:"googleAccountBlockAutoSync,omitempty"`
	// GooglePlayStoreBlocked undocumented
	GooglePlayStoreBlocked *bool `json:"googlePlayStoreBlocked,omitempty"`
	// KioskModeApps undocumented
	KioskModeApps []AppListItem `json:"kioskModeApps,omitempty"`
	// KioskModeBlockSleepButton undocumented
	KioskModeBlockSleepButton *bool `json:"kioskModeBlockSleepButton,omitempty"`
	// KioskModeBlockVolumeButtons undocumented
	KioskModeBlockVolumeButtons *bool `json:"kioskModeBlockVolumeButtons,omitempty"`
	// LocationServicesBlocked undocumented
	LocationServicesBlocked *bool `json:"locationServicesBlocked,omitempty"`
	// NfcBlocked undocumented
	NfcBlocked *bool `json:"nfcBlocked,omitempty"`
	// PasswordBlockFingerprintUnlock undocumented
	PasswordBlockFingerprintUnlock *bool `json:"passwordBlockFingerprintUnlock,omitempty"`
	// PasswordBlockTrustAgents undocumented
	PasswordBlockTrustAgents *bool `json:"passwordBlockTrustAgents,omitempty"`
	// PasswordExpirationDays undocumented
	PasswordExpirationDays *int `json:"passwordExpirationDays,omitempty"`
	// PasswordMinimumLength undocumented
	PasswordMinimumLength *int `json:"passwordMinimumLength,omitempty"`
	// PasswordMinutesOfInactivityBeforeScreenTimeout undocumented
	PasswordMinutesOfInactivityBeforeScreenTimeout *int `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
	// PasswordPreviousPasswordBlockCount undocumented
	PasswordPreviousPasswordBlockCount *int `json:"passwordPreviousPasswordBlockCount,omitempty"`
	// PasswordRequired undocumented
	PasswordRequired *bool `json:"passwordRequired,omitempty"`
	// PasswordRequiredType undocumented
	PasswordRequiredType *AndroidRequiredPasswordType `json:"passwordRequiredType,omitempty"`
	// PasswordSignInFailureCountBeforeFactoryReset undocumented
	PasswordSignInFailureCountBeforeFactoryReset *int `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`
	// PowerOffBlocked undocumented
	PowerOffBlocked *bool `json:"powerOffBlocked,omitempty"`
	// ScreenCaptureBlocked undocumented
	ScreenCaptureBlocked *bool `json:"screenCaptureBlocked,omitempty"`
	// SecurityRequireVerifyApps undocumented
	SecurityRequireVerifyApps *bool `json:"securityRequireVerifyApps,omitempty"`
	// StorageBlockGoogleBackup undocumented
	StorageBlockGoogleBackup *bool `json:"storageBlockGoogleBackup,omitempty"`
	// StorageBlockRemovableStorage undocumented
	StorageBlockRemovableStorage *bool `json:"storageBlockRemovableStorage,omitempty"`
	// StorageRequireDeviceEncryption undocumented
	StorageRequireDeviceEncryption *bool `json:"storageRequireDeviceEncryption,omitempty"`
	// StorageRequireRemovableStorageEncryption undocumented
	StorageRequireRemovableStorageEncryption *bool `json:"storageRequireRemovableStorageEncryption,omitempty"`
	// VoiceAssistantBlocked undocumented
	VoiceAssistantBlocked *bool `json:"voiceAssistantBlocked,omitempty"`
	// VoiceDialingBlocked undocumented
	VoiceDialingBlocked *bool `json:"voiceDialingBlocked,omitempty"`
	// WebBrowserBlockAutofill undocumented
	WebBrowserBlockAutofill *bool `json:"webBrowserBlockAutofill,omitempty"`
	// WebBrowserBlocked undocumented
	WebBrowserBlocked *bool `json:"webBrowserBlocked,omitempty"`
	// WebBrowserBlockJavaScript undocumented
	WebBrowserBlockJavaScript *bool `json:"webBrowserBlockJavaScript,omitempty"`
	// WebBrowserBlockPopups undocumented
	WebBrowserBlockPopups *bool `json:"webBrowserBlockPopups,omitempty"`
	// WebBrowserCookieSettings undocumented
	WebBrowserCookieSettings *WebBrowserCookieSettings `json:"webBrowserCookieSettings,omitempty"`
	// WiFiBlocked undocumented
	WiFiBlocked *bool `json:"wiFiBlocked,omitempty"`
}

func NewAndroidGeneralDeviceConfiguration() (*AndroidGeneralDeviceConfiguration, error) {
	newAndroidGeneralDeviceConfiguration := &AndroidGeneralDeviceConfiguration{
		ODataType: "#microsoft.graph.AndroidGeneralDeviceConfiguration",
	}
	return newAndroidGeneralDeviceConfiguration, nil
}

// AndroidLobApp undocumented
type AndroidLobApp struct {
	// MobileLobApp is the base model of AndroidLobApp
	MobileLobApp

	ODataType string `json:"@odata.type,omitempty"`
	// MinimumSupportedOperatingSystem undocumented
	MinimumSupportedOperatingSystem *AndroidMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`
	// PackageID undocumented
	PackageID *string `json:"packageId,omitempty"`
	// VersionCode undocumented
	VersionCode *string `json:"versionCode,omitempty"`
	// VersionName undocumented
	VersionName *string `json:"versionName,omitempty"`
}

func NewAndroidLobApp() (*AndroidLobApp, error) {
	newAndroidLobApp := &AndroidLobApp{
		ODataType: "#microsoft.graph.AndroidLobApp",
	}
	return newAndroidLobApp, nil
}

// AndroidManagedAppProtection undocumented
type AndroidManagedAppProtection struct {
	// TargetedManagedAppProtection is the base model of AndroidManagedAppProtection
	TargetedManagedAppProtection

	ODataType string `json:"@odata.type,omitempty"`
	// CustomBrowserDisplayName undocumented
	CustomBrowserDisplayName *string `json:"customBrowserDisplayName,omitempty"`
	// CustomBrowserPackageID undocumented
	CustomBrowserPackageID *string `json:"customBrowserPackageId,omitempty"`
	// DeployedAppCount undocumented
	DeployedAppCount *int `json:"deployedAppCount,omitempty"`
	// DisableAppEncryptionIfDeviceEncryptionIsEnabled undocumented
	DisableAppEncryptionIfDeviceEncryptionIsEnabled *bool `json:"disableAppEncryptionIfDeviceEncryptionIsEnabled,omitempty"`
	// EncryptAppData undocumented
	EncryptAppData *bool `json:"encryptAppData,omitempty"`
	// MinimumRequiredPatchVersion undocumented
	MinimumRequiredPatchVersion *string `json:"minimumRequiredPatchVersion,omitempty"`
	// MinimumWarningPatchVersion undocumented
	MinimumWarningPatchVersion *string `json:"minimumWarningPatchVersion,omitempty"`
	// ScreenCaptureBlocked undocumented
	ScreenCaptureBlocked *bool `json:"screenCaptureBlocked,omitempty"`
	// Apps undocumented
	Apps []ManagedMobileApp `json:"apps,omitempty"`
	// DeploymentSummary undocumented
	DeploymentSummary *ManagedAppPolicyDeploymentSummary `json:"deploymentSummary,omitempty"`
}

func NewAndroidManagedAppProtection() (*AndroidManagedAppProtection, error) {
	newAndroidManagedAppProtection := &AndroidManagedAppProtection{
		ODataType: "#microsoft.graph.AndroidManagedAppProtection",
	}
	return newAndroidManagedAppProtection, nil
}

// AndroidManagedAppRegistration undocumented
type AndroidManagedAppRegistration struct {
	// ManagedAppRegistration is the base model of AndroidManagedAppRegistration
	ManagedAppRegistration

	ODataType string `json:"@odata.type,omitempty"`
}

func NewAndroidManagedAppRegistration() (*AndroidManagedAppRegistration, error) {
	newAndroidManagedAppRegistration := &AndroidManagedAppRegistration{
		ODataType: "#microsoft.graph.AndroidManagedAppRegistration",
	}
	return newAndroidManagedAppRegistration, nil
}

// AndroidMinimumOperatingSystem undocumented
type AndroidMinimumOperatingSystem struct {
	// Object is the base model of AndroidMinimumOperatingSystem
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// V10_0 undocumented
	V10_0 *bool `json:"v10_0,omitempty"`
	// V11_0 undocumented
	V11_0 *bool `json:"v11_0,omitempty"`
	// V4_0 undocumented
	V4_0 *bool `json:"v4_0,omitempty"`
	// V4_0_3 undocumented
	V4_0_3 *bool `json:"v4_0_3,omitempty"`
	// V4_1 undocumented
	V4_1 *bool `json:"v4_1,omitempty"`
	// V4_2 undocumented
	V4_2 *bool `json:"v4_2,omitempty"`
	// V4_3 undocumented
	V4_3 *bool `json:"v4_3,omitempty"`
	// V4_4 undocumented
	V4_4 *bool `json:"v4_4,omitempty"`
	// V5_0 undocumented
	V5_0 *bool `json:"v5_0,omitempty"`
	// V5_1 undocumented
	V5_1 *bool `json:"v5_1,omitempty"`
	// V6_0 undocumented
	V6_0 *bool `json:"v6_0,omitempty"`
	// V7_0 undocumented
	V7_0 *bool `json:"v7_0,omitempty"`
	// V7_1 undocumented
	V7_1 *bool `json:"v7_1,omitempty"`
	// V8_0 undocumented
	V8_0 *bool `json:"v8_0,omitempty"`
	// V8_1 undocumented
	V8_1 *bool `json:"v8_1,omitempty"`
	// V9_0 undocumented
	V9_0 *bool `json:"v9_0,omitempty"`
}

func NewAndroidMinimumOperatingSystem() (*AndroidMinimumOperatingSystem, error) {
	newAndroidMinimumOperatingSystem := &AndroidMinimumOperatingSystem{
		ODataType: "#microsoft.graph.AndroidMinimumOperatingSystem",
	}
	return newAndroidMinimumOperatingSystem, nil
}

// AndroidMobileAppIdentifier undocumented
type AndroidMobileAppIdentifier struct {
	// MobileAppIdentifier is the base model of AndroidMobileAppIdentifier
	MobileAppIdentifier

	ODataType string `json:"@odata.type,omitempty"`
	// PackageID undocumented
	PackageID *string `json:"packageId,omitempty"`
}

func NewAndroidMobileAppIdentifier() (*AndroidMobileAppIdentifier, error) {
	newAndroidMobileAppIdentifier := &AndroidMobileAppIdentifier{
		ODataType: "#microsoft.graph.AndroidMobileAppIdentifier",
	}
	return newAndroidMobileAppIdentifier, nil
}

// AndroidStoreApp undocumented
type AndroidStoreApp struct {
	// MobileApp is the base model of AndroidStoreApp
	MobileApp

	ODataType string `json:"@odata.type,omitempty"`
	// AppStoreURL undocumented
	AppStoreURL *string `json:"appStoreUrl,omitempty"`
	// MinimumSupportedOperatingSystem undocumented
	MinimumSupportedOperatingSystem *AndroidMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`
	// PackageID undocumented
	PackageID *string `json:"packageId,omitempty"`
}

func NewAndroidStoreApp() (*AndroidStoreApp, error) {
	newAndroidStoreApp := &AndroidStoreApp{
		ODataType: "#microsoft.graph.AndroidStoreApp",
	}
	return newAndroidStoreApp, nil
}

// AndroidWorkProfileCompliancePolicy undocumented
type AndroidWorkProfileCompliancePolicy struct {
	// DeviceCompliancePolicy is the base model of AndroidWorkProfileCompliancePolicy
	DeviceCompliancePolicy

	ODataType string `json:"@odata.type,omitempty"`
	// DeviceThreatProtectionEnabled undocumented
	DeviceThreatProtectionEnabled *bool `json:"deviceThreatProtectionEnabled,omitempty"`
	// DeviceThreatProtectionRequiredSecurityLevel undocumented
	DeviceThreatProtectionRequiredSecurityLevel *DeviceThreatProtectionLevel `json:"deviceThreatProtectionRequiredSecurityLevel,omitempty"`
	// MinAndroidSecurityPatchLevel undocumented
	MinAndroidSecurityPatchLevel *string `json:"minAndroidSecurityPatchLevel,omitempty"`
	// OsMaximumVersion undocumented
	OsMaximumVersion *string `json:"osMaximumVersion,omitempty"`
	// OsMinimumVersion undocumented
	OsMinimumVersion *string `json:"osMinimumVersion,omitempty"`
	// PasswordExpirationDays undocumented
	PasswordExpirationDays *int `json:"passwordExpirationDays,omitempty"`
	// PasswordMinimumLength undocumented
	PasswordMinimumLength *int `json:"passwordMinimumLength,omitempty"`
	// PasswordMinutesOfInactivityBeforeLock undocumented
	PasswordMinutesOfInactivityBeforeLock *int `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`
	// PasswordPreviousPasswordBlockCount undocumented
	PasswordPreviousPasswordBlockCount *int `json:"passwordPreviousPasswordBlockCount,omitempty"`
	// PasswordRequired undocumented
	PasswordRequired *bool `json:"passwordRequired,omitempty"`
	// PasswordRequiredType undocumented
	PasswordRequiredType *AndroidRequiredPasswordType `json:"passwordRequiredType,omitempty"`
	// SecurityBlockJailbrokenDevices undocumented
	SecurityBlockJailbrokenDevices *bool `json:"securityBlockJailbrokenDevices,omitempty"`
	// SecurityDisableUsbDebugging undocumented
	SecurityDisableUsbDebugging *bool `json:"securityDisableUsbDebugging,omitempty"`
	// SecurityPreventInstallAppsFromUnknownSources undocumented
	SecurityPreventInstallAppsFromUnknownSources *bool `json:"securityPreventInstallAppsFromUnknownSources,omitempty"`
	// SecurityRequireCompanyPortalAppIntegrity undocumented
	SecurityRequireCompanyPortalAppIntegrity *bool `json:"securityRequireCompanyPortalAppIntegrity,omitempty"`
	// SecurityRequireGooglePlayServices undocumented
	SecurityRequireGooglePlayServices *bool `json:"securityRequireGooglePlayServices,omitempty"`
	// SecurityRequireSafetyNetAttestationBasicIntegrity undocumented
	SecurityRequireSafetyNetAttestationBasicIntegrity *bool `json:"securityRequireSafetyNetAttestationBasicIntegrity,omitempty"`
	// SecurityRequireSafetyNetAttestationCertifiedDevice undocumented
	SecurityRequireSafetyNetAttestationCertifiedDevice *bool `json:"securityRequireSafetyNetAttestationCertifiedDevice,omitempty"`
	// SecurityRequireUpToDateSecurityProviders undocumented
	SecurityRequireUpToDateSecurityProviders *bool `json:"securityRequireUpToDateSecurityProviders,omitempty"`
	// SecurityRequireVerifyApps undocumented
	SecurityRequireVerifyApps *bool `json:"securityRequireVerifyApps,omitempty"`
	// StorageRequireEncryption undocumented
	StorageRequireEncryption *bool `json:"storageRequireEncryption,omitempty"`
}

func NewAndroidWorkProfileCompliancePolicy() (*AndroidWorkProfileCompliancePolicy, error) {
	newAndroidWorkProfileCompliancePolicy := &AndroidWorkProfileCompliancePolicy{
		ODataType: "#microsoft.graph.AndroidWorkProfileCompliancePolicy",
	}
	return newAndroidWorkProfileCompliancePolicy, nil
}

// AndroidWorkProfileCustomConfiguration undocumented
type AndroidWorkProfileCustomConfiguration struct {
	// DeviceConfiguration is the base model of AndroidWorkProfileCustomConfiguration
	DeviceConfiguration

	ODataType string `json:"@odata.type,omitempty"`
	// OMASettings undocumented
	OMASettings []OMASetting `json:"omaSettings,omitempty"`
}

func NewAndroidWorkProfileCustomConfiguration() (*AndroidWorkProfileCustomConfiguration, error) {
	newAndroidWorkProfileCustomConfiguration := &AndroidWorkProfileCustomConfiguration{
		ODataType: "#microsoft.graph.AndroidWorkProfileCustomConfiguration",
	}
	return newAndroidWorkProfileCustomConfiguration, nil
}

// AndroidWorkProfileGeneralDeviceConfiguration undocumented
type AndroidWorkProfileGeneralDeviceConfiguration struct {
	// DeviceConfiguration is the base model of AndroidWorkProfileGeneralDeviceConfiguration
	DeviceConfiguration

	ODataType string `json:"@odata.type,omitempty"`
	// PasswordBlockFingerprintUnlock undocumented
	PasswordBlockFingerprintUnlock *bool `json:"passwordBlockFingerprintUnlock,omitempty"`
	// PasswordBlockTrustAgents undocumented
	PasswordBlockTrustAgents *bool `json:"passwordBlockTrustAgents,omitempty"`
	// PasswordExpirationDays undocumented
	PasswordExpirationDays *int `json:"passwordExpirationDays,omitempty"`
	// PasswordMinimumLength undocumented
	PasswordMinimumLength *int `json:"passwordMinimumLength,omitempty"`
	// PasswordMinutesOfInactivityBeforeScreenTimeout undocumented
	PasswordMinutesOfInactivityBeforeScreenTimeout *int `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
	// PasswordPreviousPasswordBlockCount undocumented
	PasswordPreviousPasswordBlockCount *int `json:"passwordPreviousPasswordBlockCount,omitempty"`
	// PasswordRequiredType undocumented
	PasswordRequiredType *AndroidWorkProfileRequiredPasswordType `json:"passwordRequiredType,omitempty"`
	// PasswordSignInFailureCountBeforeFactoryReset undocumented
	PasswordSignInFailureCountBeforeFactoryReset *int `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`
	// SecurityRequireVerifyApps undocumented
	SecurityRequireVerifyApps *bool `json:"securityRequireVerifyApps,omitempty"`
	// WorkProfileBlockAddingAccounts undocumented
	WorkProfileBlockAddingAccounts *bool `json:"workProfileBlockAddingAccounts,omitempty"`
	// WorkProfileBlockCamera undocumented
	WorkProfileBlockCamera *bool `json:"workProfileBlockCamera,omitempty"`
	// WorkProfileBlockCrossProfileCallerID undocumented
	WorkProfileBlockCrossProfileCallerID *bool `json:"workProfileBlockCrossProfileCallerId,omitempty"`
	// WorkProfileBlockCrossProfileContactsSearch undocumented
	WorkProfileBlockCrossProfileContactsSearch *bool `json:"workProfileBlockCrossProfileContactsSearch,omitempty"`
	// WorkProfileBlockCrossProfileCopyPaste undocumented
	WorkProfileBlockCrossProfileCopyPaste *bool `json:"workProfileBlockCrossProfileCopyPaste,omitempty"`
	// WorkProfileBlockNotificationsWhileDeviceLocked undocumented
	WorkProfileBlockNotificationsWhileDeviceLocked *bool `json:"workProfileBlockNotificationsWhileDeviceLocked,omitempty"`
	// WorkProfileBlockScreenCapture undocumented
	WorkProfileBlockScreenCapture *bool `json:"workProfileBlockScreenCapture,omitempty"`
	// WorkProfileBluetoothEnableContactSharing undocumented
	WorkProfileBluetoothEnableContactSharing *bool `json:"workProfileBluetoothEnableContactSharing,omitempty"`
	// WorkProfileDataSharingType undocumented
	WorkProfileDataSharingType *AndroidWorkProfileCrossProfileDataSharingType `json:"workProfileDataSharingType,omitempty"`
	// WorkProfileDefaultAppPermissionPolicy undocumented
	WorkProfileDefaultAppPermissionPolicy *AndroidWorkProfileDefaultAppPermissionPolicyType `json:"workProfileDefaultAppPermissionPolicy,omitempty"`
	// WorkProfilePasswordBlockFingerprintUnlock undocumented
	WorkProfilePasswordBlockFingerprintUnlock *bool `json:"workProfilePasswordBlockFingerprintUnlock,omitempty"`
	// WorkProfilePasswordBlockTrustAgents undocumented
	WorkProfilePasswordBlockTrustAgents *bool `json:"workProfilePasswordBlockTrustAgents,omitempty"`
	// WorkProfilePasswordExpirationDays undocumented
	WorkProfilePasswordExpirationDays *int `json:"workProfilePasswordExpirationDays,omitempty"`
	// WorkProfilePasswordMinimumLength undocumented
	WorkProfilePasswordMinimumLength *int `json:"workProfilePasswordMinimumLength,omitempty"`
	// WorkProfilePasswordMinLetterCharacters undocumented
	WorkProfilePasswordMinLetterCharacters *int `json:"workProfilePasswordMinLetterCharacters,omitempty"`
	// WorkProfilePasswordMinLowerCaseCharacters undocumented
	WorkProfilePasswordMinLowerCaseCharacters *int `json:"workProfilePasswordMinLowerCaseCharacters,omitempty"`
	// WorkProfilePasswordMinNonLetterCharacters undocumented
	WorkProfilePasswordMinNonLetterCharacters *int `json:"workProfilePasswordMinNonLetterCharacters,omitempty"`
	// WorkProfilePasswordMinNumericCharacters undocumented
	WorkProfilePasswordMinNumericCharacters *int `json:"workProfilePasswordMinNumericCharacters,omitempty"`
	// WorkProfilePasswordMinSymbolCharacters undocumented
	WorkProfilePasswordMinSymbolCharacters *int `json:"workProfilePasswordMinSymbolCharacters,omitempty"`
	// WorkProfilePasswordMinUpperCaseCharacters undocumented
	WorkProfilePasswordMinUpperCaseCharacters *int `json:"workProfilePasswordMinUpperCaseCharacters,omitempty"`
	// WorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout undocumented
	WorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout *int `json:"workProfilePasswordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
	// WorkProfilePasswordPreviousPasswordBlockCount undocumented
	WorkProfilePasswordPreviousPasswordBlockCount *int `json:"workProfilePasswordPreviousPasswordBlockCount,omitempty"`
	// WorkProfilePasswordRequiredType undocumented
	WorkProfilePasswordRequiredType *AndroidWorkProfileRequiredPasswordType `json:"workProfilePasswordRequiredType,omitempty"`
	// WorkProfilePasswordSignInFailureCountBeforeFactoryReset undocumented
	WorkProfilePasswordSignInFailureCountBeforeFactoryReset *int `json:"workProfilePasswordSignInFailureCountBeforeFactoryReset,omitempty"`
	// WorkProfileRequirePassword undocumented
	WorkProfileRequirePassword *bool `json:"workProfileRequirePassword,omitempty"`
}

func NewAndroidWorkProfileGeneralDeviceConfiguration() (*AndroidWorkProfileGeneralDeviceConfiguration, error) {
	newAndroidWorkProfileGeneralDeviceConfiguration := &AndroidWorkProfileGeneralDeviceConfiguration{
		ODataType: "#microsoft.graph.AndroidWorkProfileGeneralDeviceConfiguration",
	}
	return newAndroidWorkProfileGeneralDeviceConfiguration, nil
}
