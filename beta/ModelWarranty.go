// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// WarrantyOffer undocumented
type WarrantyOffer struct {
	// Object is the base model of WarrantyOffer
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// Type undocumented
	Type *WarrantyType `json:"type,omitempty"`
}

func NewWarrantyOffer() (*WarrantyOffer, error) {
	newWarrantyOffer := &WarrantyOffer{
		ODataType: "#microsoft.graph.WarrantyOffer",
	}
	return newWarrantyOffer, nil
}
