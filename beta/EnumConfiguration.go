// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ConfigurationManagerActionDeliveryStatus undocumented
type ConfigurationManagerActionDeliveryStatus string

const (
	// ConfigurationManagerActionDeliveryStatusVUnknown undocumented
	ConfigurationManagerActionDeliveryStatusVUnknown ConfigurationManagerActionDeliveryStatus = "unknown"
	// ConfigurationManagerActionDeliveryStatusVPendingDelivery undocumented
	ConfigurationManagerActionDeliveryStatusVPendingDelivery ConfigurationManagerActionDeliveryStatus = "pendingDelivery"
	// ConfigurationManagerActionDeliveryStatusVDeliveredToConnectorService undocumented
	ConfigurationManagerActionDeliveryStatusVDeliveredToConnectorService ConfigurationManagerActionDeliveryStatus = "deliveredToConnectorService"
	// ConfigurationManagerActionDeliveryStatusVFailedToDeliverToConnectorService undocumented
	ConfigurationManagerActionDeliveryStatusVFailedToDeliverToConnectorService ConfigurationManagerActionDeliveryStatus = "failedToDeliverToConnectorService"
	// ConfigurationManagerActionDeliveryStatusVDeliveredToOnPremisesServer undocumented
	ConfigurationManagerActionDeliveryStatusVDeliveredToOnPremisesServer ConfigurationManagerActionDeliveryStatus = "deliveredToOnPremisesServer"
)

var (
	// ConfigurationManagerActionDeliveryStatusPUnknown is a pointer to ConfigurationManagerActionDeliveryStatusVUnknown
	ConfigurationManagerActionDeliveryStatusPUnknown = &_ConfigurationManagerActionDeliveryStatusPUnknown
	// ConfigurationManagerActionDeliveryStatusPPendingDelivery is a pointer to ConfigurationManagerActionDeliveryStatusVPendingDelivery
	ConfigurationManagerActionDeliveryStatusPPendingDelivery = &_ConfigurationManagerActionDeliveryStatusPPendingDelivery
	// ConfigurationManagerActionDeliveryStatusPDeliveredToConnectorService is a pointer to ConfigurationManagerActionDeliveryStatusVDeliveredToConnectorService
	ConfigurationManagerActionDeliveryStatusPDeliveredToConnectorService = &_ConfigurationManagerActionDeliveryStatusPDeliveredToConnectorService
	// ConfigurationManagerActionDeliveryStatusPFailedToDeliverToConnectorService is a pointer to ConfigurationManagerActionDeliveryStatusVFailedToDeliverToConnectorService
	ConfigurationManagerActionDeliveryStatusPFailedToDeliverToConnectorService = &_ConfigurationManagerActionDeliveryStatusPFailedToDeliverToConnectorService
	// ConfigurationManagerActionDeliveryStatusPDeliveredToOnPremisesServer is a pointer to ConfigurationManagerActionDeliveryStatusVDeliveredToOnPremisesServer
	ConfigurationManagerActionDeliveryStatusPDeliveredToOnPremisesServer = &_ConfigurationManagerActionDeliveryStatusPDeliveredToOnPremisesServer
)

var (
	_ConfigurationManagerActionDeliveryStatusPUnknown                           = ConfigurationManagerActionDeliveryStatusVUnknown
	_ConfigurationManagerActionDeliveryStatusPPendingDelivery                   = ConfigurationManagerActionDeliveryStatusVPendingDelivery
	_ConfigurationManagerActionDeliveryStatusPDeliveredToConnectorService       = ConfigurationManagerActionDeliveryStatusVDeliveredToConnectorService
	_ConfigurationManagerActionDeliveryStatusPFailedToDeliverToConnectorService = ConfigurationManagerActionDeliveryStatusVFailedToDeliverToConnectorService
	_ConfigurationManagerActionDeliveryStatusPDeliveredToOnPremisesServer       = ConfigurationManagerActionDeliveryStatusVDeliveredToOnPremisesServer
)

// ConfigurationManagerActionType undocumented
type ConfigurationManagerActionType string

