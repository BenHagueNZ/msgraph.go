// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ProvisionedIdentity undocumented
type ProvisionedIdentity struct {
	// Identity is the base model of ProvisionedIdentity
	Identity

	ODataType string `json:"@odata.type,omitempty"`
	// Details undocumented
	Details *DetailsInfo `json:"details,omitempty"`
	// IdentityType undocumented
	IdentityType *string `json:"identityType,omitempty"`
}

func NewProvisionedIdentity() (*ProvisionedIdentity, error) {
	newProvisionedIdentity := &ProvisionedIdentity{
		ODataType: "#microsoft.graph.ProvisionedIdentity",
	}
	return newProvisionedIdentity, nil
}

// ProvisionedPlan undocumented
type ProvisionedPlan struct {
	// Object is the base model of ProvisionedPlan
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// CapabilityStatus undocumented
	CapabilityStatus *string `json:"capabilityStatus,omitempty"`
	// ProvisioningStatus undocumented
	ProvisioningStatus *string `json:"provisioningStatus,omitempty"`
	// Service undocumented
	Service *string `json:"service,omitempty"`
}

func NewProvisionedPlan() (*ProvisionedPlan, error) {
	newProvisionedPlan := &ProvisionedPlan{
		ODataType: "#microsoft.graph.ProvisionedPlan",
	}
	return newProvisionedPlan, nil
}
