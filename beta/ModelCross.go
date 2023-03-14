
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// CrossCloudAzureActiveDirectoryTenant undocumented
type CrossCloudAzureActiveDirectoryTenant struct {
    // IdentitySource is the base model of CrossCloudAzureActiveDirectoryTenant
    IdentitySource
    // CloudInstance undocumented
    CloudInstance *string `json:"cloudInstance,omitempty"`
    // DisplayName undocumented
    DisplayName *string `json:"displayName,omitempty"`
    // TenantID undocumented
    TenantID *string `json:"tenantId,omitempty"`
}

// CrossTenantAccessPolicy undocumented
type CrossTenantAccessPolicy struct {
    // TenantRelationshipAccessPolicyBase is the base model of CrossTenantAccessPolicy
    TenantRelationshipAccessPolicyBase
    // AllowedCloudEndpoints undocumented
    AllowedCloudEndpoints []string `json:"allowedCloudEndpoints,omitempty"`
    // Default undocumented
    Default *CrossTenantAccessPolicyConfigurationDefault `json:"default,omitempty"`
    // Partners undocumented
    Partners []CrossTenantAccessPolicyConfigurationPartner `json:"partners,omitempty"`
}

// CrossTenantAccessPolicyB2BSetting undocumented
type CrossTenantAccessPolicyB2BSetting struct {
    // Object is the base model of CrossTenantAccessPolicyB2BSetting
    Object
    // Applications undocumented
    Applications *CrossTenantAccessPolicyTargetConfiguration `json:"applications,omitempty"`
    // UsersAndGroups undocumented
    UsersAndGroups *CrossTenantAccessPolicyTargetConfiguration `json:"usersAndGroups,omitempty"`
}

// CrossTenantAccessPolicyConfigurationDefault undocumented
type CrossTenantAccessPolicyConfigurationDefault struct {
    // Entity is the base model of CrossTenantAccessPolicyConfigurationDefault
    Entity
    // AutomaticUserConsentSettings undocumented
    AutomaticUserConsentSettings *InboundOutboundPolicyConfiguration `json:"automaticUserConsentSettings,omitempty"`
    // B2bCollaborationInbound undocumented
    B2bCollaborationInbound *CrossTenantAccessPolicyB2BSetting `json:"b2bCollaborationInbound,omitempty"`
    // B2bCollaborationOutbound undocumented
    B2bCollaborationOutbound *CrossTenantAccessPolicyB2BSetting `json:"b2bCollaborationOutbound,omitempty"`
    // B2bDirectConnectInbound undocumented
    B2bDirectConnectInbound *CrossTenantAccessPolicyB2BSetting `json:"b2bDirectConnectInbound,omitempty"`
    // B2bDirectConnectOutbound undocumented
    B2bDirectConnectOutbound *CrossTenantAccessPolicyB2BSetting `json:"b2bDirectConnectOutbound,omitempty"`
    // InboundTrust undocumented
    InboundTrust *CrossTenantAccessPolicyInboundTrust `json:"inboundTrust,omitempty"`
    // IsServiceDefault undocumented
    IsServiceDefault *bool `json:"isServiceDefault,omitempty"`
    // TenantRestrictions undocumented
    TenantRestrictions *CrossTenantAccessPolicyTenantRestrictions `json:"tenantRestrictions,omitempty"`
}

