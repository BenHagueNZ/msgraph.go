// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// SignInCollectionConfirmCompromisedRequestParameter undocumented
type SignInCollectionConfirmCompromisedRequestParameter struct {
	// RequestIDs undocumented
	RequestIDs []string `json:"requestIds,omitempty"`
}

// SignInCollectionConfirmSafeRequestParameter undocumented
type SignInCollectionConfirmSafeRequestParameter struct {
	// RequestIDs undocumented
	RequestIDs []string `json:"requestIds,omitempty"`
}

// Entity is navigation property rn
func (b *SignInRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
