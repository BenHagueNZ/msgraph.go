// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// URIClickSecurityState undocumented
type URIClickSecurityState struct {
	// Object is the base model of URIClickSecurityState
	Object

	ODataType string `json:"@odata.type"`
	// ClickAction undocumented
	ClickAction *string `json:"clickAction,omitempty"`
	// ClickDateTime undocumented
	ClickDateTime *time.Time `json:"clickDateTime,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// SourceID undocumented
	SourceID *string `json:"sourceId,omitempty"`
	// URIDomain undocumented
	URIDomain *string `json:"uriDomain,omitempty"`
	// Verdict undocumented
	Verdict *string `json:"verdict,omitempty"`
}

func NewURIClickSecurityState() (*URIClickSecurityState, error) {
	newURIClickSecurityState := &URIClickSecurityState{
		ODataType: "#microsoft.graph.UriClickSecurityState",
	}
	return newURIClickSecurityState, nil
}