const (
	// ConfigurationManagerActionTypeVRefreshMachinePolicy undocumented
	ConfigurationManagerActionTypeVRefreshMachinePolicy ConfigurationManagerActionType = "refreshMachinePolicy"
	// ConfigurationManagerActionTypeVRefreshUserPolicy undocumented
	ConfigurationManagerActionTypeVRefreshUserPolicy ConfigurationManagerActionType = "refreshUserPolicy"
	// ConfigurationManagerActionTypeVWakeUpClient undocumented
	ConfigurationManagerActionTypeVWakeUpClient ConfigurationManagerActionType = "wakeUpClient"
	// ConfigurationManagerActionTypeVAppEvaluation undocumented
	ConfigurationManagerActionTypeVAppEvaluation ConfigurationManagerActionType = "appEvaluation"
	// ConfigurationManagerActionTypeVQuickScan undocumented
	ConfigurationManagerActionTypeVQuickScan ConfigurationManagerActionType = "quickScan"
	// ConfigurationManagerActionTypeVFullScan undocumented
	ConfigurationManagerActionTypeVFullScan ConfigurationManagerActionType = "fullScan"
	// ConfigurationManagerActionTypeVWindowsDefenderUpdateSignatures undocumented
	ConfigurationManagerActionTypeVWindowsDefenderUpdateSignatures ConfigurationManagerActionType = "windowsDefenderUpdateSignatures"
)

var (
	// ConfigurationManagerActionTypePRefreshMachinePolicy is a pointer to ConfigurationManagerActionTypeVRefreshMachinePolicy
	ConfigurationManagerActionTypePRefreshMachinePolicy = &_ConfigurationManagerActionTypePRefreshMachinePolicy
	// ConfigurationManagerActionTypePRefreshUserPolicy is a pointer to ConfigurationManagerActionTypeVRefreshUserPolicy
	ConfigurationManagerActionTypePRefreshUserPolicy = &_ConfigurationManagerActionTypePRefreshUserPolicy
	// ConfigurationManagerActionTypePWakeUpClient is a pointer to ConfigurationManagerActionTypeVWakeUpClient
	ConfigurationManagerActionTypePWakeUpClient = &_ConfigurationManagerActionTypePWakeUpClient
	// ConfigurationManagerActionTypePAppEvaluation is a pointer to ConfigurationManagerActionTypeVAppEvaluation
	ConfigurationManagerActionTypePAppEvaluation = &_ConfigurationManagerActionTypePAppEvaluation
	// ConfigurationManagerActionTypePQuickScan is a pointer to ConfigurationManagerActionTypeVQuickScan
	ConfigurationManagerActionTypePQuickScan = &_ConfigurationManagerActionTypePQuickScan
	// ConfigurationManagerActionTypePFullScan is a pointer to ConfigurationManagerActionTypeVFullScan
	ConfigurationManagerActionTypePFullScan = &_ConfigurationManagerActionTypePFullScan
	// ConfigurationManagerActionTypePWindowsDefenderUpdateSignatures is a pointer to ConfigurationManagerActionTypeVWindowsDefenderUpdateSignatures
	ConfigurationManagerActionTypePWindowsDefenderUpdateSignatures = &_ConfigurationManagerActionTypePWindowsDefenderUpdateSignatures
)

var (
	_ConfigurationManagerActionTypePRefreshMachinePolicy            = ConfigurationManagerActionTypeVRefreshMachinePolicy
	_ConfigurationManagerActionTypePRefreshUserPolicy               = ConfigurationManagerActionTypeVRefreshUserPolicy
	_ConfigurationManagerActionTypePWakeUpClient                    = ConfigurationManagerActionTypeVWakeUpClient
	_ConfigurationManagerActionTypePAppEvaluation                   = ConfigurationManagerActionTypeVAppEvaluation
	_ConfigurationManagerActionTypePQuickScan                       = ConfigurationManagerActionTypeVQuickScan
	_ConfigurationManagerActionTypePFullScan                        = ConfigurationManagerActionTypeVFullScan
	_ConfigurationManagerActionTypePWindowsDefenderUpdateSignatures = ConfigurationManagerActionTypeVWindowsDefenderUpdateSignatures
)

// ConfigurationManagerClientState undocumented
type ConfigurationManagerClientState string

