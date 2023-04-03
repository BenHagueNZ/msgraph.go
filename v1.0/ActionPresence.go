// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// PresenceClearPresenceRequestParameter undocumented
type PresenceClearPresenceRequestParameter struct {
	// SessionID undocumented
	SessionID *string `json:"sessionId,omitempty"`
}

// PresenceClearUserPreferredPresenceRequestParameter undocumented
type PresenceClearUserPreferredPresenceRequestParameter struct {
}

// PresenceSetPresenceRequestParameter undocumented
type PresenceSetPresenceRequestParameter struct {
	// SessionID undocumented
	SessionID *string `json:"sessionId,omitempty"`
	// Availability undocumented
	Availability *string `json:"availability,omitempty"`
	// Activity undocumented
	Activity *string `json:"activity,omitempty"`
	// ExpirationDuration undocumented
	ExpirationDuration *Duration `json:"expirationDuration,omitempty"`
}

// PresenceSetUserPreferredPresenceRequestParameter undocumented
type PresenceSetUserPreferredPresenceRequestParameter struct {
	// Availability undocumented
	Availability *string `json:"availability,omitempty"`
	// Activity undocumented
	Activity *string `json:"activity,omitempty"`
	// ExpirationDuration undocumented
	ExpirationDuration *Duration `json:"expirationDuration,omitempty"`
}

// Entity is navigation property rn
func (b *PresenceRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
