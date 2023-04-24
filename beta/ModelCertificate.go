// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

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

// CertificateConnectorDetails undocumented
type CertificateConnectorDetails struct {
	// Entity is the base model of CertificateConnectorDetails
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ConnectorName undocumented
	ConnectorName *string `json:"connectorName,omitempty"`
	// ConnectorVersion undocumented
	ConnectorVersion *string `json:"connectorVersion,omitempty"`
	// EnrollmentDateTime undocumented
	EnrollmentDateTime *time.Time `json:"enrollmentDateTime,omitempty"`
	// LastCheckinDateTime undocumented
	LastCheckinDateTime *time.Time `json:"lastCheckinDateTime,omitempty"`
	// MachineName undocumented
	MachineName *string `json:"machineName,omitempty"`
}

func NewCertificateConnectorDetails() (*CertificateConnectorDetails, error) {
	newCertificateConnectorDetails := &CertificateConnectorDetails{
		ODataType: "#microsoft.graph.CertificateConnectorDetails",
	}
	return newCertificateConnectorDetails, nil
}

// CertificateConnectorHealthMetricValue undocumented
type CertificateConnectorHealthMetricValue struct {
	// Object is the base model of CertificateConnectorHealthMetricValue
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// DateTime undocumented
	DateTime *time.Time `json:"dateTime,omitempty"`
	// FailureCount undocumented
	FailureCount *int `json:"failureCount,omitempty"`
	// SuccessCount undocumented
	SuccessCount *int `json:"successCount,omitempty"`
}

func NewCertificateConnectorHealthMetricValue() (*CertificateConnectorHealthMetricValue, error) {
	newCertificateConnectorHealthMetricValue := &CertificateConnectorHealthMetricValue{
		ODataType: "#microsoft.graph.CertificateConnectorHealthMetricValue",
	}
	return newCertificateConnectorHealthMetricValue, nil
}

// CertificateConnectorSetting undocumented
type CertificateConnectorSetting struct {
	// Object is the base model of CertificateConnectorSetting
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// CertExpiryTime undocumented
	CertExpiryTime *time.Time `json:"certExpiryTime,omitempty"`
	// ConnectorVersion undocumented
	ConnectorVersion *string `json:"connectorVersion,omitempty"`
	// EnrollmentError undocumented
	EnrollmentError *string `json:"enrollmentError,omitempty"`
	// LastConnectorConnectionTime undocumented
	LastConnectorConnectionTime *time.Time `json:"lastConnectorConnectionTime,omitempty"`
	// LastUploadVersion undocumented
	LastUploadVersion *int `json:"lastUploadVersion,omitempty"`
	// Status undocumented
	Status *int `json:"status,omitempty"`
}

func NewCertificateConnectorSetting() (*CertificateConnectorSetting, error) {
	newCertificateConnectorSetting := &CertificateConnectorSetting{
		ODataType: "#microsoft.graph.CertificateConnectorSetting",
	}
	return newCertificateConnectorSetting, nil
}
