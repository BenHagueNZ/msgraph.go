// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// IOSCertificateProfile undocumented
type IOSCertificateProfile struct {
	// DeviceConfiguration is the base model of IOSCertificateProfile
	DeviceConfiguration
}

// IOSCompliancePolicy undocumented
type IOSCompliancePolicy struct {
	// DeviceCompliancePolicy is the base model of IOSCompliancePolicy
	DeviceCompliancePolicy
	// DeviceThreatProtectionEnabled undocumented
	DeviceThreatProtectionEnabled *bool `json:"deviceThreatProtectionEnabled,omitempty"`
	// DeviceThreatProtectionRequiredSecurityLevel undocumented
	DeviceThreatProtectionRequiredSecurityLevel *DeviceThreatProtectionLevel `json:"deviceThreatProtectionRequiredSecurityLevel,omitempty"`
	// ManagedEmailProfileRequired undocumented
	ManagedEmailProfileRequired *bool `json:"managedEmailProfileRequired,omitempty"`
	// OsMaximumVersion undocumented
	OsMaximumVersion *string `json:"osMaximumVersion,omitempty"`
	// OsMinimumVersion undocumented
	OsMinimumVersion *string `json:"osMinimumVersion,omitempty"`
	// PasscodeBlockSimple undocumented
	PasscodeBlockSimple *bool `json:"passcodeBlockSimple,omitempty"`
	// PasscodeExpirationDays undocumented
	PasscodeExpirationDays *int `json:"passcodeExpirationDays,omitempty"`
	// PasscodeMinimumCharacterSetCount undocumented
	PasscodeMinimumCharacterSetCount *int `json:"passcodeMinimumCharacterSetCount,omitempty"`
	// PasscodeMinimumLength undocumented
	PasscodeMinimumLength *int `json:"passcodeMinimumLength,omitempty"`
	// PasscodeMinutesOfInactivityBeforeLock undocumented
	PasscodeMinutesOfInactivityBeforeLock *int `json:"passcodeMinutesOfInactivityBeforeLock,omitempty"`
	// PasscodePreviousPasscodeBlockCount undocumented
	PasscodePreviousPasscodeBlockCount *int `json:"passcodePreviousPasscodeBlockCount,omitempty"`
	// PasscodeRequired undocumented
	PasscodeRequired *bool `json:"passcodeRequired,omitempty"`
	// PasscodeRequiredType undocumented
	PasscodeRequiredType *RequiredPasswordType `json:"passcodeRequiredType,omitempty"`
	// SecurityBlockJailbrokenDevices undocumented
	SecurityBlockJailbrokenDevices *bool `json:"securityBlockJailbrokenDevices,omitempty"`
}

// IOSCustomConfiguration undocumented
type IOSCustomConfiguration struct {
	// DeviceConfiguration is the base model of IOSCustomConfiguration
	DeviceConfiguration
	// Payload undocumented
	Payload *Binary `json:"payload,omitempty"`
	// PayloadFileName undocumented
	PayloadFileName *string `json:"payloadFileName,omitempty"`
	// PayloadName undocumented
	PayloadName *string `json:"payloadName,omitempty"`
}

// IOSDeviceFeaturesConfiguration undocumented
type IOSDeviceFeaturesConfiguration struct {
	// AppleDeviceFeaturesConfigurationBase is the base model of IOSDeviceFeaturesConfiguration
	AppleDeviceFeaturesConfigurationBase
	// AssetTagTemplate undocumented
	AssetTagTemplate *string `json:"assetTagTemplate,omitempty"`
	// HomeScreenDockIcons undocumented
	HomeScreenDockIcons []IOSHomeScreenItem `json:"homeScreenDockIcons,omitempty"`
	// HomeScreenPages undocumented
	HomeScreenPages []IOSHomeScreenPage `json:"homeScreenPages,omitempty"`
	// LockScreenFootnote undocumented
	LockScreenFootnote *string `json:"lockScreenFootnote,omitempty"`
	// NotificationSettings undocumented
	NotificationSettings []IOSNotificationSettings `json:"notificationSettings,omitempty"`
}

// IOSDeviceType undocumented
type IOSDeviceType struct {
	// Object is the base model of IOSDeviceType
	Object
	// IPad undocumented
	IPad *bool `json:"iPad,omitempty"`
	// IPhoneAndIPod undocumented
	IPhoneAndIPod *bool `json:"iPhoneAndIPod,omitempty"`
}

