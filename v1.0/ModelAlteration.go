// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// AlterationResponse undocumented
type AlterationResponse struct {
	// Object is the base model of AlterationResponse
	Object
	// OriginalQueryString undocumented
	OriginalQueryString *string `json:"originalQueryString,omitempty"`
	// QueryAlteration undocumented
	QueryAlteration *SearchAlteration `json:"queryAlteration,omitempty"`
	// QueryAlterationType undocumented
	QueryAlterationType *SearchAlterationType `json:"queryAlterationType,omitempty"`
}
