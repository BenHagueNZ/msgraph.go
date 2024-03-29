// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// UpdateAllowedCombinationsResult undocumented
type UpdateAllowedCombinationsResult struct {
	// Object is the base model of UpdateAllowedCombinationsResult
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AdditionalInformation undocumented
	AdditionalInformation *string `json:"additionalInformation,omitempty"`
	// ConditionalAccessReferences undocumented
	ConditionalAccessReferences []string `json:"conditionalAccessReferences,omitempty"`
	// CurrentCombinations undocumented
	CurrentCombinations []AuthenticationMethodModes `json:"currentCombinations,omitempty"`
	// PreviousCombinations undocumented
	PreviousCombinations []AuthenticationMethodModes `json:"previousCombinations,omitempty"`
}

func NewUpdateAllowedCombinationsResult() (*UpdateAllowedCombinationsResult, error) {
	newUpdateAllowedCombinationsResult := &UpdateAllowedCombinationsResult{
		ODataType: "#microsoft.graph.UpdateAllowedCombinationsResult",
	}
	return newUpdateAllowedCombinationsResult, nil
}

// UpdateRecordingStatusOperation undocumented
type UpdateRecordingStatusOperation struct {
	// CommsOperation is the base model of UpdateRecordingStatusOperation
	CommsOperation

	ODataType string `json:"@odata.type,omitempty"`
}

func NewUpdateRecordingStatusOperation() (*UpdateRecordingStatusOperation, error) {
	newUpdateRecordingStatusOperation := &UpdateRecordingStatusOperation{
		ODataType: "#microsoft.graph.UpdateRecordingStatusOperation",
	}
	return newUpdateRecordingStatusOperation, nil
}

// UpdateWindow undocumented
type UpdateWindow struct {
	// Object is the base model of UpdateWindow
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// UpdateWindowEndTime undocumented
	UpdateWindowEndTime *TimeOfDay `json:"updateWindowEndTime,omitempty"`
	// UpdateWindowStartTime undocumented
	UpdateWindowStartTime *TimeOfDay `json:"updateWindowStartTime,omitempty"`
}

func NewUpdateWindow() (*UpdateWindow, error) {
	newUpdateWindow := &UpdateWindow{
		ODataType: "#microsoft.graph.UpdateWindow",
	}
	return newUpdateWindow, nil
}

// UpdateWindowsDeviceAccountActionParameter undocumented
type UpdateWindowsDeviceAccountActionParameter struct {
	// Object is the base model of UpdateWindowsDeviceAccountActionParameter
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// CalendarSyncEnabled undocumented
	CalendarSyncEnabled *bool `json:"calendarSyncEnabled,omitempty"`
	// DeviceAccount undocumented
	DeviceAccount *WindowsDeviceAccount `json:"deviceAccount,omitempty"`
	// DeviceAccountEmail undocumented
	DeviceAccountEmail *string `json:"deviceAccountEmail,omitempty"`
	// ExchangeServer undocumented
	ExchangeServer *string `json:"exchangeServer,omitempty"`
	// PasswordRotationEnabled undocumented
	PasswordRotationEnabled *bool `json:"passwordRotationEnabled,omitempty"`
	// SessionInitiationProtocalAddress undocumented
	SessionInitiationProtocalAddress *string `json:"sessionInitiationProtocalAddress,omitempty"`
}

func NewUpdateWindowsDeviceAccountActionParameter() (*UpdateWindowsDeviceAccountActionParameter, error) {
	newUpdateWindowsDeviceAccountActionParameter := &UpdateWindowsDeviceAccountActionParameter{
		ODataType: "#microsoft.graph.UpdateWindowsDeviceAccountActionParameter",
	}
	return newUpdateWindowsDeviceAccountActionParameter, nil
}
