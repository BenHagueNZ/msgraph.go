
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// JoinMeetingIDMeetingInfo undocumented
type JoinMeetingIDMeetingInfo struct {
    // MeetingInfo is the base model of JoinMeetingIDMeetingInfo
    MeetingInfo
    // JoinMeetingID undocumented
    JoinMeetingID *string `json:"joinMeetingId,omitempty"`
    // Passcode undocumented
    Passcode *string `json:"passcode,omitempty"`
}

// JoinMeetingIDSettings undocumented
type JoinMeetingIDSettings struct {
    // Object is the base model of JoinMeetingIDSettings
    Object
    // IsPasscodeRequired undocumented
    IsPasscodeRequired *bool `json:"isPasscodeRequired,omitempty"`
    // JoinMeetingID undocumented
    JoinMeetingID *string `json:"joinMeetingId,omitempty"`
    // Passcode undocumented
    Passcode *string `json:"passcode,omitempty"`
}
