
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// RecentNotebook undocumented
type RecentNotebook struct {
    // Object is the base model of RecentNotebook
    Object
    // DisplayName undocumented
    DisplayName *string `json:"displayName,omitempty"`
    // LastAccessedTime undocumented
    LastAccessedTime *time.Time `json:"lastAccessedTime,omitempty"`
    // Links undocumented
    Links *RecentNotebookLinks `json:"links,omitempty"`
    // SourceService undocumented
    SourceService *OnenoteSourceService `json:"sourceService,omitempty"`
}

// RecentNotebookLinks undocumented
type RecentNotebookLinks struct {
    // Object is the base model of RecentNotebookLinks
    Object
    // OneNoteClientURL undocumented
    OneNoteClientURL *ExternalLink `json:"oneNoteClientUrl,omitempty"`
    // OneNoteWebURL undocumented
    OneNoteWebURL *ExternalLink `json:"oneNoteWebUrl,omitempty"`
}
