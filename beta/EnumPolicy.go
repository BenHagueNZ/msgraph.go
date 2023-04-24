// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// PolicyPlatformType undocumented
type PolicyPlatformType string

const (
	// PolicyPlatformTypeVAndroid undocumented
	PolicyPlatformTypeVAndroid PolicyPlatformType = "android"
	// PolicyPlatformTypeVAndroidForWork undocumented
	PolicyPlatformTypeVAndroidForWork PolicyPlatformType = "androidForWork"
	// PolicyPlatformTypeVIOS undocumented
	PolicyPlatformTypeVIOS PolicyPlatformType = "iOS"
	// PolicyPlatformTypeVMacOS undocumented
	PolicyPlatformTypeVMacOS PolicyPlatformType = "macOS"
	// PolicyPlatformTypeVWindowsPhone81 undocumented
	PolicyPlatformTypeVWindowsPhone81 PolicyPlatformType = "windowsPhone81"
	// PolicyPlatformTypeVWindows81AndLater undocumented
	PolicyPlatformTypeVWindows81AndLater PolicyPlatformType = "windows81AndLater"
	// PolicyPlatformTypeVWindows10AndLater undocumented
	PolicyPlatformTypeVWindows10AndLater PolicyPlatformType = "windows10AndLater"
	// PolicyPlatformTypeVAndroidWorkProfile undocumented
	PolicyPlatformTypeVAndroidWorkProfile PolicyPlatformType = "androidWorkProfile"
	// PolicyPlatformTypeVWindows10XProfile undocumented
	PolicyPlatformTypeVWindows10XProfile PolicyPlatformType = "windows10XProfile"
	// PolicyPlatformTypeVAndroidAOSP undocumented
	PolicyPlatformTypeVAndroidAOSP PolicyPlatformType = "androidAOSP"
	// PolicyPlatformTypeVAll undocumented
	PolicyPlatformTypeVAll PolicyPlatformType = "all"
)

var (
	// PolicyPlatformTypePAndroid is a pointer to PolicyPlatformTypeVAndroid
	PolicyPlatformTypePAndroid = &_PolicyPlatformTypePAndroid
	// PolicyPlatformTypePAndroidForWork is a pointer to PolicyPlatformTypeVAndroidForWork
	PolicyPlatformTypePAndroidForWork = &_PolicyPlatformTypePAndroidForWork
	// PolicyPlatformTypePIOS is a pointer to PolicyPlatformTypeVIOS
	PolicyPlatformTypePIOS = &_PolicyPlatformTypePIOS
	// PolicyPlatformTypePMacOS is a pointer to PolicyPlatformTypeVMacOS
	PolicyPlatformTypePMacOS = &_PolicyPlatformTypePMacOS
	// PolicyPlatformTypePWindowsPhone81 is a pointer to PolicyPlatformTypeVWindowsPhone81
	PolicyPlatformTypePWindowsPhone81 = &_PolicyPlatformTypePWindowsPhone81
	// PolicyPlatformTypePWindows81AndLater is a pointer to PolicyPlatformTypeVWindows81AndLater
	PolicyPlatformTypePWindows81AndLater = &_PolicyPlatformTypePWindows81AndLater
	// PolicyPlatformTypePWindows10AndLater is a pointer to PolicyPlatformTypeVWindows10AndLater
	PolicyPlatformTypePWindows10AndLater = &_PolicyPlatformTypePWindows10AndLater
	// PolicyPlatformTypePAndroidWorkProfile is a pointer to PolicyPlatformTypeVAndroidWorkProfile
	PolicyPlatformTypePAndroidWorkProfile = &_PolicyPlatformTypePAndroidWorkProfile
	// PolicyPlatformTypePWindows10XProfile is a pointer to PolicyPlatformTypeVWindows10XProfile
	PolicyPlatformTypePWindows10XProfile = &_PolicyPlatformTypePWindows10XProfile
	// PolicyPlatformTypePAndroidAOSP is a pointer to PolicyPlatformTypeVAndroidAOSP
	PolicyPlatformTypePAndroidAOSP = &_PolicyPlatformTypePAndroidAOSP
	// PolicyPlatformTypePAll is a pointer to PolicyPlatformTypeVAll
	PolicyPlatformTypePAll = &_PolicyPlatformTypePAll
)

