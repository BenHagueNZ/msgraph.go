// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// TeamsApp undocumented
type TeamsApp struct {
	// Entity is the base model of TeamsApp
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// DistributionMethod undocumented
	DistributionMethod *TeamsAppDistributionMethod `json:"distributionMethod,omitempty"`
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// AppDefinitions undocumented
	AppDefinitions []TeamsAppDefinition `json:"appDefinitions,omitempty"`
}

func NewTeamsApp() (*TeamsApp, error) {
	newTeamsApp := &TeamsApp{
		ODataType: "#microsoft.graph.TeamsApp",
	}
	return newTeamsApp, nil
}

// TeamsAppAuthorization undocumented
type TeamsAppAuthorization struct {
	// Object is the base model of TeamsAppAuthorization
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// RequiredPermissionSet undocumented
	RequiredPermissionSet *TeamsAppPermissionSet `json:"requiredPermissionSet,omitempty"`
}

func NewTeamsAppAuthorization() (*TeamsAppAuthorization, error) {
	newTeamsAppAuthorization := &TeamsAppAuthorization{
		ODataType: "#microsoft.graph.TeamsAppAuthorization",
	}
	return newTeamsAppAuthorization, nil
}

// TeamsAppDefinition undocumented
type TeamsAppDefinition struct {
	// Entity is the base model of TeamsAppDefinition
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AllowedInstallationScopes undocumented
	AllowedInstallationScopes *TeamsAppInstallationScopes `json:"allowedInstallationScopes,omitempty"`
	// Authorization undocumented
	Authorization *TeamsAppAuthorization `json:"authorization,omitempty"`
	// AzureADAppID undocumented
	AzureADAppID *string `json:"azureADAppId,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// PublishingState undocumented
	PublishingState *TeamsAppPublishingState `json:"publishingState,omitempty"`
	// Shortdescription undocumented
	Shortdescription *string `json:"shortdescription,omitempty"`
	// TeamsAppID undocumented
	TeamsAppID *string `json:"teamsAppId,omitempty"`
	// Version undocumented
	Version *string `json:"version,omitempty"`
	// Bot undocumented
	Bot *TeamworkBot `json:"bot,omitempty"`
	// ColorIcon undocumented
	ColorIcon *TeamsAppIcon `json:"colorIcon,omitempty"`
	// OutlineIcon undocumented
	OutlineIcon *TeamsAppIcon `json:"outlineIcon,omitempty"`
}

func NewTeamsAppDefinition() (*TeamsAppDefinition, error) {
	newTeamsAppDefinition := &TeamsAppDefinition{
		ODataType: "#microsoft.graph.TeamsAppDefinition",
	}
	return newTeamsAppDefinition, nil
}

// TeamsAppIcon undocumented
type TeamsAppIcon struct {
	// Entity is the base model of TeamsAppIcon
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
	// HostedContent undocumented
	HostedContent *TeamworkHostedContent `json:"hostedContent,omitempty"`
}

func NewTeamsAppIcon() (*TeamsAppIcon, error) {
	newTeamsAppIcon := &TeamsAppIcon{
		ODataType: "#microsoft.graph.TeamsAppIcon",
	}
	return newTeamsAppIcon, nil
}

// TeamsAppInstallation undocumented
type TeamsAppInstallation struct {
	// Entity is the base model of TeamsAppInstallation
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ConsentedPermissionSet undocumented
	ConsentedPermissionSet *TeamsAppPermissionSet `json:"consentedPermissionSet,omitempty"`
	// TeamsApp undocumented
	TeamsApp *TeamsApp `json:"teamsApp,omitempty"`
	// TeamsAppDefinition undocumented
	TeamsAppDefinition *TeamsAppDefinition `json:"teamsAppDefinition,omitempty"`
}

func NewTeamsAppInstallation() (*TeamsAppInstallation, error) {
	newTeamsAppInstallation := &TeamsAppInstallation{
		ODataType: "#microsoft.graph.TeamsAppInstallation",
	}
	return newTeamsAppInstallation, nil
}

// TeamsAppInstalledEventMessageDetail undocumented
type TeamsAppInstalledEventMessageDetail struct {
	// EventMessageDetail is the base model of TeamsAppInstalledEventMessageDetail
	EventMessageDetail

	ODataType string `json:"@odata.type,omitempty"`
	// Initiator undocumented
	Initiator *IdentitySet `json:"initiator,omitempty"`
	// TeamsAppDisplayName undocumented
	TeamsAppDisplayName *string `json:"teamsAppDisplayName,omitempty"`
	// TeamsAppID undocumented
	TeamsAppID *string `json:"teamsAppId,omitempty"`
}

func NewTeamsAppInstalledEventMessageDetail() (*TeamsAppInstalledEventMessageDetail, error) {
	newTeamsAppInstalledEventMessageDetail := &TeamsAppInstalledEventMessageDetail{
		ODataType: "#microsoft.graph.TeamsAppInstalledEventMessageDetail",
	}
	return newTeamsAppInstalledEventMessageDetail, nil
}

// TeamsAppPermissionSet undocumented
type TeamsAppPermissionSet struct {
	// Object is the base model of TeamsAppPermissionSet
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ResourceSpecificPermissions undocumented
	ResourceSpecificPermissions []TeamsAppResourceSpecificPermission `json:"resourceSpecificPermissions,omitempty"`
}

func NewTeamsAppPermissionSet() (*TeamsAppPermissionSet, error) {
	newTeamsAppPermissionSet := &TeamsAppPermissionSet{
		ODataType: "#microsoft.graph.TeamsAppPermissionSet",
	}
	return newTeamsAppPermissionSet, nil
}

// TeamsAppRemovedEventMessageDetail undocumented
type TeamsAppRemovedEventMessageDetail struct {
	// EventMessageDetail is the base model of TeamsAppRemovedEventMessageDetail
	EventMessageDetail

	ODataType string `json:"@odata.type,omitempty"`
	// Initiator undocumented
	Initiator *IdentitySet `json:"initiator,omitempty"`
	// TeamsAppDisplayName undocumented
	TeamsAppDisplayName *string `json:"teamsAppDisplayName,omitempty"`
	// TeamsAppID undocumented
	TeamsAppID *string `json:"teamsAppId,omitempty"`
}

func NewTeamsAppRemovedEventMessageDetail() (*TeamsAppRemovedEventMessageDetail, error) {
	newTeamsAppRemovedEventMessageDetail := &TeamsAppRemovedEventMessageDetail{
		ODataType: "#microsoft.graph.TeamsAppRemovedEventMessageDetail",
	}
	return newTeamsAppRemovedEventMessageDetail, nil
}

// TeamsAppResourceSpecificPermission undocumented
type TeamsAppResourceSpecificPermission struct {
	// Object is the base model of TeamsAppResourceSpecificPermission
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// PermissionType undocumented
	PermissionType *TeamsAppResourceSpecificPermissionType `json:"permissionType,omitempty"`
	// PermissionValue undocumented
	PermissionValue *string `json:"permissionValue,omitempty"`
}

func NewTeamsAppResourceSpecificPermission() (*TeamsAppResourceSpecificPermission, error) {
	newTeamsAppResourceSpecificPermission := &TeamsAppResourceSpecificPermission{
		ODataType: "#microsoft.graph.TeamsAppResourceSpecificPermission",
	}
	return newTeamsAppResourceSpecificPermission, nil
}

// TeamsAppSettings undocumented
type TeamsAppSettings struct {
	// Entity is the base model of TeamsAppSettings
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AllowUserRequestsForAppAccess undocumented
	AllowUserRequestsForAppAccess *bool `json:"allowUserRequestsForAppAccess,omitempty"`
	// IsChatResourceSpecificConsentEnabled undocumented
	IsChatResourceSpecificConsentEnabled *bool `json:"isChatResourceSpecificConsentEnabled,omitempty"`
}

func NewTeamsAppSettings() (*TeamsAppSettings, error) {
	newTeamsAppSettings := &TeamsAppSettings{
		ODataType: "#microsoft.graph.TeamsAppSettings",
	}
	return newTeamsAppSettings, nil
}

// TeamsAppUpgradedEventMessageDetail undocumented
type TeamsAppUpgradedEventMessageDetail struct {
	// EventMessageDetail is the base model of TeamsAppUpgradedEventMessageDetail
	EventMessageDetail

	ODataType string `json:"@odata.type,omitempty"`
	// Initiator undocumented
	Initiator *IdentitySet `json:"initiator,omitempty"`
	// TeamsAppDisplayName undocumented
	TeamsAppDisplayName *string `json:"teamsAppDisplayName,omitempty"`
	// TeamsAppID undocumented
	TeamsAppID *string `json:"teamsAppId,omitempty"`
}

func NewTeamsAppUpgradedEventMessageDetail() (*TeamsAppUpgradedEventMessageDetail, error) {
	newTeamsAppUpgradedEventMessageDetail := &TeamsAppUpgradedEventMessageDetail{
		ODataType: "#microsoft.graph.TeamsAppUpgradedEventMessageDetail",
	}
	return newTeamsAppUpgradedEventMessageDetail, nil
}

// TeamsAsyncOperation undocumented
type TeamsAsyncOperation struct {
	// Entity is the base model of TeamsAsyncOperation
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AttemptsCount undocumented
	AttemptsCount *int `json:"attemptsCount,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Error undocumented
	Error *OperationError `json:"error,omitempty"`
	// LastActionDateTime undocumented
	LastActionDateTime *time.Time `json:"lastActionDateTime,omitempty"`
	// OperationType undocumented
	OperationType *TeamsAsyncOperationType `json:"operationType,omitempty"`
	// Status undocumented
	Status *TeamsAsyncOperationStatus `json:"status,omitempty"`
	// TargetResourceID undocumented
	TargetResourceID *string `json:"targetResourceId,omitempty"`
	// TargetResourceLocation undocumented
	TargetResourceLocation *string `json:"targetResourceLocation,omitempty"`
}

