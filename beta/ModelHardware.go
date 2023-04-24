// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// HardwareInformation undocumented
type HardwareInformation struct {
	// Object is the base model of HardwareInformation
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// BatteryChargeCycles undocumented
	BatteryChargeCycles *int `json:"batteryChargeCycles,omitempty"`
	// BatteryHealthPercentage undocumented
	BatteryHealthPercentage *int `json:"batteryHealthPercentage,omitempty"`
	// BatteryLevelPercentage undocumented
	BatteryLevelPercentage *float64 `json:"batteryLevelPercentage,omitempty"`
	// BatterySerialNumber undocumented
	BatterySerialNumber *string `json:"batterySerialNumber,omitempty"`
	// CellularTechnology undocumented
	CellularTechnology *string `json:"cellularTechnology,omitempty"`
	// DeviceFullQualifiedDomainName undocumented
	DeviceFullQualifiedDomainName *string `json:"deviceFullQualifiedDomainName,omitempty"`
	// DeviceGuardLocalSystemAuthorityCredentialGuardState undocumented
	DeviceGuardLocalSystemAuthorityCredentialGuardState *DeviceGuardLocalSystemAuthorityCredentialGuardState `json:"deviceGuardLocalSystemAuthorityCredentialGuardState,omitempty"`
	// DeviceGuardVirtualizationBasedSecurityHardwareRequirementState undocumented
	DeviceGuardVirtualizationBasedSecurityHardwareRequirementState *DeviceGuardVirtualizationBasedSecurityHardwareRequirementState `json:"deviceGuardVirtualizationBasedSecurityHardwareRequirementState,omitempty"`
	// DeviceGuardVirtualizationBasedSecurityState undocumented
	DeviceGuardVirtualizationBasedSecurityState *DeviceGuardVirtualizationBasedSecurityState `json:"deviceGuardVirtualizationBasedSecurityState,omitempty"`
	// DeviceLicensingLastErrorCode undocumented
	DeviceLicensingLastErrorCode *int `json:"deviceLicensingLastErrorCode,omitempty"`
	// DeviceLicensingLastErrorDescription undocumented
	DeviceLicensingLastErrorDescription *string `json:"deviceLicensingLastErrorDescription,omitempty"`
	// DeviceLicensingStatus undocumented
	DeviceLicensingStatus *DeviceLicensingStatus `json:"deviceLicensingStatus,omitempty"`
	// EsimIdentifier undocumented
	EsimIdentifier *string `json:"esimIdentifier,omitempty"`
	// FreeStorageSpace undocumented
	FreeStorageSpace *int `json:"freeStorageSpace,omitempty"`
	// Imei undocumented
	Imei *string `json:"imei,omitempty"`
	// IPAddressV4 undocumented
	IPAddressV4 *string `json:"ipAddressV4,omitempty"`
	// IsEncrypted undocumented
	IsEncrypted *bool `json:"isEncrypted,omitempty"`
	// IsSharedDevice undocumented
	IsSharedDevice *bool `json:"isSharedDevice,omitempty"`
	// IsSupervised undocumented
	IsSupervised *bool `json:"isSupervised,omitempty"`
	// Manufacturer undocumented
	Manufacturer *string `json:"manufacturer,omitempty"`
	// Meid undocumented
	Meid *string `json:"meid,omitempty"`
	// Model undocumented
	Model *string `json:"model,omitempty"`
	// OperatingSystemEdition undocumented
	OperatingSystemEdition *string `json:"operatingSystemEdition,omitempty"`
	// OperatingSystemLanguage undocumented
	OperatingSystemLanguage *string `json:"operatingSystemLanguage,omitempty"`
	// OperatingSystemProductType undocumented
	OperatingSystemProductType *int `json:"operatingSystemProductType,omitempty"`
	// OsBuildNumber undocumented
	OsBuildNumber *string `json:"osBuildNumber,omitempty"`
	// PhoneNumber undocumented
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// ProductName undocumented
	ProductName *string `json:"productName,omitempty"`
	// ResidentUsersCount undocumented
	ResidentUsersCount *int `json:"residentUsersCount,omitempty"`
	// SerialNumber undocumented
	SerialNumber *string `json:"serialNumber,omitempty"`
	// SharedDeviceCachedUsers undocumented
	SharedDeviceCachedUsers []SharedAppleDeviceUser `json:"sharedDeviceCachedUsers,omitempty"`
	// SubnetAddress undocumented
	SubnetAddress *string `json:"subnetAddress,omitempty"`
	// SubscriberCarrier undocumented
	SubscriberCarrier *string `json:"subscriberCarrier,omitempty"`
	// SystemManagementBIOSVersion undocumented
	SystemManagementBIOSVersion *string `json:"systemManagementBIOSVersion,omitempty"`
	// TotalStorageSpace undocumented
	TotalStorageSpace *int `json:"totalStorageSpace,omitempty"`
	// TpmManufacturer undocumented
	TpmManufacturer *string `json:"tpmManufacturer,omitempty"`
	// TpmSpecificationVersion undocumented
	TpmSpecificationVersion *string `json:"tpmSpecificationVersion,omitempty"`
	// TpmVersion undocumented
	TpmVersion *string `json:"tpmVersion,omitempty"`
	// WiFiMac undocumented
	WiFiMac *string `json:"wifiMac,omitempty"`
	// WiredIPv4Addresses undocumented
	WiredIPv4Addresses []string `json:"wiredIPv4Addresses,omitempty"`
}

func NewHardwareInformation() (*HardwareInformation, error) {
	newHardwareInformation := &HardwareInformation{
		ODataType: "#microsoft.graph.HardwareInformation",
	}
	return newHardwareInformation, nil
}