// IOSGeneralDeviceConfiguration undocumented
type IOSGeneralDeviceConfiguration struct {
	// DeviceConfiguration is the base model of IOSGeneralDeviceConfiguration
	DeviceConfiguration
	// AccountBlockModification undocumented
	AccountBlockModification *bool `json:"accountBlockModification,omitempty"`
	// ActivationLockAllowWhenSupervised undocumented
	ActivationLockAllowWhenSupervised *bool `json:"activationLockAllowWhenSupervised,omitempty"`
	// AirDropBlocked undocumented
	AirDropBlocked *bool `json:"airDropBlocked,omitempty"`
	// AirDropForceUnmanagedDropTarget undocumented
	AirDropForceUnmanagedDropTarget *bool `json:"airDropForceUnmanagedDropTarget,omitempty"`
	// AirPlayForcePairingPasswordForOutgoingRequests undocumented
	AirPlayForcePairingPasswordForOutgoingRequests *bool `json:"airPlayForcePairingPasswordForOutgoingRequests,omitempty"`
	// AppleNewsBlocked undocumented
	AppleNewsBlocked *bool `json:"appleNewsBlocked,omitempty"`
	// AppleWatchBlockPairing undocumented
	AppleWatchBlockPairing *bool `json:"appleWatchBlockPairing,omitempty"`
	// AppleWatchForceWristDetection undocumented
	AppleWatchForceWristDetection *bool `json:"appleWatchForceWristDetection,omitempty"`
	// AppsSingleAppModeList undocumented
	AppsSingleAppModeList []AppListItem `json:"appsSingleAppModeList,omitempty"`
	// AppStoreBlockAutomaticDownloads undocumented
	AppStoreBlockAutomaticDownloads *bool `json:"appStoreBlockAutomaticDownloads,omitempty"`
	// AppStoreBlocked undocumented
	AppStoreBlocked *bool `json:"appStoreBlocked,omitempty"`
	// AppStoreBlockInAppPurchases undocumented
	AppStoreBlockInAppPurchases *bool `json:"appStoreBlockInAppPurchases,omitempty"`
	// AppStoreBlockUIAppInstallation undocumented
	AppStoreBlockUIAppInstallation *bool `json:"appStoreBlockUIAppInstallation,omitempty"`
	// AppStoreRequirePassword undocumented
	AppStoreRequirePassword *bool `json:"appStoreRequirePassword,omitempty"`
	// AppsVisibilityList undocumented
	AppsVisibilityList []AppListItem `json:"appsVisibilityList,omitempty"`
	// AppsVisibilityListType undocumented
	AppsVisibilityListType *AppListType `json:"appsVisibilityListType,omitempty"`
	// BluetoothBlockModification undocumented
	BluetoothBlockModification *bool `json:"bluetoothBlockModification,omitempty"`
	// CameraBlocked undocumented
	CameraBlocked *bool `json:"cameraBlocked,omitempty"`
	// CellularBlockDataRoaming undocumented
	CellularBlockDataRoaming *bool `json:"cellularBlockDataRoaming,omitempty"`
	// CellularBlockGlobalBackgroundFetchWhileRoaming undocumented
	CellularBlockGlobalBackgroundFetchWhileRoaming *bool `json:"cellularBlockGlobalBackgroundFetchWhileRoaming,omitempty"`
	// CellularBlockPerAppDataModification undocumented
	CellularBlockPerAppDataModification *bool `json:"cellularBlockPerAppDataModification,omitempty"`
	// CellularBlockPersonalHotspot undocumented
	CellularBlockPersonalHotspot *bool `json:"cellularBlockPersonalHotspot,omitempty"`
	// CellularBlockVoiceRoaming undocumented
	CellularBlockVoiceRoaming *bool `json:"cellularBlockVoiceRoaming,omitempty"`
	// CertificatesBlockUntrustedTLSCertificates undocumented
	CertificatesBlockUntrustedTLSCertificates *bool `json:"certificatesBlockUntrustedTlsCertificates,omitempty"`
	// ClassroomAppBlockRemoteScreenObservation undocumented
	ClassroomAppBlockRemoteScreenObservation *bool `json:"classroomAppBlockRemoteScreenObservation,omitempty"`
	// ClassroomAppForceUnpromptedScreenObservation undocumented
	ClassroomAppForceUnpromptedScreenObservation *bool `json:"classroomAppForceUnpromptedScreenObservation,omitempty"`
	// CompliantAppListType undocumented
	CompliantAppListType *AppListType `json:"compliantAppListType,omitempty"`
	// CompliantAppsList undocumented
	CompliantAppsList []AppListItem `json:"compliantAppsList,omitempty"`
	// ConfigurationProfileBlockChanges undocumented
	ConfigurationProfileBlockChanges *bool `json:"configurationProfileBlockChanges,omitempty"`
	// DefinitionLookupBlocked undocumented
	DefinitionLookupBlocked *bool `json:"definitionLookupBlocked,omitempty"`
	// DeviceBlockEnableRestrictions undocumented
	DeviceBlockEnableRestrictions *bool `json:"deviceBlockEnableRestrictions,omitempty"`
	// DeviceBlockEraseContentAndSettings undocumented
	DeviceBlockEraseContentAndSettings *bool `json:"deviceBlockEraseContentAndSettings,omitempty"`
	// DeviceBlockNameModification undocumented
	DeviceBlockNameModification *bool `json:"deviceBlockNameModification,omitempty"`
	// DiagnosticDataBlockSubmission undocumented
	DiagnosticDataBlockSubmission *bool `json:"diagnosticDataBlockSubmission,omitempty"`
	// DiagnosticDataBlockSubmissionModification undocumented
	DiagnosticDataBlockSubmissionModification *bool `json:"diagnosticDataBlockSubmissionModification,omitempty"`
	// DocumentsBlockManagedDocumentsInUnmanagedApps undocumented
	DocumentsBlockManagedDocumentsInUnmanagedApps *bool `json:"documentsBlockManagedDocumentsInUnmanagedApps,omitempty"`
	// DocumentsBlockUnmanagedDocumentsInManagedApps undocumented
	DocumentsBlockUnmanagedDocumentsInManagedApps *bool `json:"documentsBlockUnmanagedDocumentsInManagedApps,omitempty"`
	// EmailInDomainSuffixes undocumented
	EmailInDomainSuffixes []string `json:"emailInDomainSuffixes,omitempty"`
	// EnterpriseAppBlockTrust undocumented
	EnterpriseAppBlockTrust *bool `json:"enterpriseAppBlockTrust,omitempty"`
	// EnterpriseAppBlockTrustModification undocumented
	EnterpriseAppBlockTrustModification *bool `json:"enterpriseAppBlockTrustModification,omitempty"`
	// FaceTimeBlocked undocumented
	FaceTimeBlocked *bool `json:"faceTimeBlocked,omitempty"`
	// FindMyFriendsBlocked undocumented
	FindMyFriendsBlocked *bool `json:"findMyFriendsBlocked,omitempty"`
	// GameCenterBlocked undocumented
	GameCenterBlocked *bool `json:"gameCenterBlocked,omitempty"`
	// GamingBlockGameCenterFriends undocumented
	GamingBlockGameCenterFriends *bool `json:"gamingBlockGameCenterFriends,omitempty"`
	// GamingBlockMultiplayer undocumented
	GamingBlockMultiplayer *bool `json:"gamingBlockMultiplayer,omitempty"`
	// HostPairingBlocked undocumented
	HostPairingBlocked *bool `json:"hostPairingBlocked,omitempty"`
	// IBooksStoreBlocked undocumented
	IBooksStoreBlocked *bool `json:"iBooksStoreBlocked,omitempty"`
	// IBooksStoreBlockErotica undocumented
	IBooksStoreBlockErotica *bool `json:"iBooksStoreBlockErotica,omitempty"`
	// ICloudBlockActivityContinuation undocumented
	ICloudBlockActivityContinuation *bool `json:"iCloudBlockActivityContinuation,omitempty"`
	// ICloudBlockBackup undocumented
	ICloudBlockBackup *bool `json:"iCloudBlockBackup,omitempty"`
	// ICloudBlockDocumentSync undocumented
	ICloudBlockDocumentSync *bool `json:"iCloudBlockDocumentSync,omitempty"`
	// ICloudBlockManagedAppsSync undocumented
	ICloudBlockManagedAppsSync *bool `json:"iCloudBlockManagedAppsSync,omitempty"`
	// ICloudBlockPhotoLibrary undocumented
	ICloudBlockPhotoLibrary *bool `json:"iCloudBlockPhotoLibrary,omitempty"`
	// ICloudBlockPhotoStreamSync undocumented
	ICloudBlockPhotoStreamSync *bool `json:"iCloudBlockPhotoStreamSync,omitempty"`
	// ICloudBlockSharedPhotoStream undocumented
	ICloudBlockSharedPhotoStream *bool `json:"iCloudBlockSharedPhotoStream,omitempty"`
	// ICloudRequireEncryptedBackup undocumented
	ICloudRequireEncryptedBackup *bool `json:"iCloudRequireEncryptedBackup,omitempty"`
	// ITunesBlockExplicitContent undocumented
	ITunesBlockExplicitContent *bool `json:"iTunesBlockExplicitContent,omitempty"`
	// ITunesBlockMusicService undocumented
	ITunesBlockMusicService *bool `json:"iTunesBlockMusicService,omitempty"`
	// ITunesBlockRadio undocumented
	ITunesBlockRadio *bool `json:"iTunesBlockRadio,omitempty"`
	// KeyboardBlockAutoCorrect undocumented
	KeyboardBlockAutoCorrect *bool `json:"keyboardBlockAutoCorrect,omitempty"`
	// KeyboardBlockDictation undocumented
	KeyboardBlockDictation *bool `json:"keyboardBlockDictation,omitempty"`
	// KeyboardBlockPredictive undocumented
	KeyboardBlockPredictive *bool `json:"keyboardBlockPredictive,omitempty"`
	// KeyboardBlockShortcuts undocumented
	KeyboardBlockShortcuts *bool `json:"keyboardBlockShortcuts,omitempty"`
	// KeyboardBlockSpellCheck undocumented
	KeyboardBlockSpellCheck *bool `json:"keyboardBlockSpellCheck,omitempty"`
	// KioskModeAllowAssistiveSpeak undocumented
	KioskModeAllowAssistiveSpeak *bool `json:"kioskModeAllowAssistiveSpeak,omitempty"`
	// KioskModeAllowAssistiveTouchSettings undocumented
	KioskModeAllowAssistiveTouchSettings *bool `json:"kioskModeAllowAssistiveTouchSettings,omitempty"`
	// KioskModeAllowAutoLock undocumented
	KioskModeAllowAutoLock *bool `json:"kioskModeAllowAutoLock,omitempty"`
	// KioskModeAllowColorInversionSettings undocumented
	KioskModeAllowColorInversionSettings *bool `json:"kioskModeAllowColorInversionSettings,omitempty"`
	// KioskModeAllowRingerSwitch undocumented
	KioskModeAllowRingerSwitch *bool `json:"kioskModeAllowRingerSwitch,omitempty"`
	// KioskModeAllowScreenRotation undocumented
	KioskModeAllowScreenRotation *bool `json:"kioskModeAllowScreenRotation,omitempty"`
	// KioskModeAllowSleepButton undocumented
	KioskModeAllowSleepButton *bool `json:"kioskModeAllowSleepButton,omitempty"`
	// KioskModeAllowTouchscreen undocumented
	KioskModeAllowTouchscreen *bool `json:"kioskModeAllowTouchscreen,omitempty"`
	// KioskModeAllowVoiceOverSettings undocumented
	KioskModeAllowVoiceOverSettings *bool `json:"kioskModeAllowVoiceOverSettings,omitempty"`
	// KioskModeAllowVolumeButtons undocumented
	KioskModeAllowVolumeButtons *bool `json:"kioskModeAllowVolumeButtons,omitempty"`
	// KioskModeAllowZoomSettings undocumented
	KioskModeAllowZoomSettings *bool `json:"kioskModeAllowZoomSettings,omitempty"`
	// KioskModeAppStoreURL undocumented
	KioskModeAppStoreURL *string `json:"kioskModeAppStoreUrl,omitempty"`
	// KioskModeBuiltInAppID undocumented
	KioskModeBuiltInAppID *string `json:"kioskModeBuiltInAppId,omitempty"`
	// KioskModeManagedAppID undocumented
	KioskModeManagedAppID *string `json:"kioskModeManagedAppId,omitempty"`
	// KioskModeRequireAssistiveTouch undocumented
	KioskModeRequireAssistiveTouch *bool `json:"kioskModeRequireAssistiveTouch,omitempty"`
	// KioskModeRequireColorInversion undocumented
	KioskModeRequireColorInversion *bool `json:"kioskModeRequireColorInversion,omitempty"`
	// KioskModeRequireMonoAudio undocumented
	KioskModeRequireMonoAudio *bool `json:"kioskModeRequireMonoAudio,omitempty"`
	// KioskModeRequireVoiceOver undocumented
	KioskModeRequireVoiceOver *bool `json:"kioskModeRequireVoiceOver,omitempty"`
	// KioskModeRequireZoom undocumented
	KioskModeRequireZoom *bool `json:"kioskModeRequireZoom,omitempty"`
	// LockScreenBlockControlCenter undocumented
	LockScreenBlockControlCenter *bool `json:"lockScreenBlockControlCenter,omitempty"`
	// LockScreenBlockNotificationView undocumented
	LockScreenBlockNotificationView *bool `json:"lockScreenBlockNotificationView,omitempty"`
	// LockScreenBlockPassbook undocumented
	LockScreenBlockPassbook *bool `json:"lockScreenBlockPassbook,omitempty"`
	// LockScreenBlockTodayView undocumented
	LockScreenBlockTodayView *bool `json:"lockScreenBlockTodayView,omitempty"`
	// MediaContentRatingApps undocumented
	MediaContentRatingApps *RatingAppsType `json:"mediaContentRatingApps,omitempty"`
	// MediaContentRatingAustralia undocumented
	MediaContentRatingAustralia *MediaContentRatingAustralia `json:"mediaContentRatingAustralia,omitempty"`
	// MediaContentRatingCanada undocumented
	MediaContentRatingCanada *MediaContentRatingCanada `json:"mediaContentRatingCanada,omitempty"`
	// MediaContentRatingFrance undocumented
	MediaContentRatingFrance *MediaContentRatingFrance `json:"mediaContentRatingFrance,omitempty"`
	// MediaContentRatingGermany undocumented
	MediaContentRatingGermany *MediaContentRatingGermany `json:"mediaContentRatingGermany,omitempty"`
	// MediaContentRatingIreland undocumented
	MediaContentRatingIreland *MediaContentRatingIreland `json:"mediaContentRatingIreland,omitempty"`
	// MediaContentRatingJapan undocumented
	MediaContentRatingJapan *MediaContentRatingJapan `json:"mediaContentRatingJapan,omitempty"`
	// MediaContentRatingNewZealand undocumented
	MediaContentRatingNewZealand *MediaContentRatingNewZealand `json:"mediaContentRatingNewZealand,omitempty"`
	// MediaContentRatingUnitedKingdom undocumented
	MediaContentRatingUnitedKingdom *MediaContentRatingUnitedKingdom `json:"mediaContentRatingUnitedKingdom,omitempty"`
	// MediaContentRatingUnitedStates undocumented
	MediaContentRatingUnitedStates *MediaContentRatingUnitedStates `json:"mediaContentRatingUnitedStates,omitempty"`
	// MessagesBlocked undocumented
	MessagesBlocked *bool `json:"messagesBlocked,omitempty"`
	// NetworkUsageRules undocumented
	NetworkUsageRules []IOSNetworkUsageRule `json:"networkUsageRules,omitempty"`
	// NotificationsBlockSettingsModification undocumented
	NotificationsBlockSettingsModification *bool `json:"notificationsBlockSettingsModification,omitempty"`
	// PasscodeBlockFingerprintModification undocumented
	PasscodeBlockFingerprintModification *bool `json:"passcodeBlockFingerprintModification,omitempty"`
	// PasscodeBlockFingerprintUnlock undocumented
	PasscodeBlockFingerprintUnlock *bool `json:"passcodeBlockFingerprintUnlock,omitempty"`
	// PasscodeBlockModification undocumented
	PasscodeBlockModification *bool `json:"passcodeBlockModification,omitempty"`
	// PasscodeBlockSimple undocumented
	PasscodeBlockSimple *bool `json:"passcodeBlockSimple,omitempty"`
	// PasscodeExpirationDays undocumented
	PasscodeExpirationDays *int `json:"passcodeExpirationDays,omitempty"`
	// PasscodeMinimumCharacterSetCount undocumented
	PasscodeMinimumCharacterSetCount *int `json:"passcodeMinimumCharacterSetCount,omitempty"`
	// PasscodeMinimumLength undocumented
	PasscodeMinimumLength *int `json:"passcodeMinimumLength,omitempty"`
	// PasscodeMinutesOfInactivityBeforeLock undocumented
	PasscodeMinutesOfInactivityBeforeLock *int `json:"passcodeMinutesOfInactivityBeforeLock,omitempty"`
	// PasscodeMinutesOfInactivityBeforeScreenTimeout undocumented
	PasscodeMinutesOfInactivityBeforeScreenTimeout *int `json:"passcodeMinutesOfInactivityBeforeScreenTimeout,omitempty"`
	// PasscodePreviousPasscodeBlockCount undocumented
	PasscodePreviousPasscodeBlockCount *int `json:"passcodePreviousPasscodeBlockCount,omitempty"`
	// PasscodeRequired undocumented
	PasscodeRequired *bool `json:"passcodeRequired,omitempty"`
	// PasscodeRequiredType undocumented
	PasscodeRequiredType *RequiredPasswordType `json:"passcodeRequiredType,omitempty"`
	// PasscodeSignInFailureCountBeforeWipe undocumented
	PasscodeSignInFailureCountBeforeWipe *int `json:"passcodeSignInFailureCountBeforeWipe,omitempty"`
	// PodcastsBlocked undocumented
	PodcastsBlocked *bool `json:"podcastsBlocked,omitempty"`
	// SafariBlockAutofill undocumented
	SafariBlockAutofill *bool `json:"safariBlockAutofill,omitempty"`
	// SafariBlocked undocumented
	SafariBlocked *bool `json:"safariBlocked,omitempty"`
	// SafariBlockJavaScript undocumented
	SafariBlockJavaScript *bool `json:"safariBlockJavaScript,omitempty"`
	// SafariBlockPopups undocumented
	SafariBlockPopups *bool `json:"safariBlockPopups,omitempty"`
	// SafariCookieSettings undocumented
	SafariCookieSettings *WebBrowserCookieSettings `json:"safariCookieSettings,omitempty"`
	// SafariManagedDomains undocumented
	SafariManagedDomains []string `json:"safariManagedDomains,omitempty"`
	// SafariPasswordAutoFillDomains undocumented
	SafariPasswordAutoFillDomains []string `json:"safariPasswordAutoFillDomains,omitempty"`
	// SafariRequireFraudWarning undocumented
	SafariRequireFraudWarning *bool `json:"safariRequireFraudWarning,omitempty"`
	// ScreenCaptureBlocked undocumented
	ScreenCaptureBlocked *bool `json:"screenCaptureBlocked,omitempty"`
	// SiriBlocked undocumented
	SiriBlocked *bool `json:"siriBlocked,omitempty"`
	// SiriBlockedWhenLocked undocumented
	SiriBlockedWhenLocked *bool `json:"siriBlockedWhenLocked,omitempty"`
	// SiriBlockUserGeneratedContent undocumented
	SiriBlockUserGeneratedContent *bool `json:"siriBlockUserGeneratedContent,omitempty"`
	// SiriRequireProfanityFilter undocumented
	SiriRequireProfanityFilter *bool `json:"siriRequireProfanityFilter,omitempty"`
	// SpotlightBlockInternetResults undocumented
	SpotlightBlockInternetResults *bool `json:"spotlightBlockInternetResults,omitempty"`
	// VoiceDialingBlocked undocumented
	VoiceDialingBlocked *bool `json:"voiceDialingBlocked,omitempty"`
	// WallpaperBlockModification undocumented
	WallpaperBlockModification *bool `json:"wallpaperBlockModification,omitempty"`
	// WiFiConnectOnlyToConfiguredNetworks undocumented
	WiFiConnectOnlyToConfiguredNetworks *bool `json:"wiFiConnectOnlyToConfiguredNetworks,omitempty"`
}

