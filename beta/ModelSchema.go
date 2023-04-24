// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Schema undocumented
type Schema struct {
	// Entity is the base model of Schema
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// BaseType undocumented
	BaseType *string `json:"baseType,omitempty"`
	// Properties undocumented
	Properties []Property `json:"properties,omitempty"`
}

func NewSchema() (*Schema, error) {
	newSchema := &Schema{
		ODataType: "#microsoft.graph.Schema",
	}
	return newSchema, nil
}

// SchemaExtension undocumented
type SchemaExtension struct {
	// Entity is the base model of SchemaExtension
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// Owner undocumented
	Owner *string `json:"owner,omitempty"`
	// Properties undocumented
	Properties []ExtensionSchemaProperty `json:"properties,omitempty"`
	// Status undocumented
	Status *string `json:"status,omitempty"`
	// TargetTypes undocumented
	TargetTypes []string `json:"targetTypes,omitempty"`
}

func NewSchemaExtension() (*SchemaExtension, error) {
	newSchemaExtension := &SchemaExtension{
		ODataType: "#microsoft.graph.SchemaExtension",
	}
	return newSchemaExtension, nil
}
