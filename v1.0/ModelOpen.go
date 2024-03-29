// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// OpenShift undocumented
type OpenShift struct {
	// ChangeTrackedEntity is the base model of OpenShift
	ChangeTrackedEntity

	ODataType string `json:"@odata.type,omitempty"`
	// DraftOpenShift undocumented
	DraftOpenShift *OpenShiftItem `json:"draftOpenShift,omitempty"`
	// SchedulingGroupID undocumented
	SchedulingGroupID *string `json:"schedulingGroupId,omitempty"`
	// SharedOpenShift undocumented
	SharedOpenShift *OpenShiftItem `json:"sharedOpenShift,omitempty"`
}

func NewOpenShift() (*OpenShift, error) {
	newOpenShift := &OpenShift{
		ODataType: "#microsoft.graph.OpenShift",
	}
	return newOpenShift, nil
}

// OpenShiftChangeRequestObject undocumented
type OpenShiftChangeRequestObject struct {
	// ScheduleChangeRequestObject is the base model of OpenShiftChangeRequestObject
	ScheduleChangeRequestObject

	ODataType string `json:"@odata.type,omitempty"`
	// OpenShiftID undocumented
	OpenShiftID *string `json:"openShiftId,omitempty"`
}

func NewOpenShiftChangeRequestObject() (*OpenShiftChangeRequestObject, error) {
	newOpenShiftChangeRequestObject := &OpenShiftChangeRequestObject{
		ODataType: "#microsoft.graph.OpenShiftChangeRequestObject",
	}
	return newOpenShiftChangeRequestObject, nil
}

// OpenShiftItem undocumented
type OpenShiftItem struct {
	// ShiftItem is the base model of OpenShiftItem
	ShiftItem

	ODataType string `json:"@odata.type,omitempty"`
	// OpenSlotCount undocumented
	OpenSlotCount *int `json:"openSlotCount,omitempty"`
}

func NewOpenShiftItem() (*OpenShiftItem, error) {
	newOpenShiftItem := &OpenShiftItem{
		ODataType: "#microsoft.graph.OpenShiftItem",
	}
	return newOpenShiftItem, nil
}

// OpenTypeExtension undocumented
type OpenTypeExtension struct {
	// Extension is the base model of OpenTypeExtension
	Extension

	ODataType string `json:"@odata.type,omitempty"`
	// ExtensionName undocumented
	ExtensionName *string `json:"extensionName,omitempty"`
}

func NewOpenTypeExtension() (*OpenTypeExtension, error) {
	newOpenTypeExtension := &OpenTypeExtension{
		ODataType: "#microsoft.graph.OpenTypeExtension",
	}
	return newOpenTypeExtension, nil
}
