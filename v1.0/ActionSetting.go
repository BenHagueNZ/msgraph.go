// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// SettingStateDeviceSummary is navigation property rn
func (b *SettingStateDeviceSummaryRequestBuilder) SettingStateDeviceSummary() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}