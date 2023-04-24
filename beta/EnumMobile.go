// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// MobileAppActionType undocumented
type MobileAppActionType string

const (
	// MobileAppActionTypeVUnknown undocumented
	MobileAppActionTypeVUnknown MobileAppActionType = "unknown"
	// MobileAppActionTypeVInstallCommandSent undocumented
	MobileAppActionTypeVInstallCommandSent MobileAppActionType = "installCommandSent"
	// MobileAppActionTypeVInstalled undocumented
	MobileAppActionTypeVInstalled MobileAppActionType = "installed"
	// MobileAppActionTypeVUninstalled undocumented
	MobileAppActionTypeVUninstalled MobileAppActionType = "uninstalled"
	// MobileAppActionTypeVUserRequestedInstall undocumented
	MobileAppActionTypeVUserRequestedInstall MobileAppActionType = "userRequestedInstall"
)

var (
	// MobileAppActionTypePUnknown is a pointer to MobileAppActionTypeVUnknown
	MobileAppActionTypePUnknown = &_MobileAppActionTypePUnknown
	// MobileAppActionTypePInstallCommandSent is a pointer to MobileAppActionTypeVInstallCommandSent
	MobileAppActionTypePInstallCommandSent = &_MobileAppActionTypePInstallCommandSent
	// MobileAppActionTypePInstalled is a pointer to MobileAppActionTypeVInstalled
	MobileAppActionTypePInstalled = &_MobileAppActionTypePInstalled
	// MobileAppActionTypePUninstalled is a pointer to MobileAppActionTypeVUninstalled
	MobileAppActionTypePUninstalled = &_MobileAppActionTypePUninstalled
	// MobileAppActionTypePUserRequestedInstall is a pointer to MobileAppActionTypeVUserRequestedInstall
	MobileAppActionTypePUserRequestedInstall = &_MobileAppActionTypePUserRequestedInstall
)

var (
	_MobileAppActionTypePUnknown              = MobileAppActionTypeVUnknown
	_MobileAppActionTypePInstallCommandSent   = MobileAppActionTypeVInstallCommandSent
	_MobileAppActionTypePInstalled            = MobileAppActionTypeVInstalled
	_MobileAppActionTypePUninstalled          = MobileAppActionTypeVUninstalled
	_MobileAppActionTypePUserRequestedInstall = MobileAppActionTypeVUserRequestedInstall
)

// MobileAppContentFileUploadState undocumented
type MobileAppContentFileUploadState string

const (
	// MobileAppContentFileUploadStateVSuccess undocumented
	MobileAppContentFileUploadStateVSuccess MobileAppContentFileUploadState = "success"
	// MobileAppContentFileUploadStateVTransientError undocumented
	MobileAppContentFileUploadStateVTransientError MobileAppContentFileUploadState = "transientError"
	// MobileAppContentFileUploadStateVError undocumented
	MobileAppContentFileUploadStateVError MobileAppContentFileUploadState = "error"
	// MobileAppContentFileUploadStateVUnknown undocumented
	MobileAppContentFileUploadStateVUnknown MobileAppContentFileUploadState = "unknown"
	// MobileAppContentFileUploadStateVAzureStorageURIRequestSuccess undocumented
	MobileAppContentFileUploadStateVAzureStorageURIRequestSuccess MobileAppContentFileUploadState = "azureStorageUriRequestSuccess"
	// MobileAppContentFileUploadStateVAzureStorageURIRequestPending undocumented
	MobileAppContentFileUploadStateVAzureStorageURIRequestPending MobileAppContentFileUploadState = "azureStorageUriRequestPending"
	// MobileAppContentFileUploadStateVAzureStorageURIRequestFailed undocumented
	MobileAppContentFileUploadStateVAzureStorageURIRequestFailed MobileAppContentFileUploadState = "azureStorageUriRequestFailed"
	// MobileAppContentFileUploadStateVAzureStorageURIRequestTimedOut undocumented
	MobileAppContentFileUploadStateVAzureStorageURIRequestTimedOut MobileAppContentFileUploadState = "azureStorageUriRequestTimedOut"
	// MobileAppContentFileUploadStateVAzureStorageURIRenewalSuccess undocumented
	MobileAppContentFileUploadStateVAzureStorageURIRenewalSuccess MobileAppContentFileUploadState = "azureStorageUriRenewalSuccess"
	// MobileAppContentFileUploadStateVAzureStorageURIRenewalPending undocumented
	MobileAppContentFileUploadStateVAzureStorageURIRenewalPending MobileAppContentFileUploadState = "azureStorageUriRenewalPending"
	// MobileAppContentFileUploadStateVAzureStorageURIRenewalFailed undocumented
	MobileAppContentFileUploadStateVAzureStorageURIRenewalFailed MobileAppContentFileUploadState = "azureStorageUriRenewalFailed"
	// MobileAppContentFileUploadStateVAzureStorageURIRenewalTimedOut undocumented
	MobileAppContentFileUploadStateVAzureStorageURIRenewalTimedOut MobileAppContentFileUploadState = "azureStorageUriRenewalTimedOut"
	// MobileAppContentFileUploadStateVCommitFileSuccess undocumented
	MobileAppContentFileUploadStateVCommitFileSuccess MobileAppContentFileUploadState = "commitFileSuccess"
	// MobileAppContentFileUploadStateVCommitFilePending undocumented
	MobileAppContentFileUploadStateVCommitFilePending MobileAppContentFileUploadState = "commitFilePending"
	// MobileAppContentFileUploadStateVCommitFileFailed undocumented
	MobileAppContentFileUploadStateVCommitFileFailed MobileAppContentFileUploadState = "commitFileFailed"
	// MobileAppContentFileUploadStateVCommitFileTimedOut undocumented
	MobileAppContentFileUploadStateVCommitFileTimedOut MobileAppContentFileUploadState = "commitFileTimedOut"
)

