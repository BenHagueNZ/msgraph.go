// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// VPPTokenSyncLicensesRequestParameter undocumented
type VPPTokenSyncLicensesRequestParameter struct {
}

// Entity is navigation property rn
func (b *VPPTokenRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