// CrossTenantAccessPolicyConfigurationPartner undocumented
type CrossTenantAccessPolicyConfigurationPartner struct {
    // Object is the base model of CrossTenantAccessPolicyConfigurationPartner
    Object
    // AutomaticUserConsentSettings undocumented
    AutomaticUserConsentSettings *InboundOutboundPolicyConfiguration `json:"automaticUserConsentSettings,omitempty"`
    // B2bCollaborationInbound undocumented
    B2bCollaborationInbound *CrossTenantAccessPolicyB2BSetting `json:"b2bCollaborationInbound,omitempty"`
    // B2bCollaborationOutbound undocumented
    B2bCollaborationOutbound *CrossTenantAccessPolicyB2BSetting `json:"b2bCollaborationOutbound,omitempty"`
    // B2bDirectConnectInbound undocumented
    B2bDirectConnectInbound *CrossTenantAccessPolicyB2BSetting `json:"b2bDirectConnectInbound,omitempty"`
    // B2bDirectConnectOutbound undocumented
    B2bDirectConnectOutbound *CrossTenantAccessPolicyB2BSetting `json:"b2bDirectConnectOutbound,omitempty"`
    // InboundTrust undocumented
    InboundTrust *CrossTenantAccessPolicyInboundTrust `json:"inboundTrust,omitempty"`
    // IsServiceProvider undocumented
    IsServiceProvider *bool `json:"isServiceProvider,omitempty"`
    // TenantID undocumented
    TenantID *string `json:"tenantId,omitempty"`
    // TenantRestrictions undocumented
    TenantRestrictions *CrossTenantAccessPolicyTenantRestrictions `json:"tenantRestrictions,omitempty"`
    // IdentitySynchronization undocumented
    IdentitySynchronization *CrossTenantIdentitySyncPolicyPartner `json:"identitySynchronization,omitempty"`
}

// CrossTenantAccessPolicyInboundTrust undocumented
type CrossTenantAccessPolicyInboundTrust struct {
    // Object is the base model of CrossTenantAccessPolicyInboundTrust
    Object
    // IsCompliantDeviceAccepted undocumented
    IsCompliantDeviceAccepted *bool `json:"isCompliantDeviceAccepted,omitempty"`
    // IsHybridAzureADJoinedDeviceAccepted undocumented
    IsHybridAzureADJoinedDeviceAccepted *bool `json:"isHybridAzureADJoinedDeviceAccepted,omitempty"`
    // IsMFAAccepted undocumented
    IsMFAAccepted *bool `json:"isMfaAccepted,omitempty"`
}

// CrossTenantAccessPolicyTarget undocumented
type CrossTenantAccessPolicyTarget struct {
    // Object is the base model of CrossTenantAccessPolicyTarget
    Object
    // Target undocumented
    Target *string `json:"target,omitempty"`
    // TargetType undocumented
    TargetType *CrossTenantAccessPolicyTargetType `json:"targetType,omitempty"`
}

// CrossTenantAccessPolicyTargetConfiguration undocumented
type CrossTenantAccessPolicyTargetConfiguration struct {
    // Object is the base model of CrossTenantAccessPolicyTargetConfiguration
    Object
    // AccessType undocumented
    AccessType *CrossTenantAccessPolicyTargetConfigurationAccessType `json:"accessType,omitempty"`
    // Targets undocumented
    Targets []CrossTenantAccessPolicyTarget `json:"targets,omitempty"`
}

// CrossTenantAccessPolicyTenantRestrictions undocumented
type CrossTenantAccessPolicyTenantRestrictions struct {
    // CrossTenantAccessPolicyB2BSetting is the base model of CrossTenantAccessPolicyTenantRestrictions
    CrossTenantAccessPolicyB2BSetting
    // Devices undocumented
    Devices *DevicesFilter `json:"devices,omitempty"`
}

// CrossTenantIdentitySyncPolicyPartner undocumented
type CrossTenantIdentitySyncPolicyPartner struct {
    // Object is the base model of CrossTenantIdentitySyncPolicyPartner
    Object
    // DisplayName undocumented
    DisplayName *string `json:"displayName,omitempty"`
    // TenantID undocumented
    TenantID *string `json:"tenantId,omitempty"`
    // UserSyncInbound undocumented
    UserSyncInbound *CrossTenantUserSyncInbound `json:"userSyncInbound,omitempty"`
}

// CrossTenantUserSyncInbound undocumented
type CrossTenantUserSyncInbound struct {
    // Object is the base model of CrossTenantUserSyncInbound
    Object
    // IsSyncAllowed undocumented
    IsSyncAllowed *bool `json:"isSyncAllowed,omitempty"`
}
