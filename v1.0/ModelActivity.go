// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// ActivityBasedTimeoutPolicy undocumented
type ActivityBasedTimeoutPolicy struct {
	// StsPolicy is the base model of ActivityBasedTimeoutPolicy
	StsPolicy

	ODataType string `json:"@odata.type"`
}

func NewActivityBasedTimeoutPolicy() (*ActivityBasedTimeoutPolicy, error) {
	newActivityBasedTimeoutPolicy := &ActivityBasedTimeoutPolicy{
		ODataType: "#microsoft.graph.ActivityBasedTimeoutPolicy",
	}
	return newActivityBasedTimeoutPolicy, nil
}

// ActivityHistoryItem undocumented
type ActivityHistoryItem struct {
	// Entity is the base model of ActivityHistoryItem
	Entity

	ODataType string `json:"@odata.type"`
	// ActiveDurationSeconds undocumented
	ActiveDurationSeconds *int `json:"activeDurationSeconds,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// ExpirationDateTime undocumented
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// LastActiveDateTime undocumented
	LastActiveDateTime *time.Time `json:"lastActiveDateTime,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// StartedDateTime undocumented
	StartedDateTime *time.Time `json:"startedDateTime,omitempty"`
	// Status undocumented
	Status *Status `json:"status,omitempty"`
	// UserTimezone undocumented
	UserTimezone *string `json:"userTimezone,omitempty"`
	// Activity undocumented
	Activity *UserActivity `json:"activity,omitempty"`
}

func NewActivityHistoryItem() (*ActivityHistoryItem, error) {
	newActivityHistoryItem := &ActivityHistoryItem{
		ODataType: "#microsoft.graph.ActivityHistoryItem",
	}
	return newActivityHistoryItem, nil
}
