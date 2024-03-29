// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// AuthorizationInfo undocumented
type AuthorizationInfo struct {
	// Object is the base model of AuthorizationInfo
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// CertificateUserIDs undocumented
	CertificateUserIDs []string `json:"certificateUserIds,omitempty"`
}

func NewAuthorizationInfo() (*AuthorizationInfo, error) {
	newAuthorizationInfo := &AuthorizationInfo{
		ODataType: "#microsoft.graph.AuthorizationInfo",
	}
	return newAuthorizationInfo, nil
}

// AuthorizationPolicy undocumented
type AuthorizationPolicy struct {
	// PolicyBase is the base model of AuthorizationPolicy
	PolicyBase

	ODataType string `json:"@odata.type,omitempty"`
	// AllowedToSignUpEmailBasedSubscriptions undocumented
	AllowedToSignUpEmailBasedSubscriptions *bool `json:"allowedToSignUpEmailBasedSubscriptions,omitempty"`
	// AllowedToUseSSPR undocumented
	AllowedToUseSSPR *bool `json:"allowedToUseSSPR,omitempty"`
	// AllowEmailVerifiedUsersToJoinOrganization undocumented
	AllowEmailVerifiedUsersToJoinOrganization *bool `json:"allowEmailVerifiedUsersToJoinOrganization,omitempty"`
	// AllowInvitesFrom undocumented
	AllowInvitesFrom *AllowInvitesFrom `json:"allowInvitesFrom,omitempty"`
	// BlockMsolPowerShell undocumented
	BlockMsolPowerShell *bool `json:"blockMsolPowerShell,omitempty"`
	// DefaultUserRolePermissions undocumented
	DefaultUserRolePermissions *DefaultUserRolePermissions `json:"defaultUserRolePermissions,omitempty"`
	// GuestUserRoleID undocumented
	GuestUserRoleID *UUID `json:"guestUserRoleId,omitempty"`
}

func NewAuthorizationPolicy() (*AuthorizationPolicy, error) {
	newAuthorizationPolicy := &AuthorizationPolicy{
		ODataType: "#microsoft.graph.AuthorizationPolicy",
	}
	return newAuthorizationPolicy, nil
}
