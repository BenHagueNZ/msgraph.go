// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Entity is navigation property rn
func (b *AgedAccountsPayableRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *AgedAccountsReceivableRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
