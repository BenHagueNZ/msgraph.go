// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// RelyingPartyDetailedSummary undocumented
type RelyingPartyDetailedSummary struct {
	// Entity is the base model of RelyingPartyDetailedSummary
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// FailedSignInCount undocumented
	FailedSignInCount *int `json:"failedSignInCount,omitempty"`
	// MigrationStatus undocumented
	MigrationStatus *MigrationStatus `json:"migrationStatus,omitempty"`
	// MigrationValidationDetails undocumented
	MigrationValidationDetails []KeyValuePair `json:"migrationValidationDetails,omitempty"`
	// RelyingPartyID undocumented
	RelyingPartyID *string `json:"relyingPartyId,omitempty"`
	// RelyingPartyName undocumented
	RelyingPartyName *string `json:"relyingPartyName,omitempty"`
	// ReplyUrls undocumented
	ReplyUrls []string `json:"replyUrls,omitempty"`
	// ServiceID undocumented
	ServiceID *string `json:"serviceId,omitempty"`
	// SignInSuccessRate undocumented
	SignInSuccessRate *float64 `json:"signInSuccessRate,omitempty"`
	// SuccessfulSignInCount undocumented
	SuccessfulSignInCount *int `json:"successfulSignInCount,omitempty"`
	// TotalSignInCount undocumented
	TotalSignInCount *int `json:"totalSignInCount,omitempty"`
	// UniqueUserCount undocumented
	UniqueUserCount *int `json:"uniqueUserCount,omitempty"`
}

func NewRelyingPartyDetailedSummary() (*RelyingPartyDetailedSummary, error) {
	newRelyingPartyDetailedSummary := &RelyingPartyDetailedSummary{
		ODataType: "#microsoft.graph.RelyingPartyDetailedSummary",
	}
	return newRelyingPartyDetailedSummary, nil
}
