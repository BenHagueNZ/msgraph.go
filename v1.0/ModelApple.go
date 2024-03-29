// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// AppleDeviceFeaturesConfigurationBase undocumented
type AppleDeviceFeaturesConfigurationBase struct {
	// DeviceConfiguration is the base model of AppleDeviceFeaturesConfigurationBase
	DeviceConfiguration

	ODataType string `json:"@odata.type,omitempty"`
}

func NewAppleDeviceFeaturesConfigurationBase() (*AppleDeviceFeaturesConfigurationBase, error) {
	newAppleDeviceFeaturesConfigurationBase := &AppleDeviceFeaturesConfigurationBase{
		ODataType: "#microsoft.graph.AppleDeviceFeaturesConfigurationBase",
	}
	return newAppleDeviceFeaturesConfigurationBase, nil
}

// AppleManagedIdentityProvider undocumented
type AppleManagedIdentityProvider struct {
	// IdentityProviderBase is the base model of AppleManagedIdentityProvider
	IdentityProviderBase

	ODataType string `json:"@odata.type,omitempty"`
	// CertificateData undocumented
	CertificateData *string `json:"certificateData,omitempty"`
	// DeveloperID undocumented
	DeveloperID *string `json:"developerId,omitempty"`
	// KeyID undocumented
	KeyID *string `json:"keyId,omitempty"`
	// ServiceID undocumented
	ServiceID *string `json:"serviceId,omitempty"`
}

func NewAppleManagedIdentityProvider() (*AppleManagedIdentityProvider, error) {
	newAppleManagedIdentityProvider := &AppleManagedIdentityProvider{
		ODataType: "#microsoft.graph.AppleManagedIdentityProvider",
	}
	return newAppleManagedIdentityProvider, nil
}

// ApplePushNotificationCertificate undocumented
type ApplePushNotificationCertificate struct {
	// Entity is the base model of ApplePushNotificationCertificate
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AppleIdentifier undocumented
	AppleIdentifier *string `json:"appleIdentifier,omitempty"`
	// Certificate undocumented
	Certificate *string `json:"certificate,omitempty"`
	// CertificateSerialNumber undocumented
	CertificateSerialNumber *string `json:"certificateSerialNumber,omitempty"`
	// CertificateUploadFailureReason undocumented
	CertificateUploadFailureReason *string `json:"certificateUploadFailureReason,omitempty"`
	// CertificateUploadStatus undocumented
	CertificateUploadStatus *string `json:"certificateUploadStatus,omitempty"`
	// ExpirationDateTime undocumented
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// TopicIdentifier undocumented
	TopicIdentifier *string `json:"topicIdentifier,omitempty"`
}

func NewApplePushNotificationCertificate() (*ApplePushNotificationCertificate, error) {
	newApplePushNotificationCertificate := &ApplePushNotificationCertificate{
		ODataType: "#microsoft.graph.ApplePushNotificationCertificate",
	}
	return newApplePushNotificationCertificate, nil
}