func NewTeamsAsyncOperation() (*TeamsAsyncOperation, error) {
	newTeamsAsyncOperation := &TeamsAsyncOperation{
		ODataType: "#microsoft.graph.TeamsAsyncOperation",
	}
	return newTeamsAsyncOperation, nil
}

// TeamsTab undocumented
type TeamsTab struct {
	// Entity is the base model of TeamsTab
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Configuration undocumented
	Configuration *TeamsTabConfiguration `json:"configuration,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// MessageID undocumented
	MessageID *string `json:"messageId,omitempty"`
	// SortOrderIndex undocumented
	SortOrderIndex *string `json:"sortOrderIndex,omitempty"`
	// TeamsAppID undocumented
	TeamsAppID *string `json:"teamsAppId,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
	// TeamsApp undocumented
	TeamsApp *TeamsApp `json:"teamsApp,omitempty"`
}

func NewTeamsTab() (*TeamsTab, error) {
	newTeamsTab := &TeamsTab{
		ODataType: "#microsoft.graph.TeamsTab",
	}
	return newTeamsTab, nil
}

// TeamsTabConfiguration undocumented
type TeamsTabConfiguration struct {
	// Object is the base model of TeamsTabConfiguration
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ContentURL undocumented
	ContentURL *string `json:"contentUrl,omitempty"`
	// EntityID undocumented
	EntityID *string `json:"entityId,omitempty"`
	// RemoveURL undocumented
	RemoveURL *string `json:"removeUrl,omitempty"`
	// WebsiteURL undocumented
	WebsiteURL *string `json:"websiteUrl,omitempty"`
}

func NewTeamsTabConfiguration() (*TeamsTabConfiguration, error) {
	newTeamsTabConfiguration := &TeamsTabConfiguration{
		ODataType: "#microsoft.graph.TeamsTabConfiguration",
	}
	return newTeamsTabConfiguration, nil
}

// TeamsTemplate undocumented
type TeamsTemplate struct {
	// Entity is the base model of TeamsTemplate
	Entity

	ODataType string `json:"@odata.type,omitempty"`
}

func NewTeamsTemplate() (*TeamsTemplate, error) {
	newTeamsTemplate := &TeamsTemplate{
		ODataType: "#microsoft.graph.TeamsTemplate",
	}
	return newTeamsTemplate, nil
}
