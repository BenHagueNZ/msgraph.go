
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// ResetPasscodeActionResult undocumented
type ResetPasscodeActionResult struct {
    // DeviceActionResult is the base model of ResetPasscodeActionResult
    DeviceActionResult
    // ErrorCode undocumented
    ErrorCode *int `json:"errorCode,omitempty"`
    // Passcode undocumented
    Passcode *string `json:"passcode,omitempty"`
}
