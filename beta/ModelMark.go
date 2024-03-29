// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// MarkContent undocumented
type MarkContent struct {
	// LabelActionBase is the base model of MarkContent
	LabelActionBase

	ODataType string `json:"@odata.type,omitempty"`
	// FontColor undocumented
	FontColor *string `json:"fontColor,omitempty"`
	// FontSize undocumented
	FontSize *int `json:"fontSize,omitempty"`
	// Text undocumented
	Text *string `json:"text,omitempty"`
}

func NewMarkContent() (*MarkContent, error) {
	newMarkContent := &MarkContent{
		ODataType: "#microsoft.graph.MarkContent",
	}
	return newMarkContent, nil
}
