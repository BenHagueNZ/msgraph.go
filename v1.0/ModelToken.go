// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// TokenIssuancePolicy undocumented
type TokenIssuancePolicy struct {
	// StsPolicy is the base model of TokenIssuancePolicy
	StsPolicy

	ODataType string `json:"@odata.type,omitempty"`
}

func NewTokenIssuancePolicy() (*TokenIssuancePolicy, error) {
	newTokenIssuancePolicy := &TokenIssuancePolicy{
		ODataType: "#microsoft.graph.TokenIssuancePolicy",
	}
	return newTokenIssuancePolicy, nil
}

// TokenLifetimePolicy undocumented
type TokenLifetimePolicy struct {
	// StsPolicy is the base model of TokenLifetimePolicy
	StsPolicy

	ODataType string `json:"@odata.type,omitempty"`
}

func NewTokenLifetimePolicy() (*TokenLifetimePolicy, error) {
	newTokenLifetimePolicy := &TokenLifetimePolicy{
		ODataType: "#microsoft.graph.TokenLifetimePolicy",
	}
	return newTokenLifetimePolicy, nil
}

// TokenMeetingInfo undocumented
type TokenMeetingInfo struct {
	// MeetingInfo is the base model of TokenMeetingInfo
	MeetingInfo

	ODataType string `json:"@odata.type,omitempty"`
	// Token undocumented
	Token *string `json:"token,omitempty"`
}

func NewTokenMeetingInfo() (*TokenMeetingInfo, error) {
	newTokenMeetingInfo := &TokenMeetingInfo{
		ODataType: "#microsoft.graph.TokenMeetingInfo",
	}
	return newTokenMeetingInfo, nil
}
