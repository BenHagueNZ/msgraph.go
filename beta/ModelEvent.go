// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Event undocumented
type Event struct {
	// OutlookItem is the base model of Event
	OutlookItem

	ODataType string `json:"@odata.type,omitempty"`
	// AllowNewTimeProposals undocumented
	AllowNewTimeProposals *bool `json:"allowNewTimeProposals,omitempty"`
	// Attendees undocumented
	Attendees []Attendee `json:"attendees,omitempty"`
	// Body undocumented
	Body *ItemBody `json:"body,omitempty"`
	// BodyPreview undocumented
	BodyPreview *string `json:"bodyPreview,omitempty"`
	// CancelledOccurrences undocumented
	CancelledOccurrences []string `json:"cancelledOccurrences,omitempty"`
	// End undocumented
	End *DateTimeTimeZone `json:"end,omitempty"`
	// HasAttachments undocumented
	HasAttachments *bool `json:"hasAttachments,omitempty"`
	// HideAttendees undocumented
	HideAttendees *bool `json:"hideAttendees,omitempty"`
	// Importance undocumented
	Importance *Importance `json:"importance,omitempty"`
	// IsAllDay undocumented
	IsAllDay *bool `json:"isAllDay,omitempty"`
	// IsCancelled undocumented
	IsCancelled *bool `json:"isCancelled,omitempty"`
	// IsDraft undocumented
	IsDraft *bool `json:"isDraft,omitempty"`
	// IsOnlineMeeting undocumented
	IsOnlineMeeting *bool `json:"isOnlineMeeting,omitempty"`
	// IsOrganizer undocumented
	IsOrganizer *bool `json:"isOrganizer,omitempty"`
	// IsReminderOn undocumented
	IsReminderOn *bool `json:"isReminderOn,omitempty"`
	// Location undocumented
	Location *Location `json:"location,omitempty"`
	// Locations undocumented
	Locations []Location `json:"locations,omitempty"`
	// OccurrenceID undocumented
	OccurrenceID *string `json:"occurrenceId,omitempty"`
	// OnlineMeeting undocumented
	OnlineMeeting *OnlineMeetingInfo `json:"onlineMeeting,omitempty"`
	// OnlineMeetingProvider undocumented
	OnlineMeetingProvider *OnlineMeetingProviderType `json:"onlineMeetingProvider,omitempty"`
	// OnlineMeetingURL undocumented
	OnlineMeetingURL *string `json:"onlineMeetingUrl,omitempty"`
	// Organizer undocumented
	Organizer *Recipient `json:"organizer,omitempty"`
	// OriginalEndTimeZone undocumented
	OriginalEndTimeZone *string `json:"originalEndTimeZone,omitempty"`
	// OriginalStart undocumented
	OriginalStart *time.Time `json:"originalStart,omitempty"`
	// OriginalStartTimeZone undocumented
	OriginalStartTimeZone *string `json:"originalStartTimeZone,omitempty"`
	// Recurrence undocumented
	Recurrence *PatternedRecurrence `json:"recurrence,omitempty"`
	// ReminderMinutesBeforeStart undocumented
	ReminderMinutesBeforeStart *int `json:"reminderMinutesBeforeStart,omitempty"`
	// ResponseRequested undocumented
	ResponseRequested *bool `json:"responseRequested,omitempty"`
	// ResponseStatus undocumented
	ResponseStatus *ResponseStatus `json:"responseStatus,omitempty"`
	// Sensitivity undocumented
	Sensitivity *Sensitivity `json:"sensitivity,omitempty"`
	// SeriesMasterID undocumented
	SeriesMasterID *string `json:"seriesMasterId,omitempty"`
	// ShowAs undocumented
	ShowAs *FreeBusyStatus `json:"showAs,omitempty"`
	// Start undocumented
	Start *DateTimeTimeZone `json:"start,omitempty"`
	// Subject undocumented
	Subject *string `json:"subject,omitempty"`
	// TransactionID undocumented
	TransactionID *string `json:"transactionId,omitempty"`
	// Type undocumented
	Type *EventType `json:"type,omitempty"`
	// UID undocumented
	UID *string `json:"uid,omitempty"`
	// WebLink undocumented
	WebLink *string `json:"webLink,omitempty"`
	// Attachments undocumented
	Attachments []Attachment `json:"attachments,omitempty"`
	// Calendar undocumented
	Calendar *Calendar `json:"calendar,omitempty"`
	// ExceptionOccurrences undocumented
	ExceptionOccurrences []Event `json:"exceptionOccurrences,omitempty"`
	// Extensions undocumented
	Extensions []Extension `json:"extensions,omitempty"`
	// Instances undocumented
	Instances []Event `json:"instances,omitempty"`
	// MultiValueExtendedProperties undocumented
	MultiValueExtendedProperties []MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`
	// SingleValueExtendedProperties undocumented
	SingleValueExtendedProperties []SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
}

func NewEvent() (*Event, error) {
	newEvent := &Event{
		ODataType: "#microsoft.graph.Event",
	}
	return newEvent, nil
}

// EventMessage undocumented
type EventMessage struct {
	// Message is the base model of EventMessage
	Message

	ODataType string `json:"@odata.type,omitempty"`
	// EndDateTime undocumented
	EndDateTime *DateTimeTimeZone `json:"endDateTime,omitempty"`
	// IsAllDay undocumented
	IsAllDay *bool `json:"isAllDay,omitempty"`
	// IsDelegated undocumented
	IsDelegated *bool `json:"isDelegated,omitempty"`
	// IsOutOfDate undocumented
	IsOutOfDate *bool `json:"isOutOfDate,omitempty"`
	// Location undocumented
	Location *Location `json:"location,omitempty"`
	// MeetingMessageType undocumented
	MeetingMessageType *MeetingMessageType `json:"meetingMessageType,omitempty"`
	// Recurrence undocumented
	Recurrence *PatternedRecurrence `json:"recurrence,omitempty"`
	// StartDateTime undocumented
	StartDateTime *DateTimeTimeZone `json:"startDateTime,omitempty"`
	// Type undocumented
	Type *EventType `json:"type,omitempty"`
	// Event undocumented
	Event *Event `json:"event,omitempty"`
}

func NewEventMessage() (*EventMessage, error) {
	newEventMessage := &EventMessage{
		ODataType: "#microsoft.graph.EventMessage",
	}
	return newEventMessage, nil
}

// EventMessageDetail undocumented
type EventMessageDetail struct {
	// Object is the base model of EventMessageDetail
	Object

	ODataType string `json:"@odata.type,omitempty"`
}

func NewEventMessageDetail() (*EventMessageDetail, error) {
	newEventMessageDetail := &EventMessageDetail{
		ODataType: "#microsoft.graph.EventMessageDetail",
	}
	return newEventMessageDetail, nil
}

// EventMessageRequestObject undocumented
type EventMessageRequestObject struct {
	// EventMessage is the base model of EventMessageRequestObject
	EventMessage

	ODataType string `json:"@odata.type,omitempty"`
	// AllowNewTimeProposals undocumented
	AllowNewTimeProposals *bool `json:"allowNewTimeProposals,omitempty"`
	// PreviousEndDateTime undocumented
	PreviousEndDateTime *DateTimeTimeZone `json:"previousEndDateTime,omitempty"`
	// PreviousLocation undocumented
	PreviousLocation *Location `json:"previousLocation,omitempty"`
	// PreviousStartDateTime undocumented
	PreviousStartDateTime *DateTimeTimeZone `json:"previousStartDateTime,omitempty"`
	// ResponseRequested undocumented
	ResponseRequested *bool `json:"responseRequested,omitempty"`
}

func NewEventMessageRequestObject() (*EventMessageRequestObject, error) {
	newEventMessageRequestObject := &EventMessageRequestObject{
		ODataType: "#microsoft.graph.EventMessageRequestObject",
	}
	return newEventMessageRequestObject, nil
}

// EventMessageResponse undocumented
type EventMessageResponse struct {
	// EventMessage is the base model of EventMessageResponse
	EventMessage

	ODataType string `json:"@odata.type,omitempty"`
	// ProposedNewTime undocumented
	ProposedNewTime *TimeSlot `json:"proposedNewTime,omitempty"`
	// ResponseType undocumented
	ResponseType *ResponseType `json:"responseType,omitempty"`
}

func NewEventMessageResponse() (*EventMessageResponse, error) {
	newEventMessageResponse := &EventMessageResponse{
		ODataType: "#microsoft.graph.EventMessageResponse",
	}
	return newEventMessageResponse, nil
}
