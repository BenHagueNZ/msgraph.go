// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Album undocumented
type Album struct {
	// Object is the base model of Album
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// CoverImageItemID undocumented
	CoverImageItemID *string `json:"coverImageItemId,omitempty"`
}

func NewAlbum() (*Album, error) {
	newAlbum := &Album{
		ODataType: "#microsoft.graph.Album",
	}
	return newAlbum, nil
}
