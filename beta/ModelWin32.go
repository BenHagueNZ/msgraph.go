// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Win32LobApp undocumented
type Win32LobApp struct {
	// MobileLobApp is the base model of Win32LobApp
	MobileLobApp

	ODataType string `json:"@odata.type,omitempty"`
	// AllowAvailableUninstall undocumented
	AllowAvailableUninstall *bool `json:"allowAvailableUninstall,omitempty"`
	// ApplicableArchitectures undocumented
	ApplicableArchitectures *WindowsArchitecture `json:"applicableArchitectures,omitempty"`
	// DetectionRules undocumented
	DetectionRules []Win32LobAppDetection `json:"detectionRules,omitempty"`
	// DisplayVersion undocumented
	DisplayVersion *string `json:"displayVersion,omitempty"`
	// InstallCommandLine undocumented
	InstallCommandLine *string `json:"installCommandLine,omitempty"`
	// InstallExperience undocumented
	InstallExperience *Win32LobAppInstallExperience `json:"installExperience,omitempty"`
	// MinimumCPUSpeedInMHz undocumented
	MinimumCPUSpeedInMHz *int `json:"minimumCpuSpeedInMHz,omitempty"`
	// MinimumFreeDiskSpaceInMB undocumented
	MinimumFreeDiskSpaceInMB *int `json:"minimumFreeDiskSpaceInMB,omitempty"`
	// MinimumMemoryInMB undocumented
	MinimumMemoryInMB *int `json:"minimumMemoryInMB,omitempty"`
	// MinimumNumberOfProcessors undocumented
	MinimumNumberOfProcessors *int `json:"minimumNumberOfProcessors,omitempty"`
	// MinimumSupportedOperatingSystem undocumented
	MinimumSupportedOperatingSystem *WindowsMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`
	// MinimumSupportedWindowsRelease undocumented
	MinimumSupportedWindowsRelease *string `json:"minimumSupportedWindowsRelease,omitempty"`
	// MsiInformation undocumented
	MsiInformation *Win32LobAppMsiInformation `json:"msiInformation,omitempty"`
	// RequirementRules undocumented
	RequirementRules []Win32LobAppRequirement `json:"requirementRules,omitempty"`
	// ReturnCodes undocumented
	ReturnCodes []Win32LobAppReturnCode `json:"returnCodes,omitempty"`
	// Rules undocumented
	Rules []Win32LobAppRule `json:"rules,omitempty"`
	// SetupFilePath undocumented
	SetupFilePath *string `json:"setupFilePath,omitempty"`
	// UninstallCommandLine undocumented
	UninstallCommandLine *string `json:"uninstallCommandLine,omitempty"`
}

func NewWin32LobApp() (*Win32LobApp, error) {
	newWin32LobApp := &Win32LobApp{
		ODataType: "#microsoft.graph.Win32LobApp",
	}
	return newWin32LobApp, nil
}

// Win32LobAppAssignmentSettings undocumented
type Win32LobAppAssignmentSettings struct {
	// MobileAppAssignmentSettings is the base model of Win32LobAppAssignmentSettings
	MobileAppAssignmentSettings

	ODataType string `json:"@odata.type,omitempty"`
	// DeliveryOptimizationPriority undocumented
	DeliveryOptimizationPriority *Win32LobAppDeliveryOptimizationPriority `json:"deliveryOptimizationPriority,omitempty"`
	// InstallTimeSettings undocumented
	InstallTimeSettings *MobileAppInstallTimeSettings `json:"installTimeSettings,omitempty"`
	// Notifications undocumented
	Notifications *Win32LobAppNotification `json:"notifications,omitempty"`
	// RestartSettings undocumented
	RestartSettings *Win32LobAppRestartSettings `json:"restartSettings,omitempty"`
}

func NewWin32LobAppAssignmentSettings() (*Win32LobAppAssignmentSettings, error) {
	newWin32LobAppAssignmentSettings := &Win32LobAppAssignmentSettings{
		ODataType: "#microsoft.graph.Win32LobAppAssignmentSettings",
	}
	return newWin32LobAppAssignmentSettings, nil
}

// Win32LobAppDetection undocumented
type Win32LobAppDetection struct {
	// Object is the base model of Win32LobAppDetection
	Object

	ODataType string `json:"@odata.type,omitempty"`
}

func NewWin32LobAppDetection() (*Win32LobAppDetection, error) {
	newWin32LobAppDetection := &Win32LobAppDetection{
		ODataType: "#microsoft.graph.Win32LobAppDetection",
	}
	return newWin32LobAppDetection, nil
}

// Win32LobAppFileSystemDetection undocumented
type Win32LobAppFileSystemDetection struct {
	// Win32LobAppDetection is the base model of Win32LobAppFileSystemDetection
	Win32LobAppDetection

	ODataType string `json:"@odata.type,omitempty"`
	// Check32BitOn64System undocumented
	Check32BitOn64System *bool `json:"check32BitOn64System,omitempty"`
	// DetectionType undocumented
	DetectionType *Win32LobAppFileSystemDetectionType `json:"detectionType,omitempty"`
	// DetectionValue undocumented
	DetectionValue *string `json:"detectionValue,omitempty"`
	// FileOrFolderName undocumented
	FileOrFolderName *string `json:"fileOrFolderName,omitempty"`
	// Operator undocumented
	Operator *Win32LobAppDetectionOperator `json:"operator,omitempty"`
	// Path undocumented
	Path *string `json:"path,omitempty"`
}

func NewWin32LobAppFileSystemDetection() (*Win32LobAppFileSystemDetection, error) {
	newWin32LobAppFileSystemDetection := &Win32LobAppFileSystemDetection{
		ODataType: "#microsoft.graph.Win32LobAppFileSystemDetection",
	}
	return newWin32LobAppFileSystemDetection, nil
}

// Win32LobAppFileSystemRequirement undocumented
type Win32LobAppFileSystemRequirement struct {
	// Win32LobAppRequirement is the base model of Win32LobAppFileSystemRequirement
	Win32LobAppRequirement

	ODataType string `json:"@odata.type,omitempty"`
	// Check32BitOn64System undocumented
	Check32BitOn64System *bool `json:"check32BitOn64System,omitempty"`
	// DetectionType undocumented
	DetectionType *Win32LobAppFileSystemDetectionType `json:"detectionType,omitempty"`
	// FileOrFolderName undocumented
	FileOrFolderName *string `json:"fileOrFolderName,omitempty"`
	// Path undocumented
	Path *string `json:"path,omitempty"`
}

func NewWin32LobAppFileSystemRequirement() (*Win32LobAppFileSystemRequirement, error) {
	newWin32LobAppFileSystemRequirement := &Win32LobAppFileSystemRequirement{
		ODataType: "#microsoft.graph.Win32LobAppFileSystemRequirement",
	}
	return newWin32LobAppFileSystemRequirement, nil
}

// Win32LobAppFileSystemRule undocumented
type Win32LobAppFileSystemRule struct {
	// Win32LobAppRule is the base model of Win32LobAppFileSystemRule
	Win32LobAppRule

	ODataType string `json:"@odata.type,omitempty"`
	// Check32BitOn64System undocumented
	Check32BitOn64System *bool `json:"check32BitOn64System,omitempty"`
	// ComparisonValue undocumented
	ComparisonValue *string `json:"comparisonValue,omitempty"`
	// FileOrFolderName undocumented
	FileOrFolderName *string `json:"fileOrFolderName,omitempty"`
	// OperationType undocumented
	OperationType *Win32LobAppFileSystemOperationType `json:"operationType,omitempty"`
	// Operator undocumented
	Operator *Win32LobAppRuleOperator `json:"operator,omitempty"`
	// Path undocumented
	Path *string `json:"path,omitempty"`
}

func NewWin32LobAppFileSystemRule() (*Win32LobAppFileSystemRule, error) {
	newWin32LobAppFileSystemRule := &Win32LobAppFileSystemRule{
		ODataType: "#microsoft.graph.Win32LobAppFileSystemRule",
	}
	return newWin32LobAppFileSystemRule, nil
}

// Win32LobAppInstallExperience undocumented
type Win32LobAppInstallExperience struct {
	// Object is the base model of Win32LobAppInstallExperience
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// DeviceRestartBehavior undocumented
	DeviceRestartBehavior *Win32LobAppRestartBehavior `json:"deviceRestartBehavior,omitempty"`
	// RunAsAccount undocumented
	RunAsAccount *RunAsAccountType `json:"runAsAccount,omitempty"`
}

func NewWin32LobAppInstallExperience() (*Win32LobAppInstallExperience, error) {
	newWin32LobAppInstallExperience := &Win32LobAppInstallExperience{
		ODataType: "#microsoft.graph.Win32LobAppInstallExperience",
	}
	return newWin32LobAppInstallExperience, nil
}

// Win32LobAppMsiInformation undocumented
type Win32LobAppMsiInformation struct {
	// Object is the base model of Win32LobAppMsiInformation
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// PackageType undocumented
	PackageType *Win32LobAppMsiPackageType `json:"packageType,omitempty"`
	// ProductCode undocumented
	ProductCode *string `json:"productCode,omitempty"`
	// ProductName undocumented
	ProductName *string `json:"productName,omitempty"`
	// ProductVersion undocumented
	ProductVersion *string `json:"productVersion,omitempty"`
	// Publisher undocumented
	Publisher *string `json:"publisher,omitempty"`
	// RequiresReboot undocumented
	RequiresReboot *bool `json:"requiresReboot,omitempty"`
	// UpgradeCode undocumented
	UpgradeCode *string `json:"upgradeCode,omitempty"`
}

func NewWin32LobAppMsiInformation() (*Win32LobAppMsiInformation, error) {
	newWin32LobAppMsiInformation := &Win32LobAppMsiInformation{
		ODataType: "#microsoft.graph.Win32LobAppMsiInformation",
	}
	return newWin32LobAppMsiInformation, nil
}

// Win32LobAppPowerShellScriptDetection undocumented
type Win32LobAppPowerShellScriptDetection struct {
	// Win32LobAppDetection is the base model of Win32LobAppPowerShellScriptDetection
	Win32LobAppDetection

	ODataType string `json:"@odata.type,omitempty"`
	// EnforceSignatureCheck undocumented
	EnforceSignatureCheck *bool `json:"enforceSignatureCheck,omitempty"`
	// RunAs32Bit undocumented
	RunAs32Bit *bool `json:"runAs32Bit,omitempty"`
	// ScriptContent undocumented
	ScriptContent *string `json:"scriptContent,omitempty"`
}

func NewWin32LobAppPowerShellScriptDetection() (*Win32LobAppPowerShellScriptDetection, error) {
	newWin32LobAppPowerShellScriptDetection := &Win32LobAppPowerShellScriptDetection{
		ODataType: "#microsoft.graph.Win32LobAppPowerShellScriptDetection",
	}
	return newWin32LobAppPowerShellScriptDetection, nil
}

// Win32LobAppPowerShellScriptRequirement undocumented
type Win32LobAppPowerShellScriptRequirement struct {
	// Win32LobAppRequirement is the base model of Win32LobAppPowerShellScriptRequirement
	Win32LobAppRequirement

	ODataType string `json:"@odata.type,omitempty"`
	// DetectionType undocumented
	DetectionType *Win32LobAppPowerShellScriptDetectionType `json:"detectionType,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// EnforceSignatureCheck undocumented
	EnforceSignatureCheck *bool `json:"enforceSignatureCheck,omitempty"`
	// RunAs32Bit undocumented
	RunAs32Bit *bool `json:"runAs32Bit,omitempty"`
	// RunAsAccount undocumented
	RunAsAccount *RunAsAccountType `json:"runAsAccount,omitempty"`
	// ScriptContent undocumented
	ScriptContent *string `json:"scriptContent,omitempty"`
}

