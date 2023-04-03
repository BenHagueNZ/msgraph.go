// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// AvailabilityItem undocumented
type AvailabilityItem struct {
	// Object is the base model of AvailabilityItem
	Object

	ODataType string `json:"@odata.type"`
	// EndDateTime undocumented
	EndDateTime *DateTimeTimeZone `json:"endDateTime,omitempty"`
	// ServiceID undocumented
	ServiceID *string `json:"serviceId,omitempty"`
	// StartDateTime undocumented
	StartDateTime *DateTimeTimeZone `json:"startDateTime,omitempty"`
	// Status undocumented
	Status *BookingsAvailabilityStatus `json:"status,omitempty"`
}

func NewAvailabilityItem() (*AvailabilityItem, error) {
	newAvailabilityItem := &AvailabilityItem{
		ODataType: "#microsoft.graph.AvailabilityItem",
	}
	return newAvailabilityItem, nil
}
