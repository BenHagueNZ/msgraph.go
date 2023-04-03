// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Subscription undocumented
type Subscription struct {
	// Entity is the base model of Subscription
	Entity

	ODataType string `json:"@odata.type"`
	// ApplicationID undocumented
	ApplicationID *string `json:"applicationId,omitempty"`
	// ChangeType undocumented
	ChangeType *string `json:"changeType,omitempty"`
	// ClientState undocumented
	ClientState *string `json:"clientState,omitempty"`
	// CreatorID undocumented
	CreatorID *string `json:"creatorId,omitempty"`
	// EncryptionCertificate undocumented
	EncryptionCertificate *string `json:"encryptionCertificate,omitempty"`
	// EncryptionCertificateID undocumented
	EncryptionCertificateID *string `json:"encryptionCertificateId,omitempty"`
	// ExpirationDateTime undocumented
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// IncludeResourceData undocumented
	IncludeResourceData *bool `json:"includeResourceData,omitempty"`
	// LatestSupportedTLSVersion undocumented
	LatestSupportedTLSVersion *string `json:"latestSupportedTlsVersion,omitempty"`
	// LifecycleNotificationURL undocumented
	LifecycleNotificationURL *string `json:"lifecycleNotificationUrl,omitempty"`
	// NotificationQueryOptions undocumented
	NotificationQueryOptions *string `json:"notificationQueryOptions,omitempty"`
	// NotificationURL undocumented
	NotificationURL *string `json:"notificationUrl,omitempty"`
	// NotificationURLAppID undocumented
	NotificationURLAppID *string `json:"notificationUrlAppId,omitempty"`
	// Resource undocumented
	Resource *string `json:"resource,omitempty"`
}

func NewSubscription() (*Subscription, error) {
	newSubscription := &Subscription{
		ODataType: "#microsoft.graph.Subscription",
	}
	return newSubscription, nil
}
