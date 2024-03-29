// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// TranslationLanguageOverride undocumented
type TranslationLanguageOverride struct {
	// Object is the base model of TranslationLanguageOverride
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// LanguageTag undocumented
	LanguageTag *string `json:"languageTag,omitempty"`
	// TranslationBehavior undocumented
	TranslationBehavior *TranslationBehavior `json:"translationBehavior,omitempty"`
}

func NewTranslationLanguageOverride() (*TranslationLanguageOverride, error) {
	newTranslationLanguageOverride := &TranslationLanguageOverride{
		ODataType: "#microsoft.graph.TranslationLanguageOverride",
	}
	return newTranslationLanguageOverride, nil
}

// TranslationPreferences undocumented
type TranslationPreferences struct {
	// Object is the base model of TranslationPreferences
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// LanguageOverrides undocumented
	LanguageOverrides []TranslationLanguageOverride `json:"languageOverrides,omitempty"`
	// TranslationBehavior undocumented
	TranslationBehavior *TranslationBehavior `json:"translationBehavior,omitempty"`
	// UntranslatedLanguages undocumented
	UntranslatedLanguages []string `json:"untranslatedLanguages,omitempty"`
}

func NewTranslationPreferences() (*TranslationPreferences, error) {
	newTranslationPreferences := &TranslationPreferences{
		ODataType: "#microsoft.graph.TranslationPreferences",
	}
	return newTranslationPreferences, nil
}
