// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// OrganizerMeetingInfo undocumented
type OrganizerMeetingInfo struct {
	// MeetingInfo is the base model of OrganizerMeetingInfo
	MeetingInfo

	ODataType string `json:"@odata.type,omitempty"`
	// Organizer undocumented
	Organizer *IdentitySet `json:"organizer,omitempty"`
}

func NewOrganizerMeetingInfo() (*OrganizerMeetingInfo, error) {
	newOrganizerMeetingInfo := &OrganizerMeetingInfo{
		ODataType: "#microsoft.graph.OrganizerMeetingInfo",
	}
	return newOrganizerMeetingInfo, nil
}
