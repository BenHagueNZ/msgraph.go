// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// SubscribedSKU undocumented
type SubscribedSKU struct {
	// Entity is the base model of SubscribedSKU
	Entity

	ODataType string `json:"@odata.type"`
	// AppliesTo undocumented
	AppliesTo *string `json:"appliesTo,omitempty"`
	// CapabilityStatus undocumented
	CapabilityStatus *string `json:"capabilityStatus,omitempty"`
	// ConsumedUnits undocumented
	ConsumedUnits *int `json:"consumedUnits,omitempty"`
	// PrepaidUnits undocumented
	PrepaidUnits *LicenseUnitsDetail `json:"prepaidUnits,omitempty"`
	// ServicePlans undocumented
	ServicePlans []ServicePlanInfo `json:"servicePlans,omitempty"`
	// SKUID undocumented
	SKUID *UUID `json:"skuId,omitempty"`
	// SKUPartNumber undocumented
	SKUPartNumber *string `json:"skuPartNumber,omitempty"`
}

func NewSubscribedSKU() (*SubscribedSKU, error) {
	newSubscribedSKU := &SubscribedSKU{
		ODataType: "#microsoft.graph.SubscribedSku",
	}
	return newSubscribedSKU, nil
}