// IOSHomeScreenApp undocumented
type IOSHomeScreenApp struct {
	// IOSHomeScreenItem is the base model of IOSHomeScreenApp
	IOSHomeScreenItem
	// BundleID undocumented
	BundleID *string `json:"bundleID,omitempty"`
}

// IOSHomeScreenFolder undocumented
type IOSHomeScreenFolder struct {
	// IOSHomeScreenItem is the base model of IOSHomeScreenFolder
	IOSHomeScreenItem
	// Pages undocumented
	Pages []IOSHomeScreenFolderPage `json:"pages,omitempty"`
}

// IOSHomeScreenFolderPage undocumented
type IOSHomeScreenFolderPage struct {
	// Object is the base model of IOSHomeScreenFolderPage
	Object
	// Apps undocumented
	Apps []IOSHomeScreenApp `json:"apps,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
}

// IOSHomeScreenItem undocumented
type IOSHomeScreenItem struct {
	// Object is the base model of IOSHomeScreenItem
	Object
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
}

// IOSHomeScreenPage undocumented
type IOSHomeScreenPage struct {
	// Object is the base model of IOSHomeScreenPage
	Object
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Icons undocumented
	Icons []IOSHomeScreenItem `json:"icons,omitempty"`
}

// IOSLobApp undocumented
type IOSLobApp struct {
	// MobileLobApp is the base model of IOSLobApp
	MobileLobApp
	// ApplicableDeviceType undocumented
	ApplicableDeviceType *IOSDeviceType `json:"applicableDeviceType,omitempty"`
	// BuildNumber undocumented
	BuildNumber *string `json:"buildNumber,omitempty"`
	// BundleID undocumented
	BundleID *string `json:"bundleId,omitempty"`
	// ExpirationDateTime undocumented
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// MinimumSupportedOperatingSystem undocumented
	MinimumSupportedOperatingSystem *IOSMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`
	// VersionNumber undocumented
	VersionNumber *string `json:"versionNumber,omitempty"`
}

