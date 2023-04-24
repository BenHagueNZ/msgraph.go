// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ResultantAppState undocumented
type ResultantAppState string

const (
	// ResultantAppStateVNotApplicable undocumented
	ResultantAppStateVNotApplicable ResultantAppState = "notApplicable"
	// ResultantAppStateVInstalled undocumented
	ResultantAppStateVInstalled ResultantAppState = "installed"
	// ResultantAppStateVFailed undocumented
	ResultantAppStateVFailed ResultantAppState = "failed"
	// ResultantAppStateVNotInstalled undocumented
	ResultantAppStateVNotInstalled ResultantAppState = "notInstalled"
	// ResultantAppStateVUninstallFailed undocumented
	ResultantAppStateVUninstallFailed ResultantAppState = "uninstallFailed"
	// ResultantAppStateVPendingInstall undocumented
	ResultantAppStateVPendingInstall ResultantAppState = "pendingInstall"
	// ResultantAppStateVUnknown undocumented
	ResultantAppStateVUnknown ResultantAppState = "unknown"
)

var (
	// ResultantAppStatePNotApplicable is a pointer to ResultantAppStateVNotApplicable
	ResultantAppStatePNotApplicable = &_ResultantAppStatePNotApplicable
	// ResultantAppStatePInstalled is a pointer to ResultantAppStateVInstalled
	ResultantAppStatePInstalled = &_ResultantAppStatePInstalled
	// ResultantAppStatePFailed is a pointer to ResultantAppStateVFailed
	ResultantAppStatePFailed = &_ResultantAppStatePFailed
	// ResultantAppStatePNotInstalled is a pointer to ResultantAppStateVNotInstalled
	ResultantAppStatePNotInstalled = &_ResultantAppStatePNotInstalled
	// ResultantAppStatePUninstallFailed is a pointer to ResultantAppStateVUninstallFailed
	ResultantAppStatePUninstallFailed = &_ResultantAppStatePUninstallFailed
	// ResultantAppStatePPendingInstall is a pointer to ResultantAppStateVPendingInstall
	ResultantAppStatePPendingInstall = &_ResultantAppStatePPendingInstall
	// ResultantAppStatePUnknown is a pointer to ResultantAppStateVUnknown
	ResultantAppStatePUnknown = &_ResultantAppStatePUnknown
)

var (
	_ResultantAppStatePNotApplicable   = ResultantAppStateVNotApplicable
	_ResultantAppStatePInstalled       = ResultantAppStateVInstalled
	_ResultantAppStatePFailed          = ResultantAppStateVFailed
	_ResultantAppStatePNotInstalled    = ResultantAppStateVNotInstalled
	_ResultantAppStatePUninstallFailed = ResultantAppStateVUninstallFailed
	_ResultantAppStatePPendingInstall  = ResultantAppStateVPendingInstall
	_ResultantAppStatePUnknown         = ResultantAppStateVUnknown
)

// ResultantAppStateDetail undocumented
type ResultantAppStateDetail string

