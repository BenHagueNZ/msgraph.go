// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ProfilePhoto undocumented
type ProfilePhoto struct {
	// Entity is the base model of ProfilePhoto
	Entity

	ODataType string `json:"@odata.type"`
	// Height undocumented
	Height *int `json:"height,omitempty"`
	// Width undocumented
	Width *int `json:"width,omitempty"`
}

func NewProfilePhoto() (*ProfilePhoto, error) {
	newProfilePhoto := &ProfilePhoto{
		ODataType: "#microsoft.graph.ProfilePhoto",
	}
	return newProfilePhoto, nil
}
