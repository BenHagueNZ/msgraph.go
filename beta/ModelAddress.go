// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// AddressBookAccountTargetContent undocumented
type AddressBookAccountTargetContent struct {
	// AccountTargetContent is the base model of AddressBookAccountTargetContent
	AccountTargetContent

	ODataType string `json:"@odata.type,omitempty"`
	// AccountTargetEmails undocumented
	AccountTargetEmails []string `json:"accountTargetEmails,omitempty"`
}

func NewAddressBookAccountTargetContent() (*AddressBookAccountTargetContent, error) {
	newAddressBookAccountTargetContent := &AddressBookAccountTargetContent{
		ODataType: "#microsoft.graph.AddressBookAccountTargetContent",
	}
	return newAddressBookAccountTargetContent, nil
}