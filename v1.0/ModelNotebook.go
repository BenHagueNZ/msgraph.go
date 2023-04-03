// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Notebook undocumented
type Notebook struct {
	// OnenoteEntityHierarchyModel is the base model of Notebook
	OnenoteEntityHierarchyModel

	ODataType string `json:"@odata.type"`
	// IsDefault undocumented
	IsDefault *bool `json:"isDefault,omitempty"`
	// IsShared undocumented
	IsShared *bool `json:"isShared,omitempty"`
	// Links undocumented
	Links *NotebookLinks `json:"links,omitempty"`
	// SectionGroupsURL undocumented
	SectionGroupsURL *string `json:"sectionGroupsUrl,omitempty"`
	// SectionsURL undocumented
	SectionsURL *string `json:"sectionsUrl,omitempty"`
	// UserRole undocumented
	UserRole *OnenoteUserRole `json:"userRole,omitempty"`
	// SectionGroups undocumented
	SectionGroups []SectionGroup `json:"sectionGroups,omitempty"`
	// Sections undocumented
	Sections []OnenoteSection `json:"sections,omitempty"`
}

func NewNotebook() (*Notebook, error) {
	newNotebook := &Notebook{
		ODataType: "#microsoft.graph.Notebook",
	}
	return newNotebook, nil
}

// NotebookLinks undocumented
type NotebookLinks struct {
	// Object is the base model of NotebookLinks
	Object

	ODataType string `json:"@odata.type"`
	// OneNoteClientURL undocumented
	OneNoteClientURL *ExternalLink `json:"oneNoteClientUrl,omitempty"`
	// OneNoteWebURL undocumented
	OneNoteWebURL *ExternalLink `json:"oneNoteWebUrl,omitempty"`
}

func NewNotebookLinks() (*NotebookLinks, error) {
	newNotebookLinks := &NotebookLinks{
		ODataType: "#microsoft.graph.NotebookLinks",
	}
	return newNotebookLinks, nil
}
