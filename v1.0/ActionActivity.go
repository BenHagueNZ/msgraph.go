// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Activity is navigation property rn
func (b *ActivityHistoryItemRequestBuilder) Activity() *UserActivityRequestBuilder {
	bb := &UserActivityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/activity"
	return bb
}

// ActivityHistoryItem is navigation property rn
func (b *ActivityHistoryItemRequestBuilder) ActivityHistoryItem() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
