// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Admin undocumented
type Admin struct {
	// Object is the base model of Admin
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Edge undocumented
	Edge *Edge `json:"edge,omitempty"`
	// Sharepoint undocumented
	Sharepoint *Sharepoint `json:"sharepoint,omitempty"`
	// ServiceAnnouncement undocumented
	ServiceAnnouncement *ServiceAnnouncement `json:"serviceAnnouncement,omitempty"`
	// ReportSettings undocumented
	ReportSettings *AdminReportSettings `json:"reportSettings,omitempty"`
	// Windows undocumented
	Windows *AdminWindows `json:"windows,omitempty"`
}

func NewAdmin() (*Admin, error) {
	newAdmin := &Admin{
		ODataType: "#microsoft.graph.Admin",
	}
	return newAdmin, nil
}

// AdminConsent undocumented
type AdminConsent struct {
	// Object is the base model of AdminConsent
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ShareAPNSData undocumented
	ShareAPNSData *AdminConsentState `json:"shareAPNSData,omitempty"`
	// ShareUserExperienceAnalyticsData undocumented
	ShareUserExperienceAnalyticsData *AdminConsentState `json:"shareUserExperienceAnalyticsData,omitempty"`
}

func NewAdminConsent() (*AdminConsent, error) {
	newAdminConsent := &AdminConsent{
		ODataType: "#microsoft.graph.AdminConsent",
	}
	return newAdminConsent, nil
}

// AdminConsentRequestPolicy undocumented
type AdminConsentRequestPolicy struct {
	// Entity is the base model of AdminConsentRequestPolicy
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// NotifyReviewers undocumented
	NotifyReviewers *bool `json:"notifyReviewers,omitempty"`
	// RemindersEnabled undocumented
	RemindersEnabled *bool `json:"remindersEnabled,omitempty"`
	// RequestDurationInDays undocumented
	RequestDurationInDays *int `json:"requestDurationInDays,omitempty"`
	// Reviewers undocumented
	Reviewers []AccessReviewReviewerScope `json:"reviewers,omitempty"`
	// Version undocumented
	Version *int `json:"version,omitempty"`
}

func NewAdminConsentRequestPolicy() (*AdminConsentRequestPolicy, error) {
	newAdminConsentRequestPolicy := &AdminConsentRequestPolicy{
		ODataType: "#microsoft.graph.AdminConsentRequestPolicy",
	}
	return newAdminConsentRequestPolicy, nil
}

// AdminReportSettings undocumented
type AdminReportSettings struct {
	// Entity is the base model of AdminReportSettings
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// DisplayConcealedNames undocumented
	DisplayConcealedNames *bool `json:"displayConcealedNames,omitempty"`
}

func NewAdminReportSettings() (*AdminReportSettings, error) {
	newAdminReportSettings := &AdminReportSettings{
		ODataType: "#microsoft.graph.AdminReportSettings",
	}
	return newAdminReportSettings, nil
}

// AdminWindows undocumented
type AdminWindows struct {
	// Entity is the base model of AdminWindows
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Updates undocumented
	Updates *AdminWindowsUpdates `json:"updates,omitempty"`
}

func NewAdminWindows() (*AdminWindows, error) {
	newAdminWindows := &AdminWindows{
		ODataType: "#microsoft.graph.AdminWindows",
	}
	return newAdminWindows, nil
}

// AdminWindowsUpdates undocumented
type AdminWindowsUpdates struct {
	// Entity is the base model of AdminWindowsUpdates
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Catalog undocumented
	Catalog *WindowsUpdatesCatalog `json:"catalog,omitempty"`
	// DeploymentAudiences undocumented
	DeploymentAudiences []WindowsUpdatesDeploymentAudience `json:"deploymentAudiences,omitempty"`
	// Deployments undocumented
	Deployments []WindowsUpdatesDeployment `json:"deployments,omitempty"`
	// ResourceConnections undocumented
	ResourceConnections []WindowsUpdatesResourceConnection `json:"resourceConnections,omitempty"`
	// UpdatableAssets undocumented
	UpdatableAssets []WindowsUpdatesUpdatableAsset `json:"updatableAssets,omitempty"`
	// UpdatePolicies undocumented
	UpdatePolicies []WindowsUpdatesUpdatePolicy `json:"updatePolicies,omitempty"`
}

func NewAdminWindowsUpdates() (*AdminWindowsUpdates, error) {
	newAdminWindowsUpdates := &AdminWindowsUpdates{
		ODataType: "#microsoft.graph.AdminWindowsUpdates",
	}
	return newAdminWindowsUpdates, nil
}