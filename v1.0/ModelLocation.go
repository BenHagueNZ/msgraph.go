// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Location undocumented
type Location struct {
	// Object is the base model of Location
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Address undocumented
	Address *PhysicalAddress `json:"address,omitempty"`
	// Coordinates undocumented
	Coordinates *OutlookGeoCoordinates `json:"coordinates,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LocationEmailAddress undocumented
	LocationEmailAddress *string `json:"locationEmailAddress,omitempty"`
	// LocationType undocumented
	LocationType *LocationType `json:"locationType,omitempty"`
	// LocationURI undocumented
	LocationURI *string `json:"locationUri,omitempty"`
	// UniqueID undocumented
	UniqueID *string `json:"uniqueId,omitempty"`
	// UniqueIDType undocumented
	UniqueIDType *LocationUniqueIDType `json:"uniqueIdType,omitempty"`
}

func NewLocation() (*Location, error) {
	newLocation := &Location{
		ODataType: "#microsoft.graph.Location",
	}
	return newLocation, nil
}

// LocationConstraint undocumented
type LocationConstraint struct {
	// Object is the base model of LocationConstraint
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// IsRequired undocumented
	IsRequired *bool `json:"isRequired,omitempty"`
	// Locations undocumented
	Locations []LocationConstraintItem `json:"locations,omitempty"`
	// SuggestLocation undocumented
	SuggestLocation *bool `json:"suggestLocation,omitempty"`
}

func NewLocationConstraint() (*LocationConstraint, error) {
	newLocationConstraint := &LocationConstraint{
		ODataType: "#microsoft.graph.LocationConstraint",
	}
	return newLocationConstraint, nil
}

// LocationConstraintItem undocumented
type LocationConstraintItem struct {
	// Location is the base model of LocationConstraintItem
	Location

	ODataType string `json:"@odata.type,omitempty"`
	// ResolveAvailability undocumented
	ResolveAvailability *bool `json:"resolveAvailability,omitempty"`
}

func NewLocationConstraintItem() (*LocationConstraintItem, error) {
	newLocationConstraintItem := &LocationConstraintItem{
		ODataType: "#microsoft.graph.LocationConstraintItem",
	}
	return newLocationConstraintItem, nil
}