var (
	// MobileAppContentFileUploadStatePSuccess is a pointer to MobileAppContentFileUploadStateVSuccess
	MobileAppContentFileUploadStatePSuccess = &_MobileAppContentFileUploadStatePSuccess
	// MobileAppContentFileUploadStatePTransientError is a pointer to MobileAppContentFileUploadStateVTransientError
	MobileAppContentFileUploadStatePTransientError = &_MobileAppContentFileUploadStatePTransientError
	// MobileAppContentFileUploadStatePError is a pointer to MobileAppContentFileUploadStateVError
	MobileAppContentFileUploadStatePError = &_MobileAppContentFileUploadStatePError
	// MobileAppContentFileUploadStatePUnknown is a pointer to MobileAppContentFileUploadStateVUnknown
	MobileAppContentFileUploadStatePUnknown = &_MobileAppContentFileUploadStatePUnknown
	// MobileAppContentFileUploadStatePAzureStorageURIRequestSuccess is a pointer to MobileAppContentFileUploadStateVAzureStorageURIRequestSuccess
	MobileAppContentFileUploadStatePAzureStorageURIRequestSuccess = &_MobileAppContentFileUploadStatePAzureStorageURIRequestSuccess
	// MobileAppContentFileUploadStatePAzureStorageURIRequestPending is a pointer to MobileAppContentFileUploadStateVAzureStorageURIRequestPending
	MobileAppContentFileUploadStatePAzureStorageURIRequestPending = &_MobileAppContentFileUploadStatePAzureStorageURIRequestPending
	// MobileAppContentFileUploadStatePAzureStorageURIRequestFailed is a pointer to MobileAppContentFileUploadStateVAzureStorageURIRequestFailed
	MobileAppContentFileUploadStatePAzureStorageURIRequestFailed = &_MobileAppContentFileUploadStatePAzureStorageURIRequestFailed
	// MobileAppContentFileUploadStatePAzureStorageURIRequestTimedOut is a pointer to MobileAppContentFileUploadStateVAzureStorageURIRequestTimedOut
	MobileAppContentFileUploadStatePAzureStorageURIRequestTimedOut = &_MobileAppContentFileUploadStatePAzureStorageURIRequestTimedOut
	// MobileAppContentFileUploadStatePAzureStorageURIRenewalSuccess is a pointer to MobileAppContentFileUploadStateVAzureStorageURIRenewalSuccess
	MobileAppContentFileUploadStatePAzureStorageURIRenewalSuccess = &_MobileAppContentFileUploadStatePAzureStorageURIRenewalSuccess
	// MobileAppContentFileUploadStatePAzureStorageURIRenewalPending is a pointer to MobileAppContentFileUploadStateVAzureStorageURIRenewalPending
	MobileAppContentFileUploadStatePAzureStorageURIRenewalPending = &_MobileAppContentFileUploadStatePAzureStorageURIRenewalPending
	// MobileAppContentFileUploadStatePAzureStorageURIRenewalFailed is a pointer to MobileAppContentFileUploadStateVAzureStorageURIRenewalFailed
	MobileAppContentFileUploadStatePAzureStorageURIRenewalFailed = &_MobileAppContentFileUploadStatePAzureStorageURIRenewalFailed
	// MobileAppContentFileUploadStatePAzureStorageURIRenewalTimedOut is a pointer to MobileAppContentFileUploadStateVAzureStorageURIRenewalTimedOut
	MobileAppContentFileUploadStatePAzureStorageURIRenewalTimedOut = &_MobileAppContentFileUploadStatePAzureStorageURIRenewalTimedOut
	// MobileAppContentFileUploadStatePCommitFileSuccess is a pointer to MobileAppContentFileUploadStateVCommitFileSuccess
	MobileAppContentFileUploadStatePCommitFileSuccess = &_MobileAppContentFileUploadStatePCommitFileSuccess
	// MobileAppContentFileUploadStatePCommitFilePending is a pointer to MobileAppContentFileUploadStateVCommitFilePending
	MobileAppContentFileUploadStatePCommitFilePending = &_MobileAppContentFileUploadStatePCommitFilePending
	// MobileAppContentFileUploadStatePCommitFileFailed is a pointer to MobileAppContentFileUploadStateVCommitFileFailed
	MobileAppContentFileUploadStatePCommitFileFailed = &_MobileAppContentFileUploadStatePCommitFileFailed
	// MobileAppContentFileUploadStatePCommitFileTimedOut is a pointer to MobileAppContentFileUploadStateVCommitFileTimedOut
	MobileAppContentFileUploadStatePCommitFileTimedOut = &_MobileAppContentFileUploadStatePCommitFileTimedOut
)

