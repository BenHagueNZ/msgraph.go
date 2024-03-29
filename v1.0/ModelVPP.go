// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// VPPLicensingType undocumented
type VPPLicensingType struct {
	// Object is the base model of VPPLicensingType
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// SupportsDeviceLicensing undocumented
	SupportsDeviceLicensing *bool `json:"supportsDeviceLicensing,omitempty"`
	// SupportsUserLicensing undocumented
	SupportsUserLicensing *bool `json:"supportsUserLicensing,omitempty"`
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
	// CountryOrRegion undocumented
	CountryOrRegion *string `json:"countryOrRegion,omitempty"`
	// ExpirationDateTime undocumented
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// LastSyncDateTime undocumented
	LastSyncDateTime *time.Time `json:"lastSyncDateTime,omitempty"`
	// LastSyncStatus undocumented
	LastSyncStatus *VPPTokenSyncStatus `json:"lastSyncStatus,omitempty"`
	// OrganizationName undocumented
	OrganizationName *string `json:"organizationName,omitempty"`
	// State undocumented
	State *VPPTokenState `json:"state,omitempty"`
	// Token undocumented
	Token *string `json:"token,omitempty"`
	// VPPTokenAccountType undocumented
	VPPTokenAccountType *VPPTokenAccountType `json:"vppTokenAccountType,omitempty"`
}

func NewVPPToken() (*VPPToken, error) {
	newVPPToken := &VPPToken{
		ODataType: "#microsoft.graph.VppToken",
	}
	return newVPPToken, nil
}
