// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// EBookInstallSummary is navigation property rn
func (b *EBookInstallSummaryRequestBuilder) EBookInstallSummary() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}