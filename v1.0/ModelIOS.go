// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// IOSCertificateProfile undocumented
type IOSCertificateProfile struct {
	// DeviceConfiguration is the base model of IOSCertificateProfile
	DeviceConfiguration

	ODataType string `json:"@odata.type,omitempty"`
}

func NewIOSCertificateProfile() (*IOSCertificateProfile, error) {
	newIOSCertificateProfile := &IOSCertificateProfile{
		ODataType: "#microsoft.graph.IosCertificateProfile",
	}
	return newIOSCertificateProfile, nil
}

// IOSCompliancePolicy undocumented
type IOSCompliancePolicy struct {
	// DeviceCompliancePolicy is the base model of IOSCompliancePolicy
	DeviceCompliancePolicy

	ODataType string `json:"@odata.type,omitempty"`
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

func NewIOSCompliancePolicy() (*IOSCompliancePolicy, error) {
	newIOSCompliancePolicy := &IOSCompliancePolicy{
		ODataType: "#microsoft.graph.IosCompliancePolicy",
	}
	return newIOSCompliancePolicy, nil
}

// IOSCustomConfiguration undocumented
type IOSCustomConfiguration struct {
	// DeviceConfiguration is the base model of IOSCustomConfiguration
	DeviceConfiguration

	ODataType string `json:"@odata.type,omitempty"`
	// Payload undocumented
	Payload *Binary `json:"payload,omitempty"`
	// PayloadFileName undocumented
	PayloadFileName *string `json:"payloadFileName,omitempty"`
	// PayloadName undocumented
	PayloadName *string `json:"payloadName,omitempty"`
}

func NewIOSCustomConfiguration() (*IOSCustomConfiguration, error) {
	newIOSCustomConfiguration := &IOSCustomConfiguration{
		ODataType: "#microsoft.graph.IosCustomConfiguration",
	}
	return newIOSCustomConfiguration, nil
}

// IOSDeviceFeaturesConfiguration undocumented
type IOSDeviceFeaturesConfiguration struct {
	// AppleDeviceFeaturesConfigurationBase is the base model of IOSDeviceFeaturesConfiguration
	AppleDeviceFeaturesConfigurationBase

	ODataType string `json:"@odata.type,omitempty"`
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

func NewIOSDeviceFeaturesConfiguration() (*IOSDeviceFeaturesConfiguration, error) {
	newIOSDeviceFeaturesConfiguration := &IOSDeviceFeaturesConfiguration{
		ODataType: "#microsoft.graph.IosDeviceFeaturesConfiguration",
	}
	return newIOSDeviceFeaturesConfiguration, nil
}

// IOSDeviceType undocumented
type IOSDeviceType struct {
	// Object is the base model of IOSDeviceType
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// IPad undocumented
	IPad *bool `json:"iPad,omitempty"`
	// IPhoneAndIPod undocumented
	IPhoneAndIPod *bool `json:"iPhoneAndIPod,omitempty"`
}

func NewIOSDeviceType() (*IOSDeviceType, error) {
	newIOSDeviceType := &IOSDeviceType{
		ODataType: "#microsoft.graph.IosDeviceType",
	}
	return newIOSDeviceType, nil
}

// IOSGeneralDeviceConfiguration undocumented
type IOSGeneralDeviceConfiguration struct {
	// DeviceConfiguration is the base model of IOSGeneralDeviceConfiguration
	DeviceConfiguration

	ODataType string `json:"@odata.type,omitempty"`
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

func NewIOSGeneralDeviceConfiguration() (*IOSGeneralDeviceConfiguration, error) {
	newIOSGeneralDeviceConfiguration := &IOSGeneralDeviceConfiguration{
		ODataType: "#microsoft.graph.IosGeneralDeviceConfiguration",
	}
	return newIOSGeneralDeviceConfiguration, nil
}

// IOSHomeScreenApp undocumented
type IOSHomeScreenApp struct {
	// IOSHomeScreenItem is the base model of IOSHomeScreenApp
	IOSHomeScreenItem

	ODataType string `json:"@odata.type,omitempty"`
	// BundleID undocumented
	BundleID *string `json:"bundleID,omitempty"`
}

func NewIOSHomeScreenApp() (*IOSHomeScreenApp, error) {
	newIOSHomeScreenApp := &IOSHomeScreenApp{
		ODataType: "#microsoft.graph.IosHomeScreenApp",
	}
	return newIOSHomeScreenApp, nil
}

// IOSHomeScreenFolder undocumented
type IOSHomeScreenFolder struct {
	// IOSHomeScreenItem is the base model of IOSHomeScreenFolder
	IOSHomeScreenItem

	ODataType string `json:"@odata.type,omitempty"`
	// Pages undocumented
	Pages []IOSHomeScreenFolderPage `json:"pages,omitempty"`
}

func NewIOSHomeScreenFolder() (*IOSHomeScreenFolder, error) {
	newIOSHomeScreenFolder := &IOSHomeScreenFolder{
		ODataType: "#microsoft.graph.IosHomeScreenFolder",
	}
	return newIOSHomeScreenFolder, nil
}

// IOSHomeScreenFolderPage undocumented
type IOSHomeScreenFolderPage struct {
	// Object is the base model of IOSHomeScreenFolderPage
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Apps undocumented
	Apps []IOSHomeScreenApp `json:"apps,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
}

func NewIOSHomeScreenFolderPage() (*IOSHomeScreenFolderPage, error) {
	newIOSHomeScreenFolderPage := &IOSHomeScreenFolderPage{
		ODataType: "#microsoft.graph.IosHomeScreenFolderPage",
	}
	return newIOSHomeScreenFolderPage, nil
}

// IOSHomeScreenItem undocumented
type IOSHomeScreenItem struct {
	// Object is the base model of IOSHomeScreenItem
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
}

func NewIOSHomeScreenItem() (*IOSHomeScreenItem, error) {
	newIOSHomeScreenItem := &IOSHomeScreenItem{
		ODataType: "#microsoft.graph.IosHomeScreenItem",
	}
	return newIOSHomeScreenItem, nil
}

// IOSHomeScreenPage undocumented
type IOSHomeScreenPage struct {
	// Object is the base model of IOSHomeScreenPage
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Icons undocumented
	Icons []IOSHomeScreenItem `json:"icons,omitempty"`
}

func NewIOSHomeScreenPage() (*IOSHomeScreenPage, error) {
	newIOSHomeScreenPage := &IOSHomeScreenPage{
		ODataType: "#microsoft.graph.IosHomeScreenPage",
	}
	return newIOSHomeScreenPage, nil
}

// IOSLobApp undocumented
type IOSLobApp struct {
	// MobileLobApp is the base model of IOSLobApp
	MobileLobApp

	ODataType string `json:"@odata.type,omitempty"`
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

func NewIOSLobApp() (*IOSLobApp, error) {
	newIOSLobApp := &IOSLobApp{
		ODataType: "#microsoft.graph.IosLobApp",
	}
	return newIOSLobApp, nil
}

// IOSLobAppAssignmentSettings undocumented
type IOSLobAppAssignmentSettings struct {
	// MobileAppAssignmentSettings is the base model of IOSLobAppAssignmentSettings
	MobileAppAssignmentSettings

	ODataType string `json:"@odata.type,omitempty"`
	// IsRemovable undocumented
	IsRemovable *bool `json:"isRemovable,omitempty"`
	// UninstallOnDeviceRemoval undocumented
	UninstallOnDeviceRemoval *bool `json:"uninstallOnDeviceRemoval,omitempty"`
	// VPNConfigurationID undocumented
	VPNConfigurationID *string `json:"vpnConfigurationId,omitempty"`
}

func NewIOSLobAppAssignmentSettings() (*IOSLobAppAssignmentSettings, error) {
	newIOSLobAppAssignmentSettings := &IOSLobAppAssignmentSettings{
		ODataType: "#microsoft.graph.IosLobAppAssignmentSettings",
	}
	return newIOSLobAppAssignmentSettings, nil
}

// IOSLobAppProvisioningConfigurationAssignment undocumented
type IOSLobAppProvisioningConfigurationAssignment struct {
	// Entity is the base model of IOSLobAppProvisioningConfigurationAssignment
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Target undocumented
	Target *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
}

func NewIOSLobAppProvisioningConfigurationAssignment() (*IOSLobAppProvisioningConfigurationAssignment, error) {
	newIOSLobAppProvisioningConfigurationAssignment := &IOSLobAppProvisioningConfigurationAssignment{
		ODataType: "#microsoft.graph.IosLobAppProvisioningConfigurationAssignment",
	}
	return newIOSLobAppProvisioningConfigurationAssignment, nil
}

// IOSManagedAppProtection undocumented
type IOSManagedAppProtection struct {
	// TargetedManagedAppProtection is the base model of IOSManagedAppProtection
	TargetedManagedAppProtection

	ODataType string `json:"@odata.type,omitempty"`
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

func NewIOSManagedAppProtection() (*IOSManagedAppProtection, error) {
	newIOSManagedAppProtection := &IOSManagedAppProtection{
		ODataType: "#microsoft.graph.IosManagedAppProtection",
	}
	return newIOSManagedAppProtection, nil
}

// IOSManagedAppRegistration undocumented
type IOSManagedAppRegistration struct {
	// ManagedAppRegistration is the base model of IOSManagedAppRegistration
	ManagedAppRegistration

	ODataType string `json:"@odata.type,omitempty"`
}

func NewIOSManagedAppRegistration() (*IOSManagedAppRegistration, error) {
	newIOSManagedAppRegistration := &IOSManagedAppRegistration{
		ODataType: "#microsoft.graph.IosManagedAppRegistration",
	}
	return newIOSManagedAppRegistration, nil
}

// IOSMinimumOperatingSystem undocumented
type IOSMinimumOperatingSystem struct {
	// Object is the base model of IOSMinimumOperatingSystem
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// V10Underscore0 undocumented
	V10Underscore0 *bool `json:"v10_0,omitempty"`
	// V11Underscore0 undocumented
	V11Underscore0 *bool `json:"v11_0,omitempty"`
	// V12Underscore0 undocumented
	V12Underscore0 *bool `json:"v12_0,omitempty"`
	// V13Underscore0 undocumented
	V13Underscore0 *bool `json:"v13_0,omitempty"`
	// V14Underscore0 undocumented
	V14Underscore0 *bool `json:"v14_0,omitempty"`
	// V15Underscore0 undocumented
	V15Underscore0 *bool `json:"v15_0,omitempty"`
	// V8Underscore0 undocumented
	V8Underscore0 *bool `json:"v8_0,omitempty"`
	// V9Underscore0 undocumented
	V9Underscore0 *bool `json:"v9_0,omitempty"`
}

func NewIOSMinimumOperatingSystem() (*IOSMinimumOperatingSystem, error) {
	newIOSMinimumOperatingSystem := &IOSMinimumOperatingSystem{
		ODataType: "#microsoft.graph.IosMinimumOperatingSystem",
	}
	return newIOSMinimumOperatingSystem, nil
}

// IOSMobileAppConfiguration undocumented
type IOSMobileAppConfiguration struct {
	// ManagedDeviceMobileAppConfiguration is the base model of IOSMobileAppConfiguration
	ManagedDeviceMobileAppConfiguration

	ODataType string `json:"@odata.type,omitempty"`
	// EncodedSettingXML undocumented
	EncodedSettingXML *Binary `json:"encodedSettingXml,omitempty"`
	// Settings undocumented
	Settings []AppConfigurationSettingItem `json:"settings,omitempty"`
}

func NewIOSMobileAppConfiguration() (*IOSMobileAppConfiguration, error) {
	newIOSMobileAppConfiguration := &IOSMobileAppConfiguration{
		ODataType: "#microsoft.graph.IosMobileAppConfiguration",
	}
	return newIOSMobileAppConfiguration, nil
}

// IOSMobileAppIdentifier undocumented
type IOSMobileAppIdentifier struct {
	// MobileAppIdentifier is the base model of IOSMobileAppIdentifier
	MobileAppIdentifier

	ODataType string `json:"@odata.type,omitempty"`
	// BundleID undocumented
	BundleID *string `json:"bundleId,omitempty"`
}

func NewIOSMobileAppIdentifier() (*IOSMobileAppIdentifier, error) {
	newIOSMobileAppIdentifier := &IOSMobileAppIdentifier{
		ODataType: "#microsoft.graph.IosMobileAppIdentifier",
	}
	return newIOSMobileAppIdentifier, nil
}

// IOSNetworkUsageRule undocumented
type IOSNetworkUsageRule struct {
	// Object is the base model of IOSNetworkUsageRule
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// CellularDataBlocked undocumented
	CellularDataBlocked *bool `json:"cellularDataBlocked,omitempty"`
	// CellularDataBlockWhenRoaming undocumented
	CellularDataBlockWhenRoaming *bool `json:"cellularDataBlockWhenRoaming,omitempty"`
	// ManagedApps undocumented
	ManagedApps []AppListItem `json:"managedApps,omitempty"`
}

func NewIOSNetworkUsageRule() (*IOSNetworkUsageRule, error) {
	newIOSNetworkUsageRule := &IOSNetworkUsageRule{
		ODataType: "#microsoft.graph.IosNetworkUsageRule",
	}
	return newIOSNetworkUsageRule, nil
}

// IOSNotificationSettings undocumented
type IOSNotificationSettings struct {
	// Object is the base model of IOSNotificationSettings
	Object

	ODataType string `json:"@odata.type,omitempty"`
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

func NewIOSNotificationSettings() (*IOSNotificationSettings, error) {
	newIOSNotificationSettings := &IOSNotificationSettings{
		ODataType: "#microsoft.graph.IosNotificationSettings",
	}
	return newIOSNotificationSettings, nil
}

// IOSStoreApp undocumented
type IOSStoreApp struct {
	// MobileApp is the base model of IOSStoreApp
	MobileApp

	ODataType string `json:"@odata.type,omitempty"`
	// ApplicableDeviceType undocumented
	ApplicableDeviceType *IOSDeviceType `json:"applicableDeviceType,omitempty"`
	// AppStoreURL undocumented
	AppStoreURL *string `json:"appStoreUrl,omitempty"`
	// BundleID undocumented
	BundleID *string `json:"bundleId,omitempty"`
	// MinimumSupportedOperatingSystem undocumented
	MinimumSupportedOperatingSystem *IOSMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`
}

func NewIOSStoreApp() (*IOSStoreApp, error) {
	newIOSStoreApp := &IOSStoreApp{
		ODataType: "#microsoft.graph.IosStoreApp",
	}
	return newIOSStoreApp, nil
}

// IOSStoreAppAssignmentSettings undocumented
type IOSStoreAppAssignmentSettings struct {
	// MobileAppAssignmentSettings is the base model of IOSStoreAppAssignmentSettings
	MobileAppAssignmentSettings

	ODataType string `json:"@odata.type,omitempty"`
	// VPNConfigurationID undocumented
	VPNConfigurationID *string `json:"vpnConfigurationId,omitempty"`
}

func NewIOSStoreAppAssignmentSettings() (*IOSStoreAppAssignmentSettings, error) {
	newIOSStoreAppAssignmentSettings := &IOSStoreAppAssignmentSettings{
		ODataType: "#microsoft.graph.IosStoreAppAssignmentSettings",
	}
	return newIOSStoreAppAssignmentSettings, nil
}

// IOSUpdateConfiguration undocumented
type IOSUpdateConfiguration struct {
	// DeviceConfiguration is the base model of IOSUpdateConfiguration
	DeviceConfiguration

	ODataType string `json:"@odata.type,omitempty"`
	// ActiveHoursEnd undocumented
	ActiveHoursEnd *TimeOfDay `json:"activeHoursEnd,omitempty"`
	// ActiveHoursStart undocumented
	ActiveHoursStart *TimeOfDay `json:"activeHoursStart,omitempty"`
	// ScheduledInstallDays undocumented
	ScheduledInstallDays []DayOfWeek `json:"scheduledInstallDays,omitempty"`
	// UtcTimeOffsetInMinutes undocumented
	UtcTimeOffsetInMinutes *int `json:"utcTimeOffsetInMinutes,omitempty"`
}

func NewIOSUpdateConfiguration() (*IOSUpdateConfiguration, error) {
	newIOSUpdateConfiguration := &IOSUpdateConfiguration{
		ODataType: "#microsoft.graph.IosUpdateConfiguration",
	}
	return newIOSUpdateConfiguration, nil
}

// IOSUpdateDeviceStatus undocumented
type IOSUpdateDeviceStatus struct {
	// Entity is the base model of IOSUpdateDeviceStatus
	Entity

	ODataType string `json:"@odata.type,omitempty"`
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

func NewIOSUpdateDeviceStatus() (*IOSUpdateDeviceStatus, error) {
	newIOSUpdateDeviceStatus := &IOSUpdateDeviceStatus{
		ODataType: "#microsoft.graph.IosUpdateDeviceStatus",
	}
	return newIOSUpdateDeviceStatus, nil
}

// IOSVPPApp undocumented
type IOSVPPApp struct {
	// MobileApp is the base model of IOSVPPApp
	MobileApp

	ODataType string `json:"@odata.type,omitempty"`
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

func NewIOSVPPApp() (*IOSVPPApp, error) {
	newIOSVPPApp := &IOSVPPApp{
		ODataType: "#microsoft.graph.IosVppApp",
	}
	return newIOSVPPApp, nil
}

// IOSVPPAppAssignmentSettings undocumented
type IOSVPPAppAssignmentSettings struct {
	// MobileAppAssignmentSettings is the base model of IOSVPPAppAssignmentSettings
	MobileAppAssignmentSettings

	ODataType string `json:"@odata.type,omitempty"`
	// UseDeviceLicensing undocumented
	UseDeviceLicensing *bool `json:"useDeviceLicensing,omitempty"`
	// VPNConfigurationID undocumented
	VPNConfigurationID *string `json:"vpnConfigurationId,omitempty"`
}

func NewIOSVPPAppAssignmentSettings() (*IOSVPPAppAssignmentSettings, error) {
	newIOSVPPAppAssignmentSettings := &IOSVPPAppAssignmentSettings{
		ODataType: "#microsoft.graph.IosVppAppAssignmentSettings",
	}
	return newIOSVPPAppAssignmentSettings, nil
}

// IOSVPPEBook undocumented
type IOSVPPEBook struct {
	// ManagedEBook is the base model of IOSVPPEBook
	ManagedEBook

	ODataType string `json:"@odata.type,omitempty"`
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

func NewIOSVPPEBook() (*IOSVPPEBook, error) {
	newIOSVPPEBook := &IOSVPPEBook{
		ODataType: "#microsoft.graph.IosVppEBook",
	}
	return newIOSVPPEBook, nil
}

// IOSVPPEBookAssignment undocumented
type IOSVPPEBookAssignment struct {
	// ManagedEBookAssignment is the base model of IOSVPPEBookAssignment
	ManagedEBookAssignment

	ODataType string `json:"@odata.type,omitempty"`
}

func NewIOSVPPEBookAssignment() (*IOSVPPEBookAssignment, error) {
	newIOSVPPEBookAssignment := &IOSVPPEBookAssignment{
		ODataType: "#microsoft.graph.IosVppEBookAssignment",
	}
	return newIOSVPPEBookAssignment, nil
}
