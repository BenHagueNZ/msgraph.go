// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ActiveDirectoryWindowsAutopilotDeploymentProfile undocumented
type ActiveDirectoryWindowsAutopilotDeploymentProfile struct {
	// WindowsAutopilotDeploymentProfile is the base model of ActiveDirectoryWindowsAutopilotDeploymentProfile
	WindowsAutopilotDeploymentProfile

	ODataType string `json:"@odata.type,omitempty"`
	// HybridAzureADJoinSkipConnectivityCheck undocumented
	HybridAzureADJoinSkipConnectivityCheck *bool `json:"hybridAzureADJoinSkipConnectivityCheck,omitempty"`
	// DomainJoinConfiguration undocumented
	DomainJoinConfiguration *WindowsDomainJoinConfiguration `json:"domainJoinConfiguration,omitempty"`
}

func NewActiveDirectoryWindowsAutopilotDeploymentProfile() (*ActiveDirectoryWindowsAutopilotDeploymentProfile, error) {
	newActiveDirectoryWindowsAutopilotDeploymentProfile := &ActiveDirectoryWindowsAutopilotDeploymentProfile{
		ODataType: "#microsoft.graph.ActiveDirectoryWindowsAutopilotDeploymentProfile",
	}
	return newActiveDirectoryWindowsAutopilotDeploymentProfile, nil
}
