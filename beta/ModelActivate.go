// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ActivateDeviceEsimActionResult undocumented
type ActivateDeviceEsimActionResult struct {
	// DeviceActionResult is the base model of ActivateDeviceEsimActionResult
	DeviceActionResult

	ODataType string `json:"@odata.type,omitempty"`
	// CarrierURL undocumented
	CarrierURL *string `json:"carrierUrl,omitempty"`
}

func NewActivateDeviceEsimActionResult() (*ActivateDeviceEsimActionResult, error) {
	newActivateDeviceEsimActionResult := &ActivateDeviceEsimActionResult{
		ODataType: "#microsoft.graph.ActivateDeviceEsimActionResult",
	}
	return newActivateDeviceEsimActionResult, nil
}
