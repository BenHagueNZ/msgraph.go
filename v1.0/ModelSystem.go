// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// SystemFacet undocumented
type SystemFacet struct {
	// Object is the base model of SystemFacet
	Object

	ODataType string `json:"@odata.type,omitempty"`
}

func NewSystemFacet() (*SystemFacet, error) {
	newSystemFacet := &SystemFacet{
		ODataType: "#microsoft.graph.SystemFacet",
	}
	return newSystemFacet, nil
}