// IOSLobAppAssignmentSettings undocumented
type IOSLobAppAssignmentSettings struct {
	// MobileAppAssignmentSettings is the base model of IOSLobAppAssignmentSettings
	MobileAppAssignmentSettings
	// IsRemovable undocumented
	IsRemovable *bool `json:"isRemovable,omitempty"`
	// UninstallOnDeviceRemoval undocumented
	UninstallOnDeviceRemoval *bool `json:"uninstallOnDeviceRemoval,omitempty"`
	// VPNConfigurationID undocumented
	VPNConfigurationID *string `json:"vpnConfigurationId,omitempty"`
}

// IOSLobAppProvisioningConfigurationAssignment undocumented
type IOSLobAppProvisioningConfigurationAssignment struct {
	// Entity is the base model of IOSLobAppProvisioningConfigurationAssignment
	Entity
	// Target undocumented
	Target *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
}

// IOSManagedAppProtection undocumented
type IOSManagedAppProtection struct {
	// TargetedManagedAppProtection is the base model of IOSManagedAppProtection
	TargetedManagedAppProtection
	// AppDataEncryptionType undocumented
	AppDataEncryptionType *ManagedAppDataEncryptionType `json:"appDataEncryptionType,omitempty"`
	// CustomBrowserProtocol undocumented
	CustomBrowserProtocol *string `json:"customBrowserProtocol,omitempty"`
	// DeployedAppCount undocumented
	DeployedAppCount *int `json:"deployedAppCount,omitempty"`
	// FaceIDBlocked undocumented
	FaceIDBlocked *bool `json:"faceIdBlocked,omitempty"`
	// MinimumRequiredSdkVersion undocumented
	MinimumRequiredSdkVersion *string `json:"minimumRequiredSdkVersion,omitempty"`
	// Apps undocumented
	Apps []ManagedMobileApp `json:"apps,omitempty"`
	// DeploymentSummary undocumented
	DeploymentSummary *ManagedAppPolicyDeploymentSummary `json:"deploymentSummary,omitempty"`
}

