
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// EBookInstallSummary undocumented
type EBookInstallSummary struct {
    // Entity is the base model of EBookInstallSummary
    Entity
    // FailedDeviceCount undocumented
    FailedDeviceCount *int `json:"failedDeviceCount,omitempty"`
    // FailedUserCount undocumented
    FailedUserCount *int `json:"failedUserCount,omitempty"`
    // InstalledDeviceCount undocumented
    InstalledDeviceCount *int `json:"installedDeviceCount,omitempty"`
    // InstalledUserCount undocumented
    InstalledUserCount *int `json:"installedUserCount,omitempty"`
    // NotInstalledDeviceCount undocumented
    NotInstalledDeviceCount *int `json:"notInstalledDeviceCount,omitempty"`
    // NotInstalledUserCount undocumented
    NotInstalledUserCount *int `json:"notInstalledUserCount,omitempty"`
}
