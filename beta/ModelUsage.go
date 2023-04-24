// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// UsageDetails undocumented
type UsageDetails struct {
	// Object is the base model of UsageDetails
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// LastAccessedDateTime undocumented
	LastAccessedDateTime *time.Time `json:"lastAccessedDateTime,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}

func NewUsageDetails() (*UsageDetails, error) {
	newUsageDetails := &UsageDetails{
		ODataType: "#microsoft.graph.UsageDetails",
	}
	return newUsageDetails, nil
}

// UsageRight undocumented
type UsageRight struct {
	// Entity is the base model of UsageRight
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// CatalogID undocumented
	CatalogID *string `json:"catalogId,omitempty"`
	// ServiceIdentifier undocumented
	ServiceIdentifier *string `json:"serviceIdentifier,omitempty"`
	// State undocumented
	State *UsageRightState `json:"state,omitempty"`
}

func NewUsageRight() (*UsageRight, error) {
	newUsageRight := &UsageRight{
		ODataType: "#microsoft.graph.UsageRight",
	}
	return newUsageRight, nil
}
