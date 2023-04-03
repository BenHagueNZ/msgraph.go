// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// OMASetting undocumented
type OMASetting struct {
	// Object is the base model of OMASetting
	Object

	ODataType string `json:"@odata.type"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// OMAURI undocumented
	OMAURI *string `json:"omaUri,omitempty"`
}

func NewOMASetting() (*OMASetting, error) {
	newOMASetting := &OMASetting{
		ODataType: "#microsoft.graph.OmaSetting",
	}
	return newOMASetting, nil
}

// OMASettingBase64 undocumented
type OMASettingBase64 struct {
	// OMASetting is the base model of OMASettingBase64
	OMASetting

	ODataType string `json:"@odata.type"`
	// FileName undocumented
	FileName *string `json:"fileName,omitempty"`
	// Value undocumented
	Value *string `json:"value,omitempty"`
}

func NewOMASettingBase64() (*OMASettingBase64, error) {
	newOMASettingBase64 := &OMASettingBase64{
		ODataType: "#microsoft.graph.OmaSettingBase64",
	}
	return newOMASettingBase64, nil
}

// OMASettingBoolean undocumented
type OMASettingBoolean struct {
	// OMASetting is the base model of OMASettingBoolean
	OMASetting

	ODataType string `json:"@odata.type"`
	// Value undocumented
	Value *bool `json:"value,omitempty"`
}

func NewOMASettingBoolean() (*OMASettingBoolean, error) {
	newOMASettingBoolean := &OMASettingBoolean{
		ODataType: "#microsoft.graph.OmaSettingBoolean",
	}
	return newOMASettingBoolean, nil
}

// OMASettingDateTime undocumented
type OMASettingDateTime struct {
	// OMASetting is the base model of OMASettingDateTime
	OMASetting

	ODataType string `json:"@odata.type"`
	// Value undocumented
	Value *time.Time `json:"value,omitempty"`
}

func NewOMASettingDateTime() (*OMASettingDateTime, error) {
	newOMASettingDateTime := &OMASettingDateTime{
		ODataType: "#microsoft.graph.OmaSettingDateTime",
	}
	return newOMASettingDateTime, nil
}

// OMASettingFloatingPoint undocumented
type OMASettingFloatingPoint struct {
	// OMASetting is the base model of OMASettingFloatingPoint
	OMASetting

	ODataType string `json:"@odata.type"`
	// Value undocumented
	Value *float64 `json:"value,omitempty"`
}

func NewOMASettingFloatingPoint() (*OMASettingFloatingPoint, error) {
	newOMASettingFloatingPoint := &OMASettingFloatingPoint{
		ODataType: "#microsoft.graph.OmaSettingFloatingPoint",
	}
	return newOMASettingFloatingPoint, nil
}

// OMASettingInteger undocumented
type OMASettingInteger struct {
	// OMASetting is the base model of OMASettingInteger
	OMASetting

	ODataType string `json:"@odata.type"`
	// Value undocumented
	Value *int `json:"value,omitempty"`
}

func NewOMASettingInteger() (*OMASettingInteger, error) {
	newOMASettingInteger := &OMASettingInteger{
		ODataType: "#microsoft.graph.OmaSettingInteger",
	}
	return newOMASettingInteger, nil
}

// OMASettingString undocumented
type OMASettingString struct {
	// OMASetting is the base model of OMASettingString
	OMASetting

	ODataType string `json:"@odata.type"`
	// Value undocumented
	Value *string `json:"value,omitempty"`
}

func NewOMASettingString() (*OMASettingString, error) {
	newOMASettingString := &OMASettingString{
		ODataType: "#microsoft.graph.OmaSettingString",
	}
	return newOMASettingString, nil
}

// OMASettingStringXML undocumented
type OMASettingStringXML struct {
	// OMASetting is the base model of OMASettingStringXML
	OMASetting

	ODataType string `json:"@odata.type"`
	// FileName undocumented
	FileName *string `json:"fileName,omitempty"`
	// Value undocumented
	Value *Binary `json:"value,omitempty"`
}

func NewOMASettingStringXML() (*OMASettingStringXML, error) {
	newOMASettingStringXML := &OMASettingStringXML{
		ODataType: "#microsoft.graph.OmaSettingStringXml",
	}
	return newOMASettingStringXML, nil
}
