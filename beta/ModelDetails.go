// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// DetailsInfo undocumented
type DetailsInfo struct {
	// Object is the base model of DetailsInfo
	Object

	ODataType string `json:"@odata.type,omitempty"`
}

func NewDetailsInfo() (*DetailsInfo, error) {
	newDetailsInfo := &DetailsInfo{
		ODataType: "#microsoft.graph.DetailsInfo",
	}
	return newDetailsInfo, nil
}