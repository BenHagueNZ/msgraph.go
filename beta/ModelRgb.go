// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// RgbColor undocumented
type RgbColor struct {
	// Object is the base model of RgbColor
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// B undocumented
	B *byte `json:"b,omitempty"`
	// G undocumented
	G *byte `json:"g,omitempty"`
	// R undocumented
	R *byte `json:"r,omitempty"`
}

func NewRgbColor() (*RgbColor, error) {
	newRgbColor := &RgbColor{
		ODataType: "#microsoft.graph.RgbColor",
	}
	return newRgbColor, nil
}