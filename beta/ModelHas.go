
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// HasPayloadLinkResultItem undocumented
type HasPayloadLinkResultItem struct {
    // Object is the base model of HasPayloadLinkResultItem
    Object
    // Error undocumented
    Error *string `json:"error,omitempty"`
    // HasLink undocumented
    HasLink *bool `json:"hasLink,omitempty"`
    // PayloadID undocumented
    PayloadID *string `json:"payloadId,omitempty"`
    // Sources undocumented
    Sources []DeviceAndAppManagementAssignmentSource `json:"sources,omitempty"`
}
