// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// HomeRealmDiscoveryPolicy undocumented
type HomeRealmDiscoveryPolicy struct {
	// StsPolicy is the base model of HomeRealmDiscoveryPolicy
	StsPolicy

	ODataType string `json:"@odata.type"`
}

func NewHomeRealmDiscoveryPolicy() (*HomeRealmDiscoveryPolicy, error) {
	newHomeRealmDiscoveryPolicy := &HomeRealmDiscoveryPolicy{
		ODataType: "#microsoft.graph.HomeRealmDiscoveryPolicy",
	}
	return newHomeRealmDiscoveryPolicy, nil
}
