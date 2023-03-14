
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// ConfigManagerCollection undocumented
type ConfigManagerCollection struct {
    // Entity is the base model of ConfigManagerCollection
    Entity
    // CollectionIdentifier undocumented
    CollectionIdentifier *string `json:"collectionIdentifier,omitempty"`
    // CreatedDateTime undocumented
    CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
    // DisplayName undocumented
    DisplayName *string `json:"displayName,omitempty"`
    // HierarchyIdentifier undocumented
    HierarchyIdentifier *string `json:"hierarchyIdentifier,omitempty"`
    // HierarchyName undocumented
    HierarchyName *string `json:"hierarchyName,omitempty"`
    // LastModifiedDateTime undocumented
    LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}

// ConfigManagerPolicySummary undocumented
type ConfigManagerPolicySummary struct {
    // Object is the base model of ConfigManagerPolicySummary
    Object
    // CompliantDeviceCount undocumented
    CompliantDeviceCount *int `json:"compliantDeviceCount,omitempty"`
    // EnforcedDeviceCount undocumented
    EnforcedDeviceCount *int `json:"enforcedDeviceCount,omitempty"`
    // FailedDeviceCount undocumented
    FailedDeviceCount *int `json:"failedDeviceCount,omitempty"`
    // NonCompliantDeviceCount undocumented
    NonCompliantDeviceCount *int `json:"nonCompliantDeviceCount,omitempty"`
    // PendingDeviceCount undocumented
    PendingDeviceCount *int `json:"pendingDeviceCount,omitempty"`
    // TargetedDeviceCount undocumented
    TargetedDeviceCount *int `json:"targetedDeviceCount,omitempty"`
}