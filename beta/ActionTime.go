// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// TimeCardCollectionClockInRequestParameter undocumented
type TimeCardCollectionClockInRequestParameter struct {
	// AtApprovedLocation undocumented
	AtApprovedLocation *bool `json:"atApprovedLocation,omitempty"`
	// OnBehalfOfUserID undocumented
	OnBehalfOfUserID *string `json:"onBehalfOfUserId,omitempty"`
	// Notes undocumented
	Notes *ItemBody `json:"notes,omitempty"`
}

// TimeCardClockOutRequestParameter undocumented
type TimeCardClockOutRequestParameter struct {
	// AtApprovedLocation undocumented
	AtApprovedLocation *bool `json:"atApprovedLocation,omitempty"`
	// Notes undocumented
	Notes *ItemBody `json:"notes,omitempty"`
}

// TimeCardConfirmRequestParameter undocumented
type TimeCardConfirmRequestParameter struct {
}

// TimeCardEndBreakRequestParameter undocumented
type TimeCardEndBreakRequestParameter struct {
	// AtApprovedLocation undocumented
	AtApprovedLocation *bool `json:"atApprovedLocation,omitempty"`
	// Notes undocumented
	Notes *ItemBody `json:"notes,omitempty"`
}

// TimeCardStartBreakRequestParameter undocumented
type TimeCardStartBreakRequestParameter struct {
	// AtApprovedLocation undocumented
	AtApprovedLocation *bool `json:"atApprovedLocation,omitempty"`
	// Notes undocumented
	Notes *ItemBody `json:"notes,omitempty"`
}

// Entity is navigation property rn
func (b *TimeCardRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *TimeOffRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *TimeOffReasonRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