var (
	_MobileAppContentFileUploadStatePSuccess                        = MobileAppContentFileUploadStateVSuccess
	_MobileAppContentFileUploadStatePTransientError                 = MobileAppContentFileUploadStateVTransientError
	_MobileAppContentFileUploadStatePError                          = MobileAppContentFileUploadStateVError
	_MobileAppContentFileUploadStatePUnknown                        = MobileAppContentFileUploadStateVUnknown
	_MobileAppContentFileUploadStatePAzureStorageURIRequestSuccess  = MobileAppContentFileUploadStateVAzureStorageURIRequestSuccess
	_MobileAppContentFileUploadStatePAzureStorageURIRequestPending  = MobileAppContentFileUploadStateVAzureStorageURIRequestPending
	_MobileAppContentFileUploadStatePAzureStorageURIRequestFailed   = MobileAppContentFileUploadStateVAzureStorageURIRequestFailed
	_MobileAppContentFileUploadStatePAzureStorageURIRequestTimedOut = MobileAppContentFileUploadStateVAzureStorageURIRequestTimedOut
	_MobileAppContentFileUploadStatePAzureStorageURIRenewalSuccess  = MobileAppContentFileUploadStateVAzureStorageURIRenewalSuccess
	_MobileAppContentFileUploadStatePAzureStorageURIRenewalPending  = MobileAppContentFileUploadStateVAzureStorageURIRenewalPending
	_MobileAppContentFileUploadStatePAzureStorageURIRenewalFailed   = MobileAppContentFileUploadStateVAzureStorageURIRenewalFailed
	_MobileAppContentFileUploadStatePAzureStorageURIRenewalTimedOut = MobileAppContentFileUploadStateVAzureStorageURIRenewalTimedOut
	_MobileAppContentFileUploadStatePCommitFileSuccess              = MobileAppContentFileUploadStateVCommitFileSuccess
	_MobileAppContentFileUploadStatePCommitFilePending              = MobileAppContentFileUploadStateVCommitFilePending
	_MobileAppContentFileUploadStatePCommitFileFailed               = MobileAppContentFileUploadStateVCommitFileFailed
	_MobileAppContentFileUploadStatePCommitFileTimedOut             = MobileAppContentFileUploadStateVCommitFileTimedOut
)

