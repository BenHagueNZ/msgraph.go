// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// DeleteAction undocumented
type DeleteAction struct {
	// Object is the base model of DeleteAction
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// ObjectType undocumented
	ObjectType *string `json:"objectType,omitempty"`
}

func NewDeleteAction() (*DeleteAction, error) {
	newDeleteAction := &DeleteAction{
		ODataType: "#microsoft.graph.DeleteAction",
	}
	return newDeleteAction, nil
}

// DeleteUserFromSharedAppleDeviceActionResult undocumented
type DeleteUserFromSharedAppleDeviceActionResult struct {
	// DeviceActionResult is the base model of DeleteUserFromSharedAppleDeviceActionResult
	DeviceActionResult

	ODataType string `json:"@odata.type,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}

func NewDeleteUserFromSharedAppleDeviceActionResult() (*DeleteUserFromSharedAppleDeviceActionResult, error) {
	newDeleteUserFromSharedAppleDeviceActionResult := &DeleteUserFromSharedAppleDeviceActionResult{
		ODataType: "#microsoft.graph.DeleteUserFromSharedAppleDeviceActionResult",
	}
	return newDeleteUserFromSharedAppleDeviceActionResult, nil
}
