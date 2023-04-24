// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// RemoteActionAudit undocumented
type RemoteActionAudit struct {
	// Entity is the base model of RemoteActionAudit
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Action undocumented
	Action *RemoteAction `json:"action,omitempty"`
	// ActionState undocumented
	ActionState *ActionState `json:"actionState,omitempty"`
	// DeviceDisplayName undocumented
	DeviceDisplayName *string `json:"deviceDisplayName,omitempty"`
	// DeviceIMEI undocumented
	DeviceIMEI *string `json:"deviceIMEI,omitempty"`
	// DeviceOwnerUserPrincipalName undocumented
	DeviceOwnerUserPrincipalName *string `json:"deviceOwnerUserPrincipalName,omitempty"`
	// InitiatedByUserPrincipalName undocumented
	InitiatedByUserPrincipalName *string `json:"initiatedByUserPrincipalName,omitempty"`
	// ManagedDeviceID undocumented
	ManagedDeviceID *string `json:"managedDeviceId,omitempty"`
	// RequestDateTime undocumented
	RequestDateTime *time.Time `json:"requestDateTime,omitempty"`
	// UserName undocumented
	UserName *string `json:"userName,omitempty"`
}

func NewRemoteActionAudit() (*RemoteActionAudit, error) {
	newRemoteActionAudit := &RemoteActionAudit{
		ODataType: "#microsoft.graph.RemoteActionAudit",
	}
	return newRemoteActionAudit, nil
}

// RemoteAssistancePartner undocumented
type RemoteAssistancePartner struct {
	// Entity is the base model of RemoteAssistancePartner
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LastConnectionDateTime undocumented
	LastConnectionDateTime *time.Time `json:"lastConnectionDateTime,omitempty"`
	// OnboardingRequestExpiryDateTime undocumented
	OnboardingRequestExpiryDateTime *time.Time `json:"onboardingRequestExpiryDateTime,omitempty"`
	// OnboardingStatus undocumented
	OnboardingStatus *RemoteAssistanceOnboardingStatus `json:"onboardingStatus,omitempty"`
	// OnboardingURL undocumented
	OnboardingURL *string `json:"onboardingUrl,omitempty"`
}

func NewRemoteAssistancePartner() (*RemoteAssistancePartner, error) {
	newRemoteAssistancePartner := &RemoteAssistancePartner{
		ODataType: "#microsoft.graph.RemoteAssistancePartner",
	}
	return newRemoteAssistancePartner, nil
}

// RemoteAssistanceSettings undocumented
type RemoteAssistanceSettings struct {
	// Entity is the base model of RemoteAssistanceSettings
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AllowSessionsToUnenrolledDevices undocumented
	AllowSessionsToUnenrolledDevices *bool `json:"allowSessionsToUnenrolledDevices,omitempty"`
	// BlockChat undocumented
	BlockChat *bool `json:"blockChat,omitempty"`
	// RemoteAssistanceState undocumented
	RemoteAssistanceState *RemoteAssistanceState `json:"remoteAssistanceState,omitempty"`
}

func NewRemoteAssistanceSettings() (*RemoteAssistanceSettings, error) {
	newRemoteAssistanceSettings := &RemoteAssistanceSettings{
		ODataType: "#microsoft.graph.RemoteAssistanceSettings",
	}
	return newRemoteAssistanceSettings, nil
}

// RemoteItem undocumented
type RemoteItem struct {
	// Object is the base model of RemoteItem
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// File undocumented
	File *File `json:"file,omitempty"`
	// FileSystemInfo undocumented
	FileSystemInfo *FileSystemInfo `json:"fileSystemInfo,omitempty"`
	// Folder undocumented
	Folder *Folder `json:"folder,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// Image undocumented
	Image *Image `json:"image,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Package undocumented
	Package *Package `json:"package,omitempty"`
	// ParentReference undocumented
	ParentReference *ItemReference `json:"parentReference,omitempty"`
	// Shared undocumented
	Shared *Shared `json:"shared,omitempty"`
	// SharepointIDs undocumented
	SharepointIDs *SharepointIDs `json:"sharepointIds,omitempty"`
	// Size undocumented
	Size *int `json:"size,omitempty"`
	// SpecialFolder undocumented
	SpecialFolder *SpecialFolder `json:"specialFolder,omitempty"`
	// Video undocumented
	Video *Video `json:"video,omitempty"`
	// WebDavURL undocumented
	WebDavURL *string `json:"webDavUrl,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
}

func NewRemoteItem() (*RemoteItem, error) {
	newRemoteItem := &RemoteItem{
		ODataType: "#microsoft.graph.RemoteItem",
	}
	return newRemoteItem, nil
}

// RemoteLockActionResult undocumented
type RemoteLockActionResult struct {
	// DeviceActionResult is the base model of RemoteLockActionResult
	DeviceActionResult

	ODataType string `json:"@odata.type,omitempty"`
	// UnlockPin undocumented
	UnlockPin *string `json:"unlockPin,omitempty"`
}

func NewRemoteLockActionResult() (*RemoteLockActionResult, error) {
	newRemoteLockActionResult := &RemoteLockActionResult{
		ODataType: "#microsoft.graph.RemoteLockActionResult",
	}
	return newRemoteLockActionResult, nil
}
