// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// ResponseStatus undocumented
type ResponseStatus struct {
	// Object is the base model of ResponseStatus
	Object

	ODataType string `json:"@odata.type"`
	// Response undocumented
	Response *ResponseType `json:"response,omitempty"`
	// Time undocumented
	Time *time.Time `json:"time,omitempty"`
}

func NewResponseStatus() (*ResponseStatus, error) {
	newResponseStatus := &ResponseStatus{
		ODataType: "#microsoft.graph.ResponseStatus",
	}
	return newResponseStatus, nil
}
