// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// InvitedUser is navigation property
func (b *InvitationRequestBuilder) InvitedUser() *UserRequestBuilder {
	bb := &UserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/invitedUser"
	return bb
}
