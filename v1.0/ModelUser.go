// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"time"
)

// User undocumented
type User struct {
	// DirectoryObject is the base model of User
	DirectoryObject

	ODataType string `json:"@odata.type"`
	// AccountEnabled undocumented
	AccountEnabled *bool `json:"accountEnabled,omitempty"`
	// AgeGroup undocumented
	AgeGroup *string `json:"ageGroup,omitempty"`
	// AssignedLicenses undocumented
	AssignedLicenses []AssignedLicense `json:"assignedLicenses,omitempty"`
	// AssignedPlans undocumented
	AssignedPlans []AssignedPlan `json:"assignedPlans,omitempty"`
	// AuthorizationInfo undocumented
	AuthorizationInfo *AuthorizationInfo `json:"authorizationInfo,omitempty"`
	// BusinessPhones undocumented
	BusinessPhones []string `json:"businessPhones,omitempty"`
	// City undocumented
	City *string `json:"city,omitempty"`
	// CompanyName undocumented
	CompanyName *string `json:"companyName,omitempty"`
	// ConsentProvidedForMinor undocumented
	ConsentProvidedForMinor *string `json:"consentProvidedForMinor,omitempty"`
	// Country undocumented
	Country *string `json:"country,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// CreationType undocumented
	CreationType *string `json:"creationType,omitempty"`
	// Department undocumented
	Department *string `json:"department,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// EmployeeHireDate undocumented
	EmployeeHireDate *time.Time `json:"employeeHireDate,omitempty"`
	// EmployeeID undocumented
	EmployeeID *string `json:"employeeId,omitempty"`
	// EmployeeLeaveDateTime undocumented
	EmployeeLeaveDateTime *time.Time `json:"employeeLeaveDateTime,omitempty"`
	// EmployeeOrgData undocumented
	EmployeeOrgData *EmployeeOrgData `json:"employeeOrgData,omitempty"`
	// EmployeeType undocumented
	EmployeeType *string `json:"employeeType,omitempty"`
	// ExternalUserState undocumented
	ExternalUserState *string `json:"externalUserState,omitempty"`
	// ExternalUserStateChangeDateTime undocumented
	ExternalUserStateChangeDateTime *time.Time `json:"externalUserStateChangeDateTime,omitempty"`
	// FaxNumber undocumented
	FaxNumber *string `json:"faxNumber,omitempty"`
	// GivenName undocumented
	GivenName *string `json:"givenName,omitempty"`
	// Identities undocumented
	Identities []ObjectIdentity `json:"identities,omitempty"`
	// ImAddresses undocumented
	ImAddresses []string `json:"imAddresses,omitempty"`
	// IsResourceAccount undocumented
	IsResourceAccount *bool `json:"isResourceAccount,omitempty"`
	// JobTitle undocumented
	JobTitle *string `json:"jobTitle,omitempty"`
	// LastPasswordChangeDateTime undocumented
	LastPasswordChangeDateTime *time.Time `json:"lastPasswordChangeDateTime,omitempty"`
	// LegalAgeGroupClassification undocumented
	LegalAgeGroupClassification *string `json:"legalAgeGroupClassification,omitempty"`
	// LicenseAssignmentStates undocumented
	LicenseAssignmentStates []LicenseAssignmentState `json:"licenseAssignmentStates,omitempty"`
	// Mail undocumented
	Mail *string `json:"mail,omitempty"`
	// MailNickname undocumented
	MailNickname *string `json:"mailNickname,omitempty"`
	// MobilePhone undocumented
	MobilePhone *string `json:"mobilePhone,omitempty"`
	// OfficeLocation undocumented
	OfficeLocation *string `json:"officeLocation,omitempty"`
	// OnPremisesDistinguishedName undocumented
	OnPremisesDistinguishedName *string `json:"onPremisesDistinguishedName,omitempty"`
	// OnPremisesDomainName undocumented
	OnPremisesDomainName *string `json:"onPremisesDomainName,omitempty"`
	// OnPremisesExtensionAttributes undocumented
	OnPremisesExtensionAttributes *OnPremisesExtensionAttributes `json:"onPremisesExtensionAttributes,omitempty"`
	// OnPremisesImmutableID undocumented
	OnPremisesImmutableID *string `json:"onPremisesImmutableId,omitempty"`
	// OnPremisesLastSyncDateTime undocumented
	OnPremisesLastSyncDateTime *time.Time `json:"onPremisesLastSyncDateTime,omitempty"`
	// OnPremisesProvisioningErrors undocumented
	OnPremisesProvisioningErrors []OnPremisesProvisioningError `json:"onPremisesProvisioningErrors,omitempty"`
	// OnPremisesSamAccountName undocumented
	OnPremisesSamAccountName *string `json:"onPremisesSamAccountName,omitempty"`
	// OnPremisesSecurityIdentifier undocumented
	OnPremisesSecurityIdentifier *string `json:"onPremisesSecurityIdentifier,omitempty"`
	// OnPremisesSyncEnabled undocumented
	OnPremisesSyncEnabled *bool `json:"onPremisesSyncEnabled,omitempty"`
	// OnPremisesUserPrincipalName undocumented
	OnPremisesUserPrincipalName *string `json:"onPremisesUserPrincipalName,omitempty"`
	// OtherMails undocumented
	OtherMails []string `json:"otherMails,omitempty"`
	// PasswordPolicies undocumented
	PasswordPolicies *string `json:"passwordPolicies,omitempty"`
	// PasswordProfile undocumented
	PasswordProfile *PasswordProfile `json:"passwordProfile,omitempty"`
	// PostalCode undocumented
	PostalCode *string `json:"postalCode,omitempty"`
	// PreferredDataLocation undocumented
	PreferredDataLocation *string `json:"preferredDataLocation,omitempty"`
	// PreferredLanguage undocumented
	PreferredLanguage *string `json:"preferredLanguage,omitempty"`
	// ProvisionedPlans undocumented
	ProvisionedPlans []ProvisionedPlan `json:"provisionedPlans,omitempty"`
	// ProxyAddresses undocumented
	ProxyAddresses []string `json:"proxyAddresses,omitempty"`
	// SecurityIdentifier undocumented
	SecurityIdentifier *string `json:"securityIdentifier,omitempty"`
	// ShowInAddressList undocumented
	ShowInAddressList *bool `json:"showInAddressList,omitempty"`
	// SignInSessionsValidFromDateTime undocumented
	SignInSessionsValidFromDateTime *time.Time `json:"signInSessionsValidFromDateTime,omitempty"`
	// State undocumented
	State *string `json:"state,omitempty"`
	// StreetAddress undocumented
	StreetAddress *string `json:"streetAddress,omitempty"`
	// Surname undocumented
	Surname *string `json:"surname,omitempty"`
	// UsageLocation undocumented
	UsageLocation *string `json:"usageLocation,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// UserType undocumented
	UserType *string `json:"userType,omitempty"`
	// MailboxSettings undocumented
	MailboxSettings *MailboxSettings `json:"mailboxSettings,omitempty"`
	// DeviceEnrollmentLimit undocumented
	DeviceEnrollmentLimit *int `json:"deviceEnrollmentLimit,omitempty"`
	// AboutMe undocumented
	AboutMe *string `json:"aboutMe,omitempty"`
	// Birthday undocumented
	Birthday *time.Time `json:"birthday,omitempty"`
	// HireDate undocumented
	HireDate *time.Time `json:"hireDate,omitempty"`
	// Interests undocumented
	Interests []string `json:"interests,omitempty"`
	// MySite undocumented
	MySite *string `json:"mySite,omitempty"`
	// PastProjects undocumented
	PastProjects []string `json:"pastProjects,omitempty"`
	// PreferredName undocumented
	PreferredName *string `json:"preferredName,omitempty"`
	// Responsibilities undocumented
	Responsibilities []string `json:"responsibilities,omitempty"`
	// Schools undocumented
	Schools []string `json:"schools,omitempty"`
	// Skills undocumented
	Skills []string `json:"skills,omitempty"`
	// AppRoleAssignments undocumented
	AppRoleAssignments []AppRoleAssignment `json:"appRoleAssignments,omitempty"`
	// CreatedObjects undocumented
	CreatedObjects []DirectoryObject `json:"createdObjects,omitempty"`
	// DirectReports undocumented
	DirectReports []DirectoryObject `json:"directReports,omitempty"`
	// LicenseDetails undocumented
	LicenseDetails []LicenseDetails `json:"licenseDetails,omitempty"`
	// Manager undocumented
	Manager *DirectoryObject `json:"manager,omitempty"`
	// MemberOf undocumented
	MemberOf []DirectoryObject `json:"memberOf,omitempty"`
	// OAuth2PermissionGrants undocumented
	OAuth2PermissionGrants []OAuth2PermissionGrant `json:"oauth2PermissionGrants,omitempty"`
	// OwnedDevices undocumented
	OwnedDevices []DirectoryObject `json:"ownedDevices,omitempty"`
	// OwnedObjects undocumented
	OwnedObjects []DirectoryObject `json:"ownedObjects,omitempty"`
	// RegisteredDevices undocumented
	RegisteredDevices []DirectoryObject `json:"registeredDevices,omitempty"`
	// ScopedRoleMemberOf undocumented
	ScopedRoleMemberOf []ScopedRoleMembership `json:"scopedRoleMemberOf,omitempty"`
	// TransitiveMemberOf undocumented
	TransitiveMemberOf []DirectoryObject `json:"transitiveMemberOf,omitempty"`
	// Calendar undocumented
	Calendar *Calendar `json:"calendar,omitempty"`
	// CalendarGroups undocumented
	CalendarGroups []CalendarGroup `json:"calendarGroups,omitempty"`
	// Calendars undocumented
	Calendars []Calendar `json:"calendars,omitempty"`
	// CalendarView undocumented
	CalendarView []Event `json:"calendarView,omitempty"`
	// ContactFolders undocumented
	ContactFolders []ContactFolder `json:"contactFolders,omitempty"`
	// Contacts undocumented
	Contacts []Contact `json:"contacts,omitempty"`
	// Events undocumented
	Events []Event `json:"events,omitempty"`
	// InferenceClassification undocumented
	InferenceClassification *InferenceClassification `json:"inferenceClassification,omitempty"`
	// MailFolders undocumented
	MailFolders []MailFolder `json:"mailFolders,omitempty"`
	// Messages undocumented
	Messages []Message `json:"messages,omitempty"`
	// Outlook undocumented
	Outlook *OutlookUser `json:"outlook,omitempty"`
	// People undocumented
	People []Person `json:"people,omitempty"`
	// Drive undocumented
	Drive *Drive `json:"drive,omitempty"`
	// Drives undocumented
	Drives []Drive `json:"drives,omitempty"`
	// FollowedSites undocumented
	FollowedSites []Site `json:"followedSites,omitempty"`
	// Extensions undocumented
	Extensions []Extension `json:"extensions,omitempty"`
	// AgreementAcceptances undocumented
	AgreementAcceptances []AgreementAcceptance `json:"agreementAcceptances,omitempty"`
	// ManagedDevices undocumented
	ManagedDevices []ManagedDevice `json:"managedDevices,omitempty"`
	// ManagedAppRegistrations undocumented
	ManagedAppRegistrations []ManagedAppRegistration `json:"managedAppRegistrations,omitempty"`
	// DeviceManagementTroubleshootingEvents undocumented
	DeviceManagementTroubleshootingEvents []DeviceManagementTroubleshootingEvent `json:"deviceManagementTroubleshootingEvents,omitempty"`
	// Planner undocumented
	Planner *PlannerUser `json:"planner,omitempty"`
	// Insights undocumented
	Insights *OfficeGraphInsights `json:"insights,omitempty"`
	// Settings undocumented
	Settings *UserSettings `json:"settings,omitempty"`
	// Onenote undocumented
	Onenote *Onenote `json:"onenote,omitempty"`
	// Photo undocumented
	Photo *ProfilePhoto `json:"photo,omitempty"`
	// Photos undocumented
	Photos []ProfilePhoto `json:"photos,omitempty"`
	// Activities undocumented
	Activities []UserActivity `json:"activities,omitempty"`
	// OnlineMeetings undocumented
	OnlineMeetings []OnlineMeeting `json:"onlineMeetings,omitempty"`
	// Presence undocumented
	Presence *Presence `json:"presence,omitempty"`
	// Authentication undocumented
	Authentication *Authentication `json:"authentication,omitempty"`
	// Chats undocumented
	Chats []Chat `json:"chats,omitempty"`
	// JoinedTeams undocumented
	JoinedTeams []Team `json:"joinedTeams,omitempty"`
	// Teamwork undocumented
	Teamwork *UserTeamwork `json:"teamwork,omitempty"`
	// Todo undocumented
	Todo *Todo `json:"todo,omitempty"`
}

func NewUser() (*User, error) {
	newUser := &User{
		ODataType: "#microsoft.graph.User",
	}
	return newUser, nil
}

// UserActivity undocumented
type UserActivity struct {
	// Entity is the base model of UserActivity
	Entity

	ODataType string `json:"@odata.type"`
	// ActivationURL undocumented
	ActivationURL *string `json:"activationUrl,omitempty"`
	// ActivitySourceHost undocumented
	ActivitySourceHost *string `json:"activitySourceHost,omitempty"`
	// AppActivityID undocumented
	AppActivityID *string `json:"appActivityId,omitempty"`
	// AppDisplayName undocumented
	AppDisplayName *string `json:"appDisplayName,omitempty"`
	// ContentInfo undocumented
	ContentInfo json.RawMessage `json:"contentInfo,omitempty"`
	// ContentURL undocumented
	ContentURL *string `json:"contentUrl,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// ExpirationDateTime undocumented
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// FallbackURL undocumented
	FallbackURL *string `json:"fallbackUrl,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Status undocumented
	Status *Status `json:"status,omitempty"`
	// UserTimezone undocumented
	UserTimezone *string `json:"userTimezone,omitempty"`
	// VisualElements undocumented
	VisualElements *VisualInfo `json:"visualElements,omitempty"`
	// HistoryItems undocumented
	HistoryItems []ActivityHistoryItem `json:"historyItems,omitempty"`
}

func NewUserActivity() (*UserActivity, error) {
	newUserActivity := &UserActivity{
		ODataType: "#microsoft.graph.UserActivity",
	}
	return newUserActivity, nil
}

// UserAttributeValuesItem undocumented
type UserAttributeValuesItem struct {
	// Object is the base model of UserAttributeValuesItem
	Object

	ODataType string `json:"@odata.type"`
	// IsDefault undocumented
	IsDefault *bool `json:"isDefault,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Value undocumented
	Value *string `json:"value,omitempty"`
}

