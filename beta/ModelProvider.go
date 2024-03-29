// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// ProviderTenantSetting undocumented
type ProviderTenantSetting struct {
	// Entity is the base model of ProviderTenantSetting
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AzureTenantID undocumented
	AzureTenantID *string `json:"azureTenantId,omitempty"`
	// Enabled undocumented
	Enabled *bool `json:"enabled,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Provider undocumented
	Provider *string `json:"provider,omitempty"`
	// Vendor undocumented
	Vendor *string `json:"vendor,omitempty"`
}

func NewProviderTenantSetting() (*ProviderTenantSetting, error) {
	newProviderTenantSetting := &ProviderTenantSetting{
		ODataType: "#microsoft.graph.ProviderTenantSetting",
	}
	return newProviderTenantSetting, nil
}