const (
	// ResultantAppStateDetailVProcessorArchitectureNotApplicable undocumented
	ResultantAppStateDetailVProcessorArchitectureNotApplicable ResultantAppStateDetail = "processorArchitectureNotApplicable"
	// ResultantAppStateDetailVMinimumDiskSpaceNotMet undocumented
	ResultantAppStateDetailVMinimumDiskSpaceNotMet ResultantAppStateDetail = "minimumDiskSpaceNotMet"
	// ResultantAppStateDetailVMinimumOsVersionNotMet undocumented
	ResultantAppStateDetailVMinimumOsVersionNotMet ResultantAppStateDetail = "minimumOsVersionNotMet"
	// ResultantAppStateDetailVMinimumPhysicalMemoryNotMet undocumented
	ResultantAppStateDetailVMinimumPhysicalMemoryNotMet ResultantAppStateDetail = "minimumPhysicalMemoryNotMet"
	// ResultantAppStateDetailVMinimumLogicalProcessorCountNotMet undocumented
	ResultantAppStateDetailVMinimumLogicalProcessorCountNotMet ResultantAppStateDetail = "minimumLogicalProcessorCountNotMet"
	// ResultantAppStateDetailVMinimumCPUSpeedNotMet undocumented
	ResultantAppStateDetailVMinimumCPUSpeedNotMet ResultantAppStateDetail = "minimumCpuSpeedNotMet"
	// ResultantAppStateDetailVPlatformNotApplicable undocumented
	ResultantAppStateDetailVPlatformNotApplicable ResultantAppStateDetail = "platformNotApplicable"
	// ResultantAppStateDetailVFileSystemRequirementNotMet undocumented
	ResultantAppStateDetailVFileSystemRequirementNotMet ResultantAppStateDetail = "fileSystemRequirementNotMet"
	// ResultantAppStateDetailVRegistryRequirementNotMet undocumented
	ResultantAppStateDetailVRegistryRequirementNotMet ResultantAppStateDetail = "registryRequirementNotMet"
	// ResultantAppStateDetailVPowerShellScriptRequirementNotMet undocumented
	ResultantAppStateDetailVPowerShellScriptRequirementNotMet ResultantAppStateDetail = "powerShellScriptRequirementNotMet"
	// ResultantAppStateDetailVSupersedingAppsNotApplicable undocumented
	ResultantAppStateDetailVSupersedingAppsNotApplicable ResultantAppStateDetail = "supersedingAppsNotApplicable"
	// ResultantAppStateDetailVNoAdditionalDetails undocumented
	ResultantAppStateDetailVNoAdditionalDetails ResultantAppStateDetail = "noAdditionalDetails"
	// ResultantAppStateDetailVDependencyFailedToInstall undocumented
	ResultantAppStateDetailVDependencyFailedToInstall ResultantAppStateDetail = "dependencyFailedToInstall"
	// ResultantAppStateDetailVDependencyWithRequirementsNotMet undocumented
	ResultantAppStateDetailVDependencyWithRequirementsNotMet ResultantAppStateDetail = "dependencyWithRequirementsNotMet"
	// ResultantAppStateDetailVDependencyPendingReboot undocumented
	ResultantAppStateDetailVDependencyPendingReboot ResultantAppStateDetail = "dependencyPendingReboot"
	// ResultantAppStateDetailVDependencyWithAutoInstallDisabled undocumented
	ResultantAppStateDetailVDependencyWithAutoInstallDisabled ResultantAppStateDetail = "dependencyWithAutoInstallDisabled"
	// ResultantAppStateDetailVSupersededAppUninstallFailed undocumented
	ResultantAppStateDetailVSupersededAppUninstallFailed ResultantAppStateDetail = "supersededAppUninstallFailed"
	// ResultantAppStateDetailVSupersededAppUninstallPendingReboot undocumented
	ResultantAppStateDetailVSupersededAppUninstallPendingReboot ResultantAppStateDetail = "supersededAppUninstallPendingReboot"
	// ResultantAppStateDetailVRemovingSupersededApps undocumented
	ResultantAppStateDetailVRemovingSupersededApps ResultantAppStateDetail = "removingSupersededApps"
	// ResultantAppStateDetailVIOSAppStoreUpdateFailedToInstall undocumented
	ResultantAppStateDetailVIOSAppStoreUpdateFailedToInstall ResultantAppStateDetail = "iosAppStoreUpdateFailedToInstall"
	// ResultantAppStateDetailVVPPAppHasUpdateAvailable undocumented
	ResultantAppStateDetailVVPPAppHasUpdateAvailable ResultantAppStateDetail = "vppAppHasUpdateAvailable"
	// ResultantAppStateDetailVUserRejectedUpdate undocumented
	ResultantAppStateDetailVUserRejectedUpdate ResultantAppStateDetail = "userRejectedUpdate"
	// ResultantAppStateDetailVUninstallPendingReboot undocumented
	ResultantAppStateDetailVUninstallPendingReboot ResultantAppStateDetail = "uninstallPendingReboot"
	// ResultantAppStateDetailVSupersedingAppsDetected undocumented
	ResultantAppStateDetailVSupersedingAppsDetected ResultantAppStateDetail = "supersedingAppsDetected"
	// ResultantAppStateDetailVSupersededAppsDetected undocumented
	ResultantAppStateDetailVSupersededAppsDetected ResultantAppStateDetail = "supersededAppsDetected"
	// ResultantAppStateDetailVSeeInstallErrorCode undocumented
	ResultantAppStateDetailVSeeInstallErrorCode ResultantAppStateDetail = "seeInstallErrorCode"
	// ResultantAppStateDetailVAutoInstallDisabled undocumented
	ResultantAppStateDetailVAutoInstallDisabled ResultantAppStateDetail = "autoInstallDisabled"
	// ResultantAppStateDetailVManagedAppNoLongerPresent undocumented
	ResultantAppStateDetailVManagedAppNoLongerPresent ResultantAppStateDetail = "managedAppNoLongerPresent"
	// ResultantAppStateDetailVUserRejectedInstall undocumented
	ResultantAppStateDetailVUserRejectedInstall ResultantAppStateDetail = "userRejectedInstall"
	// ResultantAppStateDetailVUserIsNotLoggedIntoAppStore undocumented
	ResultantAppStateDetailVUserIsNotLoggedIntoAppStore ResultantAppStateDetail = "userIsNotLoggedIntoAppStore"
	// ResultantAppStateDetailVUntargetedSupersedingAppsDetected undocumented
	ResultantAppStateDetailVUntargetedSupersedingAppsDetected ResultantAppStateDetail = "untargetedSupersedingAppsDetected"
	// ResultantAppStateDetailVAppRemovedBySupersedence undocumented
	ResultantAppStateDetailVAppRemovedBySupersedence ResultantAppStateDetail = "appRemovedBySupersedence"
	// ResultantAppStateDetailVSeeUninstallErrorCode undocumented
	ResultantAppStateDetailVSeeUninstallErrorCode ResultantAppStateDetail = "seeUninstallErrorCode"
	// ResultantAppStateDetailVPendingReboot undocumented
	ResultantAppStateDetailVPendingReboot ResultantAppStateDetail = "pendingReboot"
	// ResultantAppStateDetailVInstallingDependencies undocumented
	ResultantAppStateDetailVInstallingDependencies ResultantAppStateDetail = "installingDependencies"
	// ResultantAppStateDetailVContentDownloaded undocumented
	ResultantAppStateDetailVContentDownloaded ResultantAppStateDetail = "contentDownloaded"
)

