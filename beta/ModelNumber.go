
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// NumberColumn undocumented
type NumberColumn struct {
    // Object is the base model of NumberColumn
    Object
    // DecimalPlaces undocumented
    DecimalPlaces *string `json:"decimalPlaces,omitempty"`
    // DisplayAs undocumented
    DisplayAs *string `json:"displayAs,omitempty"`
    // Maximum undocumented
    Maximum *float64 `json:"maximum,omitempty"`
    // Minimum undocumented
    Minimum *float64 `json:"minimum,omitempty"`
}

// NumberRange undocumented
type NumberRange struct {
    // Object is the base model of NumberRange
    Object
    // LowerNumber undocumented
    LowerNumber *int `json:"lowerNumber,omitempty"`
    // UpperNumber undocumented
    UpperNumber *int `json:"upperNumber,omitempty"`
}
