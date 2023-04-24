// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// OrgContact undocumented
type OrgContact struct {
	// DirectoryObject is the base model of OrgContact
	DirectoryObject

	ODataType string `json:"@odata.type,omitempty"`
	// Addresses undocumented
	Addresses []PhysicalOfficeAddress `json:"addresses,omitempty"`
	// CompanyName undocumented
	CompanyName *string `json:"companyName,omitempty"`
	// Department undocumented
	Department *string `json:"department,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// GivenName undocumented
	GivenName *string `json:"givenName,omitempty"`
	// JobTitle undocumented
	JobTitle *string `json:"jobTitle,omitempty"`
	// Mail undocumented
	Mail *string `json:"mail,omitempty"`
	// MailNickname undocumented
	MailNickname *string `json:"mailNickname,omitempty"`
	// OnPremisesLastSyncDateTime undocumented
	OnPremisesLastSyncDateTime *time.Time `json:"onPremisesLastSyncDateTime,omitempty"`
	// OnPremisesProvisioningErrors undocumented
	OnPremisesProvisioningErrors []OnPremisesProvisioningError `json:"onPremisesProvisioningErrors,omitempty"`
	// OnPremisesSyncEnabled undocumented
	OnPremisesSyncEnabled *bool `json:"onPremisesSyncEnabled,omitempty"`
	// Phones undocumented
	Phones []Phone `json:"phones,omitempty"`
	// ProxyAddresses undocumented
	ProxyAddresses []string `json:"proxyAddresses,omitempty"`
	// ServiceProvisioningErrors undocumented
	ServiceProvisioningErrors []ServiceProvisioningError `json:"serviceProvisioningErrors,omitempty"`
	// Surname undocumented
	Surname *string `json:"surname,omitempty"`
	// DirectReports undocumented
	DirectReports []DirectoryObject `json:"directReports,omitempty"`
	// Manager undocumented
	Manager *DirectoryObject `json:"manager,omitempty"`
	// MemberOf undocumented
	MemberOf []DirectoryObject `json:"memberOf,omitempty"`
	// TransitiveMemberOf undocumented
	TransitiveMemberOf []DirectoryObject `json:"transitiveMemberOf,omitempty"`
	// TransitiveReports undocumented
	TransitiveReports []DirectoryObject `json:"transitiveReports,omitempty"`
}

func NewOrgContact() (*OrgContact, error) {
	newOrgContact := &OrgContact{
		ODataType: "#microsoft.graph.OrgContact",
	}
	return newOrgContact, nil
}