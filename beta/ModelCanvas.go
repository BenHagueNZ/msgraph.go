// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// CanvasLayout undocumented
type CanvasLayout struct {
	// Entity is the base model of CanvasLayout
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// HorizontalSections undocumented
	HorizontalSections []HorizontalSection `json:"horizontalSections,omitempty"`
	// VerticalSection undocumented
	VerticalSection *VerticalSection `json:"verticalSection,omitempty"`
}

func NewCanvasLayout() (*CanvasLayout, error) {
	newCanvasLayout := &CanvasLayout{
		ODataType: "#microsoft.graph.CanvasLayout",
	}
	return newCanvasLayout, nil
}