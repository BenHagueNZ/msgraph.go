// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ServerProcessedContent undocumented
type ServerProcessedContent struct {
	// Object is the base model of ServerProcessedContent
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ComponentDependencies undocumented
	ComponentDependencies []MetaDataKeyStringPair `json:"componentDependencies,omitempty"`
	// CustomMetadata undocumented
	CustomMetadata []MetaDataKeyValuePair `json:"customMetadata,omitempty"`
	// HTMLStrings undocumented
	HTMLStrings []MetaDataKeyStringPair `json:"htmlStrings,omitempty"`
	// ImageSources undocumented
	ImageSources []MetaDataKeyStringPair `json:"imageSources,omitempty"`
	// Links undocumented
	Links []MetaDataKeyStringPair `json:"links,omitempty"`
	// SearchablePlainTexts undocumented
	SearchablePlainTexts []MetaDataKeyStringPair `json:"searchablePlainTexts,omitempty"`
}

func NewServerProcessedContent() (*ServerProcessedContent, error) {
	newServerProcessedContent := &ServerProcessedContent{
		ODataType: "#microsoft.graph.ServerProcessedContent",
	}
	return newServerProcessedContent, nil
}
