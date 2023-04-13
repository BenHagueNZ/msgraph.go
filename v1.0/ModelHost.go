// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// HostSecurityState undocumented
type HostSecurityState struct {
	// Object is the base model of HostSecurityState
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Fqdn undocumented
	Fqdn *string `json:"fqdn,omitempty"`
	// IsAzureAdJoined undocumented
	IsAzureAdJoined *bool `json:"isAzureAdJoined,omitempty"`
	// IsAzureAdRegistered undocumented
	IsAzureAdRegistered *bool `json:"isAzureAdRegistered,omitempty"`
	// IsHybridAzureDomainJoined undocumented
	IsHybridAzureDomainJoined *bool `json:"isHybridAzureDomainJoined,omitempty"`
	// NetBiosName undocumented
	NetBiosName *string `json:"netBiosName,omitempty"`
	// Os undocumented
	Os *string `json:"os,omitempty"`
	// PrivateIPAddress undocumented
	PrivateIPAddress *string `json:"privateIpAddress,omitempty"`
	// PublicIPAddress undocumented
	PublicIPAddress *string `json:"publicIpAddress,omitempty"`
	// RiskScore undocumented
	RiskScore *string `json:"riskScore,omitempty"`
}

func NewHostSecurityState() (*HostSecurityState, error) {
	newHostSecurityState := &HostSecurityState{
		ODataType: "#microsoft.graph.HostSecurityState",
	}
	return newHostSecurityState, nil
}
