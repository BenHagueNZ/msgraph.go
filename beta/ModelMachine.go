// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// MachineLearningDetectedSensitiveContent undocumented
type MachineLearningDetectedSensitiveContent struct {
	// DetectedSensitiveContent is the base model of MachineLearningDetectedSensitiveContent
	DetectedSensitiveContent

	ODataType string `json:"@odata.type,omitempty"`
	// MatchTolerance undocumented
	MatchTolerance *MlClassificationMatchTolerance `json:"matchTolerance,omitempty"`
	// ModelVersion undocumented
	ModelVersion *string `json:"modelVersion,omitempty"`
}

func NewMachineLearningDetectedSensitiveContent() (*MachineLearningDetectedSensitiveContent, error) {
	newMachineLearningDetectedSensitiveContent := &MachineLearningDetectedSensitiveContent{
		ODataType: "#microsoft.graph.MachineLearningDetectedSensitiveContent",
	}
	return newMachineLearningDetectedSensitiveContent, nil
}