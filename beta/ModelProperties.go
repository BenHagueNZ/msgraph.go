// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Properties undocumented
type Properties struct {
	// Object is the base model of Properties
	Object

	ODataType string `json:"@odata.type,omitempty"`
}

func NewProperties() (*Properties, error) {
	newProperties := &Properties{
		ODataType: "#microsoft.graph.Properties",
	}
	return newProperties, nil
}
