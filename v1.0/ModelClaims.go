// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ClaimsMappingPolicy undocumented
type ClaimsMappingPolicy struct {
	// StsPolicy is the base model of ClaimsMappingPolicy
	StsPolicy

	ODataType string `json:"@odata.type,omitempty"`
}

func NewClaimsMappingPolicy() (*ClaimsMappingPolicy, error) {
	newClaimsMappingPolicy := &ClaimsMappingPolicy{
		ODataType: "#microsoft.graph.ClaimsMappingPolicy",
	}
	return newClaimsMappingPolicy, nil
}
