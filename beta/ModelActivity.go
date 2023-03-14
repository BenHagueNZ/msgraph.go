
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// ActivityBasedTimeoutPolicy undocumented
type ActivityBasedTimeoutPolicy struct {
    // StsPolicy is the base model of ActivityBasedTimeoutPolicy
    StsPolicy
}

// ActivityHistoryItem undocumented
type ActivityHistoryItem struct {
    // Entity is the base model of ActivityHistoryItem
    Entity
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

// ActivityStatistics undocumented
type ActivityStatistics struct {
    // Entity is the base model of ActivityStatistics
    Entity
    // Activity undocumented
    Activity *AnalyticsActivityType `json:"activity,omitempty"`
    // Duration undocumented
    Duration *Duration `json:"duration,omitempty"`
    // EndDate undocumented
    EndDate *Date `json:"endDate,omitempty"`
    // StartDate undocumented
    StartDate *Date `json:"startDate,omitempty"`
    // TimeZoneUsed undocumented
    TimeZoneUsed *string `json:"timeZoneUsed,omitempty"`
}
