// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// PrivilegeManagementElevation undocumented
type PrivilegeManagementElevation struct {
	// Entity is the base model of PrivilegeManagementElevation
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// CertificatePayload undocumented
	CertificatePayload *string `json:"certificatePayload,omitempty"`
	// CompanyName undocumented
	CompanyName *string `json:"companyName,omitempty"`
	// DeviceID undocumented
	DeviceID *string `json:"deviceId,omitempty"`
	// DeviceName undocumented
	DeviceName *string `json:"deviceName,omitempty"`
	// ElevationType undocumented
	ElevationType *PrivilegeManagementElevationType `json:"elevationType,omitempty"`
	// EventDateTime undocumented
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`
	// FileDescription undocumented
	FileDescription *string `json:"fileDescription,omitempty"`
	// FilePath undocumented
	FilePath *string `json:"filePath,omitempty"`
	// FileVersion undocumented
	FileVersion *string `json:"fileVersion,omitempty"`
	// Hash undocumented
	Hash *string `json:"hash,omitempty"`
	// InternalName undocumented
	InternalName *string `json:"internalName,omitempty"`
	// Justification undocumented
	Justification *string `json:"justification,omitempty"`
	// ProductName undocumented
	ProductName *string `json:"productName,omitempty"`
	// Result undocumented
	Result *int `json:"result,omitempty"`
	// Upn undocumented
	Upn *string `json:"upn,omitempty"`
	// UserType undocumented
	UserType *PrivilegeManagementEndUserType `json:"userType,omitempty"`
}

func NewPrivilegeManagementElevation() (*PrivilegeManagementElevation, error) {
	newPrivilegeManagementElevation := &PrivilegeManagementElevation{
		ODataType: "#microsoft.graph.PrivilegeManagementElevation",
	}
	return newPrivilegeManagementElevation, nil
}
