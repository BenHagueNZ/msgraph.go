
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// RichLongRunningOperation undocumented
type RichLongRunningOperation struct {
    // LongRunningOperation is the base model of RichLongRunningOperation
    LongRunningOperation
    // Error undocumented
    Error *PublicError `json:"error,omitempty"`
    // PercentageComplete undocumented
    PercentageComplete *int `json:"percentageComplete,omitempty"`
    // ResourceID undocumented
    ResourceID *string `json:"resourceId,omitempty"`
    // Type undocumented
    Type *string `json:"type,omitempty"`
}
