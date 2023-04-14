// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// IPNamedLocation undocumented
type IPNamedLocation struct {
	// NamedLocation is the base model of IPNamedLocation
	NamedLocation

	ODataType string `json:"@odata.type,omitempty"`
	// IPRanges undocumented
	IPRanges []IPRange `json:"ipRanges,omitempty"`
	IPv4Ranges []IPv4CidrRange `json:"ipRanges1,omitempty"`
	IPv6Ranges []IPv6CidrRange `json:"ipRanges2,omitempty"`
	// IsTrusted undocumented
	IsTrusted *bool `json:"isTrusted,omitempty"`
}

func NewIPNamedLocation() (*IPNamedLocation, error) {
	newIPNamedLocation := &IPNamedLocation{
		ODataType: "#microsoft.graph.IpNamedLocation",
	}
	return newIPNamedLocation, nil
}

// IPRange undocumented
type IPRange struct {
	// Object is the base model of IPRange
	Object

	ODataType string `json:"@odata.type,omitempty"`
}

func NewIPRange() (*IPRange, error) {
	newIPRange := &IPRange{
		ODataType: "#microsoft.graph.IpRange",
	}
	return newIPRange, nil
}
