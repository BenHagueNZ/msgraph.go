// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// SettingSource undocumented
type SettingSource struct {
	// Object is the base model of SettingSource
	Object

	ODataType string `json:"@odata.type"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// SourceType undocumented
	SourceType *SettingSourceType `json:"sourceType,omitempty"`
}

func NewSettingSource() (*SettingSource, error) {
	newSettingSource := &SettingSource{
		ODataType: "#microsoft.graph.SettingSource",
	}
	return newSettingSource, nil
}

// SettingStateDeviceSummary undocumented
type SettingStateDeviceSummary struct {
	// Entity is the base model of SettingStateDeviceSummary
	Entity

	ODataType string `json:"@odata.type"`
	// CompliantDeviceCount undocumented
	CompliantDeviceCount *int `json:"compliantDeviceCount,omitempty"`
	// ConflictDeviceCount undocumented
	ConflictDeviceCount *int `json:"conflictDeviceCount,omitempty"`
	// ErrorDeviceCount undocumented
	ErrorDeviceCount *int `json:"errorDeviceCount,omitempty"`
	// InstancePath undocumented
	InstancePath *string `json:"instancePath,omitempty"`
	// NonCompliantDeviceCount undocumented
	NonCompliantDeviceCount *int `json:"nonCompliantDeviceCount,omitempty"`
	// NotApplicableDeviceCount undocumented
	NotApplicableDeviceCount *int `json:"notApplicableDeviceCount,omitempty"`
	// RemediatedDeviceCount undocumented
	RemediatedDeviceCount *int `json:"remediatedDeviceCount,omitempty"`
	// SettingName undocumented
	SettingName *string `json:"settingName,omitempty"`
	// UnknownDeviceCount undocumented
	UnknownDeviceCount *int `json:"unknownDeviceCount,omitempty"`
}

func NewSettingStateDeviceSummary() (*SettingStateDeviceSummary, error) {
	newSettingStateDeviceSummary := &SettingStateDeviceSummary{
		ODataType: "#microsoft.graph.SettingStateDeviceSummary",
	}
	return newSettingStateDeviceSummary, nil
}

// SettingTemplateValue undocumented
type SettingTemplateValue struct {
	// Object is the base model of SettingTemplateValue
	Object

	ODataType string `json:"@odata.type"`
	// DefaultValue undocumented
	DefaultValue *string `json:"defaultValue,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
}

func NewSettingTemplateValue() (*SettingTemplateValue, error) {
	newSettingTemplateValue := &SettingTemplateValue{
		ODataType: "#microsoft.graph.SettingTemplateValue",
	}
	return newSettingTemplateValue, nil
}

// SettingValue undocumented
type SettingValue struct {
	// Object is the base model of SettingValue
	Object

	ODataType string `json:"@odata.type"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Value undocumented
	Value *string `json:"value,omitempty"`
}

func NewSettingValue() (*SettingValue, error) {
	newSettingValue := &SettingValue{
		ODataType: "#microsoft.graph.SettingValue",
	}
	return newSettingValue, nil
}
