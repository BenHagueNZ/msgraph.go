// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ShareAction undocumented
type ShareAction struct {
	// Object is the base model of ShareAction
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Recipients undocumented
	Recipients []IdentitySet `json:"recipients,omitempty"`
}

func NewShareAction() (*ShareAction, error) {
	newShareAction := &ShareAction{
		ODataType: "#microsoft.graph.ShareAction",
	}
	return newShareAction, nil
}

// SharePointIdentity undocumented
type SharePointIdentity struct {
	// Identity is the base model of SharePointIdentity
	Identity

	ODataType string `json:"@odata.type,omitempty"`
	// LoginName undocumented
	LoginName *string `json:"loginName,omitempty"`
}

func NewSharePointIdentity() (*SharePointIdentity, error) {
	newSharePointIdentity := &SharePointIdentity{
		ODataType: "#microsoft.graph.SharePointIdentity",
	}
	return newSharePointIdentity, nil
}

// SharePointIdentitySet undocumented
type SharePointIdentitySet struct {
	// IdentitySet is the base model of SharePointIdentitySet
	IdentitySet

	ODataType string `json:"@odata.type,omitempty"`
	// Group undocumented
	Group *Identity `json:"group,omitempty"`
	// SiteGroup undocumented
	SiteGroup *SharePointIdentity `json:"siteGroup,omitempty"`
	// SiteUser undocumented
	SiteUser *SharePointIdentity `json:"siteUser,omitempty"`
}

func NewSharePointIdentitySet() (*SharePointIdentitySet, error) {
	newSharePointIdentitySet := &SharePointIdentitySet{
		ODataType: "#microsoft.graph.SharePointIdentitySet",
	}
	return newSharePointIdentitySet, nil
}

// SharePointOneDriveOptions undocumented
type SharePointOneDriveOptions struct {
	// Object is the base model of SharePointOneDriveOptions
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// IncludeContent undocumented
	IncludeContent *SearchContent `json:"includeContent,omitempty"`
}

func NewSharePointOneDriveOptions() (*SharePointOneDriveOptions, error) {
	newSharePointOneDriveOptions := &SharePointOneDriveOptions{
		ODataType: "#microsoft.graph.SharePointOneDriveOptions",
	}
	return newSharePointOneDriveOptions, nil
}
