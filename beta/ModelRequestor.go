// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// RequestorManager undocumented
type RequestorManager struct {
	// UserSet is the base model of RequestorManager
	UserSet

	ODataType string `json:"@odata.type,omitempty"`
	// ManagerLevel undocumented
	ManagerLevel *int `json:"managerLevel,omitempty"`
}

func NewRequestorManager() (*RequestorManager, error) {
	newRequestorManager := &RequestorManager{
		ODataType: "#microsoft.graph.RequestorManager",
	}
	return newRequestorManager, nil
}

// RequestorSettings undocumented
type RequestorSettings struct {
	// Object is the base model of RequestorSettings
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AcceptRequests undocumented
	AcceptRequests *bool `json:"acceptRequests,omitempty"`
	// AllowedRequestors undocumented
	AllowedRequestors []UserSet `json:"allowedRequestors,omitempty"`
	// ScopeType undocumented
	ScopeType *string `json:"scopeType,omitempty"`
}

func NewRequestorSettings() (*RequestorSettings, error) {
	newRequestorSettings := &RequestorSettings{
		ODataType: "#microsoft.graph.RequestorSettings",
	}
	return newRequestorSettings, nil
}
