
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// IPv4CidrRange undocumented
type IPv4CidrRange struct {
    // IPRange is the base model of IPv4CidrRange
    IPRange
    // CIDRAddress undocumented
    CIDRAddress *string `json:"cidrAddress,omitempty"`
}

// IPv4Range undocumented
type IPv4Range struct {
    // IPRange is the base model of IPv4Range
    IPRange
    // LowerAddress undocumented
    LowerAddress *string `json:"lowerAddress,omitempty"`
    // UpperAddress undocumented
    UpperAddress *string `json:"upperAddress,omitempty"`
}
