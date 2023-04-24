// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// InstanceResourceAccess undocumented
type InstanceResourceAccess struct {
	// Object is the base model of InstanceResourceAccess
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Permissions undocumented
	Permissions []ResourcePermission `json:"permissions,omitempty"`
	// ResourceAppID undocumented
	ResourceAppID *string `json:"resourceAppId,omitempty"`
}

func NewInstanceResourceAccess() (*InstanceResourceAccess, error) {
	newInstanceResourceAccess := &InstanceResourceAccess{
		ODataType: "#microsoft.graph.InstanceResourceAccess",
	}
	return newInstanceResourceAccess, nil
}