// IOSManagedAppRegistration undocumented
type IOSManagedAppRegistration struct {
	// ManagedAppRegistration is the base model of IOSManagedAppRegistration
	ManagedAppRegistration
}

// IOSMinimumOperatingSystem undocumented
type IOSMinimumOperatingSystem struct {
	// Object is the base model of IOSMinimumOperatingSystem
	Object
	// V10_0 undocumented
	V10_0 *bool `json:"v10_0,omitempty"`
	// V11_0 undocumented
	V11_0 *bool `json:"v11_0,omitempty"`
	// V12_0 undocumented
	V12_0 *bool `json:"v12_0,omitempty"`
	// V13_0 undocumented
	V13_0 *bool `json:"v13_0,omitempty"`
	// V14_0 undocumented
	V14_0 *bool `json:"v14_0,omitempty"`
	// V15_0 undocumented
	V15_0 *bool `json:"v15_0,omitempty"`
	// V8_0 undocumented
	V8_0 *bool `json:"v8_0,omitempty"`
	// V9_0 undocumented
	V9_0 *bool `json:"v9_0,omitempty"`
}

// IOSMobileAppConfiguration undocumented
type IOSMobileAppConfiguration struct {
	// ManagedDeviceMobileAppConfiguration is the base model of IOSMobileAppConfiguration
	ManagedDeviceMobileAppConfiguration
	// EncodedSettingXML undocumented
	EncodedSettingXML *Binary `json:"encodedSettingXml,omitempty"`
	// Settings undocumented
	Settings []AppConfigurationSettingItem `json:"settings,omitempty"`
}

