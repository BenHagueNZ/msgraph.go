
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// DlpActionInfo undocumented
type DlpActionInfo struct {
    // Object is the base model of DlpActionInfo
    Object
    // Action undocumented
    Action *DlpAction `json:"action,omitempty"`
}

// DlpEvaluatePoliciesJobResponse undocumented
type DlpEvaluatePoliciesJobResponse struct {
    // JobResponseBase is the base model of DlpEvaluatePoliciesJobResponse
    JobResponseBase
    // Result undocumented
    Result *DlpPoliciesJobResult `json:"result,omitempty"`
}

// DlpEvaluatePoliciesRequestObject undocumented
type DlpEvaluatePoliciesRequestObject struct {
    // Object is the base model of DlpEvaluatePoliciesRequestObject
    Object
    // EvaluationInput undocumented
    EvaluationInput *DlpEvaluationInput `json:"evaluationInput,omitempty"`
    // NotificationInfo undocumented
    NotificationInfo *DlpNotification `json:"notificationInfo,omitempty"`
    // Target undocumented
    Target *string `json:"target,omitempty"`
}

// DlpEvaluationInput undocumented
type DlpEvaluationInput struct {
    // Object is the base model of DlpEvaluationInput
    Object
    // CurrentLabel undocumented
    CurrentLabel *CurrentLabel `json:"currentLabel,omitempty"`
    // DiscoveredSensitiveTypes undocumented
    DiscoveredSensitiveTypes []DiscoveredSensitiveType `json:"discoveredSensitiveTypes,omitempty"`
}

// DlpEvaluationWindowsDevicesInput undocumented
type DlpEvaluationWindowsDevicesInput struct {
    // DlpEvaluationInput is the base model of DlpEvaluationWindowsDevicesInput
    DlpEvaluationInput
    // ContentProperties undocumented
    ContentProperties *ContentProperties `json:"contentProperties,omitempty"`
    // SharedBy undocumented
    SharedBy *string `json:"sharedBy,omitempty"`
}

// DlpNotification undocumented
type DlpNotification struct {
    // Object is the base model of DlpNotification
    Object
    // Author undocumented
    Author *string `json:"author,omitempty"`
}

// DlpPoliciesJobResult undocumented
type DlpPoliciesJobResult struct {
    // Object is the base model of DlpPoliciesJobResult
    Object
    // AuditCorrelationID undocumented
    AuditCorrelationID *string `json:"auditCorrelationId,omitempty"`
    // EvaluationDateTime undocumented
    EvaluationDateTime *time.Time `json:"evaluationDateTime,omitempty"`
    // MatchingRules undocumented
    MatchingRules []MatchingDlpRule `json:"matchingRules,omitempty"`
}

// DlpWindowsDevicesNotification undocumented
type DlpWindowsDevicesNotification struct {
    // DlpNotification is the base model of DlpWindowsDevicesNotification
    DlpNotification
    // ContentName undocumented
    ContentName *string `json:"contentName,omitempty"`
    // LastModfiedBy undocumented
    LastModfiedBy *string `json:"lastModfiedBy,omitempty"`
}
