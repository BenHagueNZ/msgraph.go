// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Pkcs12Certificate undocumented
type Pkcs12Certificate struct {
	// APIAuthenticationConfigurationBase is the base model of Pkcs12Certificate
	APIAuthenticationConfigurationBase

	ODataType string `json:"@odata.type,omitempty"`
	// Password undocumented
	Password *string `json:"password,omitempty"`
	// Pkcs12Value undocumented
	Pkcs12Value *string `json:"pkcs12Value,omitempty"`
}

func NewPkcs12Certificate() (*Pkcs12Certificate, error) {
	newPkcs12Certificate := &Pkcs12Certificate{
		ODataType: "#microsoft.graph.Pkcs12Certificate",
	}
	return newPkcs12Certificate, nil
}

// Pkcs12CertificateInformation undocumented
type Pkcs12CertificateInformation struct {
	// Object is the base model of Pkcs12CertificateInformation
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// IsActive undocumented
	IsActive *bool `json:"isActive,omitempty"`
	// NotAfter undocumented
	NotAfter *int `json:"notAfter,omitempty"`
	// NotBefore undocumented
	NotBefore *int `json:"notBefore,omitempty"`
	// Thumbprint undocumented
	Thumbprint *string `json:"thumbprint,omitempty"`
}

func NewPkcs12CertificateInformation() (*Pkcs12CertificateInformation, error) {
	newPkcs12CertificateInformation := &Pkcs12CertificateInformation{
		ODataType: "#microsoft.graph.Pkcs12CertificateInformation",
	}
	return newPkcs12CertificateInformation, nil
}