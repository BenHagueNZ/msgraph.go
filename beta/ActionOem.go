// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// OemWarrantyInformationOnboardingDisableRequestParameter undocumented
type OemWarrantyInformationOnboardingDisableRequestParameter struct {
}

// OemWarrantyInformationOnboardingEnableRequestParameter undocumented
type OemWarrantyInformationOnboardingEnableRequestParameter struct {
}

// Entity is navigation property rn
func (b *OemWarrantyInformationOnboardingRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
