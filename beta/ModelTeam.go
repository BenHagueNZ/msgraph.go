// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Team undocumented
type Team struct {
	// Entity is the base model of Team
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Classification undocumented
	Classification *string `json:"classification,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DiscoverySettings undocumented
	DiscoverySettings *TeamDiscoverySettings `json:"discoverySettings,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// FunSettings undocumented
	FunSettings *TeamFunSettings `json:"funSettings,omitempty"`
	// GuestSettings undocumented
	GuestSettings *TeamGuestSettings `json:"guestSettings,omitempty"`
	// InternalID undocumented
	InternalID *string `json:"internalId,omitempty"`
	// IsArchived undocumented
	IsArchived *bool `json:"isArchived,omitempty"`
	// IsMembershipLimitedToOwners undocumented
	IsMembershipLimitedToOwners *bool `json:"isMembershipLimitedToOwners,omitempty"`
	// MemberSettings undocumented
	MemberSettings *TeamMemberSettings `json:"memberSettings,omitempty"`
	// MessagingSettings undocumented
	MessagingSettings *TeamMessagingSettings `json:"messagingSettings,omitempty"`
	// Specialization undocumented
	Specialization *TeamSpecialization `json:"specialization,omitempty"`
	// Summary undocumented
	Summary *TeamSummary `json:"summary,omitempty"`
	// TenantID undocumented
	TenantID *string `json:"tenantId,omitempty"`
	// Visibility undocumented
	Visibility *TeamVisibilityType `json:"visibility,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
	// AllChannels undocumented
	AllChannels []Channel `json:"allChannels,omitempty"`
	// Channels undocumented
	Channels []Channel `json:"channels,omitempty"`
	// Group undocumented
	Group *Group `json:"group,omitempty"`
	// IncomingChannels undocumented
	IncomingChannels []Channel `json:"incomingChannels,omitempty"`
	// InstalledApps undocumented
	InstalledApps []TeamsAppInstallation `json:"installedApps,omitempty"`
	// Members undocumented
	Members []ConversationMember `json:"members,omitempty"`
	// Operations undocumented
	Operations []TeamsAsyncOperation `json:"operations,omitempty"`
	// Owners undocumented
	Owners []User `json:"owners,omitempty"`
	// PermissionGrants undocumented
	PermissionGrants []ResourceSpecificPermissionGrant `json:"permissionGrants,omitempty"`
	// Photo undocumented
	Photo *ProfilePhoto `json:"photo,omitempty"`
	// PrimaryChannel undocumented
	PrimaryChannel *Channel `json:"primaryChannel,omitempty"`
	// Tags undocumented
	Tags []TeamworkTag `json:"tags,omitempty"`
	// Template undocumented
	Template *TeamsTemplate `json:"template,omitempty"`
	// TemplateDefinition undocumented
	TemplateDefinition *TeamTemplateDefinition `json:"templateDefinition,omitempty"`
	// Schedule undocumented
	Schedule *Schedule `json:"schedule,omitempty"`
}

func NewTeam() (*Team, error) {
	newTeam := &Team{
		ODataType: "#microsoft.graph.Team",
	}
	return newTeam, nil
}

// TeamArchivedEventMessageDetail undocumented
type TeamArchivedEventMessageDetail struct {
	// EventMessageDetail is the base model of TeamArchivedEventMessageDetail
	EventMessageDetail

	ODataType string `json:"@odata.type,omitempty"`
	// Initiator undocumented
	Initiator *IdentitySet `json:"initiator,omitempty"`
	// TeamID undocumented
	TeamID *string `json:"teamId,omitempty"`
}

func NewTeamArchivedEventMessageDetail() (*TeamArchivedEventMessageDetail, error) {
	newTeamArchivedEventMessageDetail := &TeamArchivedEventMessageDetail{
		ODataType: "#microsoft.graph.TeamArchivedEventMessageDetail",
	}
	return newTeamArchivedEventMessageDetail, nil
}

// TeamClassSettings undocumented
type TeamClassSettings struct {
	// Object is the base model of TeamClassSettings
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// NotifyGuardiansAboutAssignments undocumented
	NotifyGuardiansAboutAssignments *bool `json:"notifyGuardiansAboutAssignments,omitempty"`
}

func NewTeamClassSettings() (*TeamClassSettings, error) {
	newTeamClassSettings := &TeamClassSettings{
		ODataType: "#microsoft.graph.TeamClassSettings",
	}
	return newTeamClassSettings, nil
}

// TeamCreatedEventMessageDetail undocumented
type TeamCreatedEventMessageDetail struct {
	// EventMessageDetail is the base model of TeamCreatedEventMessageDetail
	EventMessageDetail

	ODataType string `json:"@odata.type,omitempty"`
	// Initiator undocumented
	Initiator *IdentitySet `json:"initiator,omitempty"`
	// TeamDescription undocumented
	TeamDescription *string `json:"teamDescription,omitempty"`
	// TeamDisplayName undocumented
	TeamDisplayName *string `json:"teamDisplayName,omitempty"`
	// TeamID undocumented
	TeamID *string `json:"teamId,omitempty"`
}

func NewTeamCreatedEventMessageDetail() (*TeamCreatedEventMessageDetail, error) {
	newTeamCreatedEventMessageDetail := &TeamCreatedEventMessageDetail{
		ODataType: "#microsoft.graph.TeamCreatedEventMessageDetail",
	}
	return newTeamCreatedEventMessageDetail, nil
}

// TeamDescriptionUpdatedEventMessageDetail undocumented
type TeamDescriptionUpdatedEventMessageDetail struct {
	// EventMessageDetail is the base model of TeamDescriptionUpdatedEventMessageDetail
	EventMessageDetail

	ODataType string `json:"@odata.type,omitempty"`
	// Initiator undocumented
	Initiator *IdentitySet `json:"initiator,omitempty"`
	// TeamDescription undocumented
	TeamDescription *string `json:"teamDescription,omitempty"`
	// TeamID undocumented
	TeamID *string `json:"teamId,omitempty"`
}

func NewTeamDescriptionUpdatedEventMessageDetail() (*TeamDescriptionUpdatedEventMessageDetail, error) {
	newTeamDescriptionUpdatedEventMessageDetail := &TeamDescriptionUpdatedEventMessageDetail{
		ODataType: "#microsoft.graph.TeamDescriptionUpdatedEventMessageDetail",
	}
	return newTeamDescriptionUpdatedEventMessageDetail, nil
}

// TeamDiscoverySettings undocumented
type TeamDiscoverySettings struct {
	// Object is the base model of TeamDiscoverySettings
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ShowInTeamsSearchAndSuggestions undocumented
	ShowInTeamsSearchAndSuggestions *bool `json:"showInTeamsSearchAndSuggestions,omitempty"`
}

func NewTeamDiscoverySettings() (*TeamDiscoverySettings, error) {
	newTeamDiscoverySettings := &TeamDiscoverySettings{
		ODataType: "#microsoft.graph.TeamDiscoverySettings",
	}
	return newTeamDiscoverySettings, nil
}

// TeamFunSettings undocumented
type TeamFunSettings struct {
	// Object is the base model of TeamFunSettings
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AllowCustomMemes undocumented
	AllowCustomMemes *bool `json:"allowCustomMemes,omitempty"`
	// AllowGiphy undocumented
	AllowGiphy *bool `json:"allowGiphy,omitempty"`
	// AllowStickersAndMemes undocumented
	AllowStickersAndMemes *bool `json:"allowStickersAndMemes,omitempty"`
	// GiphyContentRating undocumented
	GiphyContentRating *GiphyRatingType `json:"giphyContentRating,omitempty"`
}

func NewTeamFunSettings() (*TeamFunSettings, error) {
	newTeamFunSettings := &TeamFunSettings{
		ODataType: "#microsoft.graph.TeamFunSettings",
	}
	return newTeamFunSettings, nil
}

// TeamGuestSettings undocumented
type TeamGuestSettings struct {
	// Object is the base model of TeamGuestSettings
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AllowCreateUpdateChannels undocumented
	AllowCreateUpdateChannels *bool `json:"allowCreateUpdateChannels,omitempty"`
	// AllowDeleteChannels undocumented
	AllowDeleteChannels *bool `json:"allowDeleteChannels,omitempty"`
}

func NewTeamGuestSettings() (*TeamGuestSettings, error) {
	newTeamGuestSettings := &TeamGuestSettings{
		ODataType: "#microsoft.graph.TeamGuestSettings",
	}
	return newTeamGuestSettings, nil
}

// TeamInfo undocumented
type TeamInfo struct {
	// Entity is the base model of TeamInfo
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// TenantID undocumented
	TenantID *string `json:"tenantId,omitempty"`
	// Team undocumented
	Team *Team `json:"team,omitempty"`
}

func NewTeamInfo() (*TeamInfo, error) {
	newTeamInfo := &TeamInfo{
		ODataType: "#microsoft.graph.TeamInfo",
	}
	return newTeamInfo, nil
}

// TeamJoiningDisabledEventMessageDetail undocumented
type TeamJoiningDisabledEventMessageDetail struct {
	// EventMessageDetail is the base model of TeamJoiningDisabledEventMessageDetail
	EventMessageDetail

	ODataType string `json:"@odata.type,omitempty"`
	// Initiator undocumented
	Initiator *IdentitySet `json:"initiator,omitempty"`
	// TeamID undocumented
	TeamID *string `json:"teamId,omitempty"`
}

func NewTeamJoiningDisabledEventMessageDetail() (*TeamJoiningDisabledEventMessageDetail, error) {
	newTeamJoiningDisabledEventMessageDetail := &TeamJoiningDisabledEventMessageDetail{
		ODataType: "#microsoft.graph.TeamJoiningDisabledEventMessageDetail",
	}
	return newTeamJoiningDisabledEventMessageDetail, nil
}

// TeamJoiningEnabledEventMessageDetail undocumented
type TeamJoiningEnabledEventMessageDetail struct {
	// EventMessageDetail is the base model of TeamJoiningEnabledEventMessageDetail
	EventMessageDetail

	ODataType string `json:"@odata.type,omitempty"`
	// Initiator undocumented
	Initiator *IdentitySet `json:"initiator,omitempty"`
	// TeamID undocumented
	TeamID *string `json:"teamId,omitempty"`
}

func NewTeamJoiningEnabledEventMessageDetail() (*TeamJoiningEnabledEventMessageDetail, error) {
	newTeamJoiningEnabledEventMessageDetail := &TeamJoiningEnabledEventMessageDetail{
		ODataType: "#microsoft.graph.TeamJoiningEnabledEventMessageDetail",
	}
	return newTeamJoiningEnabledEventMessageDetail, nil
}

// TeamMemberSettings undocumented
type TeamMemberSettings struct {
	// Object is the base model of TeamMemberSettings
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AllowAddRemoveApps undocumented
	AllowAddRemoveApps *bool `json:"allowAddRemoveApps,omitempty"`
	// AllowCreatePrivateChannels undocumented
	AllowCreatePrivateChannels *bool `json:"allowCreatePrivateChannels,omitempty"`
	// AllowCreateUpdateChannels undocumented
	AllowCreateUpdateChannels *bool `json:"allowCreateUpdateChannels,omitempty"`
	// AllowCreateUpdateRemoveConnectors undocumented
	AllowCreateUpdateRemoveConnectors *bool `json:"allowCreateUpdateRemoveConnectors,omitempty"`
	// AllowCreateUpdateRemoveTabs undocumented
	AllowCreateUpdateRemoveTabs *bool `json:"allowCreateUpdateRemoveTabs,omitempty"`
	// AllowDeleteChannels undocumented
	AllowDeleteChannels *bool `json:"allowDeleteChannels,omitempty"`
}

func NewTeamMemberSettings() (*TeamMemberSettings, error) {
	newTeamMemberSettings := &TeamMemberSettings{
		ODataType: "#microsoft.graph.TeamMemberSettings",
	}
	return newTeamMemberSettings, nil
}

// TeamMembersNotificationRecipient undocumented
type TeamMembersNotificationRecipient struct {
	// TeamworkNotificationRecipient is the base model of TeamMembersNotificationRecipient
	TeamworkNotificationRecipient

	ODataType string `json:"@odata.type,omitempty"`
	// TeamID undocumented
	TeamID *string `json:"teamId,omitempty"`
}

func NewTeamMembersNotificationRecipient() (*TeamMembersNotificationRecipient, error) {
	newTeamMembersNotificationRecipient := &TeamMembersNotificationRecipient{
		ODataType: "#microsoft.graph.TeamMembersNotificationRecipient",
	}
	return newTeamMembersNotificationRecipient, nil
}

// TeamMessagingSettings undocumented
type TeamMessagingSettings struct {
	// Object is the base model of TeamMessagingSettings
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AllowChannelMentions undocumented
	AllowChannelMentions *bool `json:"allowChannelMentions,omitempty"`
	// AllowOwnerDeleteMessages undocumented
	AllowOwnerDeleteMessages *bool `json:"allowOwnerDeleteMessages,omitempty"`
	// AllowTeamMentions undocumented
	AllowTeamMentions *bool `json:"allowTeamMentions,omitempty"`
	// AllowUserDeleteMessages undocumented
	AllowUserDeleteMessages *bool `json:"allowUserDeleteMessages,omitempty"`
	// AllowUserEditMessages undocumented
	AllowUserEditMessages *bool `json:"allowUserEditMessages,omitempty"`
}

func NewTeamMessagingSettings() (*TeamMessagingSettings, error) {
	newTeamMessagingSettings := &TeamMessagingSettings{
		ODataType: "#microsoft.graph.TeamMessagingSettings",
	}
	return newTeamMessagingSettings, nil
}

// TeamRenamedEventMessageDetail undocumented
type TeamRenamedEventMessageDetail struct {
	// EventMessageDetail is the base model of TeamRenamedEventMessageDetail
	EventMessageDetail

	ODataType string `json:"@odata.type,omitempty"`
	// Initiator undocumented
	Initiator *IdentitySet `json:"initiator,omitempty"`
	// TeamDisplayName undocumented
	TeamDisplayName *string `json:"teamDisplayName,omitempty"`
	// TeamID undocumented
	TeamID *string `json:"teamId,omitempty"`
}

func NewTeamRenamedEventMessageDetail() (*TeamRenamedEventMessageDetail, error) {
	newTeamRenamedEventMessageDetail := &TeamRenamedEventMessageDetail{
		ODataType: "#microsoft.graph.TeamRenamedEventMessageDetail",
	}
	return newTeamRenamedEventMessageDetail, nil
}

// TeamSummary undocumented
type TeamSummary struct {
	// Object is the base model of TeamSummary
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// GuestsCount undocumented
	GuestsCount *int `json:"guestsCount,omitempty"`
	// MembersCount undocumented
	MembersCount *int `json:"membersCount,omitempty"`
	// OwnersCount undocumented
	OwnersCount *int `json:"ownersCount,omitempty"`
}

func NewTeamSummary() (*TeamSummary, error) {
	newTeamSummary := &TeamSummary{
		ODataType: "#microsoft.graph.TeamSummary",
	}
	return newTeamSummary, nil
}

// TeamTemplate undocumented
type TeamTemplate struct {
	// Entity is the base model of TeamTemplate
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Definitions undocumented
	Definitions []TeamTemplateDefinition `json:"definitions,omitempty"`
}

func NewTeamTemplate() (*TeamTemplate, error) {
	newTeamTemplate := &TeamTemplate{
		ODataType: "#microsoft.graph.TeamTemplate",
	}
	return newTeamTemplate, nil
}

// TeamTemplateDefinition undocumented
type TeamTemplateDefinition struct {
	// Entity is the base model of TeamTemplateDefinition
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Audience undocumented
	Audience *TeamTemplateAudience `json:"audience,omitempty"`
	// Categories undocumented
	Categories []string `json:"categories,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IconURL undocumented
	IconURL *string `json:"iconUrl,omitempty"`
	// LanguageTag undocumented
	LanguageTag *string `json:"languageTag,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// ParentTemplateID undocumented
	ParentTemplateID *string `json:"parentTemplateId,omitempty"`
	// PublisherName undocumented
	PublisherName *string `json:"publisherName,omitempty"`
	// ShortDescription undocumented
	ShortDescription *string `json:"shortDescription,omitempty"`
	// TeamDefinition undocumented
	TeamDefinition *Team `json:"teamDefinition,omitempty"`
}

func NewTeamTemplateDefinition() (*TeamTemplateDefinition, error) {
	newTeamTemplateDefinition := &TeamTemplateDefinition{
		ODataType: "#microsoft.graph.TeamTemplateDefinition",
	}
	return newTeamTemplateDefinition, nil
}

// TeamUnarchivedEventMessageDetail undocumented
type TeamUnarchivedEventMessageDetail struct {
	// EventMessageDetail is the base model of TeamUnarchivedEventMessageDetail
	EventMessageDetail

	ODataType string `json:"@odata.type,omitempty"`
	// Initiator undocumented
	Initiator *IdentitySet `json:"initiator,omitempty"`
	// TeamID undocumented
	TeamID *string `json:"teamId,omitempty"`
}

func NewTeamUnarchivedEventMessageDetail() (*TeamUnarchivedEventMessageDetail, error) {
	newTeamUnarchivedEventMessageDetail := &TeamUnarchivedEventMessageDetail{
		ODataType: "#microsoft.graph.TeamUnarchivedEventMessageDetail",
	}
	return newTeamUnarchivedEventMessageDetail, nil
}
