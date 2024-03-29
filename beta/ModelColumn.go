// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ColumnDefinition undocumented
type ColumnDefinition struct {
	// Entity is the base model of ColumnDefinition
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Boolean undocumented
	Boolean *BooleanColumn `json:"boolean,omitempty"`
	// Calculated undocumented
	Calculated *CalculatedColumn `json:"calculated,omitempty"`
	// Choice undocumented
	Choice *ChoiceColumn `json:"choice,omitempty"`
	// ColumnGroup undocumented
	ColumnGroup *string `json:"columnGroup,omitempty"`
	// ContentApprovalStatus undocumented
	ContentApprovalStatus *ContentApprovalStatusColumn `json:"contentApprovalStatus,omitempty"`
	// Currency undocumented
	Currency *CurrencyColumn `json:"currency,omitempty"`
	// DateTime undocumented
	DateTime *DateTimeColumn `json:"dateTime,omitempty"`
	// DefaultValue undocumented
	DefaultValue *DefaultColumnValue `json:"defaultValue,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// EnforceUniqueValues undocumented
	EnforceUniqueValues *bool `json:"enforceUniqueValues,omitempty"`
	// Geolocation undocumented
	Geolocation *GeolocationColumn `json:"geolocation,omitempty"`
	// Hidden undocumented
	Hidden *bool `json:"hidden,omitempty"`
	// HyperlinkOrPicture undocumented
	HyperlinkOrPicture *HyperlinkOrPictureColumn `json:"hyperlinkOrPicture,omitempty"`
	// Indexed undocumented
	Indexed *bool `json:"indexed,omitempty"`
	// IsDeletable undocumented
	IsDeletable *bool `json:"isDeletable,omitempty"`
	// IsReorderable undocumented
	IsReorderable *bool `json:"isReorderable,omitempty"`
	// IsSealed undocumented
	IsSealed *bool `json:"isSealed,omitempty"`
	// Lookup undocumented
	Lookup *LookupColumn `json:"lookup,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Number undocumented
	Number *NumberColumn `json:"number,omitempty"`
	// PersonOrGroup undocumented
	PersonOrGroup *PersonOrGroupColumn `json:"personOrGroup,omitempty"`
	// PropagateChanges undocumented
	PropagateChanges *bool `json:"propagateChanges,omitempty"`
	// ReadOnly undocumented
	ReadOnly *bool `json:"readOnly,omitempty"`
	// Required undocumented
	Required *bool `json:"required,omitempty"`
	// SourceContentType undocumented
	SourceContentType *ContentTypeInfo `json:"sourceContentType,omitempty"`
	// Term undocumented
	Term *TermColumn `json:"term,omitempty"`
	// Text undocumented
	Text *TextColumn `json:"text,omitempty"`
	// Thumbnail undocumented
	Thumbnail *ThumbnailColumn `json:"thumbnail,omitempty"`
	// Type undocumented
	Type *ColumnTypes `json:"type,omitempty"`
	// Validation undocumented
	Validation *ColumnValidation `json:"validation,omitempty"`
	// SourceColumn undocumented
	SourceColumn *ColumnDefinition `json:"sourceColumn,omitempty"`
}

func NewColumnDefinition() (*ColumnDefinition, error) {
	newColumnDefinition := &ColumnDefinition{
		ODataType: "#microsoft.graph.ColumnDefinition",
	}
	return newColumnDefinition, nil
}

// ColumnLink undocumented
type ColumnLink struct {
	// Entity is the base model of ColumnLink
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
}

func NewColumnLink() (*ColumnLink, error) {
	newColumnLink := &ColumnLink{
		ODataType: "#microsoft.graph.ColumnLink",
	}
	return newColumnLink, nil
}

// ColumnValidation undocumented
type ColumnValidation struct {
	// Object is the base model of ColumnValidation
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// DefaultLanguage undocumented
	DefaultLanguage *string `json:"defaultLanguage,omitempty"`
	// Descriptions undocumented
	Descriptions []DisplayNameLocalization `json:"descriptions,omitempty"`
	// Formula undocumented
	Formula *string `json:"formula,omitempty"`
}

func NewColumnValidation() (*ColumnValidation, error) {
	newColumnValidation := &ColumnValidation{
		ODataType: "#microsoft.graph.ColumnValidation",
	}
	return newColumnValidation, nil
}