const (
	// ConfigurationManagerClientStateVUnknown undocumented
	ConfigurationManagerClientStateVUnknown ConfigurationManagerClientState = "unknown"
	// ConfigurationManagerClientStateVInstalled undocumented
	ConfigurationManagerClientStateVInstalled ConfigurationManagerClientState = "installed"
	// ConfigurationManagerClientStateVHealthy undocumented
	ConfigurationManagerClientStateVHealthy ConfigurationManagerClientState = "healthy"
	// ConfigurationManagerClientStateVInstallFailed undocumented
	ConfigurationManagerClientStateVInstallFailed ConfigurationManagerClientState = "installFailed"
	// ConfigurationManagerClientStateVUpdateFailed undocumented
	ConfigurationManagerClientStateVUpdateFailed ConfigurationManagerClientState = "updateFailed"
	// ConfigurationManagerClientStateVCommunicationError undocumented
	ConfigurationManagerClientStateVCommunicationError ConfigurationManagerClientState = "communicationError"
)

var (
	// ConfigurationManagerClientStatePUnknown is a pointer to ConfigurationManagerClientStateVUnknown
	ConfigurationManagerClientStatePUnknown = &_ConfigurationManagerClientStatePUnknown
	// ConfigurationManagerClientStatePInstalled is a pointer to ConfigurationManagerClientStateVInstalled
	ConfigurationManagerClientStatePInstalled = &_ConfigurationManagerClientStatePInstalled
	// ConfigurationManagerClientStatePHealthy is a pointer to ConfigurationManagerClientStateVHealthy
	ConfigurationManagerClientStatePHealthy = &_ConfigurationManagerClientStatePHealthy
	// ConfigurationManagerClientStatePInstallFailed is a pointer to ConfigurationManagerClientStateVInstallFailed
	ConfigurationManagerClientStatePInstallFailed = &_ConfigurationManagerClientStatePInstallFailed
	// ConfigurationManagerClientStatePUpdateFailed is a pointer to ConfigurationManagerClientStateVUpdateFailed
	ConfigurationManagerClientStatePUpdateFailed = &_ConfigurationManagerClientStatePUpdateFailed
	// ConfigurationManagerClientStatePCommunicationError is a pointer to ConfigurationManagerClientStateVCommunicationError
	ConfigurationManagerClientStatePCommunicationError = &_ConfigurationManagerClientStatePCommunicationError
)

var (
	_ConfigurationManagerClientStatePUnknown            = ConfigurationManagerClientStateVUnknown
	_ConfigurationManagerClientStatePInstalled          = ConfigurationManagerClientStateVInstalled
	_ConfigurationManagerClientStatePHealthy            = ConfigurationManagerClientStateVHealthy
	_ConfigurationManagerClientStatePInstallFailed      = ConfigurationManagerClientStateVInstallFailed
	_ConfigurationManagerClientStatePUpdateFailed       = ConfigurationManagerClientStateVUpdateFailed
	_ConfigurationManagerClientStatePCommunicationError = ConfigurationManagerClientStateVCommunicationError
)

// ConfigurationUsage undocumented
type ConfigurationUsage string

const (
	// ConfigurationUsageVBlocked undocumented
	ConfigurationUsageVBlocked ConfigurationUsage = "blocked"
	// ConfigurationUsageVRequired undocumented
	ConfigurationUsageVRequired ConfigurationUsage = "required"
	// ConfigurationUsageVAllowed undocumented
	ConfigurationUsageVAllowed ConfigurationUsage = "allowed"
	// ConfigurationUsageVNotConfigured undocumented
	ConfigurationUsageVNotConfigured ConfigurationUsage = "notConfigured"
)

var (
	// ConfigurationUsagePBlocked is a pointer to ConfigurationUsageVBlocked
	ConfigurationUsagePBlocked = &_ConfigurationUsagePBlocked
	// ConfigurationUsagePRequired is a pointer to ConfigurationUsageVRequired
	ConfigurationUsagePRequired = &_ConfigurationUsagePRequired
	// ConfigurationUsagePAllowed is a pointer to ConfigurationUsageVAllowed
	ConfigurationUsagePAllowed = &_ConfigurationUsagePAllowed
	// ConfigurationUsagePNotConfigured is a pointer to ConfigurationUsageVNotConfigured
	ConfigurationUsagePNotConfigured = &_ConfigurationUsagePNotConfigured
)

var (
	_ConfigurationUsagePBlocked       = ConfigurationUsageVBlocked
	_ConfigurationUsagePRequired      = ConfigurationUsageVRequired
	_ConfigurationUsagePAllowed       = ConfigurationUsageVAllowed
	_ConfigurationUsagePNotConfigured = ConfigurationUsageVNotConfigured
)