var (
	_PolicyPlatformTypePAndroid            = PolicyPlatformTypeVAndroid
	_PolicyPlatformTypePAndroidForWork     = PolicyPlatformTypeVAndroidForWork
	_PolicyPlatformTypePIOS                = PolicyPlatformTypeVIOS
	_PolicyPlatformTypePMacOS              = PolicyPlatformTypeVMacOS
	_PolicyPlatformTypePWindowsPhone81     = PolicyPlatformTypeVWindowsPhone81
	_PolicyPlatformTypePWindows81AndLater  = PolicyPlatformTypeVWindows81AndLater
	_PolicyPlatformTypePWindows10AndLater  = PolicyPlatformTypeVWindows10AndLater
	_PolicyPlatformTypePAndroidWorkProfile = PolicyPlatformTypeVAndroidWorkProfile
	_PolicyPlatformTypePWindows10XProfile  = PolicyPlatformTypeVWindows10XProfile
	_PolicyPlatformTypePAndroidAOSP        = PolicyPlatformTypeVAndroidAOSP
	_PolicyPlatformTypePAll                = PolicyPlatformTypeVAll
)

// PolicyScope undocumented
type PolicyScope string

const (
	// PolicyScopeVNone undocumented
	PolicyScopeVNone PolicyScope = "none"
	// PolicyScopeVAll undocumented
	PolicyScopeVAll PolicyScope = "all"
	// PolicyScopeVSelected undocumented
	PolicyScopeVSelected PolicyScope = "selected"
	// PolicyScopeVUnknownFutureValue undocumented
	PolicyScopeVUnknownFutureValue PolicyScope = "unknownFutureValue"
)

var (
	// PolicyScopePNone is a pointer to PolicyScopeVNone
	PolicyScopePNone = &_PolicyScopePNone
	// PolicyScopePAll is a pointer to PolicyScopeVAll
	PolicyScopePAll = &_PolicyScopePAll
	// PolicyScopePSelected is a pointer to PolicyScopeVSelected
	PolicyScopePSelected = &_PolicyScopePSelected
	// PolicyScopePUnknownFutureValue is a pointer to PolicyScopeVUnknownFutureValue
	PolicyScopePUnknownFutureValue = &_PolicyScopePUnknownFutureValue
)

var (
	_PolicyScopePNone               = PolicyScopeVNone
	_PolicyScopePAll                = PolicyScopeVAll
	_PolicyScopePSelected           = PolicyScopeVSelected
	_PolicyScopePUnknownFutureValue = PolicyScopeVUnknownFutureValue
)

// PolicySetStatus undocumented
type PolicySetStatus string

const (
	// PolicySetStatusVUnknown undocumented
	PolicySetStatusVUnknown PolicySetStatus = "unknown"
	// PolicySetStatusVValidating undocumented
	PolicySetStatusVValidating PolicySetStatus = "validating"
	// PolicySetStatusVPartialSuccess undocumented
	PolicySetStatusVPartialSuccess PolicySetStatus = "partialSuccess"
	// PolicySetStatusVSuccess undocumented
	PolicySetStatusVSuccess PolicySetStatus = "success"
	// PolicySetStatusVError undocumented
	PolicySetStatusVError PolicySetStatus = "error"
	// PolicySetStatusVNotAssigned undocumented
	PolicySetStatusVNotAssigned PolicySetStatus = "notAssigned"
)

var (
	// PolicySetStatusPUnknown is a pointer to PolicySetStatusVUnknown
	PolicySetStatusPUnknown = &_PolicySetStatusPUnknown
	// PolicySetStatusPValidating is a pointer to PolicySetStatusVValidating
	PolicySetStatusPValidating = &_PolicySetStatusPValidating
	// PolicySetStatusPPartialSuccess is a pointer to PolicySetStatusVPartialSuccess
	PolicySetStatusPPartialSuccess = &_PolicySetStatusPPartialSuccess
	// PolicySetStatusPSuccess is a pointer to PolicySetStatusVSuccess
	PolicySetStatusPSuccess = &_PolicySetStatusPSuccess
	// PolicySetStatusPError is a pointer to PolicySetStatusVError
	PolicySetStatusPError = &_PolicySetStatusPError
	// PolicySetStatusPNotAssigned is a pointer to PolicySetStatusVNotAssigned
	PolicySetStatusPNotAssigned = &_PolicySetStatusPNotAssigned
)

var (
	_PolicySetStatusPUnknown        = PolicySetStatusVUnknown
	_PolicySetStatusPValidating     = PolicySetStatusVValidating
	_PolicySetStatusPPartialSuccess = PolicySetStatusVPartialSuccess
	_PolicySetStatusPSuccess        = PolicySetStatusVSuccess
	_PolicySetStatusPError          = PolicySetStatusVError
	_PolicySetStatusPNotAssigned    = PolicySetStatusVNotAssigned
)
