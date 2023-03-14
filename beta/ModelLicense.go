
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// LicenseAssignmentState undocumented
type LicenseAssignmentState struct {
    // Object is the base model of LicenseAssignmentState
    Object
    // AssignedByGroup undocumented
    AssignedByGroup *string `json:"assignedByGroup,omitempty"`
    // DisabledPlans undocumented
    DisabledPlans []UUID `json:"disabledPlans,omitempty"`
    // Error undocumented
    Error *string `json:"error,omitempty"`
    // LastUpdatedDateTime undocumented
    LastUpdatedDateTime *time.Time `json:"lastUpdatedDateTime,omitempty"`
    // SKUID undocumented
    SKUID *UUID `json:"skuId,omitempty"`
    // State undocumented
    State *string `json:"state,omitempty"`
}

// LicenseDetails undocumented
type LicenseDetails struct {
    // Entity is the base model of LicenseDetails
    Entity
    // ServicePlans undocumented
    ServicePlans []ServicePlanInfo `json:"servicePlans,omitempty"`
    // SKUID undocumented
    SKUID *UUID `json:"skuId,omitempty"`
    // SKUPartNumber undocumented
    SKUPartNumber *string `json:"skuPartNumber,omitempty"`
}

// LicenseProcessingState undocumented
type LicenseProcessingState struct {
    // Object is the base model of LicenseProcessingState
    Object
    // State undocumented
    State *string `json:"state,omitempty"`
}

// LicenseUnitsDetail undocumented
type LicenseUnitsDetail struct {
    // Object is the base model of LicenseUnitsDetail
    Object
    // Enabled undocumented
    Enabled *int `json:"enabled,omitempty"`
    // Suspended undocumented
    Suspended *int `json:"suspended,omitempty"`
    // Warning undocumented
    Warning *int `json:"warning,omitempty"`
}
