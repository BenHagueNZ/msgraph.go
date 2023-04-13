// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Extension undocumented
type Extension struct {
	// Entity is the base model of Extension
	Entity

	ODataType string `json:"@odata.type,omitempty"`
}

func NewExtension() (*Extension, error) {
	newExtension := &Extension{
		ODataType: "#microsoft.graph.Extension",
	}
	return newExtension, nil
}

// ExtensionProperty undocumented
type ExtensionProperty struct {
	// DirectoryObject is the base model of ExtensionProperty
	DirectoryObject

	ODataType string `json:"@odata.type,omitempty"`
	// AppDisplayName undocumented
	AppDisplayName *string `json:"appDisplayName,omitempty"`
	// DataType undocumented
	DataType *string `json:"dataType,omitempty"`
	// IsSyncedFromOnPremises undocumented
	IsSyncedFromOnPremises *bool `json:"isSyncedFromOnPremises,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// TargetObjects undocumented
	TargetObjects []string `json:"targetObjects,omitempty"`
}

func NewExtensionProperty() (*ExtensionProperty, error) {
	newExtensionProperty := &ExtensionProperty{
		ODataType: "#microsoft.graph.ExtensionProperty",
	}
	return newExtensionProperty, nil
}

// ExtensionSchemaProperty undocumented
type ExtensionSchemaProperty struct {
	// Object is the base model of ExtensionSchemaProperty
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
}

func NewExtensionSchemaProperty() (*ExtensionSchemaProperty, error) {
	newExtensionSchemaProperty := &ExtensionSchemaProperty{
		ODataType: "#microsoft.graph.ExtensionSchemaProperty",
	}
	return newExtensionSchemaProperty, nil
}