var (
	// ResultantAppStateDetailPProcessorArchitectureNotApplicable is a pointer to ResultantAppStateDetailVProcessorArchitectureNotApplicable
	ResultantAppStateDetailPProcessorArchitectureNotApplicable = &_ResultantAppStateDetailPProcessorArchitectureNotApplicable
	// ResultantAppStateDetailPMinimumDiskSpaceNotMet is a pointer to ResultantAppStateDetailVMinimumDiskSpaceNotMet
	ResultantAppStateDetailPMinimumDiskSpaceNotMet = &_ResultantAppStateDetailPMinimumDiskSpaceNotMet
	// ResultantAppStateDetailPMinimumOsVersionNotMet is a pointer to ResultantAppStateDetailVMinimumOsVersionNotMet
	ResultantAppStateDetailPMinimumOsVersionNotMet = &_ResultantAppStateDetailPMinimumOsVersionNotMet
	// ResultantAppStateDetailPMinimumPhysicalMemoryNotMet is a pointer to ResultantAppStateDetailVMinimumPhysicalMemoryNotMet
	ResultantAppStateDetailPMinimumPhysicalMemoryNotMet = &_ResultantAppStateDetailPMinimumPhysicalMemoryNotMet
	// ResultantAppStateDetailPMinimumLogicalProcessorCountNotMet is a pointer to ResultantAppStateDetailVMinimumLogicalProcessorCountNotMet
	ResultantAppStateDetailPMinimumLogicalProcessorCountNotMet = &_ResultantAppStateDetailPMinimumLogicalProcessorCountNotMet
	// ResultantAppStateDetailPMinimumCPUSpeedNotMet is a pointer to ResultantAppStateDetailVMinimumCPUSpeedNotMet
	ResultantAppStateDetailPMinimumCPUSpeedNotMet = &_ResultantAppStateDetailPMinimumCPUSpeedNotMet
	// ResultantAppStateDetailPPlatformNotApplicable is a pointer to ResultantAppStateDetailVPlatformNotApplicable
	ResultantAppStateDetailPPlatformNotApplicable = &_ResultantAppStateDetailPPlatformNotApplicable
	// ResultantAppStateDetailPFileSystemRequirementNotMet is a pointer to ResultantAppStateDetailVFileSystemRequirementNotMet
	ResultantAppStateDetailPFileSystemRequirementNotMet = &_ResultantAppStateDetailPFileSystemRequirementNotMet
	// ResultantAppStateDetailPRegistryRequirementNotMet is a pointer to ResultantAppStateDetailVRegistryRequirementNotMet
	ResultantAppStateDetailPRegistryRequirementNotMet = &_ResultantAppStateDetailPRegistryRequirementNotMet
	// ResultantAppStateDetailPPowerShellScriptRequirementNotMet is a pointer to ResultantAppStateDetailVPowerShellScriptRequirementNotMet
	ResultantAppStateDetailPPowerShellScriptRequirementNotMet = &_ResultantAppStateDetailPPowerShellScriptRequirementNotMet
	// ResultantAppStateDetailPSupersedingAppsNotApplicable is a pointer to ResultantAppStateDetailVSupersedingAppsNotApplicable
	ResultantAppStateDetailPSupersedingAppsNotApplicable = &_ResultantAppStateDetailPSupersedingAppsNotApplicable
	// ResultantAppStateDetailPNoAdditionalDetails is a pointer to ResultantAppStateDetailVNoAdditionalDetails
	ResultantAppStateDetailPNoAdditionalDetails = &_ResultantAppStateDetailPNoAdditionalDetails
	// ResultantAppStateDetailPDependencyFailedToInstall is a pointer to ResultantAppStateDetailVDependencyFailedToInstall
	ResultantAppStateDetailPDependencyFailedToInstall = &_ResultantAppStateDetailPDependencyFailedToInstall
	// ResultantAppStateDetailPDependencyWithRequirementsNotMet is a pointer to ResultantAppStateDetailVDependencyWithRequirementsNotMet
	ResultantAppStateDetailPDependencyWithRequirementsNotMet = &_ResultantAppStateDetailPDependencyWithRequirementsNotMet
	// ResultantAppStateDetailPDependencyPendingReboot is a pointer to ResultantAppStateDetailVDependencyPendingReboot
	ResultantAppStateDetailPDependencyPendingReboot = &_ResultantAppStateDetailPDependencyPendingReboot
	// ResultantAppStateDetailPDependencyWithAutoInstallDisabled is a pointer to ResultantAppStateDetailVDependencyWithAutoInstallDisabled
	ResultantAppStateDetailPDependencyWithAutoInstallDisabled = &_ResultantAppStateDetailPDependencyWithAutoInstallDisabled
	// ResultantAppStateDetailPSupersededAppUninstallFailed is a pointer to ResultantAppStateDetailVSupersededAppUninstallFailed
	ResultantAppStateDetailPSupersededAppUninstallFailed = &_ResultantAppStateDetailPSupersededAppUninstallFailed
	// ResultantAppStateDetailPSupersededAppUninstallPendingReboot is a pointer to ResultantAppStateDetailVSupersededAppUninstallPendingReboot
	ResultantAppStateDetailPSupersededAppUninstallPendingReboot = &_ResultantAppStateDetailPSupersededAppUninstallPendingReboot
	// ResultantAppStateDetailPRemovingSupersededApps is a pointer to ResultantAppStateDetailVRemovingSupersededApps
	ResultantAppStateDetailPRemovingSupersededApps = &_ResultantAppStateDetailPRemovingSupersededApps
	// ResultantAppStateDetailPIOSAppStoreUpdateFailedToInstall is a pointer to ResultantAppStateDetailVIOSAppStoreUpdateFailedToInstall
	ResultantAppStateDetailPIOSAppStoreUpdateFailedToInstall = &_ResultantAppStateDetailPIOSAppStoreUpdateFailedToInstall
	// ResultantAppStateDetailPVPPAppHasUpdateAvailable is a pointer to ResultantAppStateDetailVVPPAppHasUpdateAvailable
	ResultantAppStateDetailPVPPAppHasUpdateAvailable = &_ResultantAppStateDetailPVPPAppHasUpdateAvailable
	// ResultantAppStateDetailPUserRejectedUpdate is a pointer to ResultantAppStateDetailVUserRejectedUpdate
	ResultantAppStateDetailPUserRejectedUpdate = &_ResultantAppStateDetailPUserRejectedUpdate
	// ResultantAppStateDetailPUninstallPendingReboot is a pointer to ResultantAppStateDetailVUninstallPendingReboot
	ResultantAppStateDetailPUninstallPendingReboot = &_ResultantAppStateDetailPUninstallPendingReboot
	// ResultantAppStateDetailPSupersedingAppsDetected is a pointer to ResultantAppStateDetailVSupersedingAppsDetected
	ResultantAppStateDetailPSupersedingAppsDetected = &_ResultantAppStateDetailPSupersedingAppsDetected
	// ResultantAppStateDetailPSupersededAppsDetected is a pointer to ResultantAppStateDetailVSupersededAppsDetected
	ResultantAppStateDetailPSupersededAppsDetected = &_ResultantAppStateDetailPSupersededAppsDetected
	// ResultantAppStateDetailPSeeInstallErrorCode is a pointer to ResultantAppStateDetailVSeeInstallErrorCode
	ResultantAppStateDetailPSeeInstallErrorCode = &_ResultantAppStateDetailPSeeInstallErrorCode
	// ResultantAppStateDetailPAutoInstallDisabled is a pointer to ResultantAppStateDetailVAutoInstallDisabled
	ResultantAppStateDetailPAutoInstallDisabled = &_ResultantAppStateDetailPAutoInstallDisabled
	// ResultantAppStateDetailPManagedAppNoLongerPresent is a pointer to ResultantAppStateDetailVManagedAppNoLongerPresent
	ResultantAppStateDetailPManagedAppNoLongerPresent = &_ResultantAppStateDetailPManagedAppNoLongerPresent
	// ResultantAppStateDetailPUserRejectedInstall is a pointer to ResultantAppStateDetailVUserRejectedInstall
	ResultantAppStateDetailPUserRejectedInstall = &_ResultantAppStateDetailPUserRejectedInstall
	// ResultantAppStateDetailPUserIsNotLoggedIntoAppStore is a pointer to ResultantAppStateDetailVUserIsNotLoggedIntoAppStore
	ResultantAppStateDetailPUserIsNotLoggedIntoAppStore = &_ResultantAppStateDetailPUserIsNotLoggedIntoAppStore
	// ResultantAppStateDetailPUntargetedSupersedingAppsDetected is a pointer to ResultantAppStateDetailVUntargetedSupersedingAppsDetected
	ResultantAppStateDetailPUntargetedSupersedingAppsDetected = &_ResultantAppStateDetailPUntargetedSupersedingAppsDetected
	// ResultantAppStateDetailPAppRemovedBySupersedence is a pointer to ResultantAppStateDetailVAppRemovedBySupersedence
	ResultantAppStateDetailPAppRemovedBySupersedence = &_ResultantAppStateDetailPAppRemovedBySupersedence
	// ResultantAppStateDetailPSeeUninstallErrorCode is a pointer to ResultantAppStateDetailVSeeUninstallErrorCode
	ResultantAppStateDetailPSeeUninstallErrorCode = &_ResultantAppStateDetailPSeeUninstallErrorCode
	// ResultantAppStateDetailPPendingReboot is a pointer to ResultantAppStateDetailVPendingReboot
	ResultantAppStateDetailPPendingReboot = &_ResultantAppStateDetailPPendingReboot
	// ResultantAppStateDetailPInstallingDependencies is a pointer to ResultantAppStateDetailVInstallingDependencies
	ResultantAppStateDetailPInstallingDependencies = &_ResultantAppStateDetailPInstallingDependencies
	// ResultantAppStateDetailPContentDownloaded is a pointer to ResultantAppStateDetailVContentDownloaded
	ResultantAppStateDetailPContentDownloaded = &_ResultantAppStateDetailPContentDownloaded
)

