// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Onenote undocumented
type Onenote struct {
	// Entity is the base model of Onenote
	Entity

	ODataType string `json:"@odata.type"`
	// Notebooks undocumented
	Notebooks []Notebook `json:"notebooks,omitempty"`
	// Operations undocumented
	Operations []OnenoteOperation `json:"operations,omitempty"`
	// Pages undocumented
	Pages []OnenotePage `json:"pages,omitempty"`
	// Resources undocumented
	Resources []OnenoteResource `json:"resources,omitempty"`
	// SectionGroups undocumented
	SectionGroups []SectionGroup `json:"sectionGroups,omitempty"`
	// Sections undocumented
	Sections []OnenoteSection `json:"sections,omitempty"`
}

func NewOnenote() (*Onenote, error) {
	newOnenote := &Onenote{
		ODataType: "#microsoft.graph.Onenote",
	}
	return newOnenote, nil
}

// OnenoteEntityBaseModel undocumented
type OnenoteEntityBaseModel struct {
	// Entity is the base model of OnenoteEntityBaseModel
	Entity

	ODataType string `json:"@odata.type"`
	// Self undocumented
	Self *string `json:"self,omitempty"`
}

func NewOnenoteEntityBaseModel() (*OnenoteEntityBaseModel, error) {
	newOnenoteEntityBaseModel := &OnenoteEntityBaseModel{
		ODataType: "#microsoft.graph.OnenoteEntityBaseModel",
	}
	return newOnenoteEntityBaseModel, nil
}

// OnenoteEntityHierarchyModel undocumented
type OnenoteEntityHierarchyModel struct {
	// OnenoteEntitySchemaObjectModel is the base model of OnenoteEntityHierarchyModel
	OnenoteEntitySchemaObjectModel

	ODataType string `json:"@odata.type"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}

func NewOnenoteEntityHierarchyModel() (*OnenoteEntityHierarchyModel, error) {
	newOnenoteEntityHierarchyModel := &OnenoteEntityHierarchyModel{
		ODataType: "#microsoft.graph.OnenoteEntityHierarchyModel",
	}
	return newOnenoteEntityHierarchyModel, nil
}

// OnenoteEntitySchemaObjectModel undocumented
type OnenoteEntitySchemaObjectModel struct {
	// OnenoteEntityBaseModel is the base model of OnenoteEntitySchemaObjectModel
	OnenoteEntityBaseModel

	ODataType string `json:"@odata.type"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
}

func NewOnenoteEntitySchemaObjectModel() (*OnenoteEntitySchemaObjectModel, error) {
	newOnenoteEntitySchemaObjectModel := &OnenoteEntitySchemaObjectModel{
		ODataType: "#microsoft.graph.OnenoteEntitySchemaObjectModel",
	}
	return newOnenoteEntitySchemaObjectModel, nil
}

// OnenoteOperation undocumented
type OnenoteOperation struct {
	// Operation is the base model of OnenoteOperation
	Operation

	ODataType string `json:"@odata.type"`
	// Error undocumented
	Error *OnenoteOperationError `json:"error,omitempty"`
	// PercentComplete undocumented
	PercentComplete *string `json:"percentComplete,omitempty"`
	// ResourceID undocumented
	ResourceID *string `json:"resourceId,omitempty"`
	// ResourceLocation undocumented
	ResourceLocation *string `json:"resourceLocation,omitempty"`
}

func NewOnenoteOperation() (*OnenoteOperation, error) {
	newOnenoteOperation := &OnenoteOperation{
		ODataType: "#microsoft.graph.OnenoteOperation",
	}
	return newOnenoteOperation, nil
}

// OnenoteOperationError undocumented
type OnenoteOperationError struct {
	// Object is the base model of OnenoteOperationError
	Object

	ODataType string `json:"@odata.type"`
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// Message undocumented
	Message *string `json:"message,omitempty"`
}

func NewOnenoteOperationError() (*OnenoteOperationError, error) {
	newOnenoteOperationError := &OnenoteOperationError{
		ODataType: "#microsoft.graph.OnenoteOperationError",
	}
	return newOnenoteOperationError, nil
}