// MobileAppDependencyType undocumented
type MobileAppDependencyType string

const (
	// MobileAppDependencyTypeVDetect undocumented
	MobileAppDependencyTypeVDetect MobileAppDependencyType = "detect"
	// MobileAppDependencyTypeVAutoInstall undocumented
	MobileAppDependencyTypeVAutoInstall MobileAppDependencyType = "autoInstall"
)

var (
	// MobileAppDependencyTypePDetect is a pointer to MobileAppDependencyTypeVDetect
	MobileAppDependencyTypePDetect = &_MobileAppDependencyTypePDetect
	// MobileAppDependencyTypePAutoInstall is a pointer to MobileAppDependencyTypeVAutoInstall
	MobileAppDependencyTypePAutoInstall = &_MobileAppDependencyTypePAutoInstall
)

var (
	_MobileAppDependencyTypePDetect      = MobileAppDependencyTypeVDetect
	_MobileAppDependencyTypePAutoInstall = MobileAppDependencyTypeVAutoInstall
)

// MobileAppIntent undocumented
type MobileAppIntent string

const (
	// MobileAppIntentVAvailable undocumented
	MobileAppIntentVAvailable MobileAppIntent = "available"
	// MobileAppIntentVNotAvailable undocumented
	MobileAppIntentVNotAvailable MobileAppIntent = "notAvailable"
	// MobileAppIntentVRequiredInstall undocumented
	MobileAppIntentVRequiredInstall MobileAppIntent = "requiredInstall"
	// MobileAppIntentVRequiredUninstall undocumented
	MobileAppIntentVRequiredUninstall MobileAppIntent = "requiredUninstall"
	// MobileAppIntentVRequiredAndAvailableInstall undocumented
	MobileAppIntentVRequiredAndAvailableInstall MobileAppIntent = "requiredAndAvailableInstall"
	// MobileAppIntentVAvailableInstallWithoutEnrollment undocumented
	MobileAppIntentVAvailableInstallWithoutEnrollment MobileAppIntent = "availableInstallWithoutEnrollment"
	// MobileAppIntentVExclude undocumented
	MobileAppIntentVExclude MobileAppIntent = "exclude"
)

var (
	// MobileAppIntentPAvailable is a pointer to MobileAppIntentVAvailable
	MobileAppIntentPAvailable = &_MobileAppIntentPAvailable
	// MobileAppIntentPNotAvailable is a pointer to MobileAppIntentVNotAvailable
	MobileAppIntentPNotAvailable = &_MobileAppIntentPNotAvailable
	// MobileAppIntentPRequiredInstall is a pointer to MobileAppIntentVRequiredInstall
	MobileAppIntentPRequiredInstall = &_MobileAppIntentPRequiredInstall
	// MobileAppIntentPRequiredUninstall is a pointer to MobileAppIntentVRequiredUninstall
	MobileAppIntentPRequiredUninstall = &_MobileAppIntentPRequiredUninstall
	// MobileAppIntentPRequiredAndAvailableInstall is a pointer to MobileAppIntentVRequiredAndAvailableInstall
	MobileAppIntentPRequiredAndAvailableInstall = &_MobileAppIntentPRequiredAndAvailableInstall
	// MobileAppIntentPAvailableInstallWithoutEnrollment is a pointer to MobileAppIntentVAvailableInstallWithoutEnrollment
	MobileAppIntentPAvailableInstallWithoutEnrollment = &_MobileAppIntentPAvailableInstallWithoutEnrollment
	// MobileAppIntentPExclude is a pointer to MobileAppIntentVExclude
	MobileAppIntentPExclude = &_MobileAppIntentPExclude
)

var (
	_MobileAppIntentPAvailable                         = MobileAppIntentVAvailable
	_MobileAppIntentPNotAvailable                      = MobileAppIntentVNotAvailable
	_MobileAppIntentPRequiredInstall                   = MobileAppIntentVRequiredInstall
	_MobileAppIntentPRequiredUninstall                 = MobileAppIntentVRequiredUninstall
	_MobileAppIntentPRequiredAndAvailableInstall       = MobileAppIntentVRequiredAndAvailableInstall
	_MobileAppIntentPAvailableInstallWithoutEnrollment = MobileAppIntentVAvailableInstallWithoutEnrollment
	_MobileAppIntentPExclude                           = MobileAppIntentVExclude
)

