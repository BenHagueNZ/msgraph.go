// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// CreatedByUser is navigation property rn
func (b *BaseItemRequestBuilder) CreatedByUser() *UserRequestBuilder {
	bb := &UserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/createdByUser"
	return bb
}

// LastModifiedByUser is navigation property rn
func (b *BaseItemRequestBuilder) LastModifiedByUser() *UserRequestBuilder {
	bb := &UserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/lastModifiedByUser"
	return bb
}

// Entity is navigation property rn
func (b *BaseItemRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *BaseItemVersionRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}