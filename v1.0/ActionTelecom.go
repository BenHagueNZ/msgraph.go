// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// TelecomExpenseManagementPartner is navigation property rn
func (b *TelecomExpenseManagementPartnerRequestBuilder) TelecomExpenseManagementPartner() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
