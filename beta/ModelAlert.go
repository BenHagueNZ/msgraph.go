
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// Alert undocumented
type Alert struct {
    // Entity is the base model of Alert
    Entity
    // ActivityGroupName undocumented
    ActivityGroupName *string `json:"activityGroupName,omitempty"`
    // AlertDetections undocumented
    AlertDetections []AlertDetection `json:"alertDetections,omitempty"`
    // AssignedTo undocumented
    AssignedTo *string `json:"assignedTo,omitempty"`
    // AzureSubscriptionID undocumented
    AzureSubscriptionID *string `json:"azureSubscriptionId,omitempty"`
    // AzureTenantID undocumented
    AzureTenantID *string `json:"azureTenantId,omitempty"`
    // Category undocumented
    Category *string `json:"category,omitempty"`
    // ClosedDateTime undocumented
    ClosedDateTime *time.Time `json:"closedDateTime,omitempty"`
    // CloudAppStates undocumented
    CloudAppStates []CloudAppSecurityState `json:"cloudAppStates,omitempty"`
    // Comments undocumented
    Comments []string `json:"comments,omitempty"`
    // Confidence undocumented
    Confidence *int `json:"confidence,omitempty"`
    // CreatedDateTime undocumented
    CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
    // Description undocumented
    Description *string `json:"description,omitempty"`
    // DetectionIDs undocumented
    DetectionIDs []string `json:"detectionIds,omitempty"`
    // EventDateTime undocumented
    EventDateTime *time.Time `json:"eventDateTime,omitempty"`
    // Feedback undocumented
    Feedback *AlertFeedback `json:"feedback,omitempty"`
    // FileStates undocumented
    FileStates []FileSecurityState `json:"fileStates,omitempty"`
    // HistoryStates undocumented
    HistoryStates []AlertHistoryState `json:"historyStates,omitempty"`
    // HostStates undocumented
    HostStates []HostSecurityState `json:"hostStates,omitempty"`
    // IncidentIDs undocumented
    IncidentIDs []string `json:"incidentIds,omitempty"`
    // InvestigationSecurityStates undocumented
    InvestigationSecurityStates []InvestigationSecurityState `json:"investigationSecurityStates,omitempty"`
    // LastEventDateTime undocumented
    LastEventDateTime *time.Time `json:"lastEventDateTime,omitempty"`
    // LastModifiedDateTime undocumented
    LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
    // MalwareStates undocumented
    MalwareStates []MalwareState `json:"malwareStates,omitempty"`
    // MessageSecurityStates undocumented
    MessageSecurityStates []MessageSecurityState `json:"messageSecurityStates,omitempty"`
    // NetworkConnections undocumented
    NetworkConnections []NetworkConnection `json:"networkConnections,omitempty"`
    // Processes undocumented
    Processes []Process `json:"processes,omitempty"`
    // RecommendedActions undocumented
    RecommendedActions []string `json:"recommendedActions,omitempty"`
    // RegistryKeyStates undocumented
    RegistryKeyStates []RegistryKeyState `json:"registryKeyStates,omitempty"`
    // SecurityResources undocumented
    SecurityResources []SecurityResource `json:"securityResources,omitempty"`
    // Severity undocumented
    Severity *AlertSeverity `json:"severity,omitempty"`
    // SourceMaterials undocumented
    SourceMaterials []string `json:"sourceMaterials,omitempty"`
    // Status undocumented
    Status *AlertStatus `json:"status,omitempty"`
    // Tags undocumented
    Tags []string `json:"tags,omitempty"`
    // Title undocumented
    Title *string `json:"title,omitempty"`
    // Triggers undocumented
    Triggers []AlertTrigger `json:"triggers,omitempty"`
    // URIClickSecurityStates undocumented
    URIClickSecurityStates []URIClickSecurityState `json:"uriClickSecurityStates,omitempty"`
    // UserStates undocumented
    UserStates []UserSecurityState `json:"userStates,omitempty"`
    // VendorInformation undocumented
    VendorInformation *SecurityVendorInformation `json:"vendorInformation,omitempty"`
    // VulnerabilityStates undocumented
    VulnerabilityStates []VulnerabilityState `json:"vulnerabilityStates,omitempty"`
}

// AlertDetection undocumented
type AlertDetection struct {
    // Object is the base model of AlertDetection
    Object
    // DetectionType undocumented
    DetectionType *string `json:"detectionType,omitempty"`
    // Method undocumented
    Method *string `json:"method,omitempty"`
    // Name undocumented
    Name *string `json:"name,omitempty"`
}

// AlertHistoryState undocumented
type AlertHistoryState struct {
    // Object is the base model of AlertHistoryState
    Object
    // AppID undocumented
    AppID *string `json:"appId,omitempty"`
    // AssignedTo undocumented
    AssignedTo *string `json:"assignedTo,omitempty"`
    // Comments undocumented
    Comments []string `json:"comments,omitempty"`
    // Feedback undocumented
    Feedback *AlertFeedback `json:"feedback,omitempty"`
    // Status undocumented
    Status *AlertStatus `json:"status,omitempty"`
    // UpdatedDateTime undocumented
    UpdatedDateTime *time.Time `json:"updatedDateTime,omitempty"`
    // User undocumented
    User *string `json:"user,omitempty"`
}

// AlertTrigger undocumented
type AlertTrigger struct {
    // Object is the base model of AlertTrigger
    Object
    // Name undocumented
    Name *string `json:"name,omitempty"`
    // Type undocumented
    Type *string `json:"type,omitempty"`
    // Value undocumented
    Value *string `json:"value,omitempty"`
}
