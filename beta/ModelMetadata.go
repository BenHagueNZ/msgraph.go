// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// MetadataAction undocumented
type MetadataAction struct {
	// InformationProtectionAction is the base model of MetadataAction
	InformationProtectionAction
	// MetadataToRemove undocumented
	MetadataToRemove []string `json:"metadataToRemove,omitempty"`
	// MetadataToAdd undocumented
	MetadataToAdd []KeyValuePair `json:"metadataToAdd,omitempty"`
}

// MetadataEntry undocumented
type MetadataEntry struct {
	// Object is the base model of MetadataEntry
	Object
	// Key undocumented
	Key *string `json:"key,omitempty"`
	// Value undocumented
	Value *string `json:"value,omitempty"`
}
