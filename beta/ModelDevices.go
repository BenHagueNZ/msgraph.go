// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// DevicesFilter undocumented
type DevicesFilter struct {
	// Object is the base model of DevicesFilter
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Mode undocumented
	Mode *CrossTenantAccessPolicyTargetConfigurationAccessType `json:"mode,omitempty"`
	// Rule undocumented
	Rule *string `json:"rule,omitempty"`
}

func NewDevicesFilter() (*DevicesFilter, error) {
	newDevicesFilter := &DevicesFilter{
		ODataType: "#microsoft.graph.DevicesFilter",
	}
	return newDevicesFilter, nil
}
