// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// StaffAvailabilityItem undocumented
type StaffAvailabilityItem struct {
	// Object is the base model of StaffAvailabilityItem
	Object
	// AvailabilityItems undocumented
	AvailabilityItems []AvailabilityItem `json:"availabilityItems,omitempty"`
	// StaffID undocumented
	StaffID *string `json:"staffId,omitempty"`
}