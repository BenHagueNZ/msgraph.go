// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Resource is navigation property rn
func (b *UsedInsightRequestBuilder) Resource() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/resource"
	return bb
}

// Entity is navigation property rn
func (b *UsedInsightRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}