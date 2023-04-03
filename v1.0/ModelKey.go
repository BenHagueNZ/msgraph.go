// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// KeyCredential undocumented
type KeyCredential struct {
	// Object is the base model of KeyCredential
	Object

	ODataType string `json:"@odata.type"`
	// CustomKeyIdentifier undocumented
	CustomKeyIdentifier *Binary `json:"customKeyIdentifier,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// Key undocumented
	Key *Binary `json:"key,omitempty"`
	// KeyID undocumented
	KeyID *UUID `json:"keyId,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
	// Usage undocumented
	Usage *string `json:"usage,omitempty"`
}

func NewKeyCredential() (*KeyCredential, error) {
	newKeyCredential := &KeyCredential{
		ODataType: "#microsoft.graph.KeyCredential",
	}
	return newKeyCredential, nil
}

// KeyCredentialConfiguration undocumented
type KeyCredentialConfiguration struct {
	// Object is the base model of KeyCredentialConfiguration
	Object

	ODataType string `json:"@odata.type"`
	// MaxLifetime undocumented
	MaxLifetime *Duration `json:"maxLifetime,omitempty"`
	// RestrictForAppsCreatedAfterDateTime undocumented
	RestrictForAppsCreatedAfterDateTime *time.Time `json:"restrictForAppsCreatedAfterDateTime,omitempty"`
	// RestrictionType undocumented
	RestrictionType *AppKeyCredentialRestrictionType `json:"restrictionType,omitempty"`
}

func NewKeyCredentialConfiguration() (*KeyCredentialConfiguration, error) {
	newKeyCredentialConfiguration := &KeyCredentialConfiguration{
		ODataType: "#microsoft.graph.KeyCredentialConfiguration",
	}
	return newKeyCredentialConfiguration, nil
}

// KeyValue undocumented
type KeyValue struct {
	// Object is the base model of KeyValue
	Object

	ODataType string `json:"@odata.type"`
	// Key undocumented
	Key *string `json:"key,omitempty"`
	// Value undocumented
	Value *string `json:"value,omitempty"`
}

func NewKeyValue() (*KeyValue, error) {
	newKeyValue := &KeyValue{
		ODataType: "#microsoft.graph.KeyValue",
	}
	return newKeyValue, nil
}

// KeyValuePair undocumented
type KeyValuePair struct {
	// Object is the base model of KeyValuePair
	Object

	ODataType string `json:"@odata.type"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Value undocumented
	Value *string `json:"value,omitempty"`
}

func NewKeyValuePair() (*KeyValuePair, error) {
	newKeyValuePair := &KeyValuePair{
		ODataType: "#microsoft.graph.KeyValuePair",
	}
	return newKeyValuePair, nil
}
