// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Responsepayload is navigation property rn
func (b *CommandRequestBuilder) Responsepayload() *PayloadResponseRequestBuilder {
	bb := &PayloadResponseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/responsepayload"
	return bb
}

// Entity is navigation property rn
func (b *CommandRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
