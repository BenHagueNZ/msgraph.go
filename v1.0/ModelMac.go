// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// MacOSCompliancePolicy undocumented
type MacOSCompliancePolicy struct {
	// DeviceCompliancePolicy is the base model of MacOSCompliancePolicy
	DeviceCompliancePolicy

	ODataType string `json:"@odata.type,omitempty"`
	// DeviceThreatProtectionEnabled undocumented
	DeviceThreatProtectionEnabled *bool `json:"deviceThreatProtectionEnabled,omitempty"`
	// DeviceThreatProtectionRequiredSecurityLevel undocumented
	DeviceThreatProtectionRequiredSecurityLevel *DeviceThreatProtectionLevel `json:"deviceThreatProtectionRequiredSecurityLevel,omitempty"`
	// FirewallBlockAllIncoming undocumented
	FirewallBlockAllIncoming *bool `json:"firewallBlockAllIncoming,omitempty"`
	// FirewallEnabled undocumented
	FirewallEnabled *bool `json:"firewallEnabled,omitempty"`
	// FirewallEnableStealthMode undocumented
	FirewallEnableStealthMode *bool `json:"firewallEnableStealthMode,omitempty"`
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
	// SystemIntegrityProtectionEnabled undocumented
	SystemIntegrityProtectionEnabled *bool `json:"systemIntegrityProtectionEnabled,omitempty"`
}

func NewMacOSCompliancePolicy() (*MacOSCompliancePolicy, error) {
	newMacOSCompliancePolicy := &MacOSCompliancePolicy{
		ODataType: "#microsoft.graph.MacOSCompliancePolicy",
	}
	return newMacOSCompliancePolicy, nil
}

// MacOSCustomConfiguration undocumented
type MacOSCustomConfiguration struct {
	// DeviceConfiguration is the base model of MacOSCustomConfiguration
	DeviceConfiguration

	ODataType string `json:"@odata.type,omitempty"`
	// Payload undocumented
	Payload *Binary `json:"payload,omitempty"`
	// PayloadFileName undocumented
	PayloadFileName *string `json:"payloadFileName,omitempty"`
	// PayloadName undocumented
	PayloadName *string `json:"payloadName,omitempty"`
}

func NewMacOSCustomConfiguration() (*MacOSCustomConfiguration, error) {
	newMacOSCustomConfiguration := &MacOSCustomConfiguration{
		ODataType: "#microsoft.graph.MacOSCustomConfiguration",
	}
	return newMacOSCustomConfiguration, nil
}

// MacOSDeviceFeaturesConfiguration undocumented
type MacOSDeviceFeaturesConfiguration struct {
	// AppleDeviceFeaturesConfigurationBase is the base model of MacOSDeviceFeaturesConfiguration
	AppleDeviceFeaturesConfigurationBase

	ODataType string `json:"@odata.type,omitempty"`
}

func NewMacOSDeviceFeaturesConfiguration() (*MacOSDeviceFeaturesConfiguration, error) {
	newMacOSDeviceFeaturesConfiguration := &MacOSDeviceFeaturesConfiguration{
		ODataType: "#microsoft.graph.MacOSDeviceFeaturesConfiguration",
	}
	return newMacOSDeviceFeaturesConfiguration, nil
}

// MacOSGeneralDeviceConfiguration undocumented
type MacOSGeneralDeviceConfiguration struct {
	// DeviceConfiguration is the base model of MacOSGeneralDeviceConfiguration
	DeviceConfiguration

	ODataType string `json:"@odata.type,omitempty"`
	// CompliantAppListType undocumented
	CompliantAppListType *AppListType `json:"compliantAppListType,omitempty"`
	// CompliantAppsList undocumented
	CompliantAppsList []AppListItem `json:"compliantAppsList,omitempty"`
	// EmailInDomainSuffixes undocumented
	EmailInDomainSuffixes []string `json:"emailInDomainSuffixes,omitempty"`
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
	// PasswordMinutesOfInactivityBeforeScreenTimeout undocumented
	PasswordMinutesOfInactivityBeforeScreenTimeout *int `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
	// PasswordPreviousPasswordBlockCount undocumented
	PasswordPreviousPasswordBlockCount *int `json:"passwordPreviousPasswordBlockCount,omitempty"`
	// PasswordRequired undocumented
	PasswordRequired *bool `json:"passwordRequired,omitempty"`
	// PasswordRequiredType undocumented
	PasswordRequiredType *RequiredPasswordType `json:"passwordRequiredType,omitempty"`
}

func NewMacOSGeneralDeviceConfiguration() (*MacOSGeneralDeviceConfiguration, error) {
	newMacOSGeneralDeviceConfiguration := &MacOSGeneralDeviceConfiguration{
		ODataType: "#microsoft.graph.MacOSGeneralDeviceConfiguration",
	}
	return newMacOSGeneralDeviceConfiguration, nil
}

// MacOSLobApp undocumented
type MacOSLobApp struct {
	// MobileLobApp is the base model of MacOSLobApp
	MobileLobApp

	ODataType string `json:"@odata.type,omitempty"`
	// BuildNumber undocumented
	BuildNumber *string `json:"buildNumber,omitempty"`
	// BundleID undocumented
	BundleID *string `json:"bundleId,omitempty"`
	// ChildApps undocumented
	ChildApps []MacOSLobChildApp `json:"childApps,omitempty"`
	// IgnoreVersionDetection undocumented
	IgnoreVersionDetection *bool `json:"ignoreVersionDetection,omitempty"`
	// InstallAsManaged undocumented
	InstallAsManaged *bool `json:"installAsManaged,omitempty"`
	// Md5Hash undocumented
	Md5Hash []string `json:"md5Hash,omitempty"`
	// Md5HashChunkSize undocumented
	Md5HashChunkSize *int `json:"md5HashChunkSize,omitempty"`
	// MinimumSupportedOperatingSystem undocumented
	MinimumSupportedOperatingSystem *MacOSMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`
	// VersionNumber undocumented
	VersionNumber *string `json:"versionNumber,omitempty"`
}

func NewMacOSLobApp() (*MacOSLobApp, error) {
	newMacOSLobApp := &MacOSLobApp{
		ODataType: "#microsoft.graph.MacOSLobApp",
	}
	return newMacOSLobApp, nil
}

// MacOSLobChildApp undocumented
type MacOSLobChildApp struct {
	// Object is the base model of MacOSLobChildApp
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// BuildNumber undocumented
	BuildNumber *string `json:"buildNumber,omitempty"`
	// BundleID undocumented
	BundleID *string `json:"bundleId,omitempty"`
	// VersionNumber undocumented
	VersionNumber *string `json:"versionNumber,omitempty"`
}

func NewMacOSLobChildApp() (*MacOSLobChildApp, error) {
	newMacOSLobChildApp := &MacOSLobChildApp{
		ODataType: "#microsoft.graph.MacOSLobChildApp",
	}
	return newMacOSLobChildApp, nil
}

// MacOSMicrosoftEdgeApp undocumented
type MacOSMicrosoftEdgeApp struct {
	// MobileApp is the base model of MacOSMicrosoftEdgeApp
	MobileApp

	ODataType string `json:"@odata.type,omitempty"`
	// Channel undocumented
	Channel *MicrosoftEdgeChannel `json:"channel,omitempty"`
}

func NewMacOSMicrosoftEdgeApp() (*MacOSMicrosoftEdgeApp, error) {
	newMacOSMicrosoftEdgeApp := &MacOSMicrosoftEdgeApp{
		ODataType: "#microsoft.graph.MacOSMicrosoftEdgeApp",
	}
	return newMacOSMicrosoftEdgeApp, nil
}

// MacOSMinimumOperatingSystem undocumented
type MacOSMinimumOperatingSystem struct {
	// Object is the base model of MacOSMinimumOperatingSystem
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// V10Underscore10 undocumented
	V10Underscore10 *bool `json:"v10_10,omitempty"`
	// V10Underscore11 undocumented
	V10Underscore11 *bool `json:"v10_11,omitempty"`
	// V10Underscore12 undocumented
	V10Underscore12 *bool `json:"v10_12,omitempty"`
	// V10Underscore13 undocumented
	V10Underscore13 *bool `json:"v10_13,omitempty"`
	// V10Underscore14 undocumented
	V10Underscore14 *bool `json:"v10_14,omitempty"`
	// V10Underscore15 undocumented
	V10Underscore15 *bool `json:"v10_15,omitempty"`
	// V10Underscore7 undocumented
	V10Underscore7 *bool `json:"v10_7,omitempty"`
	// V10Underscore8 undocumented
	V10Underscore8 *bool `json:"v10_8,omitempty"`
	// V10Underscore9 undocumented
	V10Underscore9 *bool `json:"v10_9,omitempty"`
	// V11Underscore0 undocumented
	V11Underscore0 *bool `json:"v11_0,omitempty"`
	// V12Underscore0 undocumented
	V12Underscore0 *bool `json:"v12_0,omitempty"`
	// V13Underscore0 undocumented
	V13Underscore0 *bool `json:"v13_0,omitempty"`
}

func NewMacOSMinimumOperatingSystem() (*MacOSMinimumOperatingSystem, error) {
	newMacOSMinimumOperatingSystem := &MacOSMinimumOperatingSystem{
		ODataType: "#microsoft.graph.MacOSMinimumOperatingSystem",
	}
	return newMacOSMinimumOperatingSystem, nil
}

// MacOSOfficeSuiteApp undocumented
type MacOSOfficeSuiteApp struct {
	// MobileApp is the base model of MacOSOfficeSuiteApp
	MobileApp

	ODataType string `json:"@odata.type,omitempty"`
}

func NewMacOSOfficeSuiteApp() (*MacOSOfficeSuiteApp, error) {
	newMacOSOfficeSuiteApp := &MacOSOfficeSuiteApp{
		ODataType: "#microsoft.graph.MacOSOfficeSuiteApp",
	}
	return newMacOSOfficeSuiteApp, nil
}

// MacOsLobAppAssignmentSettings undocumented
type MacOsLobAppAssignmentSettings struct {
	// MobileAppAssignmentSettings is the base model of MacOsLobAppAssignmentSettings
	MobileAppAssignmentSettings

	ODataType string `json:"@odata.type,omitempty"`
	// UninstallOnDeviceRemoval undocumented
	UninstallOnDeviceRemoval *bool `json:"uninstallOnDeviceRemoval,omitempty"`
}

func NewMacOsLobAppAssignmentSettings() (*MacOsLobAppAssignmentSettings, error) {
	newMacOsLobAppAssignmentSettings := &MacOsLobAppAssignmentSettings{
		ODataType: "#microsoft.graph.MacOsLobAppAssignmentSettings",
	}
	return newMacOsLobAppAssignmentSettings, nil
}
