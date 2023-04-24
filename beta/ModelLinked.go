// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// LinkedResource undocumented
type LinkedResource struct {
	// Entity is the base model of LinkedResource
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ApplicationName undocumented
	ApplicationName *string `json:"applicationName,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
}

func NewLinkedResource() (*LinkedResource, error) {
	newLinkedResource := &LinkedResource{
		ODataType: "#microsoft.graph.LinkedResource",
	}
	return newLinkedResource, nil
}
