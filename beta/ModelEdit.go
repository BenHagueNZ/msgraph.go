// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// EditAction undocumented
type EditAction struct {
	// Object is the base model of EditAction
	Object

	ODataType string `json:"@odata.type,omitempty"`
}

func NewEditAction() (*EditAction, error) {
	newEditAction := &EditAction{
		ODataType: "#microsoft.graph.EditAction",
	}
	return newEditAction, nil
}
