// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Picture undocumented
type Picture struct {
	// Entity is the base model of Picture
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Content undocumented
	Content *Stream `json:"content,omitempty"`
	// ContentType undocumented
	ContentType *string `json:"contentType,omitempty"`
	// Height undocumented
	Height *int `json:"height,omitempty"`
	// Width undocumented
	Width *int `json:"width,omitempty"`
}

func NewPicture() (*Picture, error) {
	newPicture := &Picture{
		ODataType: "#microsoft.graph.Picture",
	}
	return newPicture, nil
}