// OnenotePage undocumented
type OnenotePage struct {
	// OnenoteEntitySchemaObjectModel is the base model of OnenotePage
	OnenoteEntitySchemaObjectModel

	ODataType string `json:"@odata.type"`
	// Content undocumented
	Content *Stream `json:"content,omitempty"`
	// ContentURL undocumented
	ContentURL *string `json:"contentUrl,omitempty"`
	// CreatedByAppID undocumented
	CreatedByAppID *string `json:"createdByAppId,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Level undocumented
	Level *int `json:"level,omitempty"`
	// Links undocumented
	Links *PageLinks `json:"links,omitempty"`
	// Order undocumented
	Order *int `json:"order,omitempty"`
	// Title undocumented
	Title *string `json:"title,omitempty"`
	// UserTags undocumented
	UserTags []string `json:"userTags,omitempty"`
	// ParentNotebook undocumented
	ParentNotebook *Notebook `json:"parentNotebook,omitempty"`
	// ParentSection undocumented
	ParentSection *OnenoteSection `json:"parentSection,omitempty"`
}

func NewOnenotePage() (*OnenotePage, error) {
	newOnenotePage := &OnenotePage{
		ODataType: "#microsoft.graph.OnenotePage",
	}
	return newOnenotePage, nil
}

// OnenotePagePreview undocumented
type OnenotePagePreview struct {
	// Object is the base model of OnenotePagePreview
	Object

	ODataType string `json:"@odata.type"`
	// Links undocumented
	Links *OnenotePagePreviewLinks `json:"links,omitempty"`
	// PreviewText undocumented
	PreviewText *string `json:"previewText,omitempty"`
}

func NewOnenotePagePreview() (*OnenotePagePreview, error) {
	newOnenotePagePreview := &OnenotePagePreview{
		ODataType: "#microsoft.graph.OnenotePagePreview",
	}
	return newOnenotePagePreview, nil
}

// OnenotePagePreviewLinks undocumented
type OnenotePagePreviewLinks struct {
	// Object is the base model of OnenotePagePreviewLinks
	Object

	ODataType string `json:"@odata.type"`
	// PreviewImageURL undocumented
	PreviewImageURL *ExternalLink `json:"previewImageUrl,omitempty"`
}

func NewOnenotePagePreviewLinks() (*OnenotePagePreviewLinks, error) {
	newOnenotePagePreviewLinks := &OnenotePagePreviewLinks{
		ODataType: "#microsoft.graph.OnenotePagePreviewLinks",
	}
	return newOnenotePagePreviewLinks, nil
}

// OnenotePatchContentCommand undocumented
type OnenotePatchContentCommand struct {
	// Object is the base model of OnenotePatchContentCommand
	Object

	ODataType string `json:"@odata.type"`
	// Action undocumented
	Action *OnenotePatchActionType `json:"action,omitempty"`
	// Content undocumented
	Content *string `json:"content,omitempty"`
	// Position undocumented
	Position *OnenotePatchInsertPosition `json:"position,omitempty"`
	// Target undocumented
	Target *string `json:"target,omitempty"`
}

func NewOnenotePatchContentCommand() (*OnenotePatchContentCommand, error) {
	newOnenotePatchContentCommand := &OnenotePatchContentCommand{
		ODataType: "#microsoft.graph.OnenotePatchContentCommand",
	}
	return newOnenotePatchContentCommand, nil
}

// OnenoteResource undocumented
type OnenoteResource struct {
	// OnenoteEntityBaseModel is the base model of OnenoteResource
	OnenoteEntityBaseModel

	ODataType string `json:"@odata.type"`
	// Content undocumented
	Content *Stream `json:"content,omitempty"`
	// ContentURL undocumented
	ContentURL *string `json:"contentUrl,omitempty"`
}

func NewOnenoteResource() (*OnenoteResource, error) {
	newOnenoteResource := &OnenoteResource{
		ODataType: "#microsoft.graph.OnenoteResource",
	}
	return newOnenoteResource, nil
}

// OnenoteSection undocumented
type OnenoteSection struct {
	// OnenoteEntityHierarchyModel is the base model of OnenoteSection
	OnenoteEntityHierarchyModel

	ODataType string `json:"@odata.type"`
	// IsDefault undocumented
	IsDefault *bool `json:"isDefault,omitempty"`
	// Links undocumented
	Links *SectionLinks `json:"links,omitempty"`
	// PagesURL undocumented
	PagesURL *string `json:"pagesUrl,omitempty"`
	// Pages undocumented
	Pages []OnenotePage `json:"pages,omitempty"`
	// ParentNotebook undocumented
	ParentNotebook *Notebook `json:"parentNotebook,omitempty"`
	// ParentSectionGroup undocumented
	ParentSectionGroup *SectionGroup `json:"parentSectionGroup,omitempty"`
}

func NewOnenoteSection() (*OnenoteSection, error) {
	newOnenoteSection := &OnenoteSection{
		ODataType: "#microsoft.graph.OnenoteSection",
	}
	return newOnenoteSection, nil
}
