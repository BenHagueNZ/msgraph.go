// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// TermColumn undocumented
type TermColumn struct {
	// Object is the base model of TermColumn
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AllowMultipleValues undocumented
	AllowMultipleValues *bool `json:"allowMultipleValues,omitempty"`
	// ShowFullyQualifiedName undocumented
	ShowFullyQualifiedName *bool `json:"showFullyQualifiedName,omitempty"`
	// ParentTerm undocumented
	ParentTerm *TermStoreTerm `json:"parentTerm,omitempty"`
	// TermSet undocumented
	TermSet *TermStoreSet `json:"termSet,omitempty"`
}

func NewTermColumn() (*TermColumn, error) {
	newTermColumn := &TermColumn{
		ODataType: "#microsoft.graph.TermColumn",
	}
	return newTermColumn, nil
}

// TermStoreGroup undocumented
type TermStoreGroup struct {
	// Entity is the base model of TermStoreGroup
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ParentSiteID undocumented
	ParentSiteID *string `json:"parentSiteId,omitempty"`
	// Scope undocumented
	Scope *TermStoreTermGroupScope `json:"scope,omitempty"`
	// Sets undocumented
	Sets []TermStoreSet `json:"sets,omitempty"`
}

func NewTermStoreGroup() (*TermStoreGroup, error) {
	newTermStoreGroup := &TermStoreGroup{
		ODataType: "#microsoft.graph.TermStoreGroup",
	}
	return newTermStoreGroup, nil
}

// TermStoreLocalizedDescription undocumented
type TermStoreLocalizedDescription struct {
	// Object is the base model of TermStoreLocalizedDescription
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// LanguageTag undocumented
	LanguageTag *string `json:"languageTag,omitempty"`
}

func NewTermStoreLocalizedDescription() (*TermStoreLocalizedDescription, error) {
	newTermStoreLocalizedDescription := &TermStoreLocalizedDescription{
		ODataType: "#microsoft.graph.TermStoreLocalizedDescription",
	}
	return newTermStoreLocalizedDescription, nil
}

// TermStoreLocalizedLabel undocumented
type TermStoreLocalizedLabel struct {
	// Object is the base model of TermStoreLocalizedLabel
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// IsDefault undocumented
	IsDefault *bool `json:"isDefault,omitempty"`
	// LanguageTag undocumented
	LanguageTag *string `json:"languageTag,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
}

func NewTermStoreLocalizedLabel() (*TermStoreLocalizedLabel, error) {
	newTermStoreLocalizedLabel := &TermStoreLocalizedLabel{
		ODataType: "#microsoft.graph.TermStoreLocalizedLabel",
	}
	return newTermStoreLocalizedLabel, nil
}

// TermStoreLocalizedName undocumented
type TermStoreLocalizedName struct {
	// Object is the base model of TermStoreLocalizedName
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// LanguageTag undocumented
	LanguageTag *string `json:"languageTag,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
}

func NewTermStoreLocalizedName() (*TermStoreLocalizedName, error) {
	newTermStoreLocalizedName := &TermStoreLocalizedName{
		ODataType: "#microsoft.graph.TermStoreLocalizedName",
	}
	return newTermStoreLocalizedName, nil
}

// TermStoreRelation undocumented
type TermStoreRelation struct {
	// Entity is the base model of TermStoreRelation
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Relationship undocumented
	Relationship *TermStoreRelationType `json:"relationship,omitempty"`
	// FromTerm undocumented
	FromTerm *TermStoreTerm `json:"fromTerm,omitempty"`
	// Set undocumented
	Set *TermStoreSet `json:"set,omitempty"`
	// ToTerm undocumented
	ToTerm *TermStoreTerm `json:"toTerm,omitempty"`
}

func NewTermStoreRelation() (*TermStoreRelation, error) {
	newTermStoreRelation := &TermStoreRelation{
		ODataType: "#microsoft.graph.TermStoreRelation",
	}
	return newTermStoreRelation, nil
}

// TermStoreSet undocumented
type TermStoreSet struct {
	// Entity is the base model of TermStoreSet
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// LocalizedNames undocumented
	LocalizedNames []TermStoreLocalizedName `json:"localizedNames,omitempty"`
	// Properties undocumented
	Properties []KeyValue `json:"properties,omitempty"`
	// Children undocumented
	Children []TermStoreTerm `json:"children,omitempty"`
	// ParentGroup undocumented
	ParentGroup *TermStoreGroup `json:"parentGroup,omitempty"`
	// Relations undocumented
	Relations []TermStoreRelation `json:"relations,omitempty"`
	// Terms undocumented
	Terms []TermStoreTerm `json:"terms,omitempty"`
}

func NewTermStoreSet() (*TermStoreSet, error) {
	newTermStoreSet := &TermStoreSet{
		ODataType: "#microsoft.graph.TermStoreSet",
	}
	return newTermStoreSet, nil
}

// TermStoreStore undocumented
type TermStoreStore struct {
	// Entity is the base model of TermStoreStore
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// DefaultLanguageTag undocumented
	DefaultLanguageTag *string `json:"defaultLanguageTag,omitempty"`
	// LanguageTags undocumented
	LanguageTags []string `json:"languageTags,omitempty"`
	// Groups undocumented
	Groups []TermStoreGroup `json:"groups,omitempty"`
	// Sets undocumented
	Sets []TermStoreSet `json:"sets,omitempty"`
}

func NewTermStoreStore() (*TermStoreStore, error) {
	newTermStoreStore := &TermStoreStore{
		ODataType: "#microsoft.graph.TermStoreStore",
	}
	return newTermStoreStore, nil
}

// TermStoreTerm undocumented
type TermStoreTerm struct {
	// Entity is the base model of TermStoreTerm
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Descriptions undocumented
	Descriptions []TermStoreLocalizedDescription `json:"descriptions,omitempty"`
	// Labels undocumented
	Labels []TermStoreLocalizedLabel `json:"labels,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Properties undocumented
	Properties []KeyValue `json:"properties,omitempty"`
	// Children undocumented
	Children []TermStoreTerm `json:"children,omitempty"`
	// Relations undocumented
	Relations []TermStoreRelation `json:"relations,omitempty"`
	// Set undocumented
	Set *TermStoreSet `json:"set,omitempty"`
}

func NewTermStoreTerm() (*TermStoreTerm, error) {
	newTermStoreTerm := &TermStoreTerm{
		ODataType: "#microsoft.graph.TermStoreTerm",
	}
	return newTermStoreTerm, nil
}
