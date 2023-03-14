
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// AuditActivityInitiator undocumented
type AuditActivityInitiator struct {
    // Object is the base model of AuditActivityInitiator
    Object
    // App undocumented
    App *AppIdentity `json:"app,omitempty"`
    // User undocumented
    User *AuditUserIdentity `json:"user,omitempty"`
}

// AuditActor undocumented
type AuditActor struct {
    // Object is the base model of AuditActor
    Object
    // ApplicationDisplayName undocumented
    ApplicationDisplayName *string `json:"applicationDisplayName,omitempty"`
    // ApplicationID undocumented
    ApplicationID *string `json:"applicationId,omitempty"`
    // AuditActorType undocumented
    AuditActorType *string `json:"auditActorType,omitempty"`
    // IPAddress undocumented
    IPAddress *string `json:"ipAddress,omitempty"`
    // RemoteTenantID undocumented
    RemoteTenantID *string `json:"remoteTenantId,omitempty"`
    // RemoteUserID undocumented
    RemoteUserID *string `json:"remoteUserId,omitempty"`
    // ServicePrincipalName undocumented
    ServicePrincipalName *string `json:"servicePrincipalName,omitempty"`
    // Type undocumented
    Type *string `json:"type,omitempty"`
    // UserID undocumented
    UserID *string `json:"userId,omitempty"`
    // UserPermissions undocumented
    UserPermissions []string `json:"userPermissions,omitempty"`
    // UserPrincipalName undocumented
    UserPrincipalName *string `json:"userPrincipalName,omitempty"`
    // UserRoleScopeTags undocumented
    UserRoleScopeTags []RoleScopeTagInfo `json:"userRoleScopeTags,omitempty"`
}

// AuditEvent undocumented
type AuditEvent struct {
    // Entity is the base model of AuditEvent
    Entity
    // Activity undocumented
    Activity *string `json:"activity,omitempty"`
    // ActivityDateTime undocumented
    ActivityDateTime *time.Time `json:"activityDateTime,omitempty"`
    // ActivityOperationType undocumented
    ActivityOperationType *string `json:"activityOperationType,omitempty"`
    // ActivityResult undocumented
    ActivityResult *string `json:"activityResult,omitempty"`
    // ActivityType undocumented
    ActivityType *string `json:"activityType,omitempty"`
    // Actor undocumented
    Actor *AuditActor `json:"actor,omitempty"`
    // Category undocumented
    Category *string `json:"category,omitempty"`
    // ComponentName undocumented
    ComponentName *string `json:"componentName,omitempty"`
    // CorrelationID undocumented
    CorrelationID *UUID `json:"correlationId,omitempty"`
    // DisplayName undocumented
    DisplayName *string `json:"displayName,omitempty"`
    // Resources undocumented
    Resources []AuditResource `json:"resources,omitempty"`
}

// AuditLogRoot undocumented
type AuditLogRoot struct {
    // Object is the base model of AuditLogRoot
    Object
    // DirectoryAudits undocumented
    DirectoryAudits []DirectoryAudit `json:"directoryAudits,omitempty"`
    // DirectoryProvisioning undocumented
    DirectoryProvisioning []ProvisioningObjectSummary `json:"directoryProvisioning,omitempty"`
    // Provisioning undocumented
    Provisioning []ProvisioningObjectSummary `json:"provisioning,omitempty"`
    // SignIns undocumented
    SignIns []SignIn `json:"signIns,omitempty"`
}

// AuditProperty undocumented
type AuditProperty struct {
    // Object is the base model of AuditProperty
    Object
    // DisplayName undocumented
    DisplayName *string `json:"displayName,omitempty"`
    // NewValue undocumented
    NewValue *string `json:"newValue,omitempty"`
    // OldValue undocumented
    OldValue *string `json:"oldValue,omitempty"`
}

// AuditResource undocumented
type AuditResource struct {
    // Object is the base model of AuditResource
    Object
    // AuditResourceType undocumented
    AuditResourceType *string `json:"auditResourceType,omitempty"`
    // DisplayName undocumented
    DisplayName *string `json:"displayName,omitempty"`
    // ModifiedProperties undocumented
    ModifiedProperties []AuditProperty `json:"modifiedProperties,omitempty"`
    // ResourceID undocumented
    ResourceID *string `json:"resourceId,omitempty"`
    // Type undocumented
    Type *string `json:"type,omitempty"`
}

// AuditUserIdentity undocumented
type AuditUserIdentity struct {
    // UserIdentity is the base model of AuditUserIdentity
    UserIdentity
    // HomeTenantID undocumented
    HomeTenantID *string `json:"homeTenantId,omitempty"`
    // HomeTenantName undocumented
    HomeTenantName *string `json:"homeTenantName,omitempty"`
}
