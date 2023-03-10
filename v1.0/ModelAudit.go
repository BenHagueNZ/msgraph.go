// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// AuditActivityInitiator undocumented
type AuditActivityInitiator struct {
	// Object is the base model of AuditActivityInitiator
	Object
	// App undocumented
	App *AppIdentity `json:"app,omitempty"`
	// User undocumented
	User *UserIdentity `json:"user,omitempty"`
}

// AuditActor A class containing the properties for Audit Actor.
type AuditActor struct {
	// Object is the base model of AuditActor
	Object
	// ApplicationDisplayName Name of the Application.
	ApplicationDisplayName *string `json:"applicationDisplayName,omitempty"`
	// ApplicationID AAD Application Id.
	ApplicationID *string `json:"applicationId,omitempty"`
	// AuditActorType Actor Type.
	AuditActorType *string `json:"auditActorType,omitempty"`
	// IPAddress IPAddress.
	IPAddress *string `json:"ipAddress,omitempty"`
	// ServicePrincipalName Service Principal Name (SPN).
	ServicePrincipalName *string `json:"servicePrincipalName,omitempty"`
	// UserID User Id.
	UserID *string `json:"userId,omitempty"`
	// UserPermissions List of user permissions when the audit was performed.
	UserPermissions []string `json:"userPermissions,omitempty"`
	// UserPrincipalName User Principal Name (UPN).
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}

// AuditEvent A class containing the properties for Audit Event.
type AuditEvent struct {
	// Entity is the base model of AuditEvent
	Entity
	// Activity Friendly name of the activity.
	Activity *string `json:"activity,omitempty"`
	// ActivityDateTime The date time in UTC when the activity was performed.
	ActivityDateTime *time.Time `json:"activityDateTime,omitempty"`
	// ActivityOperationType The HTTP operation type of the activity.
	ActivityOperationType *string `json:"activityOperationType,omitempty"`
	// ActivityResult The result of the activity.
	ActivityResult *string `json:"activityResult,omitempty"`
	// ActivityType The type of activity that was being performed.
	ActivityType *string `json:"activityType,omitempty"`
	// Actor AAD user and application that are associated with the audit event.
	Actor *AuditActor `json:"actor,omitempty"`
	// Category Audit category.
	Category *string `json:"category,omitempty"`
	// ComponentName Component name.
	ComponentName *string `json:"componentName,omitempty"`
	// CorrelationID The client request Id that is used to correlate activity within the system.
	CorrelationID *UUID `json:"correlationId,omitempty"`
	// DisplayName Event display name.
	DisplayName *string `json:"displayName,omitempty"`
	// Resources Resources being modified.
	Resources []AuditResource `json:"resources,omitempty"`
}

// AuditLogRoot undocumented
type AuditLogRoot struct {
	// Entity is the base model of AuditLogRoot
	Entity
	// DirectoryAudits undocumented
	DirectoryAudits []DirectoryAudit `json:"directoryAudits,omitempty"`
	// Provisioning undocumented
	Provisioning []ProvisioningObjectSummary `json:"provisioning,omitempty"`
	// SignIns undocumented
	SignIns []SignIn `json:"signIns,omitempty"`
}

// AuditProperty A class containing the properties for Audit Property.
type AuditProperty struct {
	// Object is the base model of AuditProperty
	Object
	// DisplayName Display name.
	DisplayName *string `json:"displayName,omitempty"`
	// NewValue New value.
	NewValue *string `json:"newValue,omitempty"`
	// OldValue Old value.
	OldValue *string `json:"oldValue,omitempty"`
}

// AuditResource A class containing the properties for Audit Resource.
type AuditResource struct {
	// Object is the base model of AuditResource
	Object
	// AuditResourceType Audit resource's type.
	AuditResourceType *string `json:"auditResourceType,omitempty"`
	// DisplayName Display name.
	DisplayName *string `json:"displayName,omitempty"`
	// ModifiedProperties List of modified properties.
	ModifiedProperties []AuditProperty `json:"modifiedProperties,omitempty"`
	// ResourceID Audit resource's Id.
	ResourceID *string `json:"resourceId,omitempty"`
}
