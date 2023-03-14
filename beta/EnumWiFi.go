
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// WiFiAuthenticationMethod undocumented
type WiFiAuthenticationMethod string

const (
    // WiFiAuthenticationMethodVCertificate undocumented
    WiFiAuthenticationMethodVCertificate WiFiAuthenticationMethod = "certificate"
    // WiFiAuthenticationMethodVUsernameAndPassword undocumented
    WiFiAuthenticationMethodVUsernameAndPassword WiFiAuthenticationMethod = "usernameAndPassword"
    // WiFiAuthenticationMethodVDerivedCredential undocumented
    WiFiAuthenticationMethodVDerivedCredential WiFiAuthenticationMethod = "derivedCredential"
)

var (
    // WiFiAuthenticationMethodPCertificate is a pointer to WiFiAuthenticationMethodVCertificate
    WiFiAuthenticationMethodPCertificate = &_WiFiAuthenticationMethodPCertificate
    // WiFiAuthenticationMethodPUsernameAndPassword is a pointer to WiFiAuthenticationMethodVUsernameAndPassword
    WiFiAuthenticationMethodPUsernameAndPassword = &_WiFiAuthenticationMethodPUsernameAndPassword
    // WiFiAuthenticationMethodPDerivedCredential is a pointer to WiFiAuthenticationMethodVDerivedCredential
    WiFiAuthenticationMethodPDerivedCredential = &_WiFiAuthenticationMethodPDerivedCredential
)

var (
    _WiFiAuthenticationMethodPCertificate = WiFiAuthenticationMethodVCertificate
    _WiFiAuthenticationMethodPUsernameAndPassword = WiFiAuthenticationMethodVUsernameAndPassword
    _WiFiAuthenticationMethodPDerivedCredential = WiFiAuthenticationMethodVDerivedCredential
)

// WiFiProxySetting undocumented
type WiFiProxySetting string

const (
    // WiFiProxySettingVNone undocumented
    WiFiProxySettingVNone WiFiProxySetting = "none"
    // WiFiProxySettingVManual undocumented
    WiFiProxySettingVManual WiFiProxySetting = "manual"
    // WiFiProxySettingVAutomatic undocumented
    WiFiProxySettingVAutomatic WiFiProxySetting = "automatic"
)

var (
    // WiFiProxySettingPNone is a pointer to WiFiProxySettingVNone
    WiFiProxySettingPNone = &_WiFiProxySettingPNone
    // WiFiProxySettingPManual is a pointer to WiFiProxySettingVManual
    WiFiProxySettingPManual = &_WiFiProxySettingPManual
    // WiFiProxySettingPAutomatic is a pointer to WiFiProxySettingVAutomatic
    WiFiProxySettingPAutomatic = &_WiFiProxySettingPAutomatic
)

var (
    _WiFiProxySettingPNone = WiFiProxySettingVNone
    _WiFiProxySettingPManual = WiFiProxySettingVManual
    _WiFiProxySettingPAutomatic = WiFiProxySettingVAutomatic
)

// WiFiSecurityType undocumented
type WiFiSecurityType string

const (
    // WiFiSecurityTypeVOpen undocumented
    WiFiSecurityTypeVOpen WiFiSecurityType = "open"
    // WiFiSecurityTypeVWpaPersonal undocumented
    WiFiSecurityTypeVWpaPersonal WiFiSecurityType = "wpaPersonal"
    // WiFiSecurityTypeVWpaEnterprise undocumented
    WiFiSecurityTypeVWpaEnterprise WiFiSecurityType = "wpaEnterprise"
    // WiFiSecurityTypeVWep undocumented
    WiFiSecurityTypeVWep WiFiSecurityType = "wep"
    // WiFiSecurityTypeVWpa2Personal undocumented
    WiFiSecurityTypeVWpa2Personal WiFiSecurityType = "wpa2Personal"
    // WiFiSecurityTypeVWpa2Enterprise undocumented
    WiFiSecurityTypeVWpa2Enterprise WiFiSecurityType = "wpa2Enterprise"
)

