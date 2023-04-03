// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ProxiedDomain undocumented
type ProxiedDomain struct {
	// Object is the base model of ProxiedDomain
	Object

	ODataType string `json:"@odata.type"`
	// IPAddressOrFQDN undocumented
	IPAddressOrFQDN *string `json:"ipAddressOrFQDN,omitempty"`
	// Proxy undocumented
	Proxy *string `json:"proxy,omitempty"`
}

func NewProxiedDomain() (*ProxiedDomain, error) {
	newProxiedDomain := &ProxiedDomain{
		ODataType: "#microsoft.graph.ProxiedDomain",
	}
	return newProxiedDomain, nil
}
