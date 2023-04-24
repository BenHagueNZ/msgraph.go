// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// VPPLicensingType undocumented
type VPPLicensingType struct {
	// Object is the base model of VPPLicensingType
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// SupportDeviceLicensing undocumented
	SupportDeviceLicensing *bool `json:"supportDeviceLicensing,omitempty"`
	// SupportsDeviceLicensing undocumented
	SupportsDeviceLicensing *bool `json:"supportsDeviceLicensing,omitempty"`
	// SupportsUserLicensing undocumented
	SupportsUserLicensing *bool `json:"supportsUserLicensing,omitempty"`
	// SupportUserLicensing undocumented
	SupportUserLicensing *bool `json:"supportUserLicensing,omitempty"`
}

func NewVPPLicensingType() (*VPPLicensingType, error) {
	newVPPLicensingType := &VPPLicensingType{
		ODataType: "#microsoft.graph.VppLicensingType",
	}
	return newVPPLicensingType, nil
}

// VPPToken undocumented
type VPPToken struct {
	// Entity is the base model of VPPToken
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AppleID undocumented
	AppleID *string `json:"appleId,omitempty"`
	// AutomaticallyUpdateApps undocumented
	AutomaticallyUpdateApps *bool `json:"automaticallyUpdateApps,omitempty"`
	// ClaimTokenManagementFromExternalMDM undocumented
	ClaimTokenManagementFromExternalMDM *bool `json:"claimTokenManagementFromExternalMdm,omitempty"`
	// CountryOrRegion undocumented
	CountryOrRegion *string `json:"countryOrRegion,omitempty"`
	// DataSharingConsentGranted undocumented
	DataSharingConsentGranted *bool `json:"dataSharingConsentGranted,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ExpirationDateTime undocumented
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// LastSyncDateTime undocumented
	LastSyncDateTime *time.Time `json:"lastSyncDateTime,omitempty"`
	// LastSyncStatus undocumented
	LastSyncStatus *VPPTokenSyncStatus `json:"lastSyncStatus,omitempty"`
	// LocationName undocumented
	LocationName *string `json:"locationName,omitempty"`
	// OrganizationName undocumented
	OrganizationName *string `json:"organizationName,omitempty"`
	// RoleScopeTagIDs undocumented
	RoleScopeTagIDs []string `json:"roleScopeTagIds,omitempty"`
	// State undocumented
	State *VPPTokenState `json:"state,omitempty"`
	// Token undocumented
	Token *string `json:"token,omitempty"`
	// TokenActionResults undocumented
	TokenActionResults []VPPTokenActionResult `json:"tokenActionResults,omitempty"`
	// VPPTokenAccountType undocumented
	VPPTokenAccountType *VPPTokenAccountType `json:"vppTokenAccountType,omitempty"`
}

func NewVPPToken() (*VPPToken, error) {
	newVPPToken := &VPPToken{
		ODataType: "#microsoft.graph.VppToken",
	}
	return newVPPToken, nil
}

// VPPTokenActionResult undocumented
type VPPTokenActionResult struct {
	// Object is the base model of VPPTokenActionResult
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ActionName undocumented
	ActionName *string `json:"actionName,omitempty"`
	// ActionState undocumented
	ActionState *ActionState `json:"actionState,omitempty"`
	// LastUpdatedDateTime undocumented
	LastUpdatedDateTime *time.Time `json:"lastUpdatedDateTime,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
}

func NewVPPTokenActionResult() (*VPPTokenActionResult, error) {
	newVPPTokenActionResult := &VPPTokenActionResult{
		ODataType: "#microsoft.graph.VppTokenActionResult",
	}
	return newVPPTokenActionResult, nil
}

// VPPTokenLicenseSummary undocumented
type VPPTokenLicenseSummary struct {
	// Object is the base model of VPPTokenLicenseSummary
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AppleID undocumented
	AppleID *string `json:"appleId,omitempty"`
	// AvailableLicenseCount undocumented
	AvailableLicenseCount *int `json:"availableLicenseCount,omitempty"`
	// OrganizationName undocumented
	OrganizationName *string `json:"organizationName,omitempty"`
	// UsedLicenseCount undocumented
	UsedLicenseCount *int `json:"usedLicenseCount,omitempty"`
	// VPPTokenID undocumented
	VPPTokenID *string `json:"vppTokenId,omitempty"`
}

func NewVPPTokenLicenseSummary() (*VPPTokenLicenseSummary, error) {
	newVPPTokenLicenseSummary := &VPPTokenLicenseSummary{
		ODataType: "#microsoft.graph.VppTokenLicenseSummary",
	}
	return newVPPTokenLicenseSummary, nil
}

// VPPTokenRevokeLicensesActionResult undocumented
type VPPTokenRevokeLicensesActionResult struct {
	// VPPTokenActionResult is the base model of VPPTokenRevokeLicensesActionResult
	VPPTokenActionResult

	ODataType string `json:"@odata.type,omitempty"`
	// ActionFailureReason undocumented
	ActionFailureReason *VPPTokenActionFailureReason `json:"actionFailureReason,omitempty"`
	// FailedLicensesCount undocumented
	FailedLicensesCount *int `json:"failedLicensesCount,omitempty"`
	// TotalLicensesCount undocumented
	TotalLicensesCount *int `json:"totalLicensesCount,omitempty"`
}

func NewVPPTokenRevokeLicensesActionResult() (*VPPTokenRevokeLicensesActionResult, error) {
	newVPPTokenRevokeLicensesActionResult := &VPPTokenRevokeLicensesActionResult{
		ODataType: "#microsoft.graph.VppTokenRevokeLicensesActionResult",
	}
	return newVPPTokenRevokeLicensesActionResult, nil
}
