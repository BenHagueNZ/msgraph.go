
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// ExpeditedWindowsQualityUpdateSettings undocumented
type ExpeditedWindowsQualityUpdateSettings struct {
    // Object is the base model of ExpeditedWindowsQualityUpdateSettings
    Object
    // DaysUntilForcedReboot undocumented
    DaysUntilForcedReboot *int `json:"daysUntilForcedReboot,omitempty"`
    // QualityUpdateRelease undocumented
    QualityUpdateRelease *string `json:"qualityUpdateRelease,omitempty"`
}