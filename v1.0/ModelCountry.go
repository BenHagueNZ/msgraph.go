// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// CountryNamedLocation undocumented
type CountryNamedLocation struct {
	// NamedLocation is the base model of CountryNamedLocation
	NamedLocation
	// CountriesAndRegions undocumented
	CountriesAndRegions []string `json:"countriesAndRegions,omitempty"`
	// CountryLookupMethod undocumented
	CountryLookupMethod *CountryLookupMethodType `json:"countryLookupMethod,omitempty"`
	// IncludeUnknownCountriesAndRegions undocumented
	IncludeUnknownCountriesAndRegions *bool `json:"includeUnknownCountriesAndRegions,omitempty"`
}
