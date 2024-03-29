// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// CertificateAuthority undocumented
type CertificateAuthority struct {
	// Object is the base model of CertificateAuthority
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Certificate undocumented
	Certificate *Binary `json:"certificate,omitempty"`
	// CertificateRevocationListURL undocumented
	CertificateRevocationListURL *string `json:"certificateRevocationListUrl,omitempty"`
	// DeltaCertificateRevocationListURL undocumented
	DeltaCertificateRevocationListURL *string `json:"deltaCertificateRevocationListUrl,omitempty"`
	// IsRootAuthority undocumented
	IsRootAuthority *bool `json:"isRootAuthority,omitempty"`
	// Issuer undocumented
	Issuer *string `json:"issuer,omitempty"`
	// IssuerSki undocumented
	IssuerSki *string `json:"issuerSki,omitempty"`
}

func NewCertificateAuthority() (*CertificateAuthority, error) {
	newCertificateAuthority := &CertificateAuthority{
		ODataType: "#microsoft.graph.CertificateAuthority",
	}
	return newCertificateAuthority, nil
}

// CertificateBasedAuthConfiguration undocumented
type CertificateBasedAuthConfiguration struct {
	// Entity is the base model of CertificateBasedAuthConfiguration
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// CertificateAuthorities undocumented
	CertificateAuthorities []CertificateAuthority `json:"certificateAuthorities,omitempty"`
}

func NewCertificateBasedAuthConfiguration() (*CertificateBasedAuthConfiguration, error) {
	newCertificateBasedAuthConfiguration := &CertificateBasedAuthConfiguration{
		ODataType: "#microsoft.graph.CertificateBasedAuthConfiguration",
	}
	return newCertificateBasedAuthConfiguration, nil
}
