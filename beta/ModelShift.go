// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Shift undocumented
type Shift struct {
	// ChangeTrackedEntity is the base model of Shift
	ChangeTrackedEntity

	ODataType string `json:"@odata.type,omitempty"`
	// DraftShift undocumented
	DraftShift *ShiftItem `json:"draftShift,omitempty"`
	// IsStagedForDeletion undocumented
	IsStagedForDeletion *bool `json:"isStagedForDeletion,omitempty"`
	// SchedulingGroupID undocumented
	SchedulingGroupID *string `json:"schedulingGroupId,omitempty"`
	// SharedShift undocumented
	SharedShift *ShiftItem `json:"sharedShift,omitempty"`
	// UserID undocumented
	UserID *string `json:"userId,omitempty"`
}

func NewShift() (*Shift, error) {
	newShift := &Shift{
		ODataType: "#microsoft.graph.Shift",
	}
	return newShift, nil
}

// ShiftActivity undocumented
type ShiftActivity struct {
	// Object is the base model of ShiftActivity
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// IsPaid undocumented
	IsPaid *bool `json:"isPaid,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// Theme undocumented
	Theme *ScheduleEntityTheme `json:"theme,omitempty"`
}

func NewShiftActivity() (*ShiftActivity, error) {
	newShiftActivity := &ShiftActivity{
		ODataType: "#microsoft.graph.ShiftActivity",
	}
	return newShiftActivity, nil
}

// ShiftAvailability undocumented
type ShiftAvailability struct {
	// Object is the base model of ShiftAvailability
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Recurrence undocumented
	Recurrence *PatternedRecurrence `json:"recurrence,omitempty"`
	// TimeSlots undocumented
	TimeSlots []TimeRange `json:"timeSlots,omitempty"`
	// TimeZone undocumented
	TimeZone *string `json:"timeZone,omitempty"`
}

func NewShiftAvailability() (*ShiftAvailability, error) {
	newShiftAvailability := &ShiftAvailability{
		ODataType: "#microsoft.graph.ShiftAvailability",
	}
	return newShiftAvailability, nil
}

// ShiftItem undocumented
type ShiftItem struct {
	// ScheduleEntity is the base model of ShiftItem
	ScheduleEntity

	ODataType string `json:"@odata.type,omitempty"`
	// Activities undocumented
	Activities []ShiftActivity `json:"activities,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Notes undocumented
	Notes *string `json:"notes,omitempty"`
}

func NewShiftItem() (*ShiftItem, error) {
	newShiftItem := &ShiftItem{
		ODataType: "#microsoft.graph.ShiftItem",
	}
	return newShiftItem, nil
}

// ShiftPreferences undocumented
type ShiftPreferences struct {
	// ChangeTrackedEntity is the base model of ShiftPreferences
	ChangeTrackedEntity

	ODataType string `json:"@odata.type,omitempty"`
	// Availability undocumented
	Availability []ShiftAvailability `json:"availability,omitempty"`
}

func NewShiftPreferences() (*ShiftPreferences, error) {
	newShiftPreferences := &ShiftPreferences{
		ODataType: "#microsoft.graph.ShiftPreferences",
	}
	return newShiftPreferences, nil
}