// MobileAppPublishingState undocumented
type MobileAppPublishingState string

const (
	// MobileAppPublishingStateVNotPublished undocumented
	MobileAppPublishingStateVNotPublished MobileAppPublishingState = "notPublished"
	// MobileAppPublishingStateVProcessing undocumented
	MobileAppPublishingStateVProcessing MobileAppPublishingState = "processing"
	// MobileAppPublishingStateVPublished undocumented
	MobileAppPublishingStateVPublished MobileAppPublishingState = "published"
)

var (
	// MobileAppPublishingStatePNotPublished is a pointer to MobileAppPublishingStateVNotPublished
	MobileAppPublishingStatePNotPublished = &_MobileAppPublishingStatePNotPublished
	// MobileAppPublishingStatePProcessing is a pointer to MobileAppPublishingStateVProcessing
	MobileAppPublishingStatePProcessing = &_MobileAppPublishingStatePProcessing
	// MobileAppPublishingStatePPublished is a pointer to MobileAppPublishingStateVPublished
	MobileAppPublishingStatePPublished = &_MobileAppPublishingStatePPublished
)

var (
	_MobileAppPublishingStatePNotPublished = MobileAppPublishingStateVNotPublished
	_MobileAppPublishingStatePProcessing   = MobileAppPublishingStateVProcessing
	_MobileAppPublishingStatePPublished    = MobileAppPublishingStateVPublished
)

// MobileAppRelationshipType undocumented
type MobileAppRelationshipType string

const (
	// MobileAppRelationshipTypeVChild undocumented
	MobileAppRelationshipTypeVChild MobileAppRelationshipType = "child"
	// MobileAppRelationshipTypeVParent undocumented
	MobileAppRelationshipTypeVParent MobileAppRelationshipType = "parent"
)

var (
	// MobileAppRelationshipTypePChild is a pointer to MobileAppRelationshipTypeVChild
	MobileAppRelationshipTypePChild = &_MobileAppRelationshipTypePChild
	// MobileAppRelationshipTypePParent is a pointer to MobileAppRelationshipTypeVParent
	MobileAppRelationshipTypePParent = &_MobileAppRelationshipTypePParent
)

var (
	_MobileAppRelationshipTypePChild  = MobileAppRelationshipTypeVChild
	_MobileAppRelationshipTypePParent = MobileAppRelationshipTypeVParent
)

// MobileAppSupersedenceType undocumented
type MobileAppSupersedenceType string

const (
	// MobileAppSupersedenceTypeVUpdate undocumented
	MobileAppSupersedenceTypeVUpdate MobileAppSupersedenceType = "update"
	// MobileAppSupersedenceTypeVReplace undocumented
	MobileAppSupersedenceTypeVReplace MobileAppSupersedenceType = "replace"
)

var (
	// MobileAppSupersedenceTypePUpdate is a pointer to MobileAppSupersedenceTypeVUpdate
	MobileAppSupersedenceTypePUpdate = &_MobileAppSupersedenceTypePUpdate
	// MobileAppSupersedenceTypePReplace is a pointer to MobileAppSupersedenceTypeVReplace
	MobileAppSupersedenceTypePReplace = &_MobileAppSupersedenceTypePReplace
)

var (
	_MobileAppSupersedenceTypePUpdate  = MobileAppSupersedenceTypeVUpdate
	_MobileAppSupersedenceTypePReplace = MobileAppSupersedenceTypeVReplace
)

// MobileThreatDefensePartnerPriority undocumented
type MobileThreatDefensePartnerPriority string

const (
	// MobileThreatDefensePartnerPriorityVDefenderOverThirdPartyPartner undocumented
	MobileThreatDefensePartnerPriorityVDefenderOverThirdPartyPartner MobileThreatDefensePartnerPriority = "defenderOverThirdPartyPartner"
	// MobileThreatDefensePartnerPriorityVThirdPartyPartnerOverDefender undocumented
	MobileThreatDefensePartnerPriorityVThirdPartyPartnerOverDefender MobileThreatDefensePartnerPriority = "thirdPartyPartnerOverDefender"
	// MobileThreatDefensePartnerPriorityVUnknownFutureValue undocumented
	MobileThreatDefensePartnerPriorityVUnknownFutureValue MobileThreatDefensePartnerPriority = "unknownFutureValue"
)

