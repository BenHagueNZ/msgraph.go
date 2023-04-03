// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// AppCatalogs undocumented
type AppCatalogs struct {
	// Entity is the base model of AppCatalogs
	Entity

	ODataType string `json:"@odata.type"`
	// TeamsApps undocumented
	TeamsApps []TeamsApp `json:"teamsApps,omitempty"`
}

func NewAppCatalogs() (*AppCatalogs, error) {
	newAppCatalogs := &AppCatalogs{
		ODataType: "#microsoft.graph.AppCatalogs",
	}
	return newAppCatalogs, nil
}

// AppConfigurationSettingItem undocumented
type AppConfigurationSettingItem struct {
	// Object is the base model of AppConfigurationSettingItem
	Object

	ODataType string `json:"@odata.type"`
	// AppConfigKey undocumented
	AppConfigKey *string `json:"appConfigKey,omitempty"`
	// AppConfigKeyType undocumented
	AppConfigKeyType *MDMAppConfigKeyType `json:"appConfigKeyType,omitempty"`
	// AppConfigKeyValue undocumented
	AppConfigKeyValue *string `json:"appConfigKeyValue,omitempty"`
}

func NewAppConfigurationSettingItem() (*AppConfigurationSettingItem, error) {
	newAppConfigurationSettingItem := &AppConfigurationSettingItem{
		ODataType: "#microsoft.graph.AppConfigurationSettingItem",
	}
	return newAppConfigurationSettingItem, nil
}

// AppConsentApprovalRoute undocumented
type AppConsentApprovalRoute struct {
	// Entity is the base model of AppConsentApprovalRoute
	Entity

	ODataType string `json:"@odata.type"`
	// AppConsentRequests undocumented
	AppConsentRequests []AppConsentRequestObject `json:"appConsentRequests,omitempty"`
}

func NewAppConsentApprovalRoute() (*AppConsentApprovalRoute, error) {
	newAppConsentApprovalRoute := &AppConsentApprovalRoute{
		ODataType: "#microsoft.graph.AppConsentApprovalRoute",
	}
	return newAppConsentApprovalRoute, nil
}

// AppConsentRequestObject undocumented
type AppConsentRequestObject struct {
	// Entity is the base model of AppConsentRequestObject
	Entity

	ODataType string `json:"@odata.type"`
	// AppDisplayName undocumented
	AppDisplayName *string `json:"appDisplayName,omitempty"`
	// AppID undocumented
	AppID *string `json:"appId,omitempty"`
	// PendingScopes undocumented
	PendingScopes []AppConsentRequestScope `json:"pendingScopes,omitempty"`
	// UserConsentRequests undocumented
	UserConsentRequests []UserConsentRequestObject `json:"userConsentRequests,omitempty"`
}

func NewAppConsentRequestObject() (*AppConsentRequestObject, error) {
	newAppConsentRequestObject := &AppConsentRequestObject{
		ODataType: "#microsoft.graph.AppConsentRequestObject",
	}
	return newAppConsentRequestObject, nil
}

// AppConsentRequestScope undocumented
type AppConsentRequestScope struct {
	// Object is the base model of AppConsentRequestScope
	Object

	ODataType string `json:"@odata.type"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
}

func NewAppConsentRequestScope() (*AppConsentRequestScope, error) {
	newAppConsentRequestScope := &AppConsentRequestScope{
		ODataType: "#microsoft.graph.AppConsentRequestScope",
	}
	return newAppConsentRequestScope, nil
}

// AppHostedMediaConfig undocumented
type AppHostedMediaConfig struct {
	// MediaConfig is the base model of AppHostedMediaConfig
	MediaConfig

	ODataType string `json:"@odata.type"`
	// Blob undocumented
	Blob *string `json:"blob,omitempty"`
}

func NewAppHostedMediaConfig() (*AppHostedMediaConfig, error) {
	newAppHostedMediaConfig := &AppHostedMediaConfig{
		ODataType: "#microsoft.graph.AppHostedMediaConfig",
	}
	return newAppHostedMediaConfig, nil
}

// AppIdentity undocumented
type AppIdentity struct {
	// Object is the base model of AppIdentity
	Object

	ODataType string `json:"@odata.type"`
	// AppID undocumented
	AppID *string `json:"appId,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ServicePrincipalID undocumented
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty"`
	// ServicePrincipalName undocumented
	ServicePrincipalName *string `json:"servicePrincipalName,omitempty"`
}

