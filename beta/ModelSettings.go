// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Settings undocumented
type Settings struct {
	// Object is the base model of Settings
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// HasGraphMailbox undocumented
	HasGraphMailbox *bool `json:"hasGraphMailbox,omitempty"`
	// HasLicense undocumented
	HasLicense *bool `json:"hasLicense,omitempty"`
	// HasOptedOut undocumented
	HasOptedOut *bool `json:"hasOptedOut,omitempty"`
}

func NewSettings() (*Settings, error) {
	newSettings := &Settings{
		ODataType: "#microsoft.graph.Settings",
	}
	return newSettings, nil
}
