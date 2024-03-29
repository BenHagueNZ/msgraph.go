// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// SymantecCodeSigningCertificate undocumented
type SymantecCodeSigningCertificate struct {
	// Entity is the base model of SymantecCodeSigningCertificate
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Content undocumented
	Content *Binary `json:"content,omitempty"`
	// ExpirationDateTime undocumented
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// Issuer undocumented
	Issuer *string `json:"issuer,omitempty"`
	// IssuerName undocumented
	IssuerName *string `json:"issuerName,omitempty"`
	// Password undocumented
	Password *string `json:"password,omitempty"`
	// Status undocumented
	Status *CertificateStatus `json:"status,omitempty"`
	// Subject undocumented
	Subject *string `json:"subject,omitempty"`
	// SubjectName undocumented
	SubjectName *string `json:"subjectName,omitempty"`
	// UploadDateTime undocumented
	UploadDateTime *time.Time `json:"uploadDateTime,omitempty"`
}

func NewSymantecCodeSigningCertificate() (*SymantecCodeSigningCertificate, error) {
	newSymantecCodeSigningCertificate := &SymantecCodeSigningCertificate{
		ODataType: "#microsoft.graph.SymantecCodeSigningCertificate",
	}
	return newSymantecCodeSigningCertificate, nil
}
