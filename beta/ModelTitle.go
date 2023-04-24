// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// TitleArea undocumented
type TitleArea struct {
	// Object is the base model of TitleArea
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AlternativeText undocumented
	AlternativeText *string `json:"alternativeText,omitempty"`
	// EnableGradientEffect undocumented
	EnableGradientEffect *bool `json:"enableGradientEffect,omitempty"`
	// ImageWebURL undocumented
	ImageWebURL *string `json:"imageWebUrl,omitempty"`
	// Layout undocumented
	Layout *TitleAreaLayoutType `json:"layout,omitempty"`
	// ServerProcessedContent undocumented
	ServerProcessedContent *ServerProcessedContent `json:"serverProcessedContent,omitempty"`
	// ShowAuthor undocumented
	ShowAuthor *bool `json:"showAuthor,omitempty"`
	// ShowPublishedDate undocumented
	ShowPublishedDate *bool `json:"showPublishedDate,omitempty"`
	// ShowTextBlockAboveTitle undocumented
	ShowTextBlockAboveTitle *bool `json:"showTextBlockAboveTitle,omitempty"`
	// TextAboveTitle undocumented
	TextAboveTitle *string `json:"textAboveTitle,omitempty"`
	// TextAlignment undocumented
	TextAlignment *TitleAreaTextAlignmentType `json:"textAlignment,omitempty"`
}

func NewTitleArea() (*TitleArea, error) {
	newTitleArea := &TitleArea{
		ODataType: "#microsoft.graph.TitleArea",
	}
	return newTitleArea, nil
}