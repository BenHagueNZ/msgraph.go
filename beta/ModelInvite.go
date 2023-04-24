// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// InviteNewBotResponse undocumented
type InviteNewBotResponse struct {
	// ParticipantJoiningResponse is the base model of InviteNewBotResponse
	ParticipantJoiningResponse

	ODataType string `json:"@odata.type,omitempty"`
	// InviteURI undocumented
	InviteURI *string `json:"inviteUri,omitempty"`
}

func NewInviteNewBotResponse() (*InviteNewBotResponse, error) {
	newInviteNewBotResponse := &InviteNewBotResponse{
		ODataType: "#microsoft.graph.InviteNewBotResponse",
	}
	return newInviteNewBotResponse, nil
}

// InviteParticipantsOperation undocumented
type InviteParticipantsOperation struct {
	// CommsOperation is the base model of InviteParticipantsOperation
	CommsOperation

	ODataType string `json:"@odata.type,omitempty"`
	// Participants undocumented
	Participants []InvitationParticipantInfo `json:"participants,omitempty"`
}

func NewInviteParticipantsOperation() (*InviteParticipantsOperation, error) {
	newInviteParticipantsOperation := &InviteParticipantsOperation{
		ODataType: "#microsoft.graph.InviteParticipantsOperation",
	}
	return newInviteParticipantsOperation, nil
}