var (
	// MobileThreatDefensePartnerPriorityPDefenderOverThirdPartyPartner is a pointer to MobileThreatDefensePartnerPriorityVDefenderOverThirdPartyPartner
	MobileThreatDefensePartnerPriorityPDefenderOverThirdPartyPartner = &_MobileThreatDefensePartnerPriorityPDefenderOverThirdPartyPartner
	// MobileThreatDefensePartnerPriorityPThirdPartyPartnerOverDefender is a pointer to MobileThreatDefensePartnerPriorityVThirdPartyPartnerOverDefender
	MobileThreatDefensePartnerPriorityPThirdPartyPartnerOverDefender = &_MobileThreatDefensePartnerPriorityPThirdPartyPartnerOverDefender
	// MobileThreatDefensePartnerPriorityPUnknownFutureValue is a pointer to MobileThreatDefensePartnerPriorityVUnknownFutureValue
	MobileThreatDefensePartnerPriorityPUnknownFutureValue = &_MobileThreatDefensePartnerPriorityPUnknownFutureValue
)

var (
	_MobileThreatDefensePartnerPriorityPDefenderOverThirdPartyPartner = MobileThreatDefensePartnerPriorityVDefenderOverThirdPartyPartner
	_MobileThreatDefensePartnerPriorityPThirdPartyPartnerOverDefender = MobileThreatDefensePartnerPriorityVThirdPartyPartnerOverDefender
	_MobileThreatDefensePartnerPriorityPUnknownFutureValue            = MobileThreatDefensePartnerPriorityVUnknownFutureValue
)

// MobileThreatPartnerTenantState undocumented
type MobileThreatPartnerTenantState string

const (
	// MobileThreatPartnerTenantStateVUnavailable undocumented
	MobileThreatPartnerTenantStateVUnavailable MobileThreatPartnerTenantState = "unavailable"
	// MobileThreatPartnerTenantStateVAvailable undocumented
	MobileThreatPartnerTenantStateVAvailable MobileThreatPartnerTenantState = "available"
	// MobileThreatPartnerTenantStateVEnabled undocumented
	MobileThreatPartnerTenantStateVEnabled MobileThreatPartnerTenantState = "enabled"
	// MobileThreatPartnerTenantStateVUnresponsive undocumented
	MobileThreatPartnerTenantStateVUnresponsive MobileThreatPartnerTenantState = "unresponsive"
)

var (
	// MobileThreatPartnerTenantStatePUnavailable is a pointer to MobileThreatPartnerTenantStateVUnavailable
	MobileThreatPartnerTenantStatePUnavailable = &_MobileThreatPartnerTenantStatePUnavailable
	// MobileThreatPartnerTenantStatePAvailable is a pointer to MobileThreatPartnerTenantStateVAvailable
	MobileThreatPartnerTenantStatePAvailable = &_MobileThreatPartnerTenantStatePAvailable
	// MobileThreatPartnerTenantStatePEnabled is a pointer to MobileThreatPartnerTenantStateVEnabled
	MobileThreatPartnerTenantStatePEnabled = &_MobileThreatPartnerTenantStatePEnabled
	// MobileThreatPartnerTenantStatePUnresponsive is a pointer to MobileThreatPartnerTenantStateVUnresponsive
	MobileThreatPartnerTenantStatePUnresponsive = &_MobileThreatPartnerTenantStatePUnresponsive
)

var (
	_MobileThreatPartnerTenantStatePUnavailable  = MobileThreatPartnerTenantStateVUnavailable
	_MobileThreatPartnerTenantStatePAvailable    = MobileThreatPartnerTenantStateVAvailable
	_MobileThreatPartnerTenantStatePEnabled      = MobileThreatPartnerTenantStateVEnabled
	_MobileThreatPartnerTenantStatePUnresponsive = MobileThreatPartnerTenantStateVUnresponsive
)