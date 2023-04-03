// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// AuthoredNote undocumented
type AuthoredNote struct {
	// Entity is the base model of AuthoredNote
	Entity

	ODataType string `json:"@odata.type"`
	// Author undocumented
	Author *Identity `json:"author,omitempty"`
	// Content undocumented
	Content *ItemBody `json:"content,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
}

func NewAuthoredNote() (*AuthoredNote, error) {
	newAuthoredNote := &AuthoredNote{
		ODataType: "#microsoft.graph.AuthoredNote",
	}
	return newAuthoredNote, nil
}
