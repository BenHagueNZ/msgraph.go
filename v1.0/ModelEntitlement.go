// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// EntitlementManagement undocumented
type EntitlementManagement struct {
	// Entity is the base model of EntitlementManagement
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AccessPackageAssignmentApprovals undocumented
	AccessPackageAssignmentApprovals []Approval `json:"accessPackageAssignmentApprovals,omitempty"`
	// AccessPackages undocumented
	AccessPackages []AccessPackage `json:"accessPackages,omitempty"`
	// AssignmentPolicies undocumented
	AssignmentPolicies []AccessPackageAssignmentPolicy `json:"assignmentPolicies,omitempty"`
	// AssignmentRequests undocumented
	AssignmentRequests []AccessPackageAssignmentRequestObject `json:"assignmentRequests,omitempty"`
	// Assignments undocumented
	Assignments []AccessPackageAssignment `json:"assignments,omitempty"`
	// Catalogs undocumented
	Catalogs []AccessPackageCatalog `json:"catalogs,omitempty"`
	// ConnectedOrganizations undocumented
	ConnectedOrganizations []ConnectedOrganization `json:"connectedOrganizations,omitempty"`
	// Settings undocumented
	Settings *EntitlementManagementSettings `json:"settings,omitempty"`
}

func NewEntitlementManagement() (*EntitlementManagement, error) {
	newEntitlementManagement := &EntitlementManagement{
		ODataType: "#microsoft.graph.EntitlementManagement",
	}
	return newEntitlementManagement, nil
}

// EntitlementManagementSchedule undocumented
type EntitlementManagementSchedule struct {
	// Object is the base model of EntitlementManagementSchedule
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Expiration undocumented
	Expiration *ExpirationPattern `json:"expiration,omitempty"`
	// Recurrence undocumented
	Recurrence *PatternedRecurrence `json:"recurrence,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
}

func NewEntitlementManagementSchedule() (*EntitlementManagementSchedule, error) {
	newEntitlementManagementSchedule := &EntitlementManagementSchedule{
		ODataType: "#microsoft.graph.EntitlementManagementSchedule",
	}
	return newEntitlementManagementSchedule, nil
}

// EntitlementManagementSettings undocumented
type EntitlementManagementSettings struct {
	// Entity is the base model of EntitlementManagementSettings
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// DurationUntilExternalUserDeletedAfterBlocked undocumented
	DurationUntilExternalUserDeletedAfterBlocked *Duration `json:"durationUntilExternalUserDeletedAfterBlocked,omitempty"`
	// ExternalUserLifecycleAction undocumented
	ExternalUserLifecycleAction *AccessPackageExternalUserLifecycleAction `json:"externalUserLifecycleAction,omitempty"`
}

func NewEntitlementManagementSettings() (*EntitlementManagementSettings, error) {
	newEntitlementManagementSettings := &EntitlementManagementSettings{
		ODataType: "#microsoft.graph.EntitlementManagementSettings",
	}
	return newEntitlementManagementSettings, nil
}
