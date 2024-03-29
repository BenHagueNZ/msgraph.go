// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// SupportedClaimConfiguration undocumented
type SupportedClaimConfiguration struct {
	// Object is the base model of SupportedClaimConfiguration
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// NameIDPolicyFormat undocumented
	NameIDPolicyFormat *string `json:"nameIdPolicyFormat,omitempty"`
}

func NewSupportedClaimConfiguration() (*SupportedClaimConfiguration, error) {
	newSupportedClaimConfiguration := &SupportedClaimConfiguration{
		ODataType: "#microsoft.graph.SupportedClaimConfiguration",
	}
	return newSupportedClaimConfiguration, nil
}
