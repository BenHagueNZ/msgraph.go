// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// BooleanColumn undocumented
type BooleanColumn struct {
	// Object is the base model of BooleanColumn
	Object

	ODataType string `json:"@odata.type"`
}

func NewBooleanColumn() (*BooleanColumn, error) {
	newBooleanColumn := &BooleanColumn{
		ODataType: "#microsoft.graph.BooleanColumn",
	}
	return newBooleanColumn, nil
}
