// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// SourceColumn is navigation property rn
func (b *ColumnDefinitionRequestBuilder) SourceColumn() *ColumnDefinitionRequestBuilder {
	bb := &ColumnDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sourceColumn"
	return bb
}

// Entity is navigation property rn
func (b *ColumnDefinitionRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *ColumnLinkRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