func NewUserAttributeValuesItem() (*UserAttributeValuesItem, error) {
	newUserAttributeValuesItem := &UserAttributeValuesItem{
		ODataType: "#microsoft.graph.UserAttributeValuesItem",
	}
	return newUserAttributeValuesItem, nil
}

// UserConsentRequestObject undocumented
type UserConsentRequestObject struct {
	// RequestObject is the base model of UserConsentRequestObject
	RequestObject

	ODataType string `json:"@odata.type"`
	// Reason undocumented
	Reason *string `json:"reason,omitempty"`
	// Approval undocumented
	Approval *Approval `json:"approval,omitempty"`
}

func NewUserConsentRequestObject() (*UserConsentRequestObject, error) {
	newUserConsentRequestObject := &UserConsentRequestObject{
		ODataType: "#microsoft.graph.UserConsentRequestObject",
	}
	return newUserConsentRequestObject, nil
}

// UserExperienceAnalyticsDevicePerformance undocumented
type UserExperienceAnalyticsDevicePerformance struct {
	// Entity is the base model of UserExperienceAnalyticsDevicePerformance
	Entity

	ODataType string `json:"@odata.type"`
	// AverageBlueScreens undocumented
	AverageBlueScreens *float64 `json:"averageBlueScreens,omitempty"`
	// AverageRestarts undocumented
	AverageRestarts *float64 `json:"averageRestarts,omitempty"`
	// BlueScreenCount undocumented
	BlueScreenCount *int `json:"blueScreenCount,omitempty"`
	// BootScore undocumented
	BootScore *int `json:"bootScore,omitempty"`
	// CoreBootTimeInMs undocumented
	CoreBootTimeInMs *int `json:"coreBootTimeInMs,omitempty"`
	// CoreLoginTimeInMs undocumented
	CoreLoginTimeInMs *int `json:"coreLoginTimeInMs,omitempty"`
	// DeviceCount undocumented
	DeviceCount *int `json:"deviceCount,omitempty"`
	// DeviceName undocumented
	DeviceName *string `json:"deviceName,omitempty"`
	// DiskType undocumented
	DiskType *DiskType `json:"diskType,omitempty"`
	// GroupPolicyBootTimeInMs undocumented
	GroupPolicyBootTimeInMs *int `json:"groupPolicyBootTimeInMs,omitempty"`
	// GroupPolicyLoginTimeInMs undocumented
	GroupPolicyLoginTimeInMs *int `json:"groupPolicyLoginTimeInMs,omitempty"`
	// HealthStatus undocumented
	HealthStatus *UserExperienceAnalyticsHealthState `json:"healthStatus,omitempty"`
	// LoginScore undocumented
	LoginScore *int `json:"loginScore,omitempty"`
	// Manufacturer undocumented
	Manufacturer *string `json:"manufacturer,omitempty"`
	// Model undocumented
	Model *string `json:"model,omitempty"`
	// ModelStartupPerformanceScore undocumented
	ModelStartupPerformanceScore *float64 `json:"modelStartupPerformanceScore,omitempty"`
	// OperatingSystemVersion undocumented
	OperatingSystemVersion *string `json:"operatingSystemVersion,omitempty"`
	// ResponsiveDesktopTimeInMs undocumented
	ResponsiveDesktopTimeInMs *int `json:"responsiveDesktopTimeInMs,omitempty"`
	// RestartCount undocumented
	RestartCount *int `json:"restartCount,omitempty"`
	// StartupPerformanceScore undocumented
	StartupPerformanceScore *float64 `json:"startupPerformanceScore,omitempty"`
}

