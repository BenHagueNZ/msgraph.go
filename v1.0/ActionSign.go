// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// SignIn is navigation property rn
func (b *SignInRequestBuilder) SignIn() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}