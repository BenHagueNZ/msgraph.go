// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Entity is navigation property rn
func (b *CartToClassAssociationRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
