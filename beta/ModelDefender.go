
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// DefenderDetectedMalwareActions undocumented
type DefenderDetectedMalwareActions struct {
    // Object is the base model of DefenderDetectedMalwareActions
    Object
    // HighSeverity undocumented
    HighSeverity *DefenderThreatAction `json:"highSeverity,omitempty"`
    // LowSeverity undocumented
    LowSeverity *DefenderThreatAction `json:"lowSeverity,omitempty"`
    // ModerateSeverity undocumented
    ModerateSeverity *DefenderThreatAction `json:"moderateSeverity,omitempty"`
    // SevereSeverity undocumented
    SevereSeverity *DefenderThreatAction `json:"severeSeverity,omitempty"`
}
