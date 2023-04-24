// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "encoding/json"

// ResultInfo undocumented
type ResultInfo struct {
	// Object is the base model of ResultInfo
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Code undocumented
	Code *int `json:"code,omitempty"`
	// Message undocumented
	Message *string `json:"message,omitempty"`
	// Subcode undocumented
	Subcode *int `json:"subcode,omitempty"`
}

func NewResultInfo() (*ResultInfo, error) {
	newResultInfo := &ResultInfo{
		ODataType: "#microsoft.graph.ResultInfo",
	}
	return newResultInfo, nil
}

// ResultTemplate undocumented
type ResultTemplate struct {
	// Object is the base model of ResultTemplate
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Body undocumented
	Body json.RawMessage `json:"body,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
}

func NewResultTemplate() (*ResultTemplate, error) {
	newResultTemplate := &ResultTemplate{
		ODataType: "#microsoft.graph.ResultTemplate",
	}
	return newResultTemplate, nil
}

// ResultTemplateDictionary undocumented
type ResultTemplateDictionary struct {
	// Dictionary is the base model of ResultTemplateDictionary
	Dictionary

	ODataType string `json:"@odata.type,omitempty"`
}

func NewResultTemplateDictionary() (*ResultTemplateDictionary, error) {
	newResultTemplateDictionary := &ResultTemplateDictionary{
		ODataType: "#microsoft.graph.ResultTemplateDictionary",
	}
	return newResultTemplateDictionary, nil
}

// ResultTemplateOption undocumented
type ResultTemplateOption struct {
	// Object is the base model of ResultTemplateOption
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// EnableResultTemplate undocumented
	EnableResultTemplate *bool `json:"enableResultTemplate,omitempty"`
}

func NewResultTemplateOption() (*ResultTemplateOption, error) {
	newResultTemplateOption := &ResultTemplateOption{
		ODataType: "#microsoft.graph.ResultTemplateOption",
	}
	return newResultTemplateOption, nil
}