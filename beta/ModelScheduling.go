// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// SchedulingGroup undocumented
type SchedulingGroup struct {
	// ChangeTrackedEntity is the base model of SchedulingGroup
	ChangeTrackedEntity

	ODataType string `json:"@odata.type,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IsActive undocumented
	IsActive *bool `json:"isActive,omitempty"`
	// UserIDs undocumented
	UserIDs []string `json:"userIds,omitempty"`
}

func NewSchedulingGroup() (*SchedulingGroup, error) {
	newSchedulingGroup := &SchedulingGroup{
		ODataType: "#microsoft.graph.SchedulingGroup",
	}
	return newSchedulingGroup, nil
}
