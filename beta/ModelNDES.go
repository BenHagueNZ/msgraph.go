
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// NDESConnector undocumented
type NDESConnector struct {
    // Entity is the base model of NDESConnector
    Entity
    // ConnectorVersion undocumented
    ConnectorVersion *string `json:"connectorVersion,omitempty"`
    // DisplayName undocumented
    DisplayName *string `json:"displayName,omitempty"`
    // EnrolledDateTime undocumented
    EnrolledDateTime *time.Time `json:"enrolledDateTime,omitempty"`
    // LastConnectionDateTime undocumented
    LastConnectionDateTime *time.Time `json:"lastConnectionDateTime,omitempty"`
    // MachineName undocumented
    MachineName *string `json:"machineName,omitempty"`
    // RoleScopeTagIDs undocumented
    RoleScopeTagIDs []string `json:"roleScopeTagIds,omitempty"`
    // State undocumented
    State *NDESConnectorState `json:"state,omitempty"`
}
