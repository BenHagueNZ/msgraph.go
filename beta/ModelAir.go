
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// AirPrintDestination undocumented
type AirPrintDestination struct {
    // Object is the base model of AirPrintDestination
    Object
    // ForceTLS undocumented
    ForceTLS *bool `json:"forceTls,omitempty"`
    // IPAddress undocumented
    IPAddress *string `json:"ipAddress,omitempty"`
    // Port undocumented
    Port *int `json:"port,omitempty"`
    // ResourcePath undocumented
    ResourcePath *string `json:"resourcePath,omitempty"`
}