func NewUserExperienceAnalyticsDevicePerformance() (*UserExperienceAnalyticsDevicePerformance, error) {
	newUserExperienceAnalyticsDevicePerformance := &UserExperienceAnalyticsDevicePerformance{
		ODataType: "#microsoft.graph.UserExperienceAnalyticsDevicePerformance",
	}
	return newUserExperienceAnalyticsDevicePerformance, nil
}

// UserFlowAPIConnectorConfiguration undocumented
type UserFlowAPIConnectorConfiguration struct {
	// Object is the base model of UserFlowAPIConnectorConfiguration
	Object

	ODataType string `json:"@odata.type"`
	// PostAttributeCollection undocumented
	PostAttributeCollection *IdentityAPIConnector `json:"postAttributeCollection,omitempty"`
	// PostFederationSignup undocumented
	PostFederationSignup *IdentityAPIConnector `json:"postFederationSignup,omitempty"`
}

func NewUserFlowAPIConnectorConfiguration() (*UserFlowAPIConnectorConfiguration, error) {
	newUserFlowAPIConnectorConfiguration := &UserFlowAPIConnectorConfiguration{
		ODataType: "#microsoft.graph.UserFlowApiConnectorConfiguration",
	}
	return newUserFlowAPIConnectorConfiguration, nil
}