// IOSMobileAppIdentifier undocumented
type IOSMobileAppIdentifier struct {
	// MobileAppIdentifier is the base model of IOSMobileAppIdentifier
	MobileAppIdentifier
	// BundleID undocumented
	BundleID *string `json:"bundleId,omitempty"`
}

// IOSNetworkUsageRule undocumented
type IOSNetworkUsageRule struct {
	// Object is the base model of IOSNetworkUsageRule
	Object
	// CellularDataBlocked undocumented
	CellularDataBlocked *bool `json:"cellularDataBlocked,omitempty"`
	// CellularDataBlockWhenRoaming undocumented
	CellularDataBlockWhenRoaming *bool `json:"cellularDataBlockWhenRoaming,omitempty"`
	// ManagedApps undocumented
	ManagedApps []AppListItem `json:"managedApps,omitempty"`
}

// IOSNotificationSettings undocumented
type IOSNotificationSettings struct {
	// Object is the base model of IOSNotificationSettings
	Object
	// AlertType undocumented
	AlertType *IOSNotificationAlertType `json:"alertType,omitempty"`
	// AppName undocumented
	AppName *string `json:"appName,omitempty"`
	// BadgesEnabled undocumented
	BadgesEnabled *bool `json:"badgesEnabled,omitempty"`
	// BundleID undocumented
	BundleID *string `json:"bundleID,omitempty"`
	// Enabled undocumented
	Enabled *bool `json:"enabled,omitempty"`
	// Publisher undocumented
	Publisher *string `json:"publisher,omitempty"`
	// ShowInNotificationCenter undocumented
	ShowInNotificationCenter *bool `json:"showInNotificationCenter,omitempty"`
	// ShowOnLockScreen undocumented
	ShowOnLockScreen *bool `json:"showOnLockScreen,omitempty"`
	// SoundsEnabled undocumented
	SoundsEnabled *bool `json:"soundsEnabled,omitempty"`
}

