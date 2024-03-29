// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// MDMWindowsInformationProtectionPolicy undocumented
type MDMWindowsInformationProtectionPolicy struct {
	// WindowsInformationProtection is the base model of MDMWindowsInformationProtectionPolicy
	WindowsInformationProtection

	ODataType string `json:"@odata.type,omitempty"`
}

func NewMDMWindowsInformationProtectionPolicy() (*MDMWindowsInformationProtectionPolicy, error) {
	newMDMWindowsInformationProtectionPolicy := &MDMWindowsInformationProtectionPolicy{
		ODataType: "#microsoft.graph.MdmWindowsInformationProtectionPolicy",
	}
	return newMDMWindowsInformationProtectionPolicy, nil
}

// MDMWindowsInformationProtectionPolicyPolicySetItem undocumented
type MDMWindowsInformationProtectionPolicyPolicySetItem struct {
	// PolicySetItem is the base model of MDMWindowsInformationProtectionPolicyPolicySetItem
	PolicySetItem

	ODataType string `json:"@odata.type,omitempty"`
}

func NewMDMWindowsInformationProtectionPolicyPolicySetItem() (*MDMWindowsInformationProtectionPolicyPolicySetItem, error) {
	newMDMWindowsInformationProtectionPolicyPolicySetItem := &MDMWindowsInformationProtectionPolicyPolicySetItem{
		ODataType: "#microsoft.graph.MdmWindowsInformationProtectionPolicyPolicySetItem",
	}
	return newMDMWindowsInformationProtectionPolicyPolicySetItem, nil
}