func NewWin32LobAppPowerShellScriptRequirement() (*Win32LobAppPowerShellScriptRequirement, error) {
	newWin32LobAppPowerShellScriptRequirement := &Win32LobAppPowerShellScriptRequirement{
		ODataType: "#microsoft.graph.Win32LobAppPowerShellScriptRequirement",
	}
	return newWin32LobAppPowerShellScriptRequirement, nil
}

// Win32LobAppPowerShellScriptRule undocumented
type Win32LobAppPowerShellScriptRule struct {
	// Win32LobAppRule is the base model of Win32LobAppPowerShellScriptRule
	Win32LobAppRule

	ODataType string `json:"@odata.type,omitempty"`
	// ComparisonValue undocumented
	ComparisonValue *string `json:"comparisonValue,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// EnforceSignatureCheck undocumented
	EnforceSignatureCheck *bool `json:"enforceSignatureCheck,omitempty"`
	// OperationType undocumented
	OperationType *Win32LobAppPowerShellScriptRuleOperationType `json:"operationType,omitempty"`
	// Operator undocumented
	Operator *Win32LobAppRuleOperator `json:"operator,omitempty"`
	// RunAs32Bit undocumented
	RunAs32Bit *bool `json:"runAs32Bit,omitempty"`
	// RunAsAccount undocumented
	RunAsAccount *RunAsAccountType `json:"runAsAccount,omitempty"`
	// ScriptContent undocumented
	ScriptContent *string `json:"scriptContent,omitempty"`
}

func NewWin32LobAppPowerShellScriptRule() (*Win32LobAppPowerShellScriptRule, error) {
	newWin32LobAppPowerShellScriptRule := &Win32LobAppPowerShellScriptRule{
		ODataType: "#microsoft.graph.Win32LobAppPowerShellScriptRule",
	}
	return newWin32LobAppPowerShellScriptRule, nil
}

// Win32LobAppProductCodeDetection undocumented
type Win32LobAppProductCodeDetection struct {
	// Win32LobAppDetection is the base model of Win32LobAppProductCodeDetection
	Win32LobAppDetection

	ODataType string `json:"@odata.type,omitempty"`
	// ProductCode undocumented
	ProductCode *string `json:"productCode,omitempty"`
	// ProductVersion undocumented
	ProductVersion *string `json:"productVersion,omitempty"`
	// ProductVersionOperator undocumented
	ProductVersionOperator *Win32LobAppDetectionOperator `json:"productVersionOperator,omitempty"`
}

func NewWin32LobAppProductCodeDetection() (*Win32LobAppProductCodeDetection, error) {
	newWin32LobAppProductCodeDetection := &Win32LobAppProductCodeDetection{
		ODataType: "#microsoft.graph.Win32LobAppProductCodeDetection",
	}
	return newWin32LobAppProductCodeDetection, nil
}

// Win32LobAppProductCodeRule undocumented
type Win32LobAppProductCodeRule struct {
	// Win32LobAppRule is the base model of Win32LobAppProductCodeRule
	Win32LobAppRule

	ODataType string `json:"@odata.type,omitempty"`
	// ProductCode undocumented
	ProductCode *string `json:"productCode,omitempty"`
	// ProductVersion undocumented
	ProductVersion *string `json:"productVersion,omitempty"`
	// ProductVersionOperator undocumented
	ProductVersionOperator *Win32LobAppRuleOperator `json:"productVersionOperator,omitempty"`
}

func NewWin32LobAppProductCodeRule() (*Win32LobAppProductCodeRule, error) {
	newWin32LobAppProductCodeRule := &Win32LobAppProductCodeRule{
		ODataType: "#microsoft.graph.Win32LobAppProductCodeRule",
	}
	return newWin32LobAppProductCodeRule, nil
}

// Win32LobAppRegistryDetection undocumented
type Win32LobAppRegistryDetection struct {
	// Win32LobAppDetection is the base model of Win32LobAppRegistryDetection
	Win32LobAppDetection

	ODataType string `json:"@odata.type,omitempty"`
	// Check32BitOn64System undocumented
	Check32BitOn64System *bool `json:"check32BitOn64System,omitempty"`
	// DetectionType undocumented
	DetectionType *Win32LobAppRegistryDetectionType `json:"detectionType,omitempty"`
	// DetectionValue undocumented
	DetectionValue *string `json:"detectionValue,omitempty"`
	// KeyPath undocumented
	KeyPath *string `json:"keyPath,omitempty"`
	// Operator undocumented
	Operator *Win32LobAppDetectionOperator `json:"operator,omitempty"`
	// ValueName undocumented
	ValueName *string `json:"valueName,omitempty"`
}

func NewWin32LobAppRegistryDetection() (*Win32LobAppRegistryDetection, error) {
	newWin32LobAppRegistryDetection := &Win32LobAppRegistryDetection{
		ODataType: "#microsoft.graph.Win32LobAppRegistryDetection",
	}
	return newWin32LobAppRegistryDetection, nil
}

// Win32LobAppRegistryRequirement undocumented
type Win32LobAppRegistryRequirement struct {
	// Win32LobAppRequirement is the base model of Win32LobAppRegistryRequirement
	Win32LobAppRequirement

	ODataType string `json:"@odata.type,omitempty"`
	// Check32BitOn64System undocumented
	Check32BitOn64System *bool `json:"check32BitOn64System,omitempty"`
	// DetectionType undocumented
	DetectionType *Win32LobAppRegistryDetectionType `json:"detectionType,omitempty"`
	// KeyPath undocumented
	KeyPath *string `json:"keyPath,omitempty"`
	// ValueName undocumented
	ValueName *string `json:"valueName,omitempty"`
}

func NewWin32LobAppRegistryRequirement() (*Win32LobAppRegistryRequirement, error) {
	newWin32LobAppRegistryRequirement := &Win32LobAppRegistryRequirement{
		ODataType: "#microsoft.graph.Win32LobAppRegistryRequirement",
	}
	return newWin32LobAppRegistryRequirement, nil
}

// Win32LobAppRegistryRule undocumented
type Win32LobAppRegistryRule struct {
	// Win32LobAppRule is the base model of Win32LobAppRegistryRule
	Win32LobAppRule

	ODataType string `json:"@odata.type,omitempty"`
	// Check32BitOn64System undocumented
	Check32BitOn64System *bool `json:"check32BitOn64System,omitempty"`
	// ComparisonValue undocumented
	ComparisonValue *string `json:"comparisonValue,omitempty"`
	// KeyPath undocumented
	KeyPath *string `json:"keyPath,omitempty"`
	// OperationType undocumented
	OperationType *Win32LobAppRegistryRuleOperationType `json:"operationType,omitempty"`
	// Operator undocumented
	Operator *Win32LobAppRuleOperator `json:"operator,omitempty"`
	// ValueName undocumented
	ValueName *string `json:"valueName,omitempty"`
}

func NewWin32LobAppRegistryRule() (*Win32LobAppRegistryRule, error) {
	newWin32LobAppRegistryRule := &Win32LobAppRegistryRule{
		ODataType: "#microsoft.graph.Win32LobAppRegistryRule",
	}
	return newWin32LobAppRegistryRule, nil
}

// Win32LobAppRequirement undocumented
type Win32LobAppRequirement struct {
	// Object is the base model of Win32LobAppRequirement
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// DetectionValue undocumented
	DetectionValue *string `json:"detectionValue,omitempty"`
	// Operator undocumented
	Operator *Win32LobAppDetectionOperator `json:"operator,omitempty"`
}

func NewWin32LobAppRequirement() (*Win32LobAppRequirement, error) {
	newWin32LobAppRequirement := &Win32LobAppRequirement{
		ODataType: "#microsoft.graph.Win32LobAppRequirement",
	}
	return newWin32LobAppRequirement, nil
}

// Win32LobAppRestartSettings undocumented
type Win32LobAppRestartSettings struct {
	// Object is the base model of Win32LobAppRestartSettings
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// CountdownDisplayBeforeRestartInMinutes undocumented
	CountdownDisplayBeforeRestartInMinutes *int `json:"countdownDisplayBeforeRestartInMinutes,omitempty"`
	// GracePeriodInMinutes undocumented
	GracePeriodInMinutes *int `json:"gracePeriodInMinutes,omitempty"`
	// RestartNotificationSnoozeDurationInMinutes undocumented
	RestartNotificationSnoozeDurationInMinutes *int `json:"restartNotificationSnoozeDurationInMinutes,omitempty"`
}

func NewWin32LobAppRestartSettings() (*Win32LobAppRestartSettings, error) {
	newWin32LobAppRestartSettings := &Win32LobAppRestartSettings{
		ODataType: "#microsoft.graph.Win32LobAppRestartSettings",
	}
	return newWin32LobAppRestartSettings, nil
}

// Win32LobAppReturnCode undocumented
type Win32LobAppReturnCode struct {
	// Object is the base model of Win32LobAppReturnCode
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ReturnCode undocumented
	ReturnCode *int `json:"returnCode,omitempty"`
	// Type undocumented
	Type *Win32LobAppReturnCodeType `json:"type,omitempty"`
}

func NewWin32LobAppReturnCode() (*Win32LobAppReturnCode, error) {
	newWin32LobAppReturnCode := &Win32LobAppReturnCode{
		ODataType: "#microsoft.graph.Win32LobAppReturnCode",
	}
	return newWin32LobAppReturnCode, nil
}

// Win32LobAppRule undocumented
type Win32LobAppRule struct {
	// Object is the base model of Win32LobAppRule
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// RuleType undocumented
	RuleType *Win32LobAppRuleType `json:"ruleType,omitempty"`
}

func NewWin32LobAppRule() (*Win32LobAppRule, error) {
	newWin32LobAppRule := &Win32LobAppRule{
		ODataType: "#microsoft.graph.Win32LobAppRule",
	}
	return newWin32LobAppRule, nil
}