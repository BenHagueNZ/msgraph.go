// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Attendee undocumented
type Attendee struct {
	// AttendeeBase is the base model of Attendee
	AttendeeBase

	ODataType string `json:"@odata.type"`
	// ProposedNewTime undocumented
	ProposedNewTime *TimeSlot `json:"proposedNewTime,omitempty"`
	// Status undocumented
	Status *ResponseStatus `json:"status,omitempty"`
}

func NewAttendee() (*Attendee, error) {
	newAttendee := &Attendee{
		ODataType: "#microsoft.graph.Attendee",
	}
	return newAttendee, nil
}

// AttendeeAvailability undocumented
type AttendeeAvailability struct {
	// Object is the base model of AttendeeAvailability
	Object

	ODataType string `json:"@odata.type"`
	// Attendee undocumented
	Attendee *AttendeeBase `json:"attendee,omitempty"`
	// Availability undocumented
	Availability *FreeBusyStatus `json:"availability,omitempty"`
}

func NewAttendeeAvailability() (*AttendeeAvailability, error) {
	newAttendeeAvailability := &AttendeeAvailability{
		ODataType: "#microsoft.graph.AttendeeAvailability",
	}
	return newAttendeeAvailability, nil
}

// AttendeeBase undocumented
type AttendeeBase struct {
	// Recipient is the base model of AttendeeBase
	Recipient

	ODataType string `json:"@odata.type"`
	// Type undocumented
	Type *AttendeeType `json:"type,omitempty"`
}

func NewAttendeeBase() (*AttendeeBase, error) {
	newAttendeeBase := &AttendeeBase{
		ODataType: "#microsoft.graph.AttendeeBase",
	}
	return newAttendeeBase, nil
}
