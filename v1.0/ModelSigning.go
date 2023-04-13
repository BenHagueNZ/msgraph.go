// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// SigningCertificateUpdateStatus undocumented
type SigningCertificateUpdateStatus struct {
	// Object is the base model of SigningCertificateUpdateStatus
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// CertificateUpdateResult undocumented
	CertificateUpdateResult *string `json:"certificateUpdateResult,omitempty"`
	// LastRunDateTime undocumented
	LastRunDateTime *time.Time `json:"lastRunDateTime,omitempty"`
}

func NewSigningCertificateUpdateStatus() (*SigningCertificateUpdateStatus, error) {
	newSigningCertificateUpdateStatus := &SigningCertificateUpdateStatus{
		ODataType: "#microsoft.graph.SigningCertificateUpdateStatus",
	}
	return newSigningCertificateUpdateStatus, nil
}
