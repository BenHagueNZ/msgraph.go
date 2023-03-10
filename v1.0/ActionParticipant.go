// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ParticipantCollectionInviteRequestParameter undocumented
type ParticipantCollectionInviteRequestParameter struct {
	// Participants undocumented
	Participants []InvitationParticipantInfo `json:"participants,omitempty"`
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
}

// ParticipantMuteRequestParameter undocumented
type ParticipantMuteRequestParameter struct {
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
}

// ParticipantStartHoldMusicRequestParameter undocumented
type ParticipantStartHoldMusicRequestParameter struct {
	// CustomPrompt undocumented
	CustomPrompt *Prompt `json:"customPrompt,omitempty"`
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
}

// ParticipantStopHoldMusicRequestParameter undocumented
type ParticipantStopHoldMusicRequestParameter struct {
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
}

// Call is navigation property
func (b *ParticipantJoiningNotificationRequestBuilder) Call() *CallRequestBuilder {
	bb := &CallRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/call"
	return bb
}

// Call is navigation property
func (b *ParticipantLeftNotificationRequestBuilder) Call() *CallRequestBuilder {
	bb := &CallRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/call"
	return bb
}
