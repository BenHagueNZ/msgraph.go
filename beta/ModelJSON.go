// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// JSON undocumented
type JSON struct {
	// Object is the base model of JSON
	Object

	ODataType string `json:"@odata.type,omitempty"`
}

func NewJSON() (*JSON, error) {
	newJSON := &JSON{
		ODataType: "#microsoft.graph.Json",
	}
	return newJSON, nil
}
