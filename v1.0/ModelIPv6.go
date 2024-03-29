// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// IPv6CidrRange undocumented
type IPv6CidrRange struct {
	// IPRange is the base model of IPv6CidrRange
	IPRange

	ODataType string `json:"@odata.type,omitempty"`
	// CIDRAddress undocumented
	CIDRAddress *string `json:"cidrAddress,omitempty"`
}

func NewIPv6CidrRange() (*IPv6CidrRange, error) {
	newIPv6CidrRange := &IPv6CidrRange{
		ODataType: "#microsoft.graph.IPv6CidrRange",
	}
	return newIPv6CidrRange, nil
}

// IPv6Range undocumented
type IPv6Range struct {
	// IPRange is the base model of IPv6Range
	IPRange

	ODataType string `json:"@odata.type,omitempty"`
	// LowerAddress undocumented
	LowerAddress *string `json:"lowerAddress,omitempty"`
	// UpperAddress undocumented
	UpperAddress *string `json:"upperAddress,omitempty"`
}

func NewIPv6Range() (*IPv6Range, error) {
	newIPv6Range := &IPv6Range{
		ODataType: "#microsoft.graph.IPv6Range",
	}
	return newIPv6Range, nil
}
