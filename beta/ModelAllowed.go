// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// AllowedDataLocation undocumented
type AllowedDataLocation struct {
	// Entity is the base model of AllowedDataLocation
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AppID undocumented
	AppID *string `json:"appId,omitempty"`
	// Domain undocumented
	Domain *string `json:"domain,omitempty"`
	// IsDefault undocumented
	IsDefault *bool `json:"isDefault,omitempty"`
	// Location undocumented
	Location *string `json:"location,omitempty"`
}

func NewAllowedDataLocation() (*AllowedDataLocation, error) {
	newAllowedDataLocation := &AllowedDataLocation{
		ODataType: "#microsoft.graph.AllowedDataLocation",
	}
	return newAllowedDataLocation, nil
}

// AllowedValue undocumented
type AllowedValue struct {
	// Entity is the base model of AllowedValue
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// IsActive undocumented
	IsActive *bool `json:"isActive,omitempty"`
}

func NewAllowedValue() (*AllowedValue, error) {
	newAllowedValue := &AllowedValue{
		ODataType: "#microsoft.graph.AllowedValue",
	}
	return newAllowedValue, nil
}
