// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// BlockAccessAction undocumented
type BlockAccessAction struct {
	// DlpActionInfo is the base model of BlockAccessAction
	DlpActionInfo

	ODataType string `json:"@odata.type,omitempty"`
}

func NewBlockAccessAction() (*BlockAccessAction, error) {
	newBlockAccessAction := &BlockAccessAction{
		ODataType: "#microsoft.graph.BlockAccessAction",
	}
	return newBlockAccessAction, nil
}
