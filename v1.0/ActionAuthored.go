// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// AuthoredNote is navigation property rn
func (b *AuthoredNoteRequestBuilder) AuthoredNote() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
