// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// SpaApplication undocumented
type SpaApplication struct {
	// Object is the base model of SpaApplication
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// RedirectUris undocumented
	RedirectUris []string `json:"redirectUris,omitempty"`
}

func NewSpaApplication() (*SpaApplication, error) {
	newSpaApplication := &SpaApplication{
		ODataType: "#microsoft.graph.SpaApplication",
	}
	return newSpaApplication, nil
}