// UserFlowLanguageConfiguration undocumented
type UserFlowLanguageConfiguration struct {
	// Entity is the base model of UserFlowLanguageConfiguration
	Entity

	ODataType string `json:"@odata.type"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// DefaultPages undocumented
	DefaultPages []UserFlowLanguagePage `json:"defaultPages,omitempty"`
	// OverridesPages undocumented
	OverridesPages []UserFlowLanguagePage `json:"overridesPages,omitempty"`
}

func NewUserFlowLanguageConfiguration() (*UserFlowLanguageConfiguration, error) {
	newUserFlowLanguageConfiguration := &UserFlowLanguageConfiguration{
		ODataType: "#microsoft.graph.UserFlowLanguageConfiguration",
	}
	return newUserFlowLanguageConfiguration, nil
}

// UserFlowLanguagePage undocumented
type UserFlowLanguagePage struct {
	// Entity is the base model of UserFlowLanguagePage
	Entity

	ODataType string `json:"@odata.type"`
}

func NewUserFlowLanguagePage() (*UserFlowLanguagePage, error) {
	newUserFlowLanguagePage := &UserFlowLanguagePage{
		ODataType: "#microsoft.graph.UserFlowLanguagePage",
	}
	return newUserFlowLanguagePage, nil
}

// UserIdentity undocumented
type UserIdentity struct {
	// Identity is the base model of UserIdentity
	Identity

	ODataType string `json:"@odata.type"`
	// IPAddress undocumented
	IPAddress *string `json:"ipAddress,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}

func NewUserIdentity() (*UserIdentity, error) {
	newUserIdentity := &UserIdentity{
		ODataType: "#microsoft.graph.UserIdentity",
	}
	return newUserIdentity, nil
}

// UserInstallStateSummary undocumented
type UserInstallStateSummary struct {
	// Entity is the base model of UserInstallStateSummary
	Entity

	ODataType string `json:"@odata.type"`
	// FailedDeviceCount undocumented
	FailedDeviceCount *int `json:"failedDeviceCount,omitempty"`
	// InstalledDeviceCount undocumented
	InstalledDeviceCount *int `json:"installedDeviceCount,omitempty"`
	// NotInstalledDeviceCount undocumented
	NotInstalledDeviceCount *int `json:"notInstalledDeviceCount,omitempty"`
	// UserName undocumented
	UserName *string `json:"userName,omitempty"`
	// DeviceStates undocumented
	DeviceStates []DeviceInstallState `json:"deviceStates,omitempty"`
}

func NewUserInstallStateSummary() (*UserInstallStateSummary, error) {
	newUserInstallStateSummary := &UserInstallStateSummary{
		ODataType: "#microsoft.graph.UserInstallStateSummary",
	}
	return newUserInstallStateSummary, nil
}

// UserScopeTeamsAppInstallation undocumented
type UserScopeTeamsAppInstallation struct {
	// TeamsAppInstallation is the base model of UserScopeTeamsAppInstallation
	TeamsAppInstallation

	ODataType string `json:"@odata.type"`
	// Chat undocumented
	Chat *Chat `json:"chat,omitempty"`
}

func NewUserScopeTeamsAppInstallation() (*UserScopeTeamsAppInstallation, error) {
	newUserScopeTeamsAppInstallation := &UserScopeTeamsAppInstallation{
		ODataType: "#microsoft.graph.UserScopeTeamsAppInstallation",
	}
	return newUserScopeTeamsAppInstallation, nil
}

// UserSecurityState undocumented
type UserSecurityState struct {
	// Object is the base model of UserSecurityState
	Object

	ODataType string `json:"@odata.type"`
	// AadUserID undocumented
	AadUserID *string `json:"aadUserId,omitempty"`
	// AccountName undocumented
	AccountName *string `json:"accountName,omitempty"`
	// DomainName undocumented
	DomainName *string `json:"domainName,omitempty"`
	// EmailRole undocumented
	EmailRole *EmailRole `json:"emailRole,omitempty"`
	// IsVPN undocumented
	IsVPN *bool `json:"isVpn,omitempty"`
	// LogonDateTime undocumented
	LogonDateTime *time.Time `json:"logonDateTime,omitempty"`
	// LogonID undocumented
	LogonID *string `json:"logonId,omitempty"`
	// LogonIP undocumented
	LogonIP *string `json:"logonIp,omitempty"`
	// LogonLocation undocumented
	LogonLocation *string `json:"logonLocation,omitempty"`
	// LogonType undocumented
	LogonType *LogonType `json:"logonType,omitempty"`
	// OnPremisesSecurityIdentifier undocumented
	OnPremisesSecurityIdentifier *string `json:"onPremisesSecurityIdentifier,omitempty"`
	// RiskScore undocumented
	RiskScore *string `json:"riskScore,omitempty"`
	// UserAccountType undocumented
	UserAccountType *UserAccountSecurityType `json:"userAccountType,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}

func NewUserSecurityState() (*UserSecurityState, error) {
	newUserSecurityState := &UserSecurityState{
		ODataType: "#microsoft.graph.UserSecurityState",
	}
	return newUserSecurityState, nil
}

// UserSettings undocumented
type UserSettings struct {
	// Entity is the base model of UserSettings
	Entity

	ODataType string `json:"@odata.type"`
	// ContributionToContentDiscoveryAsOrganizationDisabled undocumented
	ContributionToContentDiscoveryAsOrganizationDisabled *bool `json:"contributionToContentDiscoveryAsOrganizationDisabled,omitempty"`
	// ContributionToContentDiscoveryDisabled undocumented
	ContributionToContentDiscoveryDisabled *bool `json:"contributionToContentDiscoveryDisabled,omitempty"`
	// ShiftPreferences undocumented
	ShiftPreferences *ShiftPreferences `json:"shiftPreferences,omitempty"`
}

func NewUserSettings() (*UserSettings, error) {
	newUserSettings := &UserSettings{
		ODataType: "#microsoft.graph.UserSettings",
	}
	return newUserSettings, nil
}

// UserSimulationDetails undocumented
type UserSimulationDetails struct {
	// Object is the base model of UserSimulationDetails
	Object

	ODataType string `json:"@odata.type"`
	// AssignedTrainingsCount undocumented
	AssignedTrainingsCount *int `json:"assignedTrainingsCount,omitempty"`
	// CompletedTrainingsCount undocumented
	CompletedTrainingsCount *int `json:"completedTrainingsCount,omitempty"`
	// CompromisedDateTime undocumented
	CompromisedDateTime *time.Time `json:"compromisedDateTime,omitempty"`
	// InProgressTrainingsCount undocumented
	InProgressTrainingsCount *int `json:"inProgressTrainingsCount,omitempty"`
	// IsCompromised undocumented
	IsCompromised *bool `json:"isCompromised,omitempty"`
	// ReportedPhishDateTime undocumented
	ReportedPhishDateTime *time.Time `json:"reportedPhishDateTime,omitempty"`
	// SimulationEvents undocumented
	SimulationEvents []UserSimulationEventInfo `json:"simulationEvents,omitempty"`
	// SimulationUser undocumented
	SimulationUser *AttackSimulationUser `json:"simulationUser,omitempty"`
	// TrainingEvents undocumented
	TrainingEvents []UserTrainingEventInfo `json:"trainingEvents,omitempty"`
}

func NewUserSimulationDetails() (*UserSimulationDetails, error) {
	newUserSimulationDetails := &UserSimulationDetails{
		ODataType: "#microsoft.graph.UserSimulationDetails",
	}
	return newUserSimulationDetails, nil
}

// UserSimulationEventInfo undocumented
type UserSimulationEventInfo struct {
	// Object is the base model of UserSimulationEventInfo
	Object

	ODataType string `json:"@odata.type"`
	// Browser undocumented
	Browser *string `json:"browser,omitempty"`
	// EventDateTime undocumented
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`
	// EventName undocumented
	EventName *string `json:"eventName,omitempty"`
	// IPAddress undocumented
	IPAddress *string `json:"ipAddress,omitempty"`
	// OsPlatformDeviceDetails undocumented
	OsPlatformDeviceDetails *string `json:"osPlatformDeviceDetails,omitempty"`
}

func NewUserSimulationEventInfo() (*UserSimulationEventInfo, error) {
	newUserSimulationEventInfo := &UserSimulationEventInfo{
		ODataType: "#microsoft.graph.UserSimulationEventInfo",
	}
	return newUserSimulationEventInfo, nil
}

// UserTeamwork undocumented
type UserTeamwork struct {
	// Entity is the base model of UserTeamwork
	Entity

	ODataType string `json:"@odata.type"`
	// AssociatedTeams undocumented
	AssociatedTeams []AssociatedTeamInfo `json:"associatedTeams,omitempty"`
	// InstalledApps undocumented
	InstalledApps []UserScopeTeamsAppInstallation `json:"installedApps,omitempty"`
}

func NewUserTeamwork() (*UserTeamwork, error) {
	newUserTeamwork := &UserTeamwork{
		ODataType: "#microsoft.graph.UserTeamwork",
	}
	return newUserTeamwork, nil
}

// UserTrainingContentEventInfo undocumented
type UserTrainingContentEventInfo struct {
	// Object is the base model of UserTrainingContentEventInfo
	Object

	ODataType string `json:"@odata.type"`
	// Browser undocumented
	Browser *string `json:"browser,omitempty"`
	// ContentDateTime undocumented
	ContentDateTime *time.Time `json:"contentDateTime,omitempty"`
	// IPAddress undocumented
	IPAddress *string `json:"ipAddress,omitempty"`
	// OsPlatformDeviceDetails undocumented
	OsPlatformDeviceDetails *string `json:"osPlatformDeviceDetails,omitempty"`
	// PotentialScoreImpact undocumented
	PotentialScoreImpact *float64 `json:"potentialScoreImpact,omitempty"`
}

func NewUserTrainingContentEventInfo() (*UserTrainingContentEventInfo, error) {
	newUserTrainingContentEventInfo := &UserTrainingContentEventInfo{
		ODataType: "#microsoft.graph.UserTrainingContentEventInfo",
	}
	return newUserTrainingContentEventInfo, nil
}

// UserTrainingEventInfo undocumented
type UserTrainingEventInfo struct {
	// Object is the base model of UserTrainingEventInfo
	Object

	ODataType string `json:"@odata.type"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LatestTrainingStatus undocumented
	LatestTrainingStatus *TrainingStatus `json:"latestTrainingStatus,omitempty"`
	// TrainingAssignedProperties undocumented
	TrainingAssignedProperties *UserTrainingContentEventInfo `json:"trainingAssignedProperties,omitempty"`
	// TrainingCompletedProperties undocumented
	TrainingCompletedProperties *UserTrainingContentEventInfo `json:"trainingCompletedProperties,omitempty"`
	// TrainingUpdatedProperties undocumented
	TrainingUpdatedProperties *UserTrainingContentEventInfo `json:"trainingUpdatedProperties,omitempty"`
}

func NewUserTrainingEventInfo() (*UserTrainingEventInfo, error) {
	newUserTrainingEventInfo := &UserTrainingEventInfo{
		ODataType: "#microsoft.graph.UserTrainingEventInfo",
	}
	return newUserTrainingEventInfo, nil
}

// UserTrainingStatusInfo undocumented
type UserTrainingStatusInfo struct {
	// Object is the base model of UserTrainingStatusInfo
	Object

	ODataType string `json:"@odata.type"`
	// AssignedDateTime undocumented
	AssignedDateTime *time.Time `json:"assignedDateTime,omitempty"`
	// CompletionDateTime undocumented
	CompletionDateTime *time.Time `json:"completionDateTime,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// TrainingStatus undocumented
	TrainingStatus *TrainingStatus `json:"trainingStatus,omitempty"`
}

func NewUserTrainingStatusInfo() (*UserTrainingStatusInfo, error) {
	newUserTrainingStatusInfo := &UserTrainingStatusInfo{
		ODataType: "#microsoft.graph.UserTrainingStatusInfo",
	}
	return newUserTrainingStatusInfo, nil
}
