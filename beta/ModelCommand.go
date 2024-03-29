// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Command undocumented
type Command struct {
	// Entity is the base model of Command
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AppServiceName undocumented
	AppServiceName *string `json:"appServiceName,omitempty"`
	// Error undocumented
	Error *string `json:"error,omitempty"`
	// PackageFamilyName undocumented
	PackageFamilyName *string `json:"packageFamilyName,omitempty"`
	// Payload undocumented
	Payload *PayloadRequest `json:"payload,omitempty"`
	// PermissionTicket undocumented
	PermissionTicket *string `json:"permissionTicket,omitempty"`
	// PostBackURI undocumented
	PostBackURI *string `json:"postBackUri,omitempty"`
	// Status undocumented
	Status *string `json:"status,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
	// Responsepayload undocumented
	Responsepayload *PayloadResponse `json:"responsepayload,omitempty"`
}

func NewCommand() (*Command, error) {
	newCommand := &Command{
		ODataType: "#microsoft.graph.Command",
	}
	return newCommand, nil
}
