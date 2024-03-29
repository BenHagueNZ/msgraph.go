// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Package undocumented
type Package struct {
	// Object is the base model of Package
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
}

func NewPackage() (*Package, error) {
	newPackage := &Package{
		ODataType: "#microsoft.graph.Package",
	}
	return newPackage, nil
}
