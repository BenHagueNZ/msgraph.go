// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// AlertCollectionUpdateAlertsRequestParameter undocumented
type AlertCollectionUpdateAlertsRequestParameter struct {
	// Value undocumented
	Value []Alert `json:"value,omitempty"`
}

// Entity is navigation property rn
func (b *AlertRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}