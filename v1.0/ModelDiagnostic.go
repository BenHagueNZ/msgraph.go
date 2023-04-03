// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Diagnostic undocumented
type Diagnostic struct {
	// Object is the base model of Diagnostic
	Object

	ODataType string `json:"@odata.type"`
	// Message undocumented
	Message *string `json:"message,omitempty"`
	// URL undocumented
	URL *string `json:"url,omitempty"`
}

func NewDiagnostic() (*Diagnostic, error) {
	newDiagnostic := &Diagnostic{
		ODataType: "#microsoft.graph.Diagnostic",
	}
	return newDiagnostic, nil
}
