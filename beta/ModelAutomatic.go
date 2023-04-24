// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// AutomaticRepliesMailTips undocumented
type AutomaticRepliesMailTips struct {
	// Object is the base model of AutomaticRepliesMailTips
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Message undocumented
	Message *string `json:"message,omitempty"`
	// MessageLanguage undocumented
	MessageLanguage *LocaleInfo `json:"messageLanguage,omitempty"`
	// ScheduledEndTime undocumented
	ScheduledEndTime *DateTimeTimeZone `json:"scheduledEndTime,omitempty"`
	// ScheduledStartTime undocumented
	ScheduledStartTime *DateTimeTimeZone `json:"scheduledStartTime,omitempty"`
}

func NewAutomaticRepliesMailTips() (*AutomaticRepliesMailTips, error) {
	newAutomaticRepliesMailTips := &AutomaticRepliesMailTips{
		ODataType: "#microsoft.graph.AutomaticRepliesMailTips",
	}
	return newAutomaticRepliesMailTips, nil
}

// AutomaticRepliesSetting undocumented
type AutomaticRepliesSetting struct {
	// Object is the base model of AutomaticRepliesSetting
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ExternalAudience undocumented
	ExternalAudience *ExternalAudienceScope `json:"externalAudience,omitempty"`
	// ExternalReplyMessage undocumented
	ExternalReplyMessage *string `json:"externalReplyMessage,omitempty"`
	// InternalReplyMessage undocumented
	InternalReplyMessage *string `json:"internalReplyMessage,omitempty"`
	// ScheduledEndDateTime undocumented
	ScheduledEndDateTime *DateTimeTimeZone `json:"scheduledEndDateTime,omitempty"`
	// ScheduledStartDateTime undocumented
	ScheduledStartDateTime *DateTimeTimeZone `json:"scheduledStartDateTime,omitempty"`
	// Status undocumented
	Status *AutomaticRepliesStatus `json:"status,omitempty"`
}

func NewAutomaticRepliesSetting() (*AutomaticRepliesSetting, error) {
	newAutomaticRepliesSetting := &AutomaticRepliesSetting{
		ODataType: "#microsoft.graph.AutomaticRepliesSetting",
	}
	return newAutomaticRepliesSetting, nil
}