func NewAppIdentity() (*AppIdentity, error) {
	newAppIdentity := &AppIdentity{
		ODataType: "#microsoft.graph.AppIdentity",
	}
	return newAppIdentity, nil
}

// AppListItem undocumented
type AppListItem struct {
	// Object is the base model of AppListItem
	Object

	ODataType string `json:"@odata.type"`
	// AppID undocumented
	AppID *string `json:"appId,omitempty"`
	// AppStoreURL undocumented
	AppStoreURL *string `json:"appStoreUrl,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Publisher undocumented
	Publisher *string `json:"publisher,omitempty"`
}

func NewAppListItem() (*AppListItem, error) {
	newAppListItem := &AppListItem{
		ODataType: "#microsoft.graph.AppListItem",
	}
	return newAppListItem, nil
}

// AppManagementConfiguration undocumented
type AppManagementConfiguration struct {
	// Object is the base model of AppManagementConfiguration
	Object

	ODataType string `json:"@odata.type"`
	// KeyCredentials undocumented
	KeyCredentials []KeyCredentialConfiguration `json:"keyCredentials,omitempty"`
	// PasswordCredentials undocumented
	PasswordCredentials []PasswordCredentialConfiguration `json:"passwordCredentials,omitempty"`
}

func NewAppManagementConfiguration() (*AppManagementConfiguration, error) {
	newAppManagementConfiguration := &AppManagementConfiguration{
		ODataType: "#microsoft.graph.AppManagementConfiguration",
	}
	return newAppManagementConfiguration, nil
}

// AppManagementPolicy undocumented
type AppManagementPolicy struct {
	// PolicyBase is the base model of AppManagementPolicy
	PolicyBase

	ODataType string `json:"@odata.type"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Restrictions undocumented
	Restrictions *AppManagementConfiguration `json:"restrictions,omitempty"`
	// AppliesTo undocumented
	AppliesTo []DirectoryObject `json:"appliesTo,omitempty"`
}

func NewAppManagementPolicy() (*AppManagementPolicy, error) {
	newAppManagementPolicy := &AppManagementPolicy{
		ODataType: "#microsoft.graph.AppManagementPolicy",
	}
	return newAppManagementPolicy, nil
}

// AppRole undocumented
type AppRole struct {
	// Object is the base model of AppRole
	Object

	ODataType string `json:"@odata.type"`
	// AllowedMemberTypes undocumented
	AllowedMemberTypes []string `json:"allowedMemberTypes,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ID undocumented
	ID *UUID `json:"id,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Origin undocumented
	Origin *string `json:"origin,omitempty"`
	// Value undocumented
	Value *string `json:"value,omitempty"`
}

func NewAppRole() (*AppRole, error) {
	newAppRole := &AppRole{
		ODataType: "#microsoft.graph.AppRole",
	}
	return newAppRole, nil
}

// AppRoleAssignment undocumented
type AppRoleAssignment struct {
	// DirectoryObject is the base model of AppRoleAssignment
	DirectoryObject

	ODataType string `json:"@odata.type"`
	// AppRoleID undocumented
	AppRoleID *UUID `json:"appRoleId,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// PrincipalDisplayName undocumented
	PrincipalDisplayName *string `json:"principalDisplayName,omitempty"`
	// PrincipalID undocumented
	PrincipalID *UUID `json:"principalId,omitempty"`
	// PrincipalType undocumented
	PrincipalType *string `json:"principalType,omitempty"`
	// ResourceDisplayName undocumented
	ResourceDisplayName *string `json:"resourceDisplayName,omitempty"`
	// ResourceID undocumented
	ResourceID *UUID `json:"resourceId,omitempty"`
}

func NewAppRoleAssignment() (*AppRoleAssignment, error) {
	newAppRoleAssignment := &AppRoleAssignment{
		ODataType: "#microsoft.graph.AppRoleAssignment",
	}
	return newAppRoleAssignment, nil
}

// AppScope undocumented
type AppScope struct {
	// Entity is the base model of AppScope
	Entity

	ODataType string `json:"@odata.type"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
}

func NewAppScope() (*AppScope, error) {
	newAppScope := &AppScope{
		ODataType: "#microsoft.graph.AppScope",
	}
	return newAppScope, nil
}
