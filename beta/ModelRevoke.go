// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// RevokeAppleVPPLicensesActionResult undocumented
type RevokeAppleVPPLicensesActionResult struct {
	// DeviceActionResult is the base model of RevokeAppleVPPLicensesActionResult
	DeviceActionResult

	ODataType string `json:"@odata.type,omitempty"`
	// FailedLicensesCount undocumented
	FailedLicensesCount *int `json:"failedLicensesCount,omitempty"`
	// TotalLicensesCount undocumented
	TotalLicensesCount *int `json:"totalLicensesCount,omitempty"`
}

func NewRevokeAppleVPPLicensesActionResult() (*RevokeAppleVPPLicensesActionResult, error) {
	newRevokeAppleVPPLicensesActionResult := &RevokeAppleVPPLicensesActionResult{
		ODataType: "#microsoft.graph.RevokeAppleVppLicensesActionResult",
	}
	return newRevokeAppleVPPLicensesActionResult, nil
}
