// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// RequestObjectStopRequestParameter undocumented
type RequestObjectStopRequestParameter struct {
}

// RequestObjectRecordDecisionsRequestParameter undocumented
type RequestObjectRecordDecisionsRequestParameter struct {
	// ReviewResult undocumented
	ReviewResult *string `json:"reviewResult,omitempty"`
	// Justification undocumented
	Justification *string `json:"justification,omitempty"`
}

// Entity is navigation property rn
func (b *RequestObjectRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
