// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ImplicitGrantSettings undocumented
type ImplicitGrantSettings struct {
	// Object is the base model of ImplicitGrantSettings
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// EnableAccessTokenIssuance undocumented
	EnableAccessTokenIssuance *bool `json:"enableAccessTokenIssuance,omitempty"`
	// EnableIDTokenIssuance undocumented
	EnableIDTokenIssuance *bool `json:"enableIdTokenIssuance,omitempty"`
}

func NewImplicitGrantSettings() (*ImplicitGrantSettings, error) {
	newImplicitGrantSettings := &ImplicitGrantSettings{
		ODataType: "#microsoft.graph.ImplicitGrantSettings",
	}
	return newImplicitGrantSettings, nil
}
