// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// InsightIdentity undocumented
type InsightIdentity struct {
	// Object is the base model of InsightIdentity
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Address undocumented
	Address *string `json:"address,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
}

func NewInsightIdentity() (*InsightIdentity, error) {
	newInsightIdentity := &InsightIdentity{
		ODataType: "#microsoft.graph.InsightIdentity",
	}
	return newInsightIdentity, nil
}
