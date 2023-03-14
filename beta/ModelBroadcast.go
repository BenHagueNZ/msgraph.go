
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// BroadcastMeetingCaptionSettings undocumented
type BroadcastMeetingCaptionSettings struct {
    // Object is the base model of BroadcastMeetingCaptionSettings
    Object
    // IsCaptionEnabled undocumented
    IsCaptionEnabled *bool `json:"isCaptionEnabled,omitempty"`
    // SpokenLanguage undocumented
    SpokenLanguage *string `json:"spokenLanguage,omitempty"`
    // TranslationLanguages undocumented
    TranslationLanguages []string `json:"translationLanguages,omitempty"`
}

// BroadcastMeetingSettings undocumented
type BroadcastMeetingSettings struct {
    // Object is the base model of BroadcastMeetingSettings
    Object
    // AllowedAudience undocumented
    AllowedAudience *BroadcastMeetingAudience `json:"allowedAudience,omitempty"`
    // Captions undocumented
    Captions *BroadcastMeetingCaptionSettings `json:"captions,omitempty"`
    // IsAttendeeReportEnabled undocumented
    IsAttendeeReportEnabled *bool `json:"isAttendeeReportEnabled,omitempty"`
    // IsQuestionAndAnswerEnabled undocumented
    IsQuestionAndAnswerEnabled *bool `json:"isQuestionAndAnswerEnabled,omitempty"`
    // IsRecordingEnabled undocumented
    IsRecordingEnabled *bool `json:"isRecordingEnabled,omitempty"`
    // IsVideoOnDemandEnabled undocumented
    IsVideoOnDemandEnabled *bool `json:"isVideoOnDemandEnabled,omitempty"`
}
