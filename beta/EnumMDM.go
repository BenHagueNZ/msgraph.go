// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// MDMAppConfigKeyType undocumented
type MDMAppConfigKeyType string

const (
	// MDMAppConfigKeyTypeVStringType undocumented
	MDMAppConfigKeyTypeVStringType MDMAppConfigKeyType = "stringType"
	// MDMAppConfigKeyTypeVIntegerType undocumented
	MDMAppConfigKeyTypeVIntegerType MDMAppConfigKeyType = "integerType"
	// MDMAppConfigKeyTypeVRealType undocumented
	MDMAppConfigKeyTypeVRealType MDMAppConfigKeyType = "realType"
	// MDMAppConfigKeyTypeVBooleanType undocumented
	MDMAppConfigKeyTypeVBooleanType MDMAppConfigKeyType = "booleanType"
	// MDMAppConfigKeyTypeVTokenType undocumented
	MDMAppConfigKeyTypeVTokenType MDMAppConfigKeyType = "tokenType"
)

var (
	// MDMAppConfigKeyTypePStringType is a pointer to MDMAppConfigKeyTypeVStringType
	MDMAppConfigKeyTypePStringType = &_MDMAppConfigKeyTypePStringType
	// MDMAppConfigKeyTypePIntegerType is a pointer to MDMAppConfigKeyTypeVIntegerType
	MDMAppConfigKeyTypePIntegerType = &_MDMAppConfigKeyTypePIntegerType
	// MDMAppConfigKeyTypePRealType is a pointer to MDMAppConfigKeyTypeVRealType
	MDMAppConfigKeyTypePRealType = &_MDMAppConfigKeyTypePRealType
	// MDMAppConfigKeyTypePBooleanType is a pointer to MDMAppConfigKeyTypeVBooleanType
	MDMAppConfigKeyTypePBooleanType = &_MDMAppConfigKeyTypePBooleanType
	// MDMAppConfigKeyTypePTokenType is a pointer to MDMAppConfigKeyTypeVTokenType
	MDMAppConfigKeyTypePTokenType = &_MDMAppConfigKeyTypePTokenType
)

var (
	_MDMAppConfigKeyTypePStringType  = MDMAppConfigKeyTypeVStringType
	_MDMAppConfigKeyTypePIntegerType = MDMAppConfigKeyTypeVIntegerType
	_MDMAppConfigKeyTypePRealType    = MDMAppConfigKeyTypeVRealType
	_MDMAppConfigKeyTypePBooleanType = MDMAppConfigKeyTypeVBooleanType
	_MDMAppConfigKeyTypePTokenType   = MDMAppConfigKeyTypeVTokenType
)

// MDMAuthority undocumented
type MDMAuthority string

const (
	// MDMAuthorityVUnknown undocumented
	MDMAuthorityVUnknown MDMAuthority = "unknown"
	// MDMAuthorityVIntune undocumented
	MDMAuthorityVIntune MDMAuthority = "intune"
	// MDMAuthorityVSccm undocumented
	MDMAuthorityVSccm MDMAuthority = "sccm"
	// MDMAuthorityVOffice365 undocumented
	MDMAuthorityVOffice365 MDMAuthority = "office365"
)

var (
	// MDMAuthorityPUnknown is a pointer to MDMAuthorityVUnknown
	MDMAuthorityPUnknown = &_MDMAuthorityPUnknown
	// MDMAuthorityPIntune is a pointer to MDMAuthorityVIntune
	MDMAuthorityPIntune = &_MDMAuthorityPIntune
	// MDMAuthorityPSccm is a pointer to MDMAuthorityVSccm
	MDMAuthorityPSccm = &_MDMAuthorityPSccm
	// MDMAuthorityPOffice365 is a pointer to MDMAuthorityVOffice365
	MDMAuthorityPOffice365 = &_MDMAuthorityPOffice365
)

var (
	_MDMAuthorityPUnknown   = MDMAuthorityVUnknown
	_MDMAuthorityPIntune    = MDMAuthorityVIntune
	_MDMAuthorityPSccm      = MDMAuthorityVSccm
	_MDMAuthorityPOffice365 = MDMAuthorityVOffice365
)

// MDMSupportedState undocumented
type MDMSupportedState string

const (
	// MDMSupportedStateVUnknown undocumented
	MDMSupportedStateVUnknown MDMSupportedState = "unknown"
	// MDMSupportedStateVSupported undocumented
	MDMSupportedStateVSupported MDMSupportedState = "supported"
	// MDMSupportedStateVUnsupported undocumented
	MDMSupportedStateVUnsupported MDMSupportedState = "unsupported"
	// MDMSupportedStateVDeprecated undocumented
	MDMSupportedStateVDeprecated MDMSupportedState = "deprecated"
)

var (
	// MDMSupportedStatePUnknown is a pointer to MDMSupportedStateVUnknown
	MDMSupportedStatePUnknown = &_MDMSupportedStatePUnknown
	// MDMSupportedStatePSupported is a pointer to MDMSupportedStateVSupported
	MDMSupportedStatePSupported = &_MDMSupportedStatePSupported
	// MDMSupportedStatePUnsupported is a pointer to MDMSupportedStateVUnsupported
	MDMSupportedStatePUnsupported = &_MDMSupportedStatePUnsupported
	// MDMSupportedStatePDeprecated is a pointer to MDMSupportedStateVDeprecated
	MDMSupportedStatePDeprecated = &_MDMSupportedStatePDeprecated
)

var (
	_MDMSupportedStatePUnknown     = MDMSupportedStateVUnknown
	_MDMSupportedStatePSupported   = MDMSupportedStateVSupported
	_MDMSupportedStatePUnsupported = MDMSupportedStateVUnsupported
	_MDMSupportedStatePDeprecated  = MDMSupportedStateVDeprecated
)