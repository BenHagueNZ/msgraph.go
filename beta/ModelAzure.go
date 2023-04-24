// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// AzureADRegistrationPolicy undocumented
type AzureADRegistrationPolicy struct {
	// Object is the base model of AzureADRegistrationPolicy
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AllowedGroups undocumented
	AllowedGroups []string `json:"allowedGroups,omitempty"`
	// AllowedUsers undocumented
	AllowedUsers []string `json:"allowedUsers,omitempty"`
	// AppliesTo undocumented
	AppliesTo *PolicyScope `json:"appliesTo,omitempty"`
	// IsAdminConfigurable undocumented
	IsAdminConfigurable *bool `json:"isAdminConfigurable,omitempty"`
}

func NewAzureADRegistrationPolicy() (*AzureADRegistrationPolicy, error) {
	newAzureADRegistrationPolicy := &AzureADRegistrationPolicy{
		ODataType: "#microsoft.graph.AzureADRegistrationPolicy",
	}
	return newAzureADRegistrationPolicy, nil
}

// AzureADWindowsAutopilotDeploymentProfile undocumented
type AzureADWindowsAutopilotDeploymentProfile struct {
	// WindowsAutopilotDeploymentProfile is the base model of AzureADWindowsAutopilotDeploymentProfile
	WindowsAutopilotDeploymentProfile

	ODataType string `json:"@odata.type,omitempty"`
}

func NewAzureADWindowsAutopilotDeploymentProfile() (*AzureADWindowsAutopilotDeploymentProfile, error) {
	newAzureADWindowsAutopilotDeploymentProfile := &AzureADWindowsAutopilotDeploymentProfile{
		ODataType: "#microsoft.graph.AzureADWindowsAutopilotDeploymentProfile",
	}
	return newAzureADWindowsAutopilotDeploymentProfile, nil
}

// AzureActiveDirectoryTenant undocumented
type AzureActiveDirectoryTenant struct {
	// IdentitySource is the base model of AzureActiveDirectoryTenant
	IdentitySource

	ODataType string `json:"@odata.type,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// TenantID undocumented
	TenantID *string `json:"tenantId,omitempty"`
}

func NewAzureActiveDirectoryTenant() (*AzureActiveDirectoryTenant, error) {
	newAzureActiveDirectoryTenant := &AzureActiveDirectoryTenant{
		ODataType: "#microsoft.graph.AzureActiveDirectoryTenant",
	}
	return newAzureActiveDirectoryTenant, nil
}

// AzureAdJoinPolicy undocumented
type AzureAdJoinPolicy struct {
	// Object is the base model of AzureAdJoinPolicy
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AllowedGroups undocumented
	AllowedGroups []string `json:"allowedGroups,omitempty"`
	// AllowedUsers undocumented
	AllowedUsers []string `json:"allowedUsers,omitempty"`
	// AppliesTo undocumented
	AppliesTo *PolicyScope `json:"appliesTo,omitempty"`
	// IsAdminConfigurable undocumented
	IsAdminConfigurable *bool `json:"isAdminConfigurable,omitempty"`
}

func NewAzureAdJoinPolicy() (*AzureAdJoinPolicy, error) {
	newAzureAdJoinPolicy := &AzureAdJoinPolicy{
		ODataType: "#microsoft.graph.AzureAdJoinPolicy",
	}
	return newAzureAdJoinPolicy, nil
}

// AzureAdPopTokenAuthentication undocumented
type AzureAdPopTokenAuthentication struct {
	// CustomExtensionAuthenticationConfiguration is the base model of AzureAdPopTokenAuthentication
	CustomExtensionAuthenticationConfiguration

	ODataType string `json:"@odata.type,omitempty"`
}

func NewAzureAdPopTokenAuthentication() (*AzureAdPopTokenAuthentication, error) {
	newAzureAdPopTokenAuthentication := &AzureAdPopTokenAuthentication{
		ODataType: "#microsoft.graph.AzureAdPopTokenAuthentication",
	}
	return newAzureAdPopTokenAuthentication, nil
}

// AzureAdTokenAuthentication undocumented
type AzureAdTokenAuthentication struct {
	// CustomExtensionAuthenticationConfiguration is the base model of AzureAdTokenAuthentication
	CustomExtensionAuthenticationConfiguration

	ODataType string `json:"@odata.type,omitempty"`
	// ResourceID undocumented
	ResourceID *string `json:"resourceId,omitempty"`
}

func NewAzureAdTokenAuthentication() (*AzureAdTokenAuthentication, error) {
	newAzureAdTokenAuthentication := &AzureAdTokenAuthentication{
		ODataType: "#microsoft.graph.AzureAdTokenAuthentication",
	}
	return newAzureAdTokenAuthentication, nil
}

// AzureCommunicationServicesUserConversationMember undocumented
type AzureCommunicationServicesUserConversationMember struct {
	// ConversationMember is the base model of AzureCommunicationServicesUserConversationMember
	ConversationMember

	ODataType string `json:"@odata.type,omitempty"`
	// AzureCommunicationServicesID undocumented
	AzureCommunicationServicesID *string `json:"azureCommunicationServicesId,omitempty"`
}

func NewAzureCommunicationServicesUserConversationMember() (*AzureCommunicationServicesUserConversationMember, error) {
	newAzureCommunicationServicesUserConversationMember := &AzureCommunicationServicesUserConversationMember{
		ODataType: "#microsoft.graph.AzureCommunicationServicesUserConversationMember",
	}
	return newAzureCommunicationServicesUserConversationMember, nil
}

// AzureCommunicationServicesUserIdentity undocumented
type AzureCommunicationServicesUserIdentity struct {
	// Identity is the base model of AzureCommunicationServicesUserIdentity
	Identity

	ODataType string `json:"@odata.type,omitempty"`
	// AzureCommunicationServicesResourceID undocumented
	AzureCommunicationServicesResourceID *string `json:"azureCommunicationServicesResourceId,omitempty"`
}

func NewAzureCommunicationServicesUserIdentity() (*AzureCommunicationServicesUserIdentity, error) {
	newAzureCommunicationServicesUserIdentity := &AzureCommunicationServicesUserIdentity{
		ODataType: "#microsoft.graph.AzureCommunicationServicesUserIdentity",
	}
	return newAzureCommunicationServicesUserIdentity, nil
}