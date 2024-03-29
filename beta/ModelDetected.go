// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// DetectedApp undocumented
type DetectedApp struct {
	// Entity is the base model of DetectedApp
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// DeviceCount undocumented
	DeviceCount *int `json:"deviceCount,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Platform undocumented
	Platform *DetectedAppPlatformType `json:"platform,omitempty"`
	// Publisher undocumented
	Publisher *string `json:"publisher,omitempty"`
	// SizeInByte undocumented
	SizeInByte *int `json:"sizeInByte,omitempty"`
	// Version undocumented
	Version *string `json:"version,omitempty"`
	// ManagedDevices undocumented
	ManagedDevices []ManagedDevice `json:"managedDevices,omitempty"`
}

func NewDetectedApp() (*DetectedApp, error) {
	newDetectedApp := &DetectedApp{
		ODataType: "#microsoft.graph.DetectedApp",
	}
	return newDetectedApp, nil
}

// DetectedSensitiveContent undocumented
type DetectedSensitiveContent struct {
	// DetectedSensitiveContentBase is the base model of DetectedSensitiveContent
	DetectedSensitiveContentBase

	ODataType string `json:"@odata.type,omitempty"`
	// ClassificationAttributes undocumented
	ClassificationAttributes []ClassificationAttribute `json:"classificationAttributes,omitempty"`
	// ClassificationMethod undocumented
	ClassificationMethod *ClassificationMethod `json:"classificationMethod,omitempty"`
	// Matches undocumented
	Matches []SensitiveContentLocation `json:"matches,omitempty"`
	// Scope undocumented
	Scope *SensitiveTypeScope `json:"scope,omitempty"`
	// SensitiveTypeSource undocumented
	SensitiveTypeSource *SensitiveTypeSource `json:"sensitiveTypeSource,omitempty"`
}

func NewDetectedSensitiveContent() (*DetectedSensitiveContent, error) {
	newDetectedSensitiveContent := &DetectedSensitiveContent{
		ODataType: "#microsoft.graph.DetectedSensitiveContent",
	}
	return newDetectedSensitiveContent, nil
}

// DetectedSensitiveContentBase undocumented
type DetectedSensitiveContentBase struct {
	// Object is the base model of DetectedSensitiveContentBase
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Confidence undocumented
	Confidence *int `json:"confidence,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ID undocumented
	ID *UUID `json:"id,omitempty"`
	// RecommendedConfidence undocumented
	RecommendedConfidence *int `json:"recommendedConfidence,omitempty"`
	// UniqueCount undocumented
	UniqueCount *int `json:"uniqueCount,omitempty"`
}

func NewDetectedSensitiveContentBase() (*DetectedSensitiveContentBase, error) {
	newDetectedSensitiveContentBase := &DetectedSensitiveContentBase{
		ODataType: "#microsoft.graph.DetectedSensitiveContentBase",
	}
	return newDetectedSensitiveContentBase, nil
}

// DetectedSensitiveContentWrapper undocumented
type DetectedSensitiveContentWrapper struct {
	// Object is the base model of DetectedSensitiveContentWrapper
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Classification undocumented
	Classification []DetectedSensitiveContent `json:"classification,omitempty"`
}

func NewDetectedSensitiveContentWrapper() (*DetectedSensitiveContentWrapper, error) {
	newDetectedSensitiveContentWrapper := &DetectedSensitiveContentWrapper{
		ODataType: "#microsoft.graph.DetectedSensitiveContentWrapper",
	}
	return newDetectedSensitiveContentWrapper, nil
}
