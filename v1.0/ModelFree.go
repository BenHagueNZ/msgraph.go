// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// FreeBusyError undocumented
type FreeBusyError struct {
	// Object is the base model of FreeBusyError
	Object

	ODataType string `json:"@odata.type"`
	// Message undocumented
	Message *string `json:"message,omitempty"`
	// ResponseCode undocumented
	ResponseCode *string `json:"responseCode,omitempty"`
}

func NewFreeBusyError() (*FreeBusyError, error) {
	newFreeBusyError := &FreeBusyError{
		ODataType: "#microsoft.graph.FreeBusyError",
	}
	return newFreeBusyError, nil
}
