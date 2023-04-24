// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// RedirectSingleSignOnExtension undocumented
type RedirectSingleSignOnExtension struct {
	// SingleSignOnExtension is the base model of RedirectSingleSignOnExtension
	SingleSignOnExtension

	ODataType string `json:"@odata.type,omitempty"`
	// Configurations undocumented
	Configurations []KeyTypedValuePair `json:"configurations,omitempty"`
	// ExtensionIdentifier undocumented
	ExtensionIdentifier *string `json:"extensionIdentifier,omitempty"`
	// TeamIdentifier undocumented
	TeamIdentifier *string `json:"teamIdentifier,omitempty"`
	// URLPrefixes undocumented
	URLPrefixes []string `json:"urlPrefixes,omitempty"`
}

func NewRedirectSingleSignOnExtension() (*RedirectSingleSignOnExtension, error) {
	newRedirectSingleSignOnExtension := &RedirectSingleSignOnExtension{
		ODataType: "#microsoft.graph.RedirectSingleSignOnExtension",
	}
	return newRedirectSingleSignOnExtension, nil
}

// RedirectURISettings undocumented
type RedirectURISettings struct {
	// Object is the base model of RedirectURISettings
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Index undocumented
	Index *int `json:"index,omitempty"`
	// URI undocumented
	URI *string `json:"uri,omitempty"`
}

func NewRedirectURISettings() (*RedirectURISettings, error) {
	newRedirectURISettings := &RedirectURISettings{
		ODataType: "#microsoft.graph.RedirectUriSettings",
	}
	return newRedirectURISettings, nil
}