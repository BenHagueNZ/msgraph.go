// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Win32LobApp Contains properties and inherited properties for Win32 apps.
type Win32LobApp struct {
	// MobileLobApp is the base model of Win32LobApp
	MobileLobApp
	// ApplicableArchitectures The Windows architecture(s) for which this app can run on.
	ApplicableArchitectures *WindowsArchitecture `json:"applicableArchitectures,omitempty"`
	// InstallCommandLine The command line to install this app
	InstallCommandLine *string `json:"installCommandLine,omitempty"`
	// InstallExperience The install experience for this app.
	InstallExperience *Win32LobAppInstallExperience `json:"installExperience,omitempty"`
	// MinimumCPUSpeedInMHz The value for the minimum CPU speed which is required to install this app.
	MinimumCPUSpeedInMHz *int `json:"minimumCpuSpeedInMHz,omitempty"`
	// MinimumFreeDiskSpaceInMB The value for the minimum free disk space which is required to install this app.
	MinimumFreeDiskSpaceInMB *int `json:"minimumFreeDiskSpaceInMB,omitempty"`
	// MinimumMemoryInMB The value for the minimum physical memory which is required to install this app.
	MinimumMemoryInMB *int `json:"minimumMemoryInMB,omitempty"`
	// MinimumNumberOfProcessors The value for the minimum number of processors which is required to install this app.
	MinimumNumberOfProcessors *int `json:"minimumNumberOfProcessors,omitempty"`
	// MinimumSupportedWindowsRelease The value for the minimum supported windows release.
	MinimumSupportedWindowsRelease *string `json:"minimumSupportedWindowsRelease,omitempty"`
	// MsiInformation The MSI details if this Win32 app is an MSI app.
	MsiInformation *Win32LobAppMsiInformation `json:"msiInformation,omitempty"`
	// ReturnCodes The return codes for post installation behavior.
	ReturnCodes []Win32LobAppReturnCode `json:"returnCodes,omitempty"`
	// Rules The detection and requirement rules for this app.
	Rules []Win32LobAppRule `json:"rules,omitempty"`
	// SetupFilePath The relative path of the setup file in the encrypted Win32LobApp package.
	SetupFilePath *string `json:"setupFilePath,omitempty"`
	// UninstallCommandLine The command line to uninstall this app
	UninstallCommandLine *string `json:"uninstallCommandLine,omitempty"`
}

// Win32LobAppAssignmentSettings Contains properties used to assign an Win32 LOB mobile app to a group.
type Win32LobAppAssignmentSettings struct {
	// MobileAppAssignmentSettings is the base model of Win32LobAppAssignmentSettings
	MobileAppAssignmentSettings
	// DeliveryOptimizationPriority The delivery optimization priority for this app assignment. This setting is not supported in National Cloud environments.
	DeliveryOptimizationPriority *Win32LobAppDeliveryOptimizationPriority `json:"deliveryOptimizationPriority,omitempty"`
	// InstallTimeSettings The install time settings to apply for this app assignment.
	InstallTimeSettings *MobileAppInstallTimeSettings `json:"installTimeSettings,omitempty"`
	// Notifications The notification status for this app assignment.
	Notifications *Win32LobAppNotification `json:"notifications,omitempty"`
	// RestartSettings The reboot settings to apply for this app assignment.
	RestartSettings *Win32LobAppRestartSettings `json:"restartSettings,omitempty"`
}

// Win32LobAppFileSystemRule A complex type to store file or folder rule data for a Win32 LOB app.
type Win32LobAppFileSystemRule struct {
	// Win32LobAppRule is the base model of Win32LobAppFileSystemRule
	Win32LobAppRule
	// Check32BitOn64System A value indicating whether to expand environment variables in the 32-bit context on 64-bit systems.
	Check32BitOn64System *bool `json:"check32BitOn64System,omitempty"`
	// ComparisonValue The file or folder comparison value.
	ComparisonValue *string `json:"comparisonValue,omitempty"`
	// FileOrFolderName The file or folder name to look up.
	FileOrFolderName *string `json:"fileOrFolderName,omitempty"`
	// OperationType The file system operation type.
	OperationType *Win32LobAppFileSystemOperationType `json:"operationType,omitempty"`
	// Operator The operator for file or folder detection.
	Operator *Win32LobAppRuleOperator `json:"operator,omitempty"`
	// Path The file or folder path to look up.
	Path *string `json:"path,omitempty"`
}

// Win32LobAppInstallExperience Contains installation experience properties for a Win32 App
type Win32LobAppInstallExperience struct {
	// Object is the base model of Win32LobAppInstallExperience
	Object
	// DeviceRestartBehavior Device restart behavior.
	DeviceRestartBehavior *Win32LobAppRestartBehavior `json:"deviceRestartBehavior,omitempty"`
	// RunAsAccount Indicates the type of execution context the app runs in.
	RunAsAccount *RunAsAccountType `json:"runAsAccount,omitempty"`
}

