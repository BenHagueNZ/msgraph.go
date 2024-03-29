// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// OemWarranty undocumented
type OemWarranty struct {
	// Object is the base model of OemWarranty
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AdditionalWarranties undocumented
	AdditionalWarranties []WarrantyOffer `json:"additionalWarranties,omitempty"`
	// BaseWarranties undocumented
	BaseWarranties []WarrantyOffer `json:"baseWarranties,omitempty"`
	// DeviceConfigurationURL undocumented
	DeviceConfigurationURL *string `json:"deviceConfigurationUrl,omitempty"`
	// DeviceWarrantyURL undocumented
	DeviceWarrantyURL *string `json:"deviceWarrantyUrl,omitempty"`
}

func NewOemWarranty() (*OemWarranty, error) {
	newOemWarranty := &OemWarranty{
		ODataType: "#microsoft.graph.OemWarranty",
	}
	return newOemWarranty, nil
}

// OemWarrantyInformationOnboarding undocumented
type OemWarrantyInformationOnboarding struct {
	// Entity is the base model of OemWarrantyInformationOnboarding
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Available undocumented
	Available *bool `json:"available,omitempty"`
	// Enabled undocumented
	Enabled *bool `json:"enabled,omitempty"`
	// OemName undocumented
	OemName *string `json:"oemName,omitempty"`
}

func NewOemWarrantyInformationOnboarding() (*OemWarrantyInformationOnboarding, error) {
	newOemWarrantyInformationOnboarding := &OemWarrantyInformationOnboarding{
		ODataType: "#microsoft.graph.OemWarrantyInformationOnboarding",
	}
	return newOemWarrantyInformationOnboarding, nil
}
