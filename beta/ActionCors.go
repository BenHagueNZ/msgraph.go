// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Entity is navigation property rn
func (b *CorsConfiguration_v2RequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}