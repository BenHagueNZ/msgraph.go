// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// X509CertificateAuthenticationMethodConfiguration undocumented
type X509CertificateAuthenticationMethodConfiguration struct {
	// AuthenticationMethodConfiguration is the base model of X509CertificateAuthenticationMethodConfiguration
	AuthenticationMethodConfiguration

	ODataType string `json:"@odata.type"`
	// AuthenticationModeConfiguration undocumented
	AuthenticationModeConfiguration *X509CertificateAuthenticationModeConfiguration `json:"authenticationModeConfiguration,omitempty"`
	// CertificateUserBindings undocumented
	CertificateUserBindings []X509CertificateUserBinding `json:"certificateUserBindings,omitempty"`
	// IncludeTargets undocumented
	IncludeTargets []AuthenticationMethodTarget `json:"includeTargets,omitempty"`
}

func NewX509CertificateAuthenticationMethodConfiguration() (*X509CertificateAuthenticationMethodConfiguration, error) {
	newX509CertificateAuthenticationMethodConfiguration := &X509CertificateAuthenticationMethodConfiguration{
		ODataType: "#microsoft.graph.X509CertificateAuthenticationMethodConfiguration",
	}
	return newX509CertificateAuthenticationMethodConfiguration, nil
}

// X509CertificateAuthenticationModeConfiguration undocumented
type X509CertificateAuthenticationModeConfiguration struct {
	// Object is the base model of X509CertificateAuthenticationModeConfiguration
	Object

	ODataType string `json:"@odata.type"`
	// Rules undocumented
	Rules []X509CertificateRule `json:"rules,omitempty"`
	// X509CertificateAuthenticationDefaultMode undocumented
	X509CertificateAuthenticationDefaultMode *X509CertificateAuthenticationMode `json:"x509CertificateAuthenticationDefaultMode,omitempty"`
}

func NewX509CertificateAuthenticationModeConfiguration() (*X509CertificateAuthenticationModeConfiguration, error) {
	newX509CertificateAuthenticationModeConfiguration := &X509CertificateAuthenticationModeConfiguration{
		ODataType: "#microsoft.graph.X509CertificateAuthenticationModeConfiguration",
	}
	return newX509CertificateAuthenticationModeConfiguration, nil
}

// X509CertificateRule undocumented
type X509CertificateRule struct {
	// Object is the base model of X509CertificateRule
	Object

	ODataType string `json:"@odata.type"`
	// Identifier undocumented
	Identifier *string `json:"identifier,omitempty"`
	// X509CertificateAuthenticationMode undocumented
	X509CertificateAuthenticationMode *X509CertificateAuthenticationMode `json:"x509CertificateAuthenticationMode,omitempty"`
	// X509CertificateRuleType undocumented
	X509CertificateRuleType *X509CertificateRuleType `json:"x509CertificateRuleType,omitempty"`
}

func NewX509CertificateRule() (*X509CertificateRule, error) {
	newX509CertificateRule := &X509CertificateRule{
		ODataType: "#microsoft.graph.X509CertificateRule",
	}
	return newX509CertificateRule, nil
}

// X509CertificateUserBinding undocumented
type X509CertificateUserBinding struct {
	// Object is the base model of X509CertificateUserBinding
	Object

	ODataType string `json:"@odata.type"`
	// Priority undocumented
	Priority *int `json:"priority,omitempty"`
	// UserProperty undocumented
	UserProperty *string `json:"userProperty,omitempty"`
	// X509CertificateField undocumented
	X509CertificateField *string `json:"x509CertificateField,omitempty"`
}

func NewX509CertificateUserBinding() (*X509CertificateUserBinding, error) {
	newX509CertificateUserBinding := &X509CertificateUserBinding{
		ODataType: "#microsoft.graph.X509CertificateUserBinding",
	}
	return newX509CertificateUserBinding, nil
}
