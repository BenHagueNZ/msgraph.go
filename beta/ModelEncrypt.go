// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// EncryptContent undocumented
type EncryptContent struct {
	// LabelActionBase is the base model of EncryptContent
	LabelActionBase

	ODataType string `json:"@odata.type,omitempty"`
	// EncryptWith undocumented
	EncryptWith *EncryptWith `json:"encryptWith,omitempty"`
}

func NewEncryptContent() (*EncryptContent, error) {
	newEncryptContent := &EncryptContent{
		ODataType: "#microsoft.graph.EncryptContent",
	}
	return newEncryptContent, nil
}

// EncryptWithTemplate undocumented
type EncryptWithTemplate struct {
	// EncryptContent is the base model of EncryptWithTemplate
	EncryptContent

	ODataType string `json:"@odata.type,omitempty"`
	// AvailableForEncryption undocumented
	AvailableForEncryption *bool `json:"availableForEncryption,omitempty"`
	// TemplateID undocumented
	TemplateID *string `json:"templateId,omitempty"`
}

func NewEncryptWithTemplate() (*EncryptWithTemplate, error) {
	newEncryptWithTemplate := &EncryptWithTemplate{
		ODataType: "#microsoft.graph.EncryptWithTemplate",
	}
	return newEncryptWithTemplate, nil
}

// EncryptWithUserDefinedRights undocumented
type EncryptWithUserDefinedRights struct {
	// EncryptContent is the base model of EncryptWithUserDefinedRights
	EncryptContent

	ODataType string `json:"@odata.type,omitempty"`
	// AllowAdHocPermissions undocumented
	AllowAdHocPermissions *bool `json:"allowAdHocPermissions,omitempty"`
	// AllowMailForwarding undocumented
	AllowMailForwarding *bool `json:"allowMailForwarding,omitempty"`
	// DecryptionRightsManagementTemplateID undocumented
	DecryptionRightsManagementTemplateID *string `json:"decryptionRightsManagementTemplateId,omitempty"`
}

func NewEncryptWithUserDefinedRights() (*EncryptWithUserDefinedRights, error) {
	newEncryptWithUserDefinedRights := &EncryptWithUserDefinedRights{
		ODataType: "#microsoft.graph.EncryptWithUserDefinedRights",
	}
	return newEncryptWithUserDefinedRights, nil
}
