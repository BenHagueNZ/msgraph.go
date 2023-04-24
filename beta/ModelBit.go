// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// BitLockerFixedDrivePolicy undocumented
type BitLockerFixedDrivePolicy struct {
	// Object is the base model of BitLockerFixedDrivePolicy
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// EncryptionMethod undocumented
	EncryptionMethod *BitLockerEncryptionMethod `json:"encryptionMethod,omitempty"`
	// RecoveryOptions undocumented
	RecoveryOptions *BitLockerRecoveryOptions `json:"recoveryOptions,omitempty"`
	// RequireEncryptionForWriteAccess undocumented
	RequireEncryptionForWriteAccess *bool `json:"requireEncryptionForWriteAccess,omitempty"`
}

func NewBitLockerFixedDrivePolicy() (*BitLockerFixedDrivePolicy, error) {
	newBitLockerFixedDrivePolicy := &BitLockerFixedDrivePolicy{
		ODataType: "#microsoft.graph.BitLockerFixedDrivePolicy",
	}
	return newBitLockerFixedDrivePolicy, nil
}

// BitLockerRecoveryOptions undocumented
type BitLockerRecoveryOptions struct {
	// Object is the base model of BitLockerRecoveryOptions
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// BlockDataRecoveryAgent undocumented
	BlockDataRecoveryAgent *bool `json:"blockDataRecoveryAgent,omitempty"`
	// EnableBitLockerAfterRecoveryInformationToStore undocumented
	EnableBitLockerAfterRecoveryInformationToStore *bool `json:"enableBitLockerAfterRecoveryInformationToStore,omitempty"`
	// EnableRecoveryInformationSaveToStore undocumented
	EnableRecoveryInformationSaveToStore *bool `json:"enableRecoveryInformationSaveToStore,omitempty"`
	// HideRecoveryOptions undocumented
	HideRecoveryOptions *bool `json:"hideRecoveryOptions,omitempty"`
	// RecoveryInformationToStore undocumented
	RecoveryInformationToStore *BitLockerRecoveryInformationType `json:"recoveryInformationToStore,omitempty"`
	// RecoveryKeyUsage undocumented
	RecoveryKeyUsage *ConfigurationUsage `json:"recoveryKeyUsage,omitempty"`
	// RecoveryPasswordUsage undocumented
	RecoveryPasswordUsage *ConfigurationUsage `json:"recoveryPasswordUsage,omitempty"`
}

func NewBitLockerRecoveryOptions() (*BitLockerRecoveryOptions, error) {
	newBitLockerRecoveryOptions := &BitLockerRecoveryOptions{
		ODataType: "#microsoft.graph.BitLockerRecoveryOptions",
	}
	return newBitLockerRecoveryOptions, nil
}

// BitLockerRemovableDrivePolicy undocumented
type BitLockerRemovableDrivePolicy struct {
	// Object is the base model of BitLockerRemovableDrivePolicy
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// BlockCrossOrganizationWriteAccess undocumented
	BlockCrossOrganizationWriteAccess *bool `json:"blockCrossOrganizationWriteAccess,omitempty"`
	// EncryptionMethod undocumented
	EncryptionMethod *BitLockerEncryptionMethod `json:"encryptionMethod,omitempty"`
	// RequireEncryptionForWriteAccess undocumented
	RequireEncryptionForWriteAccess *bool `json:"requireEncryptionForWriteAccess,omitempty"`
}

func NewBitLockerRemovableDrivePolicy() (*BitLockerRemovableDrivePolicy, error) {
	newBitLockerRemovableDrivePolicy := &BitLockerRemovableDrivePolicy{
		ODataType: "#microsoft.graph.BitLockerRemovableDrivePolicy",
	}
	return newBitLockerRemovableDrivePolicy, nil
}

// BitLockerSystemDrivePolicy undocumented
type BitLockerSystemDrivePolicy struct {
	// Object is the base model of BitLockerSystemDrivePolicy
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// EncryptionMethod undocumented
	EncryptionMethod *BitLockerEncryptionMethod `json:"encryptionMethod,omitempty"`
	// MinimumPinLength undocumented
	MinimumPinLength *int `json:"minimumPinLength,omitempty"`
	// PrebootRecoveryEnableMessageAndURL undocumented
	PrebootRecoveryEnableMessageAndURL *bool `json:"prebootRecoveryEnableMessageAndUrl,omitempty"`
	// PrebootRecoveryMessage undocumented
	PrebootRecoveryMessage *string `json:"prebootRecoveryMessage,omitempty"`
	// PrebootRecoveryURL undocumented
	PrebootRecoveryURL *string `json:"prebootRecoveryUrl,omitempty"`
	// RecoveryOptions undocumented
	RecoveryOptions *BitLockerRecoveryOptions `json:"recoveryOptions,omitempty"`
	// StartupAuthenticationBlockWithoutTpmChip undocumented
	StartupAuthenticationBlockWithoutTpmChip *bool `json:"startupAuthenticationBlockWithoutTpmChip,omitempty"`
	// StartupAuthenticationRequired undocumented
	StartupAuthenticationRequired *bool `json:"startupAuthenticationRequired,omitempty"`
	// StartupAuthenticationTpmKeyUsage undocumented
	StartupAuthenticationTpmKeyUsage *ConfigurationUsage `json:"startupAuthenticationTpmKeyUsage,omitempty"`
	// StartupAuthenticationTpmPinAndKeyUsage undocumented
	StartupAuthenticationTpmPinAndKeyUsage *ConfigurationUsage `json:"startupAuthenticationTpmPinAndKeyUsage,omitempty"`
	// StartupAuthenticationTpmPinUsage undocumented
	StartupAuthenticationTpmPinUsage *ConfigurationUsage `json:"startupAuthenticationTpmPinUsage,omitempty"`
	// StartupAuthenticationTpmUsage undocumented
	StartupAuthenticationTpmUsage *ConfigurationUsage `json:"startupAuthenticationTpmUsage,omitempty"`
}

func NewBitLockerSystemDrivePolicy() (*BitLockerSystemDrivePolicy, error) {
	newBitLockerSystemDrivePolicy := &BitLockerSystemDrivePolicy{
		ODataType: "#microsoft.graph.BitLockerSystemDrivePolicy",
	}
	return newBitLockerSystemDrivePolicy, nil
}