
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// OathTokenMetadata undocumented
type OathTokenMetadata struct {
    // Object is the base model of OathTokenMetadata
    Object
    // Enabled undocumented
    Enabled *bool `json:"enabled,omitempty"`
    // Manufacturer undocumented
    Manufacturer *string `json:"manufacturer,omitempty"`
    // ManufacturerProperties undocumented
    ManufacturerProperties []KeyValue `json:"manufacturerProperties,omitempty"`
    // SerialNumber undocumented
    SerialNumber *string `json:"serialNumber,omitempty"`
    // TokenType undocumented
    TokenType *string `json:"tokenType,omitempty"`
}