// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ConnectionInfo undocumented
type ConnectionInfo struct {
	// Object is the base model of ConnectionInfo
	Object

	ODataType string `json:"@odata.type"`
	// URL undocumented
	URL *string `json:"url,omitempty"`
}

func NewConnectionInfo() (*ConnectionInfo, error) {
	newConnectionInfo := &ConnectionInfo{
		ODataType: "#microsoft.graph.ConnectionInfo",
	}
	return newConnectionInfo, nil
}
