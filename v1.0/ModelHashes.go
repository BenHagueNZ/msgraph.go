// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Hashes undocumented
type Hashes struct {
	// Object is the base model of Hashes
	Object

	ODataType string `json:"@odata.type"`
	// Crc32Hash undocumented
	Crc32Hash *string `json:"crc32Hash,omitempty"`
	// QuickXorHash undocumented
	QuickXorHash *string `json:"quickXorHash,omitempty"`
	// Sha1Hash undocumented
	Sha1Hash *string `json:"sha1Hash,omitempty"`
	// Sha256Hash undocumented
	Sha256Hash *string `json:"sha256Hash,omitempty"`
}

func NewHashes() (*Hashes, error) {
	newHashes := &Hashes{
		ODataType: "#microsoft.graph.Hashes",
	}
	return newHashes, nil
}