// Win32LobAppMsiInformation Contains MSI app properties for a Win32 App.
type Win32LobAppMsiInformation struct {
	// Object is the base model of Win32LobAppMsiInformation
	Object
	// PackageType The MSI package type.
	PackageType *Win32LobAppMsiPackageType `json:"packageType,omitempty"`
	// ProductCode The MSI product code.
	ProductCode *string `json:"productCode,omitempty"`
	// ProductName The MSI product name.
	ProductName *string `json:"productName,omitempty"`
	// ProductVersion The MSI product version.
	ProductVersion *string `json:"productVersion,omitempty"`
	// Publisher The MSI publisher.
	Publisher *string `json:"publisher,omitempty"`
	// RequiresReboot Whether the MSI app requires the machine to reboot to complete installation.
	RequiresReboot *bool `json:"requiresReboot,omitempty"`
	// UpgradeCode The MSI upgrade code.
	UpgradeCode *string `json:"upgradeCode,omitempty"`
}

// Win32LobAppPowerShellScriptRule A complex type to store the PowerShell script rule data for a Win32 LOB app.
type Win32LobAppPowerShellScriptRule struct {
	// Win32LobAppRule is the base model of Win32LobAppPowerShellScriptRule
	Win32LobAppRule
	// ComparisonValue The script output comparison value. Do not specify a value if the rule is used for detection.
	ComparisonValue *string `json:"comparisonValue,omitempty"`
	// DisplayName The display name for the rule. Do not specify this value if the rule is used for detection.
	DisplayName *string `json:"displayName,omitempty"`
	// EnforceSignatureCheck A value indicating whether a signature check is enforced.
	EnforceSignatureCheck *bool `json:"enforceSignatureCheck,omitempty"`
	// OperationType The script output comparison operation type. Use NotConfigured (the default value) if the rule is used for detection.
	OperationType *Win32LobAppPowerShellScriptRuleOperationType `json:"operationType,omitempty"`
	// Operator The script output operator. Use NotConfigured (the default value) if the rule is used for detection.
	Operator *Win32LobAppRuleOperator `json:"operator,omitempty"`
	// RunAs32Bit A value indicating whether the script should run as 32-bit.
	RunAs32Bit *bool `json:"runAs32Bit,omitempty"`
	// RunAsAccount The execution context of the script. Do not specify this value if the rule is used for detection. Script detection rules will run in the same context as the associated app install context.
	RunAsAccount *RunAsAccountType `json:"runAsAccount,omitempty"`
	// ScriptContent The base64-encoded script content.
	ScriptContent *string `json:"scriptContent,omitempty"`
}

// Win32LobAppProductCodeRule A complex type to store the product code and version rule data for a Win32 LOB app. This rule is not supported as a requirement rule.
type Win32LobAppProductCodeRule struct {
	// Win32LobAppRule is the base model of Win32LobAppProductCodeRule
	Win32LobAppRule
	// ProductCode The product code of the app.
	ProductCode *string `json:"productCode,omitempty"`
	// ProductVersion The product version comparison value.
	ProductVersion *string `json:"productVersion,omitempty"`
	// ProductVersionOperator The product version comparison operator.
	ProductVersionOperator *Win32LobAppRuleOperator `json:"productVersionOperator,omitempty"`
}

// Win32LobAppRegistryRule A complex type to store registry rule data for a Win32 LOB app.
type Win32LobAppRegistryRule struct {
	// Win32LobAppRule is the base model of Win32LobAppRegistryRule
	Win32LobAppRule
	// Check32BitOn64System A value indicating whether to search the 32-bit registry on 64-bit systems.
	Check32BitOn64System *bool `json:"check32BitOn64System,omitempty"`
	// ComparisonValue The registry comparison value.
	ComparisonValue *string `json:"comparisonValue,omitempty"`
	// KeyPath The full path of the registry entry containing the value to detect.
	KeyPath *string `json:"keyPath,omitempty"`
	// OperationType The registry operation type.
	OperationType *Win32LobAppRegistryRuleOperationType `json:"operationType,omitempty"`
	// Operator The operator for registry detection.
	Operator *Win32LobAppRuleOperator `json:"operator,omitempty"`
	// ValueName The name of the registry value to detect.
	ValueName *string `json:"valueName,omitempty"`
}

// Win32LobAppRestartSettings Contains properties describing restart coordination following an app installation.
type Win32LobAppRestartSettings struct {
	// Object is the base model of Win32LobAppRestartSettings
	Object
	// CountdownDisplayBeforeRestartInMinutes The number of minutes before the restart time to display the countdown dialog for pending restarts.
	CountdownDisplayBeforeRestartInMinutes *int `json:"countdownDisplayBeforeRestartInMinutes,omitempty"`
	// GracePeriodInMinutes The number of minutes to wait before restarting the device after an app installation.
	GracePeriodInMinutes *int `json:"gracePeriodInMinutes,omitempty"`
	// RestartNotificationSnoozeDurationInMinutes The number of minutes to snooze the restart notification dialog when the snooze button is selected.
	RestartNotificationSnoozeDurationInMinutes *int `json:"restartNotificationSnoozeDurationInMinutes,omitempty"`
}

// Win32LobAppReturnCode Contains return code properties for a Win32 App
type Win32LobAppReturnCode struct {
	// Object is the base model of Win32LobAppReturnCode
	Object
	// ReturnCode Return code.
	ReturnCode *int `json:"returnCode,omitempty"`
	// Type The type of return code.
	Type *Win32LobAppReturnCodeType `json:"type,omitempty"`
}

// Win32LobAppRule A base complex type to store the detection or requirement rule data for a Win32 LOB app.
type Win32LobAppRule struct {
	// Object is the base model of Win32LobAppRule
	Object
	// RuleType The rule type indicating the purpose of the rule.
	RuleType *Win32LobAppRuleType `json:"ruleType,omitempty"`
}