var (
	_ResultantAppStateDetailPProcessorArchitectureNotApplicable  = ResultantAppStateDetailVProcessorArchitectureNotApplicable
	_ResultantAppStateDetailPMinimumDiskSpaceNotMet              = ResultantAppStateDetailVMinimumDiskSpaceNotMet
	_ResultantAppStateDetailPMinimumOsVersionNotMet              = ResultantAppStateDetailVMinimumOsVersionNotMet
	_ResultantAppStateDetailPMinimumPhysicalMemoryNotMet         = ResultantAppStateDetailVMinimumPhysicalMemoryNotMet
	_ResultantAppStateDetailPMinimumLogicalProcessorCountNotMet  = ResultantAppStateDetailVMinimumLogicalProcessorCountNotMet
	_ResultantAppStateDetailPMinimumCPUSpeedNotMet               = ResultantAppStateDetailVMinimumCPUSpeedNotMet
	_ResultantAppStateDetailPPlatformNotApplicable               = ResultantAppStateDetailVPlatformNotApplicable
	_ResultantAppStateDetailPFileSystemRequirementNotMet         = ResultantAppStateDetailVFileSystemRequirementNotMet
	_ResultantAppStateDetailPRegistryRequirementNotMet           = ResultantAppStateDetailVRegistryRequirementNotMet
	_ResultantAppStateDetailPPowerShellScriptRequirementNotMet   = ResultantAppStateDetailVPowerShellScriptRequirementNotMet
	_ResultantAppStateDetailPSupersedingAppsNotApplicable        = ResultantAppStateDetailVSupersedingAppsNotApplicable
	_ResultantAppStateDetailPNoAdditionalDetails                 = ResultantAppStateDetailVNoAdditionalDetails
	_ResultantAppStateDetailPDependencyFailedToInstall           = ResultantAppStateDetailVDependencyFailedToInstall
	_ResultantAppStateDetailPDependencyWithRequirementsNotMet    = ResultantAppStateDetailVDependencyWithRequirementsNotMet
	_ResultantAppStateDetailPDependencyPendingReboot             = ResultantAppStateDetailVDependencyPendingReboot
	_ResultantAppStateDetailPDependencyWithAutoInstallDisabled   = ResultantAppStateDetailVDependencyWithAutoInstallDisabled
	_ResultantAppStateDetailPSupersededAppUninstallFailed        = ResultantAppStateDetailVSupersededAppUninstallFailed
	_ResultantAppStateDetailPSupersededAppUninstallPendingReboot = ResultantAppStateDetailVSupersededAppUninstallPendingReboot
	_ResultantAppStateDetailPRemovingSupersededApps              = ResultantAppStateDetailVRemovingSupersededApps
	_ResultantAppStateDetailPIOSAppStoreUpdateFailedToInstall    = ResultantAppStateDetailVIOSAppStoreUpdateFailedToInstall
	_ResultantAppStateDetailPVPPAppHasUpdateAvailable            = ResultantAppStateDetailVVPPAppHasUpdateAvailable
	_ResultantAppStateDetailPUserRejectedUpdate                  = ResultantAppStateDetailVUserRejectedUpdate
	_ResultantAppStateDetailPUninstallPendingReboot              = ResultantAppStateDetailVUninstallPendingReboot
	_ResultantAppStateDetailPSupersedingAppsDetected             = ResultantAppStateDetailVSupersedingAppsDetected
	_ResultantAppStateDetailPSupersededAppsDetected              = ResultantAppStateDetailVSupersededAppsDetected
	_ResultantAppStateDetailPSeeInstallErrorCode                 = ResultantAppStateDetailVSeeInstallErrorCode
	_ResultantAppStateDetailPAutoInstallDisabled                 = ResultantAppStateDetailVAutoInstallDisabled
	_ResultantAppStateDetailPManagedAppNoLongerPresent           = ResultantAppStateDetailVManagedAppNoLongerPresent
	_ResultantAppStateDetailPUserRejectedInstall                 = ResultantAppStateDetailVUserRejectedInstall
	_ResultantAppStateDetailPUserIsNotLoggedIntoAppStore         = ResultantAppStateDetailVUserIsNotLoggedIntoAppStore
	_ResultantAppStateDetailPUntargetedSupersedingAppsDetected   = ResultantAppStateDetailVUntargetedSupersedingAppsDetected
	_ResultantAppStateDetailPAppRemovedBySupersedence            = ResultantAppStateDetailVAppRemovedBySupersedence
	_ResultantAppStateDetailPSeeUninstallErrorCode               = ResultantAppStateDetailVSeeUninstallErrorCode
	_ResultantAppStateDetailPPendingReboot                       = ResultantAppStateDetailVPendingReboot
	_ResultantAppStateDetailPInstallingDependencies              = ResultantAppStateDetailVInstallingDependencies
	_ResultantAppStateDetailPContentDownloaded                   = ResultantAppStateDetailVContentDownloaded
)
