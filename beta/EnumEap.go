// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// EapFastConfiguration undocumented
type EapFastConfiguration string

const (
	// EapFastConfigurationVNoProtectedAccessCredential undocumented
	EapFastConfigurationVNoProtectedAccessCredential EapFastConfiguration = "noProtectedAccessCredential"
	// EapFastConfigurationVUseProtectedAccessCredential undocumented
	EapFastConfigurationVUseProtectedAccessCredential EapFastConfiguration = "useProtectedAccessCredential"
	// EapFastConfigurationVUseProtectedAccessCredentialAndProvision undocumented
	EapFastConfigurationVUseProtectedAccessCredentialAndProvision EapFastConfiguration = "useProtectedAccessCredentialAndProvision"
	// EapFastConfigurationVUseProtectedAccessCredentialAndProvisionAnonymously undocumented
	EapFastConfigurationVUseProtectedAccessCredentialAndProvisionAnonymously EapFastConfiguration = "useProtectedAccessCredentialAndProvisionAnonymously"
)

var (
	// EapFastConfigurationPNoProtectedAccessCredential is a pointer to EapFastConfigurationVNoProtectedAccessCredential
	EapFastConfigurationPNoProtectedAccessCredential = &_EapFastConfigurationPNoProtectedAccessCredential
	// EapFastConfigurationPUseProtectedAccessCredential is a pointer to EapFastConfigurationVUseProtectedAccessCredential
	EapFastConfigurationPUseProtectedAccessCredential = &_EapFastConfigurationPUseProtectedAccessCredential
	// EapFastConfigurationPUseProtectedAccessCredentialAndProvision is a pointer to EapFastConfigurationVUseProtectedAccessCredentialAndProvision
	EapFastConfigurationPUseProtectedAccessCredentialAndProvision = &_EapFastConfigurationPUseProtectedAccessCredentialAndProvision
	// EapFastConfigurationPUseProtectedAccessCredentialAndProvisionAnonymously is a pointer to EapFastConfigurationVUseProtectedAccessCredentialAndProvisionAnonymously
	EapFastConfigurationPUseProtectedAccessCredentialAndProvisionAnonymously = &_EapFastConfigurationPUseProtectedAccessCredentialAndProvisionAnonymously
)

var (
	_EapFastConfigurationPNoProtectedAccessCredential                         = EapFastConfigurationVNoProtectedAccessCredential
	_EapFastConfigurationPUseProtectedAccessCredential                        = EapFastConfigurationVUseProtectedAccessCredential
	_EapFastConfigurationPUseProtectedAccessCredentialAndProvision            = EapFastConfigurationVUseProtectedAccessCredentialAndProvision
	_EapFastConfigurationPUseProtectedAccessCredentialAndProvisionAnonymously = EapFastConfigurationVUseProtectedAccessCredentialAndProvisionAnonymously
)

// EapType undocumented
type EapType string

const (
	// EapTypeVEapTLS undocumented
	EapTypeVEapTLS EapType = "eapTls"
	// EapTypeVLeap undocumented
	EapTypeVLeap EapType = "leap"
	// EapTypeVEapSim undocumented
	EapTypeVEapSim EapType = "eapSim"
	// EapTypeVEapTtls undocumented
	EapTypeVEapTtls EapType = "eapTtls"
	// EapTypeVPeap undocumented
	EapTypeVPeap EapType = "peap"
	// EapTypeVEapFast undocumented
	EapTypeVEapFast EapType = "eapFast"
	// EapTypeVTeap undocumented
	EapTypeVTeap EapType = "teap"
)

var (
	// EapTypePEapTLS is a pointer to EapTypeVEapTLS
	EapTypePEapTLS = &_EapTypePEapTLS
	// EapTypePLeap is a pointer to EapTypeVLeap
	EapTypePLeap = &_EapTypePLeap
	// EapTypePEapSim is a pointer to EapTypeVEapSim
	EapTypePEapSim = &_EapTypePEapSim
	// EapTypePEapTtls is a pointer to EapTypeVEapTtls
	EapTypePEapTtls = &_EapTypePEapTtls
	// EapTypePPeap is a pointer to EapTypeVPeap
	EapTypePPeap = &_EapTypePPeap
	// EapTypePEapFast is a pointer to EapTypeVEapFast
	EapTypePEapFast = &_EapTypePEapFast
	// EapTypePTeap is a pointer to EapTypeVTeap
	EapTypePTeap = &_EapTypePTeap
)

var (
	_EapTypePEapTLS  = EapTypeVEapTLS
	_EapTypePLeap    = EapTypeVLeap
	_EapTypePEapSim  = EapTypeVEapSim
	_EapTypePEapTtls = EapTypeVEapTtls
	_EapTypePPeap    = EapTypeVPeap
	_EapTypePEapFast = EapTypeVEapFast
	_EapTypePTeap    = EapTypeVTeap
)
