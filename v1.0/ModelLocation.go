// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Location undocumented
type Location struct {
	// Object is the base model of Location
	Object
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LocationEmailAddress undocumented
	LocationEmailAddress *string `json:"locationEmailAddress,omitempty"`
	// Address undocumented
	Address *PhysicalAddress `json:"address,omitempty"`
	// LocationURI undocumented
	LocationURI *string `json:"locationUri,omitempty"`
	// Coordinates undocumented
	Coordinates *OutlookGeoCoordinates `json:"coordinates,omitempty"`
	// LocationType undocumented
	LocationType *LocationType `json:"locationType,omitempty"`
	// UniqueID undocumented
	UniqueID *string `json:"uniqueId,omitempty"`
	// UniqueIDType undocumented
	UniqueIDType *LocationUniqueIDType `json:"uniqueIdType,omitempty"`
}

// LocationConstraint undocumented
type LocationConstraint struct {
	// Object is the base model of LocationConstraint
	Object
	// Locations undocumented
	Locations []LocationConstraintItem `json:"locations,omitempty"`
	// IsRequired undocumented
	IsRequired *bool `json:"isRequired,omitempty"`
	// SuggestLocation undocumented
	SuggestLocation *bool `json:"suggestLocation,omitempty"`
}

// LocationConstraintItem undocumented
type LocationConstraintItem struct {
	// Location is the base model of LocationConstraintItem
	Location
	// ResolveAvailability undocumented
	ResolveAvailability *bool `json:"resolveAvailability,omitempty"`
}