var (
    // WiFiSecurityTypePOpen is a pointer to WiFiSecurityTypeVOpen
    WiFiSecurityTypePOpen = &_WiFiSecurityTypePOpen
    // WiFiSecurityTypePWpaPersonal is a pointer to WiFiSecurityTypeVWpaPersonal
    WiFiSecurityTypePWpaPersonal = &_WiFiSecurityTypePWpaPersonal
    // WiFiSecurityTypePWpaEnterprise is a pointer to WiFiSecurityTypeVWpaEnterprise
    WiFiSecurityTypePWpaEnterprise = &_WiFiSecurityTypePWpaEnterprise
    // WiFiSecurityTypePWep is a pointer to WiFiSecurityTypeVWep
    WiFiSecurityTypePWep = &_WiFiSecurityTypePWep
    // WiFiSecurityTypePWpa2Personal is a pointer to WiFiSecurityTypeVWpa2Personal
    WiFiSecurityTypePWpa2Personal = &_WiFiSecurityTypePWpa2Personal
    // WiFiSecurityTypePWpa2Enterprise is a pointer to WiFiSecurityTypeVWpa2Enterprise
    WiFiSecurityTypePWpa2Enterprise = &_WiFiSecurityTypePWpa2Enterprise
)

var (
    _WiFiSecurityTypePOpen = WiFiSecurityTypeVOpen
    _WiFiSecurityTypePWpaPersonal = WiFiSecurityTypeVWpaPersonal
    _WiFiSecurityTypePWpaEnterprise = WiFiSecurityTypeVWpaEnterprise
    _WiFiSecurityTypePWep = WiFiSecurityTypeVWep
    _WiFiSecurityTypePWpa2Personal = WiFiSecurityTypeVWpa2Personal
    _WiFiSecurityTypePWpa2Enterprise = WiFiSecurityTypeVWpa2Enterprise
)

// WiFiAuthenticationType undocumented
type WiFiAuthenticationType string

const (
    // WiFiAuthenticationTypeVNone undocumented
    WiFiAuthenticationTypeVNone WiFiAuthenticationType = "none"
    // WiFiAuthenticationTypeVUser undocumented
    WiFiAuthenticationTypeVUser WiFiAuthenticationType = "user"
    // WiFiAuthenticationTypeVMachine undocumented
    WiFiAuthenticationTypeVMachine WiFiAuthenticationType = "machine"
    // WiFiAuthenticationTypeVMachineOrUser undocumented
    WiFiAuthenticationTypeVMachineOrUser WiFiAuthenticationType = "machineOrUser"
    // WiFiAuthenticationTypeVGuest undocumented
    WiFiAuthenticationTypeVGuest WiFiAuthenticationType = "guest"
)

var (
    // WiFiAuthenticationTypePNone is a pointer to WiFiAuthenticationTypeVNone
    WiFiAuthenticationTypePNone = &_WiFiAuthenticationTypePNone
    // WiFiAuthenticationTypePUser is a pointer to WiFiAuthenticationTypeVUser
    WiFiAuthenticationTypePUser = &_WiFiAuthenticationTypePUser
    // WiFiAuthenticationTypePMachine is a pointer to WiFiAuthenticationTypeVMachine
    WiFiAuthenticationTypePMachine = &_WiFiAuthenticationTypePMachine
    // WiFiAuthenticationTypePMachineOrUser is a pointer to WiFiAuthenticationTypeVMachineOrUser
    WiFiAuthenticationTypePMachineOrUser = &_WiFiAuthenticationTypePMachineOrUser
    // WiFiAuthenticationTypePGuest is a pointer to WiFiAuthenticationTypeVGuest
    WiFiAuthenticationTypePGuest = &_WiFiAuthenticationTypePGuest
)

var (
    _WiFiAuthenticationTypePNone = WiFiAuthenticationTypeVNone
    _WiFiAuthenticationTypePUser = WiFiAuthenticationTypeVUser
    _WiFiAuthenticationTypePMachine = WiFiAuthenticationTypeVMachine
    _WiFiAuthenticationTypePMachineOrUser = WiFiAuthenticationTypeVMachineOrUser
    _WiFiAuthenticationTypePGuest = WiFiAuthenticationTypeVGuest
)
