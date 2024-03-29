// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// CountryNamedLocation undocumented
type CountryNamedLocation struct {
	// NamedLocation is the base model of CountryNamedLocation
	NamedLocation

	ODataType string `json:"@odata.type,omitempty"`
	// CountriesAndRegions undocumented
	CountriesAndRegions []string `json:"countriesAndRegions,omitempty"`
	// CountryLookupMethod undocumented
	CountryLookupMethod *CountryLookupMethodType `json:"countryLookupMethod,omitempty"`
	// IncludeUnknownCountriesAndRegions undocumented
	IncludeUnknownCountriesAndRegions *bool `json:"includeUnknownCountriesAndRegions,omitempty"`
}

func NewCountryNamedLocation() (*CountryNamedLocation, error) {
	newCountryNamedLocation := &CountryNamedLocation{
		ODataType: "#microsoft.graph.CountryNamedLocation",
	}
	return newCountryNamedLocation, nil
}

// CountryRegion undocumented
type CountryRegion struct {
	// Entity is the base model of CountryRegion
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AddressFormat undocumented
	AddressFormat *string `json:"addressFormat,omitempty"`
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}

func NewCountryRegion() (*CountryRegion, error) {
	newCountryRegion := &CountryRegion{
		ODataType: "#microsoft.graph.CountryRegion",
	}
	return newCountryRegion, nil
}
