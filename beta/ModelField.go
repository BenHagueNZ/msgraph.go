// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// FieldValueSet undocumented
type FieldValueSet struct {
	// Entity is the base model of FieldValueSet
	Entity

	ODataType string `json:"@odata.type,omitempty"`
}

func NewFieldValueSet() (*FieldValueSet, error) {
	newFieldValueSet := &FieldValueSet{
		ODataType: "#microsoft.graph.FieldValueSet",
	}
	return newFieldValueSet, nil
}