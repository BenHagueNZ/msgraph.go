// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Certification undocumented
type Certification struct {
	// Object is the base model of Certification
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// CertificationDetailsURL undocumented
	CertificationDetailsURL *string `json:"certificationDetailsUrl,omitempty"`
	// CertificationExpirationDateTime undocumented
	CertificationExpirationDateTime *time.Time `json:"certificationExpirationDateTime,omitempty"`
	// IsCertifiedByMicrosoft undocumented
	IsCertifiedByMicrosoft *bool `json:"isCertifiedByMicrosoft,omitempty"`
	// IsPublisherAttested undocumented
	IsPublisherAttested *bool `json:"isPublisherAttested,omitempty"`
	// LastCertificationDateTime undocumented
	LastCertificationDateTime *time.Time `json:"lastCertificationDateTime,omitempty"`
}

func NewCertification() (*Certification, error) {
	newCertification := &Certification{
		ODataType: "#microsoft.graph.Certification",
	}
	return newCertification, nil
}

// CertificationControl undocumented
type CertificationControl struct {
	// Object is the base model of CertificationControl
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// URL undocumented
	URL *string `json:"url,omitempty"`
}

func NewCertificationControl() (*CertificationControl, error) {
	newCertificationControl := &CertificationControl{
		ODataType: "#microsoft.graph.CertificationControl",
	}
	return newCertificationControl, nil
}