// IOSStoreApp undocumented
type IOSStoreApp struct {
	// MobileApp is the base model of IOSStoreApp
	MobileApp
	// ApplicableDeviceType undocumented
	ApplicableDeviceType *IOSDeviceType `json:"applicableDeviceType,omitempty"`
	// AppStoreURL undocumented
	AppStoreURL *string `json:"appStoreUrl,omitempty"`
	// BundleID undocumented
	BundleID *string `json:"bundleId,omitempty"`
	// MinimumSupportedOperatingSystem undocumented
	MinimumSupportedOperatingSystem *IOSMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`
}

// IOSStoreAppAssignmentSettings undocumented
type IOSStoreAppAssignmentSettings struct {
	// MobileAppAssignmentSettings is the base model of IOSStoreAppAssignmentSettings
	MobileAppAssignmentSettings
	// VPNConfigurationID undocumented
	VPNConfigurationID *string `json:"vpnConfigurationId,omitempty"`
}

// IOSUpdateConfiguration undocumented
type IOSUpdateConfiguration struct {
	// DeviceConfiguration is the base model of IOSUpdateConfiguration
	DeviceConfiguration
	// ActiveHoursEnd undocumented
	ActiveHoursEnd *TimeOfDay `json:"activeHoursEnd,omitempty"`
	// ActiveHoursStart undocumented
	ActiveHoursStart *TimeOfDay `json:"activeHoursStart,omitempty"`
	// ScheduledInstallDays undocumented
	ScheduledInstallDays []DayOfWeek `json:"scheduledInstallDays,omitempty"`
	// UtcTimeOffsetInMinutes undocumented
	UtcTimeOffsetInMinutes *int `json:"utcTimeOffsetInMinutes,omitempty"`
}

// IOSUpdateDeviceStatus undocumented
type IOSUpdateDeviceStatus struct {
	// Entity is the base model of IOSUpdateDeviceStatus
	Entity
	// ComplianceGracePeriodExpirationDateTime undocumented
	ComplianceGracePeriodExpirationDateTime *time.Time `json:"complianceGracePeriodExpirationDateTime,omitempty"`
	// DeviceDisplayName undocumented
	DeviceDisplayName *string `json:"deviceDisplayName,omitempty"`
	// DeviceID undocumented
	DeviceID *string `json:"deviceId,omitempty"`
	// DeviceModel undocumented
	DeviceModel *string `json:"deviceModel,omitempty"`
	// InstallStatus undocumented
	InstallStatus *IOSUpdatesInstallStatus `json:"installStatus,omitempty"`
	// LastReportedDateTime undocumented
	LastReportedDateTime *time.Time `json:"lastReportedDateTime,omitempty"`
	// OsVersion undocumented
	OsVersion *string `json:"osVersion,omitempty"`
	// Status undocumented
	Status *ComplianceStatus `json:"status,omitempty"`
	// UserID undocumented
	UserID *string `json:"userId,omitempty"`
	// UserName undocumented
	UserName *string `json:"userName,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}

// IOSVPPApp undocumented
type IOSVPPApp struct {
	// MobileApp is the base model of IOSVPPApp
	MobileApp
	// ApplicableDeviceType undocumented
	ApplicableDeviceType *IOSDeviceType `json:"applicableDeviceType,omitempty"`
	// AppStoreURL undocumented
	AppStoreURL *string `json:"appStoreUrl,omitempty"`
	// BundleID undocumented
	BundleID *string `json:"bundleId,omitempty"`
	// LicensingType undocumented
	LicensingType *VPPLicensingType `json:"licensingType,omitempty"`
	// ReleaseDateTime undocumented
	ReleaseDateTime *time.Time `json:"releaseDateTime,omitempty"`
	// TotalLicenseCount undocumented
	TotalLicenseCount *int `json:"totalLicenseCount,omitempty"`
	// UsedLicenseCount undocumented
	UsedLicenseCount *int `json:"usedLicenseCount,omitempty"`
	// VPPTokenAccountType undocumented
	VPPTokenAccountType *VPPTokenAccountType `json:"vppTokenAccountType,omitempty"`
	// VPPTokenAppleID undocumented
	VPPTokenAppleID *string `json:"vppTokenAppleId,omitempty"`
	// VPPTokenOrganizationName undocumented
	VPPTokenOrganizationName *string `json:"vppTokenOrganizationName,omitempty"`
}

// IOSVPPAppAssignmentSettings undocumented
type IOSVPPAppAssignmentSettings struct {
	// MobileAppAssignmentSettings is the base model of IOSVPPAppAssignmentSettings
	MobileAppAssignmentSettings
	// UseDeviceLicensing undocumented
	UseDeviceLicensing *bool `json:"useDeviceLicensing,omitempty"`
	// VPNConfigurationID undocumented
	VPNConfigurationID *string `json:"vpnConfigurationId,omitempty"`
}

// IOSVPPEBook undocumented
type IOSVPPEBook struct {
	// ManagedEBook is the base model of IOSVPPEBook
	ManagedEBook
	// AppleID undocumented
	AppleID *string `json:"appleId,omitempty"`
	// Genres undocumented
	Genres []string `json:"genres,omitempty"`
	// Language undocumented
	Language *string `json:"language,omitempty"`
	// Seller undocumented
	Seller *string `json:"seller,omitempty"`
	// TotalLicenseCount undocumented
	TotalLicenseCount *int `json:"totalLicenseCount,omitempty"`
	// UsedLicenseCount undocumented
	UsedLicenseCount *int `json:"usedLicenseCount,omitempty"`
	// VPPOrganizationName undocumented
	VPPOrganizationName *string `json:"vppOrganizationName,omitempty"`
	// VPPTokenID undocumented
	VPPTokenID *UUID `json:"vppTokenId,omitempty"`
}

// IOSVPPEBookAssignment undocumented
type IOSVPPEBookAssignment struct {
	// ManagedEBookAssignment is the base model of IOSVPPEBookAssignment
	ManagedEBookAssignment
}
