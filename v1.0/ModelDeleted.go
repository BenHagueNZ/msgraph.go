// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Deleted undocumented
type Deleted struct {
	// Object is the base model of Deleted
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// State undocumented
	State *string `json:"state,omitempty"`
}

func NewDeleted() (*Deleted, error) {
	newDeleted := &Deleted{
		ODataType: "#microsoft.graph.Deleted",
	}
	return newDeleted, nil
}

// DeletedTeam undocumented
type DeletedTeam struct {
	// Entity is the base model of DeletedTeam
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Channels undocumented
	Channels []Channel `json:"channels,omitempty"`
}

func NewDeletedTeam() (*DeletedTeam, error) {
	newDeletedTeam := &DeletedTeam{
		ODataType: "#microsoft.graph.DeletedTeam",
	}
	return newDeletedTeam, nil
}
