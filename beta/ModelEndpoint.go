// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Endpoint undocumented
type Endpoint struct {
	// DirectoryObject is the base model of Endpoint
	DirectoryObject

	ODataType string `json:"@odata.type,omitempty"`
	// Capability undocumented
	Capability *string `json:"capability,omitempty"`
	// ProviderID undocumented
	ProviderID *string `json:"providerId,omitempty"`
	// ProviderName undocumented
	ProviderName *string `json:"providerName,omitempty"`
	// ProviderResourceID undocumented
	ProviderResourceID *string `json:"providerResourceId,omitempty"`
	// URI undocumented
	URI *string `json:"uri,omitempty"`
}

func NewEndpoint() (*Endpoint, error) {
	newEndpoint := &Endpoint{
		ODataType: "#microsoft.graph.Endpoint",
	}
	return newEndpoint